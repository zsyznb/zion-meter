package core

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/dylenfu/zion-meter/config"
	"github.com/dylenfu/zion-meter/pkg/log"
	"github.com/dylenfu/zion-meter/pkg/sdk"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ETH1, _  = new(big.Int).SetString("1000000000000000000", 10)
	gasUsage = new(big.Int).Mul(big.NewInt(1), ETH1)
)

// TPS try to test hotstuff tps, params nodeList represents multiple ethereum rpc url addresses,
// and num denote that this test will use multi account to send simple transaction
func TPS() bool {
	log.Info("start to handle tps")

	// generate master account
	log.Info("try to generate master account...")
	num, master, err := masterAccount()
	log.Info("num:%v", num)
	if err != nil {
		log.Errorf("load master account failed, err: %v", err)
		return false
	}
	log.Split("generate master account success!")

	// create account
	log.Info("try to generate multi test accounts...")
	total, groups, err := generateAccounts()
	if err != nil {
		log.Errorf("generate multi testing accounts failed, err: %v", err)
		return false
	}
	log.Split("generated multi test accounts success!")

	// leader reset contract
	leader := IsLeader()
	contract := common.HexToAddress(config.Conf.Contract)
	startTime := uint64(time.Now().Unix())
	if leader {
		log.Info("try to reset contract...")
		if contract, err = resetOrDeployContract(master, contract, startTime); err != nil {
			log.Errorf("reset test contract failed, err: %v", err)
			return false
		}
		log.Splitf("reset contract %s success", contract.Hex())
	}

	box := &Box{
		startTps:   config.MinTPS,
		startTime:  startTime,
		contract:   contract,
		userCnt:    total,
		master:     master,
		groups:     groups,
		txsRecords: make(map[int]int),
		quit:       make(chan struct{}),
	}

	// prepare balance
	log.Info("try to prepare test accounts balance...")
	if err := box.Deposit(); err != nil {
		log.Errorf("prepare testing accounts balance failed, err: %v", err)
		return false
	}
	log.Split("prepare test accounts balance success!")

	// send transactions continuously and calculate tps
	log.Infof("start to send tx and calculate tps...")
	lastTime, err := time.ParseDuration(config.Conf.LastTime)
	if err != nil {
		log.Errorf("parse last time failed, err: %v", err)
		return false
	}
	box.Start()
	time.AfterFunc(lastTime, func() {
		box.Stop()
	})
	if IsLeader() {
		log.Infof("the first machine will calculate tps")
		go box.Simulate()
		box.CalculateTPS()
	} else {
		log.Infof("machine do not need to calculate tps")
		box.Simulate()
	}

	return true
}

const userSigChanCap = 100

type User struct {
	accs []*sdk.Account
	sig  chan struct{}
	halt chan struct{}
	lock int32
	quit chan struct{}
	flag int
}

func newUser(accs []*sdk.Account) *User {
	return &User{
		accs: accs,
		sig:  make(chan struct{}, userSigChanCap),
		quit: make(chan struct{}, 1),
		flag: 0,
	}
}

func (u *User) run(contract common.Address) {
	for {
		select {
		case <-u.sig:
			if tx, nonce, err := u.accs[u.flag].Add(contract); err != nil {
				log.Errorf("user: %v,send tx %s failed, nonce %d, err: %v", u.flag, tx.Hex(), nonce, err)
				atomic.StoreInt32(&u.lock, 1)
				u.accs[u.flag].ResetNonce(nonce)
				u.flag = (u.flag + 1) % config.Conf.AccoutsPerUser
				time.Sleep(240 * time.Second)
				atomic.StoreInt32(&u.lock, 0)
			}
		case <-u.quit:
			return
		}
	}
}

type Box struct {
	startTps   int
	startTime  uint64
	master     *sdk.Account
	contract   common.Address
	groups     [][]*User
	txsRecords map[int]int
	userCnt    int
	quit       chan struct{}
}

func (b *Box) Deposit() error {
	allUser := make(map[common.Address]struct{})

	// 先用master给某个新用户转出本次测试总共所需要的gas
	userCnt := config.Conf.Groups * config.Conf.UsrsPerGroup
	sender, err := singleAccount()
	if err != nil {
		return err
	}
	amount := new(big.Int).Add(new(big.Int).Mul(big.NewInt(int64(userCnt)), gasUsage), ETH1)
	amount.Mul(amount, big.NewInt(int64(config.Conf.AccoutsPerUser)))
	if _, err := b.master.TransferWithConfirm(sender.Address(), amount); err != nil {
		return err
	}
	time.Sleep(6 * time.Second)

	// 然后由新用户独立给每个用户转gasusage
	for _, group := range b.groups {
		for _, user := range group {
			for i := 0; i < config.Conf.AccoutsPerUser; i++ {
				if _, err := sender.Transfer(user.accs[i].Address(), gasUsage); err != nil {
					return err
				}
				allUser[user.accs[i].Address()] = struct{}{}
			}
		}
	}

	time.Sleep(10 * time.Second)

retry:
	for addr, _ := range allUser {
		balance, err := b.master.BalanceOf(addr, nil)
		if err != nil {
			return err
		}
		if balance.Cmp(gasUsage) >= 0 {
			delete(allUser, addr)
		}
	}

	if len(allUser) > 0 {
		time.Sleep(5 * time.Second)
		log.Infof("there are %d account need to preparing", len(allUser))
		goto retry
	}

	return nil
}

// 每个用户在接收到信号后发送交易
func (b *Box) Start() {
	for _, group := range b.groups {
		for _, user := range group {
			go user.run(b.contract)
		}
	}
}

func (b *Box) Stop() {
	for _, group := range b.groups {
		for _, user := range group {
			close(user.quit)
		}
	}
	close(b.quit)
}

// 每秒选择一组用户发送信号
func (b *Box) Simulate() {
	// reset start time
	b.startTime = uint64(time.Now().Unix())

	ticker := time.NewTicker(1 * time.Second)
	cnt := 0
	for {
		select {
		case <-ticker.C:
			idx := cnt % len(b.groups)
			group := b.groups[idx]

			// 交易量逐步增加
			if _, ok := b.txsRecords[idx]; !ok {
				b.txsRecords[idx] = b.startTps
			} else {
				b.txsRecords[idx] += 1
			}
			txn := b.txsRecords[idx]
			if txn > len(group) {
				txn = len(group)
			}

			for _, user := range group[:txn] {
				if user.lock == 0 {
					user.sig <- struct{}{}
				}
			}

			cnt += 1
			log.Infof("group: %d transaction: %d", idx, txn)
		case <-b.quit:
			return
		}
	}
}

func (b *Box) CalculateTPS() {
	// n组全部轮完算一轮tps
	ticker := time.NewTicker(time.Duration(config.Conf.Groups) * time.Second)

	maxTPS := uint64(0)
	lastTxn := uint64(0)
	lastEndTime := uint64(time.Now().Unix())
	for {
		select {
		case <-ticker.C:
			total, err := b.master.TxNum(b.contract)
			if err != nil {
				log.Errorf("get total tx number failed, err: %v", err)
			} else if total == 0 {
				log.Debugf("get total tx number 0")
			} else {
				endTime := uint64(time.Now().Unix())
				txAdded := total - lastTxn
				duration := endTime - lastEndTime
				tps := txAdded / duration
				if tps >= maxTPS {
					maxTPS = tps
				}
				average := total / (endTime - b.startTime)
				log.Infof("start time %d, end time %d, total tx number %d, duration %d, tx added %d, tps %d,max tps %d, average %d",
					lastEndTime, endTime, total, duration, txAdded, tps, maxTPS, average)
				lastTxn = total
				lastEndTime = endTime
			}
		case <-b.quit:
			return
		}
	}
}

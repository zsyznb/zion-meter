package core

import (
	"math/big"
	"time"

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
	master, err := masterAccount()
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

	// deploy contract
	startTime := uint64(time.Now().Unix())
	log.Info("try to deploy contract...")
	contract, err := master.Deploy(startTime)
	if err != nil {
		log.Errorf("deploy test contract failed, err: %v", err)
		return false
	}
	log.Splitf("deploy contract %s success", contract.Hex())

	box := &Box{
		startTime: startTime,
		contract:  contract,
		userCnt:   total,
		master:    master,
		groups:    groups,
		quit:      make(chan struct{}),
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
	box.Start()
	go box.Simulate()
	box.CalculateTPS()

	return true
}

const userSigChanCap = 100

type User struct {
	acc  *sdk.Account
	sig  chan struct{}
	quit chan struct{}
}

func newUser(acc *sdk.Account) *User {
	return &User{
		acc:  acc,
		sig:  make(chan struct{}, userSigChanCap),
		quit: make(chan struct{}, 1),
	}
}

func (u *User) run(contract common.Address) {
	for {
		select {
		case <-u.sig:
			if _, err := u.acc.Add(contract); err != nil {
				log.Errorf("send tx failed, err: %v", err)
			}
		case <-u.quit:
			return
		}
	}
}

type Box struct {
	startTime uint64
	master    *sdk.Account
	contract  common.Address
	groups    [][]*User
	userCnt   int
	quit      chan struct{}
}

func (b *Box) Deposit() error {
	allUser := make(map[common.Address]struct{})
	for _, group := range b.groups {
		for _, user := range group {
			if _, err := b.master.Transfer(user.acc.Address(), gasUsage); err != nil {
				return err
			}
			allUser[user.acc.Address()] = struct{}{}
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
	ticker := time.NewTicker(1 * time.Second)
	cnt := 0
	for {
		select {
		case <-ticker.C:
			group := b.groups[cnt%len(b.groups)]
			for _, user := range group {
				user.sig <- struct{}{}
			}
			cnt += 1
		case <-b.quit:
			return
		}
	}
}

func (b *Box) CalculateTPS() {
	getTotal := func() (uint64, error) {
		return b.master.TxNum(b.contract)
	}

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			total, err := getTotal()
			if err != nil {
				log.Errorf("get total tx number failed, err: %v", err)
			} else {
				endTime := uint64(time.Now().Unix())
				tps := total / (endTime - b.startTime)
				log.Infof("start time %d, end time %d, total tx number %d, tps %d", b.startTime, endTime, total, tps)
			}
		case <-b.quit:
			return
		}
	}
}

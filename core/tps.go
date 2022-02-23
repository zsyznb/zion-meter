package core

import (
	"math/big"
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
	master, err := masterAccount()
	if err != nil {
		log.Errorf("load master account failed, err: %v", err)
		return false
	}
	log.Split("generate master account success!")

	// deploy contract
	startTime := uint64(time.Now().Unix())
	log.Info("try to deploy contract...")
	contract, err := master.Deploy(startTime)
	if err != nil {
		log.Errorf("deploy test contract failed, err: %v", err)
		return false
	}
	log.Splitf("deploy contract %s success", contract.Hex())

	// create account
	log.Info("try to generate multi test accounts...")
	accounts, err := generateAccounts()
	if err != nil {
		log.Errorf("generate multi testing accounts failed, err: %v", err)
		return false
	}
	log.Split("generated multi test accounts success!")

	// prepare balance
	log.Info("try to prepare test accounts balance...")
	if err := prepareTestingAccountsBalance(master, accounts); err != nil {
		log.Errorf("prepare testing accounts balance failed, err: %v", err)
		return false
	}
	log.Split("prepare test accounts balance success!")

	// send transactions continuously
	log.Infof("start to send tx and calculate tps...")
	for _, acc := range accounts {
		go sendTxs(acc, contract)
	}

	// calculate and print tps
	calculateTPS(master, contract, startTime)
	return true
}

func prepareTestingAccountsBalance(master *sdk.Account, accounts []*sdk.Account) error {
	for _, acc := range accounts {
		if _, err := master.Transfer(acc.Address(), gasUsage); err != nil {
			return err
		}
	}
	time.Sleep(15 * time.Second)
	for _, acc := range accounts {
		balance, err := master.BalanceOf(acc.Address(), nil)
		if err != nil {
			return err
		}
		log.Infof("%s balance %s", acc.Address().Hex(), balance.String())
	}
	return nil
}

func sendTxs(acc *sdk.Account, contract common.Address) {
	ticker := time.NewTicker(time.Second)
	txn := config.Conf.TxsPerSecond
	for range ticker.C {
		for i := 0; i < txn; i++ {
			if _, err := acc.Add(contract); err != nil {
				log.Errorf("send tx failed, err: %v", err)
			}
		}
	}
}

func calculateTPS(master *sdk.Account, contract common.Address, startTime uint64) {
	getTotal := func() (uint64, error) {
		return master.TxNum(contract)
	}

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			total, err := getTotal()
			if err != nil {
				log.Errorf("get total tx number failed, err: %v", err)
			} else {
				endTime := uint64(time.Now().Unix())
				tps := total / (endTime - startTime)
				log.Infof("start time %d, end time %d, total tx number %d, tps %d", startTime, endTime, total, tps)
			}
		}
	}
}

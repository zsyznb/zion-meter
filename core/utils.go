package core

import (
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dylenfu/zion-meter/config"
	"github.com/dylenfu/zion-meter/pkg/sdk"
	"github.com/ethereum/go-ethereum/common"
)

func masterAccount() (int, *sdk.Account, error) {
	chainID := config.Conf.ChainID
	node := config.Conf.Nodes[0]
	sender, err := sdk.NewSender(node, chainID)
	if err != nil {
		return 1, nil, err
	}
	filename := filepath.Base(config.Conf.Chainspace)
	num, _ := strconv.Atoi(filename[len(filename)-1:])
	acc, err := sdk.MasterAccount(sender, config.Conf.NodeKey[num])
	return num, acc, err
}

func singleAccount() (*sdk.Account, error) {
	chainID := config.Conf.ChainID
	node := config.Conf.Nodes[0]
	sender, err := sdk.NewSender(node, chainID)
	if err != nil {
		return nil, err
	}
	account, err := sdk.NewAccount()
	if err != nil {
		return nil, err
	}
	account.SetSender(sender)
	return account, nil
}

// 每组N个账户，总共G组
func generateAccounts() (int, [][]*User, error) {
	chainID := config.Conf.ChainID
	connNum := config.Conf.UsrsPerGroup
	nodeNum := len(config.Conf.Nodes)
	senderList := make([][]*sdk.Sender, 0)
	for i := 0; i < connNum; i++ {
		node := config.Conf.Nodes[i%nodeNum]
		senders := make([]*sdk.Sender, 0)
		for j := 0; j < config.Conf.AccoutsPerUser; j++ {
			sender, err := sdk.NewSender(node, chainID)
			if err != nil {
				return 0, nil, err
			}
			senders = append(senders, sender)
		}

		senderList = append(senderList, senders)
	}

	groupNo := config.Conf.Groups
	AccNoPerGroup := config.Conf.UsrsPerGroup
	total := groupNo * AccNoPerGroup
	list := make([][]*User, 0)
	for i := 0; i < groupNo; i++ {
		group := make([]*User, 0)
		for j := 0; j < AccNoPerGroup; j++ {
			accs := sdk.NewAccounts(config.Conf.AccoutsPerUser)
			index := (i*AccNoPerGroup + j) % connNum
			for num, acc := range accs {
				acc.SetSender(senderList[index][num])
			}
			user := newUser(accs)
			group = append(group, user)
		}
		list = append(list, group)
	}

	return total, list, nil
}

func resetOrDeployContract(acc *sdk.Account, oldContract common.Address, startTime uint64) (common.Address, error) {
	if acc.Exist(oldContract) {
		if _, err := acc.Reset(oldContract, startTime); err != nil {
			return common.EmptyAddress, err
		}
		return oldContract, nil
	}

	newContract, err := acc.Deploy()
	if err != nil {
		return common.EmptyAddress, err
	}
	if _, err := acc.Reset(newContract, startTime); err != nil {
		return newContract, err
	}
	return newContract, nil
}

func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}

// 如果不需要多台机器，则直接返回true，本机需要统计tps;
// 如果需要多台机器，则判断本机是否为第一台机器地址，是则返回true
func IsLeader() bool {
	if !config.Conf.Sharding {
		return true
	}

	local := InternalIP()
	if local == config.Conf.FirstMachine {
		return true
	}

	return false
}

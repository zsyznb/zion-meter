package core

import (
	"github.com/dylenfu/zion-meter/config"
	"github.com/dylenfu/zion-meter/pkg/sdk"
)

func masterAccount() (*sdk.Account, error) {
	chainID := config.Conf.ChainID
	node := config.Conf.Nodes[0]
	sender, err := sdk.NewSender(node.Url, chainID)
	if err != nil {
		return nil, err
	}
	return sdk.MasterAccount(sender, node.NodeKey)
}

// 每组N个账户，总共G组
func generateAccounts() (int, [][]*User, error) {
	chainID := config.Conf.ChainID
	connNum := config.Conf.ConnNum
	nodeNum := len(config.Conf.Nodes)
	senderList := make([]*sdk.Sender, 0)
	for i := 0; i < connNum; i++ {
		node := config.Conf.Nodes[i%nodeNum]
		sender, err := sdk.NewSender(node.Url, chainID)
		if err != nil {
			return 0, nil, err
		}
		senderList = append(senderList, sender)
	}

	groupNo := config.Conf.AccountGroupNum
	AccNoPerGroup := config.Conf.AccountNumPerGroup
	total := groupNo * AccNoPerGroup
	list := make([][]*User, 0)
	for i := 0; i < groupNo; i++ {
		group := make([]*User, 0)
		for j := 0; j < AccNoPerGroup; j++ {
			acc, err := sdk.NewAccount()
			if err != nil {
				return 0, nil, err
			}
			index := (i*AccNoPerGroup + j) % connNum
			acc.SetSender(senderList[index])
			user := newUser(acc)
			group = append(group, user)
		}
		list = append(list, group)
	}

	return total, list, nil
}

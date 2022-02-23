package core

import (
	"github.com/dylenfu/zion-meter/config"
	"github.com/dylenfu/zion-meter/pkg/sdk"
)

func masterAccount() (*sdk.Account, error) {
	chainID := config.Conf.ChainID
	node := config.Conf.Nodes[0]
	return sdk.MasterAccount(node.Url, chainID, node.NodeKey)
}

func generateAccounts() ([]*sdk.Account, error) {
	chainID := config.Conf.ChainID
	num := config.Conf.AccsNum
	list := make([]*sdk.Account, 0)
	nodesNum := len(config.Conf.Nodes)
	for i := 0; i < num; i++ {
		node := config.Conf.Nodes[i%nodesNum]
		acc, err := sdk.NewAccount(node.Url, chainID)
		if err != nil {
			return nil, err
		}
		list = append(list, acc)
	}
	return list, nil
}

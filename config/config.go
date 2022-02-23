/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/dylenfu/zion-meter/pkg/files"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Conf *Config
)

type Config struct {
	Workspace    string
	ChainID      uint64
	Groups       int // 账户分组，总组数为`Groups`，每秒切换不同组用户发送交易
	AccsPerGroup int // 每组账户数量

	Nodes []*Node
}

type Node struct {
	NodeKey    string            `json:"NodeKey"`
	Url        string            `json:"Url"`
	Address    common.Address    `json:"Address,omitempty"`
	PrivateKey *ecdsa.PrivateKey `json:"PrivateKey,omitempty"`
	PublicKey  *ecdsa.PublicKey  `json:"PublicKey,omitempty"`
}

func LoadConfig(filepath string) {
	data, err := files.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &Conf); err != nil {
		panic(err)
	}

	for _, v := range Conf.Nodes {
		key := v.NodeKey
		if !strings.Contains(key, "0x") {
			key = "0x" + key
		}

		enc, err := hexutil.Decode(key)
		if err != nil {
			panic(err)
		}

		privKey, err := crypto.ToECDSA(enc)
		if err != nil {
			panic(err)
		}
		v.PrivateKey = privKey
		v.PublicKey = &privKey.PublicKey
		v.Address = crypto.PubkeyToAddress(*v.PublicKey)
	}
}

func LoadParams(fileName string, data interface{}) error {
	filepath := files.FullPath(Conf.Workspace, "cases", fileName)
	bz, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bz, data)
}

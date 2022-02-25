// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stat_abi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var (
	MethodAdd = "add"

	MethodReset = "reset"

	MethodStartTime = "startTime"

	MethodTxNum = "txNum"
)

// StatABI is the input ABI used to generate the binding from.
const StatABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startTime\",\"type\":\"uint64\"}],\"name\":\"reset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"txNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StatFuncSigs maps the 4-byte function signature to its string representation.
var StatFuncSigs = map[string]string{
	"4f2be91f": "add()",
	"e365981c": "reset(uint64)",
	"78e97925": "startTime()",
	"3a4ef544": "txNum()",
}

// StatBin is the compiled bytecode used for deploying new contracts.
var StatBin = "0x608060405234801561001057600080fd5b50610130806100206000396000f3fe6080604052348015600f57600080fd5b506004361060465760003560e01c80633a4ef54414604b5780634f2be91f14606357806378e9792514606b578063e365981c14608e575b600080fd5b605160b2565b60408051918252519081900360200190f35b606960b8565b005b607160c3565b6040805167ffffffffffffffff9092168252519081900360200190f35b60696004803603602081101560a257600080fd5b503567ffffffffffffffff1660d3565b60005481565b600080546001019055565b60015467ffffffffffffffff1681565b6001805467ffffffffffffffff191667ffffffffffffffff929092169190911790556000805556fea265627a7a72315820827e69ce9ab1a1c925abcf663cc31e5d94e26b78ad14e0035ca9eeeabfeef7d564736f6c63430005110032"

// DeployStat deploys a new Ethereum contract, binding an instance of Stat to it.
func DeployStat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stat, error) {
	parsed, err := abi.JSON(strings.NewReader(StatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// Stat is an auto generated Go binding around an Ethereum contract.
type Stat struct {
	StatCaller     // Read-only binding to the contract
	StatTransactor // Write-only binding to the contract
	StatFilterer   // Log filterer for contract events
}

// StatCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatSession struct {
	Contract     *Stat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatCallerSession struct {
	Contract *StatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatTransactorSession struct {
	Contract     *StatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatRaw struct {
	Contract *Stat // Generic contract binding to access the raw methods on
}

// StatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatCallerRaw struct {
	Contract *StatCaller // Generic read-only contract binding to access the raw methods on
}

// StatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatTransactorRaw struct {
	Contract *StatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStat creates a new instance of Stat, bound to a specific deployed contract.
func NewStat(address common.Address, backend bind.ContractBackend) (*Stat, error) {
	contract, err := bindStat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// NewStatCaller creates a new read-only instance of Stat, bound to a specific deployed contract.
func NewStatCaller(address common.Address, caller bind.ContractCaller) (*StatCaller, error) {
	contract, err := bindStat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatCaller{contract: contract}, nil
}

// NewStatTransactor creates a new write-only instance of Stat, bound to a specific deployed contract.
func NewStatTransactor(address common.Address, transactor bind.ContractTransactor) (*StatTransactor, error) {
	contract, err := bindStat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatTransactor{contract: contract}, nil
}

// NewStatFilterer creates a new log filterer instance of Stat, bound to a specific deployed contract.
func NewStatFilterer(address common.Address, filterer bind.ContractFilterer) (*StatFilterer, error) {
	contract, err := bindStat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatFilterer{contract: contract}, nil
}

// bindStat binds a generic wrapper to an already deployed contract.
func bindStat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.StatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transact(opts, method, params...)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Stat *StatCaller) StartTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Stat *StatSession) StartTime() (uint64, error) {
	return _Stat.Contract.StartTime(&_Stat.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Stat *StatCallerSession) StartTime() (uint64, error) {
	return _Stat.Contract.StartTime(&_Stat.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Stat *StatCaller) TxNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "txNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Stat *StatSession) TxNum() (*big.Int, error) {
	return _Stat.Contract.TxNum(&_Stat.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Stat *StatCallerSession) TxNum() (*big.Int, error) {
	return _Stat.Contract.TxNum(&_Stat.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Stat *StatTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Stat *StatSession) Add() (*types.Transaction, error) {
	return _Stat.Contract.Add(&_Stat.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Stat *StatTransactorSession) Add() (*types.Transaction, error) {
	return _Stat.Contract.Add(&_Stat.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Stat *StatTransactor) Reset(opts *bind.TransactOpts, _startTime uint64) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "reset", _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Stat *StatSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _Stat.Contract.Reset(&_Stat.TransactOpts, _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Stat *StatTransactorSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _Stat.Contract.Reset(&_Stat.TransactOpts, _startTime)
}


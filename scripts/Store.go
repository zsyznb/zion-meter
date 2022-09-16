// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scripts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ScriptsMetaData contains all meta data concerning the Scripts contract.
var ScriptsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startTime\",\"type\":\"uint64\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061028a806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633a4ef544146100515780634f2be91f1461006f57806378e9792514610079578063e365981c14610097575b600080fd5b6100596100b3565b604051610066919061013a565b60405180910390f35b6100776100b9565b005b6100816100d4565b60405161008e9190610178565b60405180910390f35b6100b160048036038101906100ac91906101c4565b6100ee565b005b60005481565b60016000808282546100cb9190610220565b92505081905550565b600160009054906101000a900467ffffffffffffffff1681565b80600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000808190555050565b6000819050919050565b61013481610121565b82525050565b600060208201905061014f600083018461012b565b92915050565b600067ffffffffffffffff82169050919050565b61017281610155565b82525050565b600060208201905061018d6000830184610169565b92915050565b600080fd5b6101a181610155565b81146101ac57600080fd5b50565b6000813590506101be81610198565b92915050565b6000602082840312156101da576101d9610193565b5b60006101e8848285016101af565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061022b82610121565b915061023683610121565b925082820190508082111561024e5761024d6101f1565b5b9291505056fea26469706673582212204958e1a8b97be55e6abf5c3efa4cdcde0a4113f7f35fface3a4f56ca0781040564736f6c63430008110033",
}

// ScriptsABI is the input ABI used to generate the binding from.
// Deprecated: Use ScriptsMetaData.ABI instead.
var ScriptsABI = ScriptsMetaData.ABI

// ScriptsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScriptsMetaData.Bin instead.
var ScriptsBin = ScriptsMetaData.Bin

// DeployScripts deploys a new Ethereum contract, binding an instance of Scripts to it.
func DeployScripts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Scripts, error) {
	parsed, err := ScriptsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScriptsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Scripts{ScriptsCaller: ScriptsCaller{contract: contract}, ScriptsTransactor: ScriptsTransactor{contract: contract}, ScriptsFilterer: ScriptsFilterer{contract: contract}}, nil
}

// Scripts is an auto generated Go binding around an Ethereum contract.
type Scripts struct {
	ScriptsCaller     // Read-only binding to the contract
	ScriptsTransactor // Write-only binding to the contract
	ScriptsFilterer   // Log filterer for contract events
}

// ScriptsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScriptsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScriptsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScriptsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScriptsSession struct {
	Contract     *Scripts          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScriptsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScriptsCallerSession struct {
	Contract *ScriptsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ScriptsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScriptsTransactorSession struct {
	Contract     *ScriptsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ScriptsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScriptsRaw struct {
	Contract *Scripts // Generic contract binding to access the raw methods on
}

// ScriptsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScriptsCallerRaw struct {
	Contract *ScriptsCaller // Generic read-only contract binding to access the raw methods on
}

// ScriptsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScriptsTransactorRaw struct {
	Contract *ScriptsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScripts creates a new instance of Scripts, bound to a specific deployed contract.
func NewScripts(address common.Address, backend bind.ContractBackend) (*Scripts, error) {
	contract, err := bindScripts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scripts{ScriptsCaller: ScriptsCaller{contract: contract}, ScriptsTransactor: ScriptsTransactor{contract: contract}, ScriptsFilterer: ScriptsFilterer{contract: contract}}, nil
}

// NewScriptsCaller creates a new read-only instance of Scripts, bound to a specific deployed contract.
func NewScriptsCaller(address common.Address, caller bind.ContractCaller) (*ScriptsCaller, error) {
	contract, err := bindScripts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptsCaller{contract: contract}, nil
}

// NewScriptsTransactor creates a new write-only instance of Scripts, bound to a specific deployed contract.
func NewScriptsTransactor(address common.Address, transactor bind.ContractTransactor) (*ScriptsTransactor, error) {
	contract, err := bindScripts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptsTransactor{contract: contract}, nil
}

// NewScriptsFilterer creates a new log filterer instance of Scripts, bound to a specific deployed contract.
func NewScriptsFilterer(address common.Address, filterer bind.ContractFilterer) (*ScriptsFilterer, error) {
	contract, err := bindScripts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScriptsFilterer{contract: contract}, nil
}

// bindScripts binds a generic wrapper to an already deployed contract.
func bindScripts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScriptsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scripts *ScriptsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scripts.Contract.ScriptsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scripts *ScriptsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scripts.Contract.ScriptsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scripts *ScriptsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scripts.Contract.ScriptsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scripts *ScriptsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scripts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scripts *ScriptsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scripts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scripts *ScriptsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scripts.Contract.contract.Transact(opts, method, params...)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Scripts *ScriptsCaller) StartTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Scripts.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Scripts *ScriptsSession) StartTime() (uint64, error) {
	return _Scripts.Contract.StartTime(&_Scripts.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_Scripts *ScriptsCallerSession) StartTime() (uint64, error) {
	return _Scripts.Contract.StartTime(&_Scripts.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Scripts *ScriptsCaller) TxNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scripts.contract.Call(opts, &out, "txNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Scripts *ScriptsSession) TxNum() (*big.Int, error) {
	return _Scripts.Contract.TxNum(&_Scripts.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_Scripts *ScriptsCallerSession) TxNum() (*big.Int, error) {
	return _Scripts.Contract.TxNum(&_Scripts.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Scripts *ScriptsTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scripts.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Scripts *ScriptsSession) Add() (*types.Transaction, error) {
	return _Scripts.Contract.Add(&_Scripts.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Scripts *ScriptsTransactorSession) Add() (*types.Transaction, error) {
	return _Scripts.Contract.Add(&_Scripts.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Scripts *ScriptsTransactor) Reset(opts *bind.TransactOpts, _startTime uint64) (*types.Transaction, error) {
	return _Scripts.contract.Transact(opts, "reset", _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Scripts *ScriptsSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _Scripts.Contract.Reset(&_Scripts.TransactOpts, _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_Scripts *ScriptsTransactorSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _Scripts.Contract.Reset(&_Scripts.TransactOpts, _startTime)
}

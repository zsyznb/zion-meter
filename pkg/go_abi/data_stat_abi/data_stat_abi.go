// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package data_stat_abi

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
	MethodCostManyGas = "costManyGas"

	MethodReset = "reset"

	MethodStartTime = "startTime"

	MethodTxNum = "txNum"
)

// DataStatABI is the input ABI used to generate the binding from.
const DataStatABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"complexity\",\"type\":\"uint64\"}],\"name\":\"costManyGas\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startTime\",\"type\":\"uint64\"}],\"name\":\"reset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"txNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataStatFuncSigs maps the 4-byte function signature to its string representation.
var DataStatFuncSigs = map[string]string{
	"c6039136": "costManyGas(bytes,uint64)",
	"e365981c": "reset(uint64)",
	"78e97925": "startTime()",
	"3a4ef544": "txNum()",
}

// DataStatBin is the compiled bytecode used for deploying new contracts.
var DataStatBin = "0x608060405234801561001057600080fd5b506102cb806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633a4ef5441461005157806378e979251461006b578063c603913614610090578063e365981c14610144575b600080fd5b61005961016b565b60408051918252519081900360200190f35b610073610171565b6040805167ffffffffffffffff9092168252519081900360200190f35b610142600480360360408110156100a657600080fd5b8101906020810181356401000000008111156100c157600080fd5b8201836020820111156100d357600080fd5b803590602001918460018302840111640100000000831117156100f557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff1691506101819050565b005b6101426004803603602081101561015a57600080fd5b503567ffffffffffffffff166101d2565b60015481565b60025467ffffffffffffffff1681565b60005b8167ffffffffffffffff168110156101c55760008054600101808255815260036020908152604090912084516101bc928601906101fb565b50600101610184565b5050600180548101905550565b6002805467ffffffffffffffff191667ffffffffffffffff929092169190911790556000600155565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061023c57805160ff1916838001178555610269565b82800160010185558215610269579182015b8281111561026957825182559160200191906001019061024e565b50610275929150610279565b5090565b61029391905b80821115610275576000815560010161027f565b9056fea265627a7a723158208b43be0f3d2a6593ac1c0095d921bd61083ec5b6f5d0bef1674a6686cca9cf1d64736f6c63430005110032"

// DeployDataStat deploys a new Ethereum contract, binding an instance of DataStat to it.
func DeployDataStat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataStat, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataStatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataStat{DataStatCaller: DataStatCaller{contract: contract}, DataStatTransactor: DataStatTransactor{contract: contract}, DataStatFilterer: DataStatFilterer{contract: contract}}, nil
}

// DataStat is an auto generated Go binding around an Ethereum contract.
type DataStat struct {
	DataStatCaller     // Read-only binding to the contract
	DataStatTransactor // Write-only binding to the contract
	DataStatFilterer   // Log filterer for contract events
}

// DataStatCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataStatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataStatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataStatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataStatSession struct {
	Contract     *DataStat         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataStatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataStatCallerSession struct {
	Contract *DataStatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DataStatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataStatTransactorSession struct {
	Contract     *DataStatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataStatRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataStatRaw struct {
	Contract *DataStat // Generic contract binding to access the raw methods on
}

// DataStatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataStatCallerRaw struct {
	Contract *DataStatCaller // Generic read-only contract binding to access the raw methods on
}

// DataStatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataStatTransactorRaw struct {
	Contract *DataStatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataStat creates a new instance of DataStat, bound to a specific deployed contract.
func NewDataStat(address common.Address, backend bind.ContractBackend) (*DataStat, error) {
	contract, err := bindDataStat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataStat{DataStatCaller: DataStatCaller{contract: contract}, DataStatTransactor: DataStatTransactor{contract: contract}, DataStatFilterer: DataStatFilterer{contract: contract}}, nil
}

// NewDataStatCaller creates a new read-only instance of DataStat, bound to a specific deployed contract.
func NewDataStatCaller(address common.Address, caller bind.ContractCaller) (*DataStatCaller, error) {
	contract, err := bindDataStat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataStatCaller{contract: contract}, nil
}

// NewDataStatTransactor creates a new write-only instance of DataStat, bound to a specific deployed contract.
func NewDataStatTransactor(address common.Address, transactor bind.ContractTransactor) (*DataStatTransactor, error) {
	contract, err := bindDataStat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataStatTransactor{contract: contract}, nil
}

// NewDataStatFilterer creates a new log filterer instance of DataStat, bound to a specific deployed contract.
func NewDataStatFilterer(address common.Address, filterer bind.ContractFilterer) (*DataStatFilterer, error) {
	contract, err := bindDataStat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataStatFilterer{contract: contract}, nil
}

// bindDataStat binds a generic wrapper to an already deployed contract.
func bindDataStat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataStat *DataStatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataStat.Contract.DataStatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataStat *DataStatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataStat.Contract.DataStatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataStat *DataStatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataStat.Contract.DataStatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataStat *DataStatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataStat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataStat *DataStatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataStat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataStat *DataStatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataStat.Contract.contract.Transact(opts, method, params...)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataStat *DataStatCaller) StartTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DataStat.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataStat *DataStatSession) StartTime() (uint64, error) {
	return _DataStat.Contract.StartTime(&_DataStat.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataStat *DataStatCallerSession) StartTime() (uint64, error) {
	return _DataStat.Contract.StartTime(&_DataStat.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataStat *DataStatCaller) TxNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataStat.contract.Call(opts, &out, "txNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataStat *DataStatSession) TxNum() (*big.Int, error) {
	return _DataStat.Contract.TxNum(&_DataStat.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataStat *DataStatCallerSession) TxNum() (*big.Int, error) {
	return _DataStat.Contract.TxNum(&_DataStat.CallOpts)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xc6039136.
//
// Solidity: function costManyGas(bytes input, uint64 complexity) returns()
func (_DataStat *DataStatTransactor) CostManyGas(opts *bind.TransactOpts, input []byte, complexity uint64) (*types.Transaction, error) {
	return _DataStat.contract.Transact(opts, "costManyGas", input, complexity)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xc6039136.
//
// Solidity: function costManyGas(bytes input, uint64 complexity) returns()
func (_DataStat *DataStatSession) CostManyGas(input []byte, complexity uint64) (*types.Transaction, error) {
	return _DataStat.Contract.CostManyGas(&_DataStat.TransactOpts, input, complexity)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xc6039136.
//
// Solidity: function costManyGas(bytes input, uint64 complexity) returns()
func (_DataStat *DataStatTransactorSession) CostManyGas(input []byte, complexity uint64) (*types.Transaction, error) {
	return _DataStat.Contract.CostManyGas(&_DataStat.TransactOpts, input, complexity)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataStat *DataStatTransactor) Reset(opts *bind.TransactOpts, _startTime uint64) (*types.Transaction, error) {
	return _DataStat.contract.Transact(opts, "reset", _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataStat *DataStatSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _DataStat.Contract.Reset(&_DataStat.TransactOpts, _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataStat *DataStatTransactorSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _DataStat.Contract.Reset(&_DataStat.TransactOpts, _startTime)
}


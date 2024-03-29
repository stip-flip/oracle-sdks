// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.ConvertType
)

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8[8]\",\"name\":\"decimals_\",\"type\":\"uint8[8]\"},{\"internalType\":\"uint8[]\",\"name\":\"drops_\",\"type\":\"uint8[]\"},{\"internalType\":\"uint8\",\"name\":\"modulo_\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"offset\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"frequency_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"roundDuration_\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"minStake_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedBounty\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frequency\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"slot\",\"type\":\"uint8\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"lastRound\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"getRoundTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"getSlots\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"slots\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"isRoundAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"slot\",\"type\":\"uint8\"}],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"price_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"slot\",\"type\":\"uint8\"}],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"price_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"lastPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"manas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"slot\",\"type\":\"uint8\"}],\"name\":\"nextPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"price_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"priceToMana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundDuration\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prices_\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"slots\",\"type\":\"uint256[8]\"}],\"name\":\"setSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"submitters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMana\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address ) view returns(uint256)
func (_Oracle *OracleCaller) Debt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "debt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address ) view returns(uint256)
func (_Oracle *OracleSession) Debt(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Debt(&_Oracle.CallOpts, arg0)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address ) view returns(uint256)
func (_Oracle *OracleCallerSession) Debt(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Debt(&_Oracle.CallOpts, arg0)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Oracle *OracleCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Oracle *OracleSession) Description() (string, error) {
	return _Oracle.Contract.Description(&_Oracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Oracle *OracleCallerSession) Description() (string, error) {
	return _Oracle.Contract.Description(&_Oracle.CallOpts)
}

// Frequency is a free data retrieval call binding the contract method 0xead50da3.
//
// Solidity: function frequency() view returns(uint24)
func (_Oracle *OracleCaller) Frequency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "frequency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Frequency is a free data retrieval call binding the contract method 0xead50da3.
//
// Solidity: function frequency() view returns(uint24)
func (_Oracle *OracleSession) Frequency() (*big.Int, error) {
	return _Oracle.Contract.Frequency(&_Oracle.CallOpts)
}

// Frequency is a free data retrieval call binding the contract method 0xead50da3.
//
// Solidity: function frequency() view returns(uint24)
func (_Oracle *OracleCallerSession) Frequency() (*big.Int, error) {
	return _Oracle.Contract.Frequency(&_Oracle.CallOpts)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint64 currentRound)
func (_Oracle *OracleCaller) GetCurrentRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrentRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint64 currentRound)
func (_Oracle *OracleSession) GetCurrentRound() (uint64, error) {
	return _Oracle.Contract.GetCurrentRound(&_Oracle.CallOpts)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint64 currentRound)
func (_Oracle *OracleCallerSession) GetCurrentRound() (uint64, error) {
	return _Oracle.Contract.GetCurrentRound(&_Oracle.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0x7d4950d1.
//
// Solidity: function getDecimals(uint8 slot) view returns(uint8)
func (_Oracle *OracleCaller) GetDecimals(opts *bind.CallOpts, slot uint8) (uint8, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getDecimals", slot)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x7d4950d1.
//
// Solidity: function getDecimals(uint8 slot) view returns(uint8)
func (_Oracle *OracleSession) GetDecimals(slot uint8) (uint8, error) {
	return _Oracle.Contract.GetDecimals(&_Oracle.CallOpts, slot)
}

// GetDecimals is a free data retrieval call binding the contract method 0x7d4950d1.
//
// Solidity: function getDecimals(uint8 slot) view returns(uint8)
func (_Oracle *OracleCallerSession) GetDecimals(slot uint8) (uint8, error) {
	return _Oracle.Contract.GetDecimals(&_Oracle.CallOpts, slot)
}

// GetLastRound is a free data retrieval call binding the contract method 0x4231a2c3.
//
// Solidity: function getLastRound() view returns(uint64 lastRound)
func (_Oracle *OracleCaller) GetLastRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getLastRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastRound is a free data retrieval call binding the contract method 0x4231a2c3.
//
// Solidity: function getLastRound() view returns(uint64 lastRound)
func (_Oracle *OracleSession) GetLastRound() (uint64, error) {
	return _Oracle.Contract.GetLastRound(&_Oracle.CallOpts)
}

// GetLastRound is a free data retrieval call binding the contract method 0x4231a2c3.
//
// Solidity: function getLastRound() view returns(uint64 lastRound)
func (_Oracle *OracleCallerSession) GetLastRound() (uint64, error) {
	return _Oracle.Contract.GetLastRound(&_Oracle.CallOpts)
}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x95ea80a2.
//
// Solidity: function getRoundTimestamp(uint64 round) view returns(uint64)
func (_Oracle *OracleCaller) GetRoundTimestamp(opts *bind.CallOpts, round uint64) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRoundTimestamp", round)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x95ea80a2.
//
// Solidity: function getRoundTimestamp(uint64 round) view returns(uint64)
func (_Oracle *OracleSession) GetRoundTimestamp(round uint64) (uint64, error) {
	return _Oracle.Contract.GetRoundTimestamp(&_Oracle.CallOpts, round)
}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x95ea80a2.
//
// Solidity: function getRoundTimestamp(uint64 round) view returns(uint64)
func (_Oracle *OracleCallerSession) GetRoundTimestamp(round uint64) (uint64, error) {
	return _Oracle.Contract.GetRoundTimestamp(&_Oracle.CallOpts, round)
}

// GetSlots is a free data retrieval call binding the contract method 0xfc0b98b1.
//
// Solidity: function getSlots(uint256 data) pure returns(uint256[8] slots)
func (_Oracle *OracleCaller) GetSlots(opts *bind.CallOpts, data *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getSlots", data)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetSlots is a free data retrieval call binding the contract method 0xfc0b98b1.
//
// Solidity: function getSlots(uint256 data) pure returns(uint256[8] slots)
func (_Oracle *OracleSession) GetSlots(data *big.Int) ([8]*big.Int, error) {
	return _Oracle.Contract.GetSlots(&_Oracle.CallOpts, data)
}

// GetSlots is a free data retrieval call binding the contract method 0xfc0b98b1.
//
// Solidity: function getSlots(uint256 data) pure returns(uint256[8] slots)
func (_Oracle *OracleCallerSession) GetSlots(data *big.Int) ([8]*big.Int, error) {
	return _Oracle.Contract.GetSlots(&_Oracle.CallOpts, data)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(uint64)
func (_Oracle *OracleCaller) Initialized(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(uint64)
func (_Oracle *OracleSession) Initialized() (uint64, error) {
	return _Oracle.Contract.Initialized(&_Oracle.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(uint64)
func (_Oracle *OracleCallerSession) Initialized() (uint64, error) {
	return _Oracle.Contract.Initialized(&_Oracle.CallOpts)
}

// IsRoundAllowed is a free data retrieval call binding the contract method 0x4a73e507.
//
// Solidity: function isRoundAllowed(uint64 round) view returns(bool isAllowed)
func (_Oracle *OracleCaller) IsRoundAllowed(opts *bind.CallOpts, round uint64) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isRoundAllowed", round)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoundAllowed is a free data retrieval call binding the contract method 0x4a73e507.
//
// Solidity: function isRoundAllowed(uint64 round) view returns(bool isAllowed)
func (_Oracle *OracleSession) IsRoundAllowed(round uint64) (bool, error) {
	return _Oracle.Contract.IsRoundAllowed(&_Oracle.CallOpts, round)
}

// IsRoundAllowed is a free data retrieval call binding the contract method 0x4a73e507.
//
// Solidity: function isRoundAllowed(uint64 round) view returns(bool isAllowed)
func (_Oracle *OracleCallerSession) IsRoundAllowed(round uint64) (bool, error) {
	return _Oracle.Contract.IsRoundAllowed(&_Oracle.CallOpts, round)
}

// LastPrice is a free data retrieval call binding the contract method 0xb1e0ce76.
//
// Solidity: function lastPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCaller) LastPrice(opts *bind.CallOpts, round uint64, slot uint8) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "lastPrice", round, slot)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0xb1e0ce76.
//
// Solidity: function lastPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleSession) LastPrice(round uint64, slot uint8) (uint64, error) {
	return _Oracle.Contract.LastPrice(&_Oracle.CallOpts, round, slot)
}

// LastPrice is a free data retrieval call binding the contract method 0xb1e0ce76.
//
// Solidity: function lastPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCallerSession) LastPrice(round uint64, slot uint8) (uint64, error) {
	return _Oracle.Contract.LastPrice(&_Oracle.CallOpts, round, slot)
}

// LastPrice0 is a free data retrieval call binding the contract method 0xbf192e49.
//
// Solidity: function lastPrice(uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCaller) LastPrice0(opts *bind.CallOpts, slot uint8) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "lastPrice0", slot)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPrice0 is a free data retrieval call binding the contract method 0xbf192e49.
//
// Solidity: function lastPrice(uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleSession) LastPrice0(slot uint8) (uint64, error) {
	return _Oracle.Contract.LastPrice0(&_Oracle.CallOpts, slot)
}

// LastPrice0 is a free data retrieval call binding the contract method 0xbf192e49.
//
// Solidity: function lastPrice(uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCallerSession) LastPrice0(slot uint8) (uint64, error) {
	return _Oracle.Contract.LastPrice0(&_Oracle.CallOpts, slot)
}

// LastPrices is a free data retrieval call binding the contract method 0x6c695b96.
//
// Solidity: function lastPrices(uint64 ) view returns(uint256)
func (_Oracle *OracleCaller) LastPrices(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "lastPrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0x6c695b96.
//
// Solidity: function lastPrices(uint64 ) view returns(uint256)
func (_Oracle *OracleSession) LastPrices(arg0 uint64) (*big.Int, error) {
	return _Oracle.Contract.LastPrices(&_Oracle.CallOpts, arg0)
}

// LastPrices is a free data retrieval call binding the contract method 0x6c695b96.
//
// Solidity: function lastPrices(uint64 ) view returns(uint256)
func (_Oracle *OracleCallerSession) LastPrices(arg0 uint64) (*big.Int, error) {
	return _Oracle.Contract.LastPrices(&_Oracle.CallOpts, arg0)
}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(address ) view returns(uint256)
func (_Oracle *OracleCaller) Mana(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "mana", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(address ) view returns(uint256)
func (_Oracle *OracleSession) Mana(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Mana(&_Oracle.CallOpts, arg0)
}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(address ) view returns(uint256)
func (_Oracle *OracleCallerSession) Mana(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Mana(&_Oracle.CallOpts, arg0)
}

// Manas is a free data retrieval call binding the contract method 0x611e9ca7.
//
// Solidity: function manas(uint64 ) view returns(uint256)
func (_Oracle *OracleCaller) Manas(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "manas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Manas is a free data retrieval call binding the contract method 0x611e9ca7.
//
// Solidity: function manas(uint64 ) view returns(uint256)
func (_Oracle *OracleSession) Manas(arg0 uint64) (*big.Int, error) {
	return _Oracle.Contract.Manas(&_Oracle.CallOpts, arg0)
}

// Manas is a free data retrieval call binding the contract method 0x611e9ca7.
//
// Solidity: function manas(uint64 ) view returns(uint256)
func (_Oracle *OracleCallerSession) Manas(arg0 uint64) (*big.Int, error) {
	return _Oracle.Contract.Manas(&_Oracle.CallOpts, arg0)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Oracle *OracleCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Oracle *OracleSession) MinStake() (*big.Int, error) {
	return _Oracle.Contract.MinStake(&_Oracle.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Oracle *OracleCallerSession) MinStake() (*big.Int, error) {
	return _Oracle.Contract.MinStake(&_Oracle.CallOpts)
}

// NextPrice is a free data retrieval call binding the contract method 0x726b6bcc.
//
// Solidity: function nextPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCaller) NextPrice(opts *bind.CallOpts, round uint64, slot uint8) (uint64, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "nextPrice", round, slot)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextPrice is a free data retrieval call binding the contract method 0x726b6bcc.
//
// Solidity: function nextPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleSession) NextPrice(round uint64, slot uint8) (uint64, error) {
	return _Oracle.Contract.NextPrice(&_Oracle.CallOpts, round, slot)
}

// NextPrice is a free data retrieval call binding the contract method 0x726b6bcc.
//
// Solidity: function nextPrice(uint64 round, uint8 slot) view returns(uint64 price_)
func (_Oracle *OracleCallerSession) NextPrice(round uint64, slot uint8) (uint64, error) {
	return _Oracle.Contract.NextPrice(&_Oracle.CallOpts, round, slot)
}

// PriceToMana is a free data retrieval call binding the contract method 0x649a2678.
//
// Solidity: function priceToMana(uint64 , uint256 ) view returns(uint256)
func (_Oracle *OracleCaller) PriceToMana(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "priceToMana", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceToMana is a free data retrieval call binding the contract method 0x649a2678.
//
// Solidity: function priceToMana(uint64 , uint256 ) view returns(uint256)
func (_Oracle *OracleSession) PriceToMana(arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.PriceToMana(&_Oracle.CallOpts, arg0, arg1)
}

// PriceToMana is a free data retrieval call binding the contract method 0x649a2678.
//
// Solidity: function priceToMana(uint64 , uint256 ) view returns(uint256)
func (_Oracle *OracleCallerSession) PriceToMana(arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.PriceToMana(&_Oracle.CallOpts, arg0, arg1)
}

// RoundDuration is a free data retrieval call binding the contract method 0xf7cb789a.
//
// Solidity: function roundDuration() view returns(uint24)
func (_Oracle *OracleCaller) RoundDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "roundDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundDuration is a free data retrieval call binding the contract method 0xf7cb789a.
//
// Solidity: function roundDuration() view returns(uint24)
func (_Oracle *OracleSession) RoundDuration() (*big.Int, error) {
	return _Oracle.Contract.RoundDuration(&_Oracle.CallOpts)
}

// RoundDuration is a free data retrieval call binding the contract method 0xf7cb789a.
//
// Solidity: function roundDuration() view returns(uint24)
func (_Oracle *OracleCallerSession) RoundDuration() (*big.Int, error) {
	return _Oracle.Contract.RoundDuration(&_Oracle.CallOpts)
}

// SetSlots is a free data retrieval call binding the contract method 0x549d665e.
//
// Solidity: function setSlots(uint256[8] slots) pure returns(uint256 data)
func (_Oracle *OracleCaller) SetSlots(opts *bind.CallOpts, slots [8]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "setSlots", slots)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SetSlots is a free data retrieval call binding the contract method 0x549d665e.
//
// Solidity: function setSlots(uint256[8] slots) pure returns(uint256 data)
func (_Oracle *OracleSession) SetSlots(slots [8]*big.Int) (*big.Int, error) {
	return _Oracle.Contract.SetSlots(&_Oracle.CallOpts, slots)
}

// SetSlots is a free data retrieval call binding the contract method 0x549d665e.
//
// Solidity: function setSlots(uint256[8] slots) pure returns(uint256 data)
func (_Oracle *OracleCallerSession) SetSlots(slots [8]*big.Int) (*big.Int, error) {
	return _Oracle.Contract.SetSlots(&_Oracle.CallOpts, slots)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Oracle *OracleCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "stakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Oracle *OracleSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Stakes(&_Oracle.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Oracle *OracleCallerSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Stakes(&_Oracle.CallOpts, arg0)
}

// Submitters is a free data retrieval call binding the contract method 0x425d507b.
//
// Solidity: function submitters(uint64 , address ) view returns(uint256)
func (_Oracle *OracleCaller) Submitters(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "submitters", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Submitters is a free data retrieval call binding the contract method 0x425d507b.
//
// Solidity: function submitters(uint64 , address ) view returns(uint256)
func (_Oracle *OracleSession) Submitters(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Submitters(&_Oracle.CallOpts, arg0, arg1)
}

// Submitters is a free data retrieval call binding the contract method 0x425d507b.
//
// Solidity: function submitters(uint64 , address ) view returns(uint256)
func (_Oracle *OracleCallerSession) Submitters(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Submitters(&_Oracle.CallOpts, arg0, arg1)
}

// TotalMana is a free data retrieval call binding the contract method 0x1810d2d9.
//
// Solidity: function totalMana() view returns(uint256)
func (_Oracle *OracleCaller) TotalMana(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "totalMana")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMana is a free data retrieval call binding the contract method 0x1810d2d9.
//
// Solidity: function totalMana() view returns(uint256)
func (_Oracle *OracleSession) TotalMana() (*big.Int, error) {
	return _Oracle.Contract.TotalMana(&_Oracle.CallOpts)
}

// TotalMana is a free data retrieval call binding the contract method 0x1810d2d9.
//
// Solidity: function totalMana() view returns(uint256)
func (_Oracle *OracleCallerSession) TotalMana() (*big.Int, error) {
	return _Oracle.Contract.TotalMana(&_Oracle.CallOpts)
}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_Oracle *OracleCaller) TotalStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "totalStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_Oracle *OracleSession) TotalStakes() (*big.Int, error) {
	return _Oracle.Contract.TotalStakes(&_Oracle.CallOpts)
}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_Oracle *OracleCallerSession) TotalStakes() (*big.Int, error) {
	return _Oracle.Contract.TotalStakes(&_Oracle.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address recipient) returns(uint256 accumulatedBounty)
func (_Oracle *OracleTransactor) Claim(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "claim", recipient)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address recipient) returns(uint256 accumulatedBounty)
func (_Oracle *OracleSession) Claim(recipient common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Claim(&_Oracle.TransactOpts, recipient)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address recipient) returns(uint256 accumulatedBounty)
func (_Oracle *OracleTransactorSession) Claim(recipient common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Claim(&_Oracle.TransactOpts, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Oracle *OracleTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Oracle *OracleSession) Deposit() (*types.Transaction, error) {
	return _Oracle.Contract.Deposit(&_Oracle.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Oracle *OracleTransactorSession) Deposit() (*types.Transaction, error) {
	return _Oracle.Contract.Deposit(&_Oracle.TransactOpts)
}

// SetPrices is a paid mutator transaction binding the contract method 0xf97b32f6.
//
// Solidity: function setPrices(uint256 prices_, uint64 round) returns()
func (_Oracle *OracleTransactor) SetPrices(opts *bind.TransactOpts, prices_ *big.Int, round uint64) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setPrices", prices_, round)
}

// SetPrices is a paid mutator transaction binding the contract method 0xf97b32f6.
//
// Solidity: function setPrices(uint256 prices_, uint64 round) returns()
func (_Oracle *OracleSession) SetPrices(prices_ *big.Int, round uint64) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrices(&_Oracle.TransactOpts, prices_, round)
}

// SetPrices is a paid mutator transaction binding the contract method 0xf97b32f6.
//
// Solidity: function setPrices(uint256 prices_, uint64 round) returns()
func (_Oracle *OracleTransactorSession) SetPrices(prices_ *big.Int, round uint64) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrices(&_Oracle.TransactOpts, prices_, round)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address recipient) returns()
func (_Oracle *OracleTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "withdraw", amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address recipient) returns()
func (_Oracle *OracleSession) Withdraw(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Withdraw(&_Oracle.TransactOpts, amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address recipient) returns()
func (_Oracle *OracleTransactorSession) Withdraw(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Withdraw(&_Oracle.TransactOpts, amount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactorSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

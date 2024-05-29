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

// TraderMetaData contains all meta data concerning the Trader contract.
var TraderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"mintees\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnees\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"synth\",\"type\":\"address\"}],\"name\":\"claimAllPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"enterees\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"exitees\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"synth\",\"type\":\"address\"}],\"name\":\"claimAllSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TraderABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderMetaData.ABI instead.
var TraderABI = TraderMetaData.ABI

// Trader is an auto generated Go binding around an Ethereum contract.
type Trader struct {
	TraderCaller     // Read-only binding to the contract
	TraderTransactor // Write-only binding to the contract
	TraderFilterer   // Log filterer for contract events
}

// TraderCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderSession struct {
	Contract     *Trader           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderCallerSession struct {
	Contract *TraderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderTransactorSession struct {
	Contract     *TraderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderRaw struct {
	Contract *Trader // Generic contract binding to access the raw methods on
}

// TraderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderCallerRaw struct {
	Contract *TraderCaller // Generic read-only contract binding to access the raw methods on
}

// TraderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderTransactorRaw struct {
	Contract *TraderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrader creates a new instance of Trader, bound to a specific deployed contract.
func NewTrader(address common.Address, backend bind.ContractBackend) (*Trader, error) {
	contract, err := bindTrader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trader{TraderCaller: TraderCaller{contract: contract}, TraderTransactor: TraderTransactor{contract: contract}, TraderFilterer: TraderFilterer{contract: contract}}, nil
}

// NewTraderCaller creates a new read-only instance of Trader, bound to a specific deployed contract.
func NewTraderCaller(address common.Address, caller bind.ContractCaller) (*TraderCaller, error) {
	contract, err := bindTrader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderCaller{contract: contract}, nil
}

// NewTraderTransactor creates a new write-only instance of Trader, bound to a specific deployed contract.
func NewTraderTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderTransactor, error) {
	contract, err := bindTrader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderTransactor{contract: contract}, nil
}

// NewTraderFilterer creates a new log filterer instance of Trader, bound to a specific deployed contract.
func NewTraderFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderFilterer, error) {
	contract, err := bindTrader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderFilterer{contract: contract}, nil
}

// bindTrader binds a generic wrapper to an already deployed contract.
func bindTrader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trader *TraderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trader.Contract.TraderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trader *TraderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trader.Contract.TraderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trader *TraderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trader.Contract.TraderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trader *TraderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trader *TraderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trader *TraderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trader.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trader *TraderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trader.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trader *TraderSession) Owner() (common.Address, error) {
	return _Trader.Contract.Owner(&_Trader.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trader *TraderCallerSession) Owner() (common.Address, error) {
	return _Trader.Contract.Owner(&_Trader.CallOpts)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0xe89ce427.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, address synth) returns()
func (_Trader *TraderTransactor) ClaimAllPosition(opts *bind.TransactOpts, mintees [][32]byte, burnees [][32]byte, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.contract.Transact(opts, "claimAllPosition", mintees, burnees, round, synth)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0xe89ce427.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, address synth) returns()
func (_Trader *TraderSession) ClaimAllPosition(mintees [][32]byte, burnees [][32]byte, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.Contract.ClaimAllPosition(&_Trader.TransactOpts, mintees, burnees, round, synth)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0xe89ce427.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, address synth) returns()
func (_Trader *TraderTransactorSession) ClaimAllPosition(mintees [][32]byte, burnees [][32]byte, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.Contract.ClaimAllPosition(&_Trader.TransactOpts, mintees, burnees, round, synth)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x4bfa4d75.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, address synth) returns()
func (_Trader *TraderTransactor) ClaimAllSwap(opts *bind.TransactOpts, enterees []common.Address, exitees []common.Address, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.contract.Transact(opts, "claimAllSwap", enterees, exitees, round, synth)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x4bfa4d75.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, address synth) returns()
func (_Trader *TraderSession) ClaimAllSwap(enterees []common.Address, exitees []common.Address, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.Contract.ClaimAllSwap(&_Trader.TransactOpts, enterees, exitees, round, synth)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x4bfa4d75.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, address synth) returns()
func (_Trader *TraderTransactorSession) ClaimAllSwap(enterees []common.Address, exitees []common.Address, round uint64, synth common.Address) (*types.Transaction, error) {
	return _Trader.Contract.ClaimAllSwap(&_Trader.TransactOpts, enterees, exitees, round, synth)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address recipient) returns()
func (_Trader *TraderTransactor) Harvest(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Trader.contract.Transact(opts, "harvest", recipient)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address recipient) returns()
func (_Trader *TraderSession) Harvest(recipient common.Address) (*types.Transaction, error) {
	return _Trader.Contract.Harvest(&_Trader.TransactOpts, recipient)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address recipient) returns()
func (_Trader *TraderTransactorSession) Harvest(recipient common.Address) (*types.Transaction, error) {
	return _Trader.Contract.Harvest(&_Trader.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trader *TraderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trader.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trader *TraderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trader.Contract.RenounceOwnership(&_Trader.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trader *TraderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trader.Contract.RenounceOwnership(&_Trader.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trader *TraderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trader.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trader *TraderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trader.Contract.TransferOwnership(&_Trader.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trader *TraderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trader.Contract.TransferOwnership(&_Trader.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trader *TraderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trader.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trader *TraderSession) Receive() (*types.Transaction, error) {
	return _Trader.Contract.Receive(&_Trader.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trader *TraderTransactorSession) Receive() (*types.Transaction, error) {
	return _Trader.Contract.Receive(&_Trader.TransactOpts)
}

// TraderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trader contract.
type TraderOwnershipTransferredIterator struct {
	Event *TraderOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TraderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TraderOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TraderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderOwnershipTransferred represents a OwnershipTransferred event raised by the Trader contract.
type TraderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trader *TraderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TraderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trader.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TraderOwnershipTransferredIterator{contract: _Trader.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trader *TraderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TraderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trader.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderOwnershipTransferred)
				if err := _Trader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trader *TraderFilterer) ParseOwnershipTransferred(log types.Log) (*TraderOwnershipTransferred, error) {
	event := new(TraderOwnershipTransferred)
	if err := _Trader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

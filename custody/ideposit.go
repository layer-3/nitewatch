// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custody

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

// IDepositMetaData contains all meta data concerning the IDeposit contract.
var IDepositMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MsgValueMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonZeroMsgValueForERC20\",\"inputs\":[]}]",
}

// IDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use IDepositMetaData.ABI instead.
var IDepositABI = IDepositMetaData.ABI

// IDeposit is an auto generated Go binding around an Ethereum contract.
type IDeposit struct {
	IDepositCaller     // Read-only binding to the contract
	IDepositTransactor // Write-only binding to the contract
	IDepositFilterer   // Log filterer for contract events
}

// IDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositSession struct {
	Contract     *IDeposit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositCallerSession struct {
	Contract *IDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositTransactorSession struct {
	Contract     *IDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositRaw struct {
	Contract *IDeposit // Generic contract binding to access the raw methods on
}

// IDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositCallerRaw struct {
	Contract *IDepositCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositTransactorRaw struct {
	Contract *IDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDeposit creates a new instance of IDeposit, bound to a specific deployed contract.
func NewIDeposit(address common.Address, backend bind.ContractBackend) (*IDeposit, error) {
	contract, err := bindIDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDeposit{IDepositCaller: IDepositCaller{contract: contract}, IDepositTransactor: IDepositTransactor{contract: contract}, IDepositFilterer: IDepositFilterer{contract: contract}}, nil
}

// NewIDepositCaller creates a new read-only instance of IDeposit, bound to a specific deployed contract.
func NewIDepositCaller(address common.Address, caller bind.ContractCaller) (*IDepositCaller, error) {
	contract, err := bindIDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositCaller{contract: contract}, nil
}

// NewIDepositTransactor creates a new write-only instance of IDeposit, bound to a specific deployed contract.
func NewIDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositTransactor, error) {
	contract, err := bindIDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositTransactor{contract: contract}, nil
}

// NewIDepositFilterer creates a new log filterer instance of IDeposit, bound to a specific deployed contract.
func NewIDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositFilterer, error) {
	contract, err := bindIDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositFilterer{contract: contract}, nil
}

// bindIDeposit binds a generic wrapper to an already deployed contract.
func bindIDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDeposit *IDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDeposit.Contract.IDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDeposit *IDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDeposit.Contract.IDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDeposit *IDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDeposit.Contract.IDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDeposit *IDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDeposit *IDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDeposit *IDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDeposit.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_IDeposit *IDepositTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDeposit.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_IDeposit *IDepositSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDeposit.Contract.Deposit(&_IDeposit.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_IDeposit *IDepositTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDeposit.Contract.Deposit(&_IDeposit.TransactOpts, token, amount)
}

// IDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IDeposit contract.
type IDepositDepositedIterator struct {
	Event *IDepositDeposited // Event containing the contract specifics and raw log

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
func (it *IDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDepositDeposited)
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
		it.Event = new(IDepositDeposited)
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
func (it *IDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDepositDeposited represents a Deposited event raised by the IDeposit contract.
type IDepositDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_IDeposit *IDepositFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*IDepositDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IDeposit.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IDepositDepositedIterator{contract: _IDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_IDeposit *IDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IDepositDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IDeposit.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDepositDeposited)
				if err := _IDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_IDeposit *IDepositFilterer) ParseDeposited(log types.Log) (*IDepositDeposited, error) {
	event := new(IDepositDeposited)
	if err := _IDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

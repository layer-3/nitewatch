// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

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

// ICustodyMetaData contains all meta data concerning the ICustody contract.
var ICustodyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFinalized\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawStarted\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ICustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use ICustodyMetaData.ABI instead.
var ICustodyABI = ICustodyMetaData.ABI

// ICustody is an auto generated Go binding around an Ethereum contract.
type ICustody struct {
	ICustodyCaller     // Read-only binding to the contract
	ICustodyTransactor // Write-only binding to the contract
	ICustodyFilterer   // Log filterer for contract events
}

// ICustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICustodySession struct {
	Contract     *ICustody         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICustodyCallerSession struct {
	Contract *ICustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ICustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICustodyTransactorSession struct {
	Contract     *ICustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICustodyRaw struct {
	Contract *ICustody // Generic contract binding to access the raw methods on
}

// ICustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICustodyCallerRaw struct {
	Contract *ICustodyCaller // Generic read-only contract binding to access the raw methods on
}

// ICustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICustodyTransactorRaw struct {
	Contract *ICustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICustody creates a new instance of ICustody, bound to a specific deployed contract.
func NewICustody(address common.Address, backend bind.ContractBackend) (*ICustody, error) {
	contract, err := bindICustody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICustody{ICustodyCaller: ICustodyCaller{contract: contract}, ICustodyTransactor: ICustodyTransactor{contract: contract}, ICustodyFilterer: ICustodyFilterer{contract: contract}}, nil
}

// NewICustodyCaller creates a new read-only instance of ICustody, bound to a specific deployed contract.
func NewICustodyCaller(address common.Address, caller bind.ContractCaller) (*ICustodyCaller, error) {
	contract, err := bindICustody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICustodyCaller{contract: contract}, nil
}

// NewICustodyTransactor creates a new write-only instance of ICustody, bound to a specific deployed contract.
func NewICustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*ICustodyTransactor, error) {
	contract, err := bindICustody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICustodyTransactor{contract: contract}, nil
}

// NewICustodyFilterer creates a new log filterer instance of ICustody, bound to a specific deployed contract.
func NewICustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*ICustodyFilterer, error) {
	contract, err := bindICustody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICustodyFilterer{contract: contract}, nil
}

// bindICustody binds a generic wrapper to an already deployed contract.
func bindICustody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICustody *ICustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICustody.Contract.ICustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICustody *ICustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICustody.Contract.ICustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICustody *ICustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICustody.Contract.ICustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICustody *ICustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICustody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICustody *ICustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICustody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICustody *ICustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICustody.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_ICustody *ICustodyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICustody.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_ICustody *ICustodySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICustody.Contract.Deposit(&_ICustody.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_ICustody *ICustodyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ICustody.Contract.Deposit(&_ICustody.TransactOpts, token, amount)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodyTransactor) FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.contract.Transact(opts, "finalizeWithdraw", withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodySession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.Contract.FinalizeWithdraw(&_ICustody.TransactOpts, withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodyTransactorSession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.Contract.FinalizeWithdraw(&_ICustody.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodyTransactor) RejectWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.contract.Transact(opts, "rejectWithdraw", withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodySession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.Contract.RejectWithdraw(&_ICustody.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_ICustody *ICustodyTransactorSession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _ICustody.Contract.RejectWithdraw(&_ICustody.TransactOpts, withdrawalId)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_ICustody *ICustodyTransactor) StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ICustody.contract.Transact(opts, "startWithdraw", user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_ICustody *ICustodySession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ICustody.Contract.StartWithdraw(&_ICustody.TransactOpts, user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_ICustody *ICustodyTransactorSession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _ICustody.Contract.StartWithdraw(&_ICustody.TransactOpts, user, token, amount, nonce)
}

// ICustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ICustody contract.
type ICustodyDepositedIterator struct {
	Event *ICustodyDeposited // Event containing the contract specifics and raw log

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
func (it *ICustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICustodyDeposited)
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
		it.Event = new(ICustodyDeposited)
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
func (it *ICustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICustodyDeposited represents a Deposited event raised by the ICustody contract.
type ICustodyDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_ICustody *ICustodyFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*ICustodyDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICustody.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ICustodyDepositedIterator{contract: _ICustody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_ICustody *ICustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ICustodyDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICustody.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICustodyDeposited)
				if err := _ICustody.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_ICustody *ICustodyFilterer) ParseDeposited(log types.Log) (*ICustodyDeposited, error) {
	event := new(ICustodyDeposited)
	if err := _ICustody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICustodyWithdrawFinalizedIterator is returned from FilterWithdrawFinalized and is used to iterate over the raw logs and unpacked data for WithdrawFinalized events raised by the ICustody contract.
type ICustodyWithdrawFinalizedIterator struct {
	Event *ICustodyWithdrawFinalized // Event containing the contract specifics and raw log

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
func (it *ICustodyWithdrawFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICustodyWithdrawFinalized)
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
		it.Event = new(ICustodyWithdrawFinalized)
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
func (it *ICustodyWithdrawFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICustodyWithdrawFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICustodyWithdrawFinalized represents a WithdrawFinalized event raised by the ICustody contract.
type ICustodyWithdrawFinalized struct {
	WithdrawalId [32]byte
	Success      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFinalized is a free log retrieval operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_ICustody *ICustodyFilterer) FilterWithdrawFinalized(opts *bind.FilterOpts, withdrawalId [][32]byte) (*ICustodyWithdrawFinalizedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _ICustody.contract.FilterLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return &ICustodyWithdrawFinalizedIterator{contract: _ICustody.contract, event: "WithdrawFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawFinalized is a free log subscription operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_ICustody *ICustodyFilterer) WatchWithdrawFinalized(opts *bind.WatchOpts, sink chan<- *ICustodyWithdrawFinalized, withdrawalId [][32]byte) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _ICustody.contract.WatchLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICustodyWithdrawFinalized)
				if err := _ICustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
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

// ParseWithdrawFinalized is a log parse operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_ICustody *ICustodyFilterer) ParseWithdrawFinalized(log types.Log) (*ICustodyWithdrawFinalized, error) {
	event := new(ICustodyWithdrawFinalized)
	if err := _ICustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICustodyWithdrawStartedIterator is returned from FilterWithdrawStarted and is used to iterate over the raw logs and unpacked data for WithdrawStarted events raised by the ICustody contract.
type ICustodyWithdrawStartedIterator struct {
	Event *ICustodyWithdrawStarted // Event containing the contract specifics and raw log

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
func (it *ICustodyWithdrawStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICustodyWithdrawStarted)
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
		it.Event = new(ICustodyWithdrawStarted)
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
func (it *ICustodyWithdrawStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICustodyWithdrawStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICustodyWithdrawStarted represents a WithdrawStarted event raised by the ICustody contract.
type ICustodyWithdrawStarted struct {
	WithdrawalId [32]byte
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	Nonce        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawStarted is a free log retrieval operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_ICustody *ICustodyFilterer) FilterWithdrawStarted(opts *bind.FilterOpts, withdrawalId [][32]byte, user []common.Address, token []common.Address) (*ICustodyWithdrawStartedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICustody.contract.FilterLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ICustodyWithdrawStartedIterator{contract: _ICustody.contract, event: "WithdrawStarted", logs: logs, sub: sub}, nil
}

// WatchWithdrawStarted is a free log subscription operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_ICustody *ICustodyFilterer) WatchWithdrawStarted(opts *bind.WatchOpts, sink chan<- *ICustodyWithdrawStarted, withdrawalId [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ICustody.contract.WatchLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICustodyWithdrawStarted)
				if err := _ICustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
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

// ParseWithdrawStarted is a log parse operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_ICustody *ICustodyFilterer) ParseWithdrawStarted(log types.Log) (*ICustodyWithdrawStarted, error) {
	event := new(ICustodyWithdrawStarted)
	if err := _ICustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// IWithdrawMetaData contains all meta data concerning the IWithdraw contract.
var IWithdrawMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"finalizeWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"WithdrawFinalized\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawStarted\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ETHTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
}

// IWithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use IWithdrawMetaData.ABI instead.
var IWithdrawABI = IWithdrawMetaData.ABI

// IWithdraw is an auto generated Go binding around an Ethereum contract.
type IWithdraw struct {
	IWithdrawCaller     // Read-only binding to the contract
	IWithdrawTransactor // Write-only binding to the contract
	IWithdrawFilterer   // Log filterer for contract events
}

// IWithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWithdrawSession struct {
	Contract     *IWithdraw        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWithdrawCallerSession struct {
	Contract *IWithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IWithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWithdrawTransactorSession struct {
	Contract     *IWithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IWithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWithdrawRaw struct {
	Contract *IWithdraw // Generic contract binding to access the raw methods on
}

// IWithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWithdrawCallerRaw struct {
	Contract *IWithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// IWithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWithdrawTransactorRaw struct {
	Contract *IWithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWithdraw creates a new instance of IWithdraw, bound to a specific deployed contract.
func NewIWithdraw(address common.Address, backend bind.ContractBackend) (*IWithdraw, error) {
	contract, err := bindIWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWithdraw{IWithdrawCaller: IWithdrawCaller{contract: contract}, IWithdrawTransactor: IWithdrawTransactor{contract: contract}, IWithdrawFilterer: IWithdrawFilterer{contract: contract}}, nil
}

// NewIWithdrawCaller creates a new read-only instance of IWithdraw, bound to a specific deployed contract.
func NewIWithdrawCaller(address common.Address, caller bind.ContractCaller) (*IWithdrawCaller, error) {
	contract, err := bindIWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWithdrawCaller{contract: contract}, nil
}

// NewIWithdrawTransactor creates a new write-only instance of IWithdraw, bound to a specific deployed contract.
func NewIWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*IWithdrawTransactor, error) {
	contract, err := bindIWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWithdrawTransactor{contract: contract}, nil
}

// NewIWithdrawFilterer creates a new log filterer instance of IWithdraw, bound to a specific deployed contract.
func NewIWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*IWithdrawFilterer, error) {
	contract, err := bindIWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWithdrawFilterer{contract: contract}, nil
}

// bindIWithdraw binds a generic wrapper to an already deployed contract.
func bindIWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWithdrawMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWithdraw *IWithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWithdraw.Contract.IWithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWithdraw *IWithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWithdraw.Contract.IWithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWithdraw *IWithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWithdraw.Contract.IWithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWithdraw *IWithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWithdraw *IWithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWithdraw *IWithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWithdraw.Contract.contract.Transact(opts, method, params...)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawTransactor) FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.contract.Transact(opts, "finalizeWithdraw", withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawSession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.Contract.FinalizeWithdraw(&_IWithdraw.TransactOpts, withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawTransactorSession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.Contract.FinalizeWithdraw(&_IWithdraw.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawTransactor) RejectWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.contract.Transact(opts, "rejectWithdraw", withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawSession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.Contract.RejectWithdraw(&_IWithdraw.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_IWithdraw *IWithdrawTransactorSession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _IWithdraw.Contract.RejectWithdraw(&_IWithdraw.TransactOpts, withdrawalId)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_IWithdraw *IWithdrawTransactor) StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IWithdraw.contract.Transact(opts, "startWithdraw", user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_IWithdraw *IWithdrawSession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IWithdraw.Contract.StartWithdraw(&_IWithdraw.TransactOpts, user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_IWithdraw *IWithdrawTransactorSession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IWithdraw.Contract.StartWithdraw(&_IWithdraw.TransactOpts, user, token, amount, nonce)
}

// IWithdrawWithdrawFinalizedIterator is returned from FilterWithdrawFinalized and is used to iterate over the raw logs and unpacked data for WithdrawFinalized events raised by the IWithdraw contract.
type IWithdrawWithdrawFinalizedIterator struct {
	Event *IWithdrawWithdrawFinalized // Event containing the contract specifics and raw log

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
func (it *IWithdrawWithdrawFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWithdrawWithdrawFinalized)
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
		it.Event = new(IWithdrawWithdrawFinalized)
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
func (it *IWithdrawWithdrawFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWithdrawWithdrawFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWithdrawWithdrawFinalized represents a WithdrawFinalized event raised by the IWithdraw contract.
type IWithdrawWithdrawFinalized struct {
	WithdrawalId [32]byte
	Success      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFinalized is a free log retrieval operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_IWithdraw *IWithdrawFilterer) FilterWithdrawFinalized(opts *bind.FilterOpts, withdrawalId [][32]byte) (*IWithdrawWithdrawFinalizedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _IWithdraw.contract.FilterLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return &IWithdrawWithdrawFinalizedIterator{contract: _IWithdraw.contract, event: "WithdrawFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawFinalized is a free log subscription operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_IWithdraw *IWithdrawFilterer) WatchWithdrawFinalized(opts *bind.WatchOpts, sink chan<- *IWithdrawWithdrawFinalized, withdrawalId [][32]byte) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _IWithdraw.contract.WatchLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWithdrawWithdrawFinalized)
				if err := _IWithdraw.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
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
func (_IWithdraw *IWithdrawFilterer) ParseWithdrawFinalized(log types.Log) (*IWithdrawWithdrawFinalized, error) {
	event := new(IWithdrawWithdrawFinalized)
	if err := _IWithdraw.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWithdrawWithdrawStartedIterator is returned from FilterWithdrawStarted and is used to iterate over the raw logs and unpacked data for WithdrawStarted events raised by the IWithdraw contract.
type IWithdrawWithdrawStartedIterator struct {
	Event *IWithdrawWithdrawStarted // Event containing the contract specifics and raw log

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
func (it *IWithdrawWithdrawStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWithdrawWithdrawStarted)
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
		it.Event = new(IWithdrawWithdrawStarted)
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
func (it *IWithdrawWithdrawStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWithdrawWithdrawStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWithdrawWithdrawStarted represents a WithdrawStarted event raised by the IWithdraw contract.
type IWithdrawWithdrawStarted struct {
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
func (_IWithdraw *IWithdrawFilterer) FilterWithdrawStarted(opts *bind.FilterOpts, withdrawalId [][32]byte, user []common.Address, token []common.Address) (*IWithdrawWithdrawStartedIterator, error) {

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

	logs, sub, err := _IWithdraw.contract.FilterLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IWithdrawWithdrawStartedIterator{contract: _IWithdraw.contract, event: "WithdrawStarted", logs: logs, sub: sub}, nil
}

// WatchWithdrawStarted is a free log subscription operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_IWithdraw *IWithdrawFilterer) WatchWithdrawStarted(opts *bind.WatchOpts, sink chan<- *IWithdrawWithdrawStarted, withdrawalId [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IWithdraw.contract.WatchLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWithdrawWithdrawStarted)
				if err := _IWithdraw.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
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
func (_IWithdraw *IWithdrawFilterer) ParseWithdrawStarted(log types.Log) (*IWithdrawWithdrawStarted, error) {
	event := new(IWithdrawWithdrawStarted)
	if err := _IWithdraw.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

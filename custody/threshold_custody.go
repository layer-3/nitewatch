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
)

// ThresholdCustodyMetaData contains all meta data concerning the ThresholdCustody contract.
var ThresholdCustodyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ThresholdCustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use ThresholdCustodyMetaData.ABI instead.
var ThresholdCustodyABI = ThresholdCustodyMetaData.ABI

// ThresholdCustody is an auto generated Go binding around an Ethereum contract.
type ThresholdCustody struct {
	ThresholdCustodyCaller     // Read-only binding to the contract
	ThresholdCustodyTransactor // Write-only binding to the contract
	ThresholdCustodyFilterer   // Log filterer for contract events
}

// ThresholdCustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThresholdCustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdCustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThresholdCustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdCustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThresholdCustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdCustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThresholdCustodySession struct {
	Contract     *ThresholdCustody // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThresholdCustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThresholdCustodyCallerSession struct {
	Contract *ThresholdCustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ThresholdCustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThresholdCustodyTransactorSession struct {
	Contract     *ThresholdCustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ThresholdCustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThresholdCustodyRaw struct {
	Contract *ThresholdCustody // Generic contract binding to access the raw methods on
}

// ThresholdCustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThresholdCustodyCallerRaw struct {
	Contract *ThresholdCustodyCaller // Generic read-only contract binding to access the raw methods on
}

// ThresholdCustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThresholdCustodyTransactorRaw struct {
	Contract *ThresholdCustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThresholdCustody creates a new instance of ThresholdCustody, bound to a specific deployed contract.
func NewThresholdCustody(address common.Address, backend bind.ContractBackend) (*ThresholdCustody, error) {
	contract, err := bindThresholdCustody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThresholdCustody{ThresholdCustodyCaller: ThresholdCustodyCaller{contract: contract}, ThresholdCustodyTransactor: ThresholdCustodyTransactor{contract: contract}, ThresholdCustodyFilterer: ThresholdCustodyFilterer{contract: contract}}, nil
}

// NewThresholdCustodyCaller creates a new read-only instance of ThresholdCustody, bound to a specific deployed contract.
func NewThresholdCustodyCaller(address common.Address, caller bind.ContractCaller) (*ThresholdCustodyCaller, error) {
	contract, err := bindThresholdCustody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThresholdCustodyCaller{contract: contract}, nil
}

// NewThresholdCustodyTransactor creates a new write-only instance of ThresholdCustody, bound to a specific deployed contract.
func NewThresholdCustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*ThresholdCustodyTransactor, error) {
	contract, err := bindThresholdCustody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThresholdCustodyTransactor{contract: contract}, nil
}

// NewThresholdCustodyFilterer creates a new log filterer instance of ThresholdCustody, bound to a specific deployed contract.
func NewThresholdCustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*ThresholdCustodyFilterer, error) {
	contract, err := bindThresholdCustody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThresholdCustodyFilterer{contract: contract}, nil
}

// bindThresholdCustody binds a generic wrapper to an already deployed contract.
func bindThresholdCustody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThresholdCustodyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThresholdCustody *ThresholdCustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThresholdCustody.Contract.ThresholdCustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThresholdCustody *ThresholdCustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThresholdCustody.Contract.ThresholdCustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThresholdCustody *ThresholdCustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThresholdCustody.Contract.ThresholdCustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThresholdCustody *ThresholdCustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThresholdCustody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThresholdCustody *ThresholdCustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThresholdCustody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThresholdCustody *ThresholdCustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThresholdCustody.Contract.contract.Transact(opts, method, params...)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_ThresholdCustody *ThresholdCustodyCaller) IsSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _ThresholdCustody.contract.Call(opts, &out, "isSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_ThresholdCustody *ThresholdCustodySession) IsSigner(signer common.Address) (bool, error) {
	return _ThresholdCustody.Contract.IsSigner(&_ThresholdCustody.CallOpts, signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool)
func (_ThresholdCustody *ThresholdCustodyCallerSession) IsSigner(signer common.Address) (bool, error) {
	return _ThresholdCustody.Contract.IsSigner(&_ThresholdCustody.CallOpts, signer)
}

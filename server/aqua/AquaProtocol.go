// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aqua

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

// AquaABI is the input ABI used to generate the binding from.
const AquaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"ethAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"callbackWithdrawFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"callbackDepositFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"_typeOf\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"buyInsurance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"buyArt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"ethAddress\",\"type\":\"address\"},{\"name\":\"bankAccountAddress\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"withdrawFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"bankAccountAddress\",\"type\":\"string\"},{\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"depositFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"issuerId\",\"type\":\"string\"},{\"name\":\"series\",\"type\":\"string\"},{\"name\":\"numberFrom\",\"type\":\"uint256\"},{\"name\":\"numberTo\",\"type\":\"uint256\"}],\"name\":\"buyPaper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracleInstanceAddress\",\"type\":\"address\"}],\"name\":\"setOracleInstanceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_result\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"CallbackDepositFiatEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bankAccountAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFiatEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"NewOracleAddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"issuerId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"series\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"numberFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BuyPaperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BuyArtEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Aqua is an auto generated Go binding around an Ethereum contract.
type Aqua struct {
	AquaCaller     // Read-only binding to the contract
	AquaTransactor // Write-only binding to the contract
	AquaFilterer   // Log filterer for contract events
}

// AquaCaller is an auto generated read-only Go binding around an Ethereum contract.
type AquaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AquaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AquaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AquaSession struct {
	Contract     *Aqua             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AquaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AquaCallerSession struct {
	Contract *AquaCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AquaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AquaTransactorSession struct {
	Contract     *AquaTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AquaRaw is an auto generated low-level Go binding around an Ethereum contract.
type AquaRaw struct {
	Contract *Aqua // Generic contract binding to access the raw methods on
}

// AquaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AquaCallerRaw struct {
	Contract *AquaCaller // Generic read-only contract binding to access the raw methods on
}

// AquaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AquaTransactorRaw struct {
	Contract *AquaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAqua creates a new instance of Aqua, bound to a specific deployed contract.
func NewAqua(address common.Address, backend bind.ContractBackend) (*Aqua, error) {
	contract, err := bindAqua(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aqua{AquaCaller: AquaCaller{contract: contract}, AquaTransactor: AquaTransactor{contract: contract}, AquaFilterer: AquaFilterer{contract: contract}}, nil
}

// NewAquaCaller creates a new read-only instance of Aqua, bound to a specific deployed contract.
func NewAquaCaller(address common.Address, caller bind.ContractCaller) (*AquaCaller, error) {
	contract, err := bindAqua(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AquaCaller{contract: contract}, nil
}

// NewAquaTransactor creates a new write-only instance of Aqua, bound to a specific deployed contract.
func NewAquaTransactor(address common.Address, transactor bind.ContractTransactor) (*AquaTransactor, error) {
	contract, err := bindAqua(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AquaTransactor{contract: contract}, nil
}

// NewAquaFilterer creates a new log filterer instance of Aqua, bound to a specific deployed contract.
func NewAquaFilterer(address common.Address, filterer bind.ContractFilterer) (*AquaFilterer, error) {
	contract, err := bindAqua(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AquaFilterer{contract: contract}, nil
}

// bindAqua binds a generic wrapper to an already deployed contract.
func bindAqua(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AquaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aqua *AquaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aqua.Contract.AquaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aqua *AquaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aqua.Contract.AquaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aqua *AquaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aqua.Contract.AquaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aqua *AquaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aqua.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aqua *AquaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aqua.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aqua *AquaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aqua.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Aqua *AquaCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Aqua.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Aqua *AquaSession) IsOwner() (bool, error) {
	return _Aqua.Contract.IsOwner(&_Aqua.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Aqua *AquaCallerSession) IsOwner() (bool, error) {
	return _Aqua.Contract.IsOwner(&_Aqua.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aqua *AquaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Aqua.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aqua *AquaSession) Owner() (common.Address, error) {
	return _Aqua.Contract.Owner(&_Aqua.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Aqua *AquaCallerSession) Owner() (common.Address, error) {
	return _Aqua.Contract.Owner(&_Aqua.CallOpts)
}

// BuyArt is a paid mutator transaction binding the contract method 0x8e116998.
//
// Solidity: function buyArt(asset address, tokenAddress address, operator address, amount uint256, name string) returns()
func (_Aqua *AquaTransactor) BuyArt(opts *bind.TransactOpts, asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "buyArt", asset, tokenAddress, operator, amount, name)
}

// BuyArt is a paid mutator transaction binding the contract method 0x8e116998.
//
// Solidity: function buyArt(asset address, tokenAddress address, operator address, amount uint256, name string) returns()
func (_Aqua *AquaSession) BuyArt(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string) (*types.Transaction, error) {
	return _Aqua.Contract.BuyArt(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, name)
}

// BuyArt is a paid mutator transaction binding the contract method 0x8e116998.
//
// Solidity: function buyArt(asset address, tokenAddress address, operator address, amount uint256, name string) returns()
func (_Aqua *AquaTransactorSession) BuyArt(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string) (*types.Transaction, error) {
	return _Aqua.Contract.BuyArt(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, name)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x7a8e4cb2.
//
// Solidity: function buyInsurance(asset address, tokenAddress address, operator address, amount uint256, name string, _typeOf string, _date string) returns()
func (_Aqua *AquaTransactor) BuyInsurance(opts *bind.TransactOpts, asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string, _typeOf string, _date string) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "buyInsurance", asset, tokenAddress, operator, amount, name, _typeOf, _date)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x7a8e4cb2.
//
// Solidity: function buyInsurance(asset address, tokenAddress address, operator address, amount uint256, name string, _typeOf string, _date string) returns()
func (_Aqua *AquaSession) BuyInsurance(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string, _typeOf string, _date string) (*types.Transaction, error) {
	return _Aqua.Contract.BuyInsurance(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, name, _typeOf, _date)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x7a8e4cb2.
//
// Solidity: function buyInsurance(asset address, tokenAddress address, operator address, amount uint256, name string, _typeOf string, _date string) returns()
func (_Aqua *AquaTransactorSession) BuyInsurance(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string, _typeOf string, _date string) (*types.Transaction, error) {
	return _Aqua.Contract.BuyInsurance(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, name, _typeOf, _date)
}

// BuyPaper is a paid mutator transaction binding the contract method 0xcefc1a05.
//
// Solidity: function buyPaper(asset address, tokenAddress address, operator address, amount uint256, issuerId string, series string, numberFrom uint256, numberTo uint256) returns()
func (_Aqua *AquaTransactor) BuyPaper(opts *bind.TransactOpts, asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, issuerId string, series string, numberFrom *big.Int, numberTo *big.Int) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "buyPaper", asset, tokenAddress, operator, amount, issuerId, series, numberFrom, numberTo)
}

// BuyPaper is a paid mutator transaction binding the contract method 0xcefc1a05.
//
// Solidity: function buyPaper(asset address, tokenAddress address, operator address, amount uint256, issuerId string, series string, numberFrom uint256, numberTo uint256) returns()
func (_Aqua *AquaSession) BuyPaper(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, issuerId string, series string, numberFrom *big.Int, numberTo *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.BuyPaper(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, issuerId, series, numberFrom, numberTo)
}

// BuyPaper is a paid mutator transaction binding the contract method 0xcefc1a05.
//
// Solidity: function buyPaper(asset address, tokenAddress address, operator address, amount uint256, issuerId string, series string, numberFrom uint256, numberTo uint256) returns()
func (_Aqua *AquaTransactorSession) BuyPaper(asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, issuerId string, series string, numberFrom *big.Int, numberTo *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.BuyPaper(&_Aqua.TransactOpts, asset, tokenAddress, operator, amount, issuerId, series, numberFrom, numberTo)
}

// CallbackDepositFiat is a paid mutator transaction binding the contract method 0x4fb5dd98.
//
// Solidity: function callbackDepositFiat(_result string, _id uint256) returns()
func (_Aqua *AquaTransactor) CallbackDepositFiat(opts *bind.TransactOpts, _result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "callbackDepositFiat", _result, _id)
}

// CallbackDepositFiat is a paid mutator transaction binding the contract method 0x4fb5dd98.
//
// Solidity: function callbackDepositFiat(_result string, _id uint256) returns()
func (_Aqua *AquaSession) CallbackDepositFiat(_result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.CallbackDepositFiat(&_Aqua.TransactOpts, _result, _id)
}

// CallbackDepositFiat is a paid mutator transaction binding the contract method 0x4fb5dd98.
//
// Solidity: function callbackDepositFiat(_result string, _id uint256) returns()
func (_Aqua *AquaTransactorSession) CallbackDepositFiat(_result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.CallbackDepositFiat(&_Aqua.TransactOpts, _result, _id)
}

// CallbackWithdrawFiat is a paid mutator transaction binding the contract method 0x4d3959ac.
//
// Solidity: function callbackWithdrawFiat(_result string, _id uint256) returns()
func (_Aqua *AquaTransactor) CallbackWithdrawFiat(opts *bind.TransactOpts, _result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "callbackWithdrawFiat", _result, _id)
}

// CallbackWithdrawFiat is a paid mutator transaction binding the contract method 0x4d3959ac.
//
// Solidity: function callbackWithdrawFiat(_result string, _id uint256) returns()
func (_Aqua *AquaSession) CallbackWithdrawFiat(_result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.CallbackWithdrawFiat(&_Aqua.TransactOpts, _result, _id)
}

// CallbackWithdrawFiat is a paid mutator transaction binding the contract method 0x4d3959ac.
//
// Solidity: function callbackWithdrawFiat(_result string, _id uint256) returns()
func (_Aqua *AquaTransactorSession) CallbackWithdrawFiat(_result string, _id *big.Int) (*types.Transaction, error) {
	return _Aqua.Contract.CallbackWithdrawFiat(&_Aqua.TransactOpts, _result, _id)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(asset address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "deposit", asset, amount, onBehalfOf)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(asset address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.Deposit(&_Aqua.TransactOpts, asset, amount, onBehalfOf)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(asset address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.Deposit(&_Aqua.TransactOpts, asset, amount, onBehalfOf)
}

// DepositFiat is a paid mutator transaction binding the contract method 0xc2f6ca27.
//
// Solidity: function depositFiat(asset address, amount uint256, bankAccountAddress string, ethAddress address) returns()
func (_Aqua *AquaTransactor) DepositFiat(opts *bind.TransactOpts, asset common.Address, amount *big.Int, bankAccountAddress string, ethAddress common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "depositFiat", asset, amount, bankAccountAddress, ethAddress)
}

// DepositFiat is a paid mutator transaction binding the contract method 0xc2f6ca27.
//
// Solidity: function depositFiat(asset address, amount uint256, bankAccountAddress string, ethAddress address) returns()
func (_Aqua *AquaSession) DepositFiat(asset common.Address, amount *big.Int, bankAccountAddress string, ethAddress common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.DepositFiat(&_Aqua.TransactOpts, asset, amount, bankAccountAddress, ethAddress)
}

// DepositFiat is a paid mutator transaction binding the contract method 0xc2f6ca27.
//
// Solidity: function depositFiat(asset address, amount uint256, bankAccountAddress string, ethAddress address) returns()
func (_Aqua *AquaTransactorSession) DepositFiat(asset common.Address, amount *big.Int, bankAccountAddress string, ethAddress common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.DepositFiat(&_Aqua.TransactOpts, asset, amount, bankAccountAddress, ethAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aqua *AquaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aqua *AquaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aqua.Contract.RenounceOwnership(&_Aqua.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aqua *AquaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aqua.Contract.RenounceOwnership(&_Aqua.TransactOpts)
}

// SetOracleInstanceAddress is a paid mutator transaction binding the contract method 0xe9e17a9e.
//
// Solidity: function setOracleInstanceAddress(_oracleInstanceAddress address) returns()
func (_Aqua *AquaTransactor) SetOracleInstanceAddress(opts *bind.TransactOpts, _oracleInstanceAddress common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "setOracleInstanceAddress", _oracleInstanceAddress)
}

// SetOracleInstanceAddress is a paid mutator transaction binding the contract method 0xe9e17a9e.
//
// Solidity: function setOracleInstanceAddress(_oracleInstanceAddress address) returns()
func (_Aqua *AquaSession) SetOracleInstanceAddress(_oracleInstanceAddress common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.SetOracleInstanceAddress(&_Aqua.TransactOpts, _oracleInstanceAddress)
}

// SetOracleInstanceAddress is a paid mutator transaction binding the contract method 0xe9e17a9e.
//
// Solidity: function setOracleInstanceAddress(_oracleInstanceAddress address) returns()
func (_Aqua *AquaTransactorSession) SetOracleInstanceAddress(_oracleInstanceAddress common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.SetOracleInstanceAddress(&_Aqua.TransactOpts, _oracleInstanceAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Aqua *AquaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Aqua *AquaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.TransferOwnership(&_Aqua.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Aqua *AquaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.TransferOwnership(&_Aqua.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0e917f76.
//
// Solidity: function withdraw(asset address, ethAddress address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, ethAddress common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "withdraw", asset, ethAddress, amount, onBehalfOf)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0e917f76.
//
// Solidity: function withdraw(asset address, ethAddress address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaSession) Withdraw(asset common.Address, ethAddress common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.Withdraw(&_Aqua.TransactOpts, asset, ethAddress, amount, onBehalfOf)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0e917f76.
//
// Solidity: function withdraw(asset address, ethAddress address, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactorSession) Withdraw(asset common.Address, ethAddress common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.Withdraw(&_Aqua.TransactOpts, asset, ethAddress, amount, onBehalfOf)
}

// WithdrawFiat is a paid mutator transaction binding the contract method 0xb384947e.
//
// Solidity: function withdrawFiat(asset address, ethAddress address, bankAccountAddress string, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactor) WithdrawFiat(opts *bind.TransactOpts, asset common.Address, ethAddress common.Address, bankAccountAddress string, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.contract.Transact(opts, "withdrawFiat", asset, ethAddress, bankAccountAddress, amount, onBehalfOf)
}

// WithdrawFiat is a paid mutator transaction binding the contract method 0xb384947e.
//
// Solidity: function withdrawFiat(asset address, ethAddress address, bankAccountAddress string, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaSession) WithdrawFiat(asset common.Address, ethAddress common.Address, bankAccountAddress string, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.WithdrawFiat(&_Aqua.TransactOpts, asset, ethAddress, bankAccountAddress, amount, onBehalfOf)
}

// WithdrawFiat is a paid mutator transaction binding the contract method 0xb384947e.
//
// Solidity: function withdrawFiat(asset address, ethAddress address, bankAccountAddress string, amount uint256, onBehalfOf address) returns()
func (_Aqua *AquaTransactorSession) WithdrawFiat(asset common.Address, ethAddress common.Address, bankAccountAddress string, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aqua.Contract.WithdrawFiat(&_Aqua.TransactOpts, asset, ethAddress, bankAccountAddress, amount, onBehalfOf)
}

// AquaBuyArtEventIterator is returned from FilterBuyArtEvent and is used to iterate over the raw logs and unpacked data for BuyArtEvent events raised by the Aqua contract.
type AquaBuyArtEventIterator struct {
	Event *AquaBuyArtEvent // Event containing the contract specifics and raw log

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
func (it *AquaBuyArtEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBuyArtEvent)
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
		it.Event = new(AquaBuyArtEvent)
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
func (it *AquaBuyArtEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBuyArtEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBuyArtEvent represents a BuyArtEvent event raised by the Aqua contract.
type AquaBuyArtEvent struct {
	Owner  common.Address
	Name   string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyArtEvent is a free log retrieval operation binding the contract event 0xe31bdf128f9b27b2727a6cdf5e2a20b2da9a077eedd779f9d2f048327ebf5f1b.
//
// Solidity: e BuyArtEvent(owner address, name string, amount uint256)
func (_Aqua *AquaFilterer) FilterBuyArtEvent(opts *bind.FilterOpts) (*AquaBuyArtEventIterator, error) {

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "BuyArtEvent")
	if err != nil {
		return nil, err
	}
	return &AquaBuyArtEventIterator{contract: _Aqua.contract, event: "BuyArtEvent", logs: logs, sub: sub}, nil
}

// WatchBuyArtEvent is a free log subscription operation binding the contract event 0xe31bdf128f9b27b2727a6cdf5e2a20b2da9a077eedd779f9d2f048327ebf5f1b.
//
// Solidity: e BuyArtEvent(owner address, name string, amount uint256)
func (_Aqua *AquaFilterer) WatchBuyArtEvent(opts *bind.WatchOpts, sink chan<- *AquaBuyArtEvent) (event.Subscription, error) {

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "BuyArtEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBuyArtEvent)
				if err := _Aqua.contract.UnpackLog(event, "BuyArtEvent", log); err != nil {
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

// AquaBuyPaperEventIterator is returned from FilterBuyPaperEvent and is used to iterate over the raw logs and unpacked data for BuyPaperEvent events raised by the Aqua contract.
type AquaBuyPaperEventIterator struct {
	Event *AquaBuyPaperEvent // Event containing the contract specifics and raw log

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
func (it *AquaBuyPaperEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBuyPaperEvent)
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
		it.Event = new(AquaBuyPaperEvent)
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
func (it *AquaBuyPaperEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBuyPaperEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBuyPaperEvent represents a BuyPaperEvent event raised by the Aqua contract.
type AquaBuyPaperEvent struct {
	Owner      common.Address
	IssuerId   string
	Series     string
	NumberFrom *big.Int
	NumberTo   *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBuyPaperEvent is a free log retrieval operation binding the contract event 0x8321378d68103fb0329c1d1b5ad3684cb77630de035d3afae5698d9d5b243c76.
//
// Solidity: e BuyPaperEvent(owner address, issuerId string, series string, numberFrom uint256, numberTo uint256, amount uint256)
func (_Aqua *AquaFilterer) FilterBuyPaperEvent(opts *bind.FilterOpts) (*AquaBuyPaperEventIterator, error) {

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "BuyPaperEvent")
	if err != nil {
		return nil, err
	}
	return &AquaBuyPaperEventIterator{contract: _Aqua.contract, event: "BuyPaperEvent", logs: logs, sub: sub}, nil
}

// WatchBuyPaperEvent is a free log subscription operation binding the contract event 0x8321378d68103fb0329c1d1b5ad3684cb77630de035d3afae5698d9d5b243c76.
//
// Solidity: e BuyPaperEvent(owner address, issuerId string, series string, numberFrom uint256, numberTo uint256, amount uint256)
func (_Aqua *AquaFilterer) WatchBuyPaperEvent(opts *bind.WatchOpts, sink chan<- *AquaBuyPaperEvent) (event.Subscription, error) {

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "BuyPaperEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBuyPaperEvent)
				if err := _Aqua.contract.UnpackLog(event, "BuyPaperEvent", log); err != nil {
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

// AquaCallbackDepositFiatEventIterator is returned from FilterCallbackDepositFiatEvent and is used to iterate over the raw logs and unpacked data for CallbackDepositFiatEvent events raised by the Aqua contract.
type AquaCallbackDepositFiatEventIterator struct {
	Event *AquaCallbackDepositFiatEvent // Event containing the contract specifics and raw log

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
func (it *AquaCallbackDepositFiatEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaCallbackDepositFiatEvent)
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
		it.Event = new(AquaCallbackDepositFiatEvent)
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
func (it *AquaCallbackDepositFiatEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaCallbackDepositFiatEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaCallbackDepositFiatEvent represents a CallbackDepositFiatEvent event raised by the Aqua contract.
type AquaCallbackDepositFiatEvent struct {
	Result string
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallbackDepositFiatEvent is a free log retrieval operation binding the contract event 0x5d63e8ac2b3cb65246a77bbe8640dbabb46eabc773657d4ca9065973b9270569.
//
// Solidity: e CallbackDepositFiatEvent(_result string, _id uint256)
func (_Aqua *AquaFilterer) FilterCallbackDepositFiatEvent(opts *bind.FilterOpts) (*AquaCallbackDepositFiatEventIterator, error) {

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "CallbackDepositFiatEvent")
	if err != nil {
		return nil, err
	}
	return &AquaCallbackDepositFiatEventIterator{contract: _Aqua.contract, event: "CallbackDepositFiatEvent", logs: logs, sub: sub}, nil
}

// WatchCallbackDepositFiatEvent is a free log subscription operation binding the contract event 0x5d63e8ac2b3cb65246a77bbe8640dbabb46eabc773657d4ca9065973b9270569.
//
// Solidity: e CallbackDepositFiatEvent(_result string, _id uint256)
func (_Aqua *AquaFilterer) WatchCallbackDepositFiatEvent(opts *bind.WatchOpts, sink chan<- *AquaCallbackDepositFiatEvent) (event.Subscription, error) {

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "CallbackDepositFiatEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaCallbackDepositFiatEvent)
				if err := _Aqua.contract.UnpackLog(event, "CallbackDepositFiatEvent", log); err != nil {
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

// AquaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Aqua contract.
type AquaDepositIterator struct {
	Event *AquaDeposit // Event containing the contract specifics and raw log

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
func (it *AquaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaDeposit)
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
		it.Event = new(AquaDeposit)
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
func (it *AquaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaDeposit represents a Deposit event raised by the Aqua contract.
type AquaDeposit struct {
	Reserve    common.Address
	User       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: e Deposit(reserve indexed address, user address, onBehalfOf indexed address, amount uint256)
func (_Aqua *AquaFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address) (*AquaDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AquaDepositIterator{contract: _Aqua.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: e Deposit(reserve indexed address, user address, onBehalfOf indexed address, amount uint256)
func (_Aqua *AquaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AquaDeposit, reserve []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaDeposit)
				if err := _Aqua.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// AquaDepositFiatEventIterator is returned from FilterDepositFiatEvent and is used to iterate over the raw logs and unpacked data for DepositFiatEvent events raised by the Aqua contract.
type AquaDepositFiatEventIterator struct {
	Event *AquaDepositFiatEvent // Event containing the contract specifics and raw log

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
func (it *AquaDepositFiatEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaDepositFiatEvent)
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
		it.Event = new(AquaDepositFiatEvent)
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
func (it *AquaDepositFiatEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaDepositFiatEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaDepositFiatEvent represents a DepositFiatEvent event raised by the Aqua contract.
type AquaDepositFiatEvent struct {
	BankAccountAddress string
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositFiatEvent is a free log retrieval operation binding the contract event 0xda7bc68554258058a7c5b10c5c94d094db73ff500e63bd941eb353260de508aa.
//
// Solidity: e DepositFiatEvent(bankAccountAddress string, amount uint256)
func (_Aqua *AquaFilterer) FilterDepositFiatEvent(opts *bind.FilterOpts) (*AquaDepositFiatEventIterator, error) {

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "DepositFiatEvent")
	if err != nil {
		return nil, err
	}
	return &AquaDepositFiatEventIterator{contract: _Aqua.contract, event: "DepositFiatEvent", logs: logs, sub: sub}, nil
}

// WatchDepositFiatEvent is a free log subscription operation binding the contract event 0xda7bc68554258058a7c5b10c5c94d094db73ff500e63bd941eb353260de508aa.
//
// Solidity: e DepositFiatEvent(bankAccountAddress string, amount uint256)
func (_Aqua *AquaFilterer) WatchDepositFiatEvent(opts *bind.WatchOpts, sink chan<- *AquaDepositFiatEvent) (event.Subscription, error) {

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "DepositFiatEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaDepositFiatEvent)
				if err := _Aqua.contract.UnpackLog(event, "DepositFiatEvent", log); err != nil {
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

// AquaNewOracleAddressEventIterator is returned from FilterNewOracleAddressEvent and is used to iterate over the raw logs and unpacked data for NewOracleAddressEvent events raised by the Aqua contract.
type AquaNewOracleAddressEventIterator struct {
	Event *AquaNewOracleAddressEvent // Event containing the contract specifics and raw log

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
func (it *AquaNewOracleAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaNewOracleAddressEvent)
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
		it.Event = new(AquaNewOracleAddressEvent)
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
func (it *AquaNewOracleAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaNewOracleAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaNewOracleAddressEvent represents a NewOracleAddressEvent event raised by the Aqua contract.
type AquaNewOracleAddressEvent struct {
	OracleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewOracleAddressEvent is a free log retrieval operation binding the contract event 0x51de03c3383d5c428fe200dd5e9561e7018b49e1e836059d2c7e66bc37f09e01.
//
// Solidity: e NewOracleAddressEvent(oracleAddress address)
func (_Aqua *AquaFilterer) FilterNewOracleAddressEvent(opts *bind.FilterOpts) (*AquaNewOracleAddressEventIterator, error) {

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "NewOracleAddressEvent")
	if err != nil {
		return nil, err
	}
	return &AquaNewOracleAddressEventIterator{contract: _Aqua.contract, event: "NewOracleAddressEvent", logs: logs, sub: sub}, nil
}

// WatchNewOracleAddressEvent is a free log subscription operation binding the contract event 0x51de03c3383d5c428fe200dd5e9561e7018b49e1e836059d2c7e66bc37f09e01.
//
// Solidity: e NewOracleAddressEvent(oracleAddress address)
func (_Aqua *AquaFilterer) WatchNewOracleAddressEvent(opts *bind.WatchOpts, sink chan<- *AquaNewOracleAddressEvent) (event.Subscription, error) {

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "NewOracleAddressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaNewOracleAddressEvent)
				if err := _Aqua.contract.UnpackLog(event, "NewOracleAddressEvent", log); err != nil {
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

// AquaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aqua contract.
type AquaOwnershipTransferredIterator struct {
	Event *AquaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AquaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaOwnershipTransferred)
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
		it.Event = new(AquaOwnershipTransferred)
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
func (it *AquaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaOwnershipTransferred represents a OwnershipTransferred event raised by the Aqua contract.
type AquaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Aqua *AquaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AquaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AquaOwnershipTransferredIterator{contract: _Aqua.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Aqua *AquaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AquaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaOwnershipTransferred)
				if err := _Aqua.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AquaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Aqua contract.
type AquaWithdrawIterator struct {
	Event *AquaWithdraw // Event containing the contract specifics and raw log

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
func (it *AquaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaWithdraw)
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
		it.Event = new(AquaWithdraw)
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
func (it *AquaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaWithdraw represents a Withdraw event raised by the Aqua contract.
type AquaWithdraw struct {
	Reserve    common.Address
	User       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: e Withdraw(reserve indexed address, user address, onBehalfOf indexed address, amount uint256)
func (_Aqua *AquaFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address) (*AquaWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Aqua.contract.FilterLogs(opts, "Withdraw", reserveRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AquaWithdrawIterator{contract: _Aqua.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: e Withdraw(reserve indexed address, user address, onBehalfOf indexed address, amount uint256)
func (_Aqua *AquaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AquaWithdraw, reserve []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Aqua.contract.WatchLogs(opts, "Withdraw", reserveRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaWithdraw)
				if err := _Aqua.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
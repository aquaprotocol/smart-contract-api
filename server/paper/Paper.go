// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paper

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

// PaperABI is the input ABI used to generate the binding from.
const PaperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_approved\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paperToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"papers\",\"outputs\":[{\"name\":\"issuerId\",\"type\":\"string\"},{\"name\":\"series\",\"type\":\"string\"},{\"name\":\"numberFrom\",\"type\":\"uint256\"},{\"name\":\"numberTo\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_typeOf\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_issuerId\",\"type\":\"string\"},{\"name\":\"_series\",\"type\":\"string\"},{\"name\":\"_numberFrom\",\"type\":\"uint256\"},{\"name\":\"_numberTo\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_issuerId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_series\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_numberFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_numberTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewPaper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Paper is an auto generated Go binding around an Ethereum contract.
type Paper struct {
	PaperCaller     // Read-only binding to the contract
	PaperTransactor // Write-only binding to the contract
	PaperFilterer   // Log filterer for contract events
}

// PaperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaperSession struct {
	Contract     *Paper            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaperCallerSession struct {
	Contract *PaperCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PaperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaperTransactorSession struct {
	Contract     *PaperTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaperRaw struct {
	Contract *Paper // Generic contract binding to access the raw methods on
}

// PaperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaperCallerRaw struct {
	Contract *PaperCaller // Generic read-only contract binding to access the raw methods on
}

// PaperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaperTransactorRaw struct {
	Contract *PaperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaper creates a new instance of Paper, bound to a specific deployed contract.
func NewPaper(address common.Address, backend bind.ContractBackend) (*Paper, error) {
	contract, err := bindPaper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paper{PaperCaller: PaperCaller{contract: contract}, PaperTransactor: PaperTransactor{contract: contract}, PaperFilterer: PaperFilterer{contract: contract}}, nil
}

// NewPaperCaller creates a new read-only instance of Paper, bound to a specific deployed contract.
func NewPaperCaller(address common.Address, caller bind.ContractCaller) (*PaperCaller, error) {
	contract, err := bindPaper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaperCaller{contract: contract}, nil
}

// NewPaperTransactor creates a new write-only instance of Paper, bound to a specific deployed contract.
func NewPaperTransactor(address common.Address, transactor bind.ContractTransactor) (*PaperTransactor, error) {
	contract, err := bindPaper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaperTransactor{contract: contract}, nil
}

// NewPaperFilterer creates a new log filterer instance of Paper, bound to a specific deployed contract.
func NewPaperFilterer(address common.Address, filterer bind.ContractFilterer) (*PaperFilterer, error) {
	contract, err := bindPaper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaperFilterer{contract: contract}, nil
}

// bindPaper binds a generic wrapper to an already deployed contract.
func bindPaper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paper *PaperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Paper.Contract.PaperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paper *PaperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paper.Contract.PaperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paper *PaperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paper.Contract.PaperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paper *PaperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Paper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paper *PaperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paper *PaperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paper.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Paper *PaperCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Paper.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Paper *PaperSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Paper.Contract.BalanceOf(&_Paper.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Paper *PaperCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Paper.Contract.BalanceOf(&_Paper.CallOpts, _owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Paper *PaperCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Paper.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Paper *PaperSession) IsOwner() (bool, error) {
	return _Paper.Contract.IsOwner(&_Paper.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Paper *PaperCallerSession) IsOwner() (bool, error) {
	return _Paper.Contract.IsOwner(&_Paper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Paper *PaperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Paper.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Paper *PaperSession) Owner() (common.Address, error) {
	return _Paper.Contract.Owner(&_Paper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Paper *PaperCallerSession) Owner() (common.Address, error) {
	return _Paper.Contract.Owner(&_Paper.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Paper *PaperCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Paper.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Paper *PaperSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Paper.Contract.OwnerOf(&_Paper.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Paper *PaperCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Paper.Contract.OwnerOf(&_Paper.CallOpts, _tokenId)
}

// PaperToOwner is a free data retrieval call binding the contract method 0x4cb835ed.
//
// Solidity: function paperToOwner( uint256) constant returns(address)
func (_Paper *PaperCaller) PaperToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Paper.contract.Call(opts, out, "paperToOwner", arg0)
	return *ret0, err
}

// PaperToOwner is a free data retrieval call binding the contract method 0x4cb835ed.
//
// Solidity: function paperToOwner( uint256) constant returns(address)
func (_Paper *PaperSession) PaperToOwner(arg0 *big.Int) (common.Address, error) {
	return _Paper.Contract.PaperToOwner(&_Paper.CallOpts, arg0)
}

// PaperToOwner is a free data retrieval call binding the contract method 0x4cb835ed.
//
// Solidity: function paperToOwner( uint256) constant returns(address)
func (_Paper *PaperCallerSession) PaperToOwner(arg0 *big.Int) (common.Address, error) {
	return _Paper.Contract.PaperToOwner(&_Paper.CallOpts, arg0)
}

// Papers is a free data retrieval call binding the contract method 0x5e775a27.
//
// Solidity: function papers( uint256) constant returns(issuerId string, series string, numberFrom uint256, numberTo uint256, price uint256, id uint256)
func (_Paper *PaperCaller) Papers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IssuerId   string
	Series     string
	NumberFrom *big.Int
	NumberTo   *big.Int
	Price      *big.Int
	Id         *big.Int
}, error) {
	ret := new(struct {
		IssuerId   string
		Series     string
		NumberFrom *big.Int
		NumberTo   *big.Int
		Price      *big.Int
		Id         *big.Int
	})
	out := ret
	err := _Paper.contract.Call(opts, out, "papers", arg0)
	return *ret, err
}

// Papers is a free data retrieval call binding the contract method 0x5e775a27.
//
// Solidity: function papers( uint256) constant returns(issuerId string, series string, numberFrom uint256, numberTo uint256, price uint256, id uint256)
func (_Paper *PaperSession) Papers(arg0 *big.Int) (struct {
	IssuerId   string
	Series     string
	NumberFrom *big.Int
	NumberTo   *big.Int
	Price      *big.Int
	Id         *big.Int
}, error) {
	return _Paper.Contract.Papers(&_Paper.CallOpts, arg0)
}

// Papers is a free data retrieval call binding the contract method 0x5e775a27.
//
// Solidity: function papers( uint256) constant returns(issuerId string, series string, numberFrom uint256, numberTo uint256, price uint256, id uint256)
func (_Paper *PaperCallerSession) Papers(arg0 *big.Int) (struct {
	IssuerId   string
	Series     string
	NumberFrom *big.Int
	NumberTo   *big.Int
	Price      *big.Int
	Id         *big.Int
}, error) {
	return _Paper.Contract.Papers(&_Paper.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Paper *PaperTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Paper *PaperSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.Approve(&_Paper.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Paper *PaperTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.Approve(&_Paper.TransactOpts, _approved, _tokenId)
}

// Create is a paid mutator transaction binding the contract method 0xfb24f697.
//
// Solidity: function create(_issuerId string, _series string, _numberFrom uint256, _numberTo uint256, _price uint256) returns()
func (_Paper *PaperTransactor) Create(opts *bind.TransactOpts, _issuerId string, _series string, _numberFrom *big.Int, _numberTo *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paper.contract.Transact(opts, "create", _issuerId, _series, _numberFrom, _numberTo, _price)
}

// Create is a paid mutator transaction binding the contract method 0xfb24f697.
//
// Solidity: function create(_issuerId string, _series string, _numberFrom uint256, _numberTo uint256, _price uint256) returns()
func (_Paper *PaperSession) Create(_issuerId string, _series string, _numberFrom *big.Int, _numberTo *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.Create(&_Paper.TransactOpts, _issuerId, _series, _numberFrom, _numberTo, _price)
}

// Create is a paid mutator transaction binding the contract method 0xfb24f697.
//
// Solidity: function create(_issuerId string, _series string, _numberFrom uint256, _numberTo uint256, _price uint256) returns()
func (_Paper *PaperTransactorSession) Create(_issuerId string, _series string, _numberFrom *big.Int, _numberTo *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.Create(&_Paper.TransactOpts, _issuerId, _series, _numberFrom, _numberTo, _price)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paper *PaperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paper *PaperSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paper.Contract.RenounceOwnership(&_Paper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paper *PaperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paper.Contract.RenounceOwnership(&_Paper.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Paper *PaperTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Paper *PaperSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.TransferFrom(&_Paper.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Paper *PaperTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paper.Contract.TransferFrom(&_Paper.TransactOpts, _from, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Paper *PaperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Paper *PaperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paper.Contract.TransferOwnership(&_Paper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Paper *PaperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paper.Contract.TransferOwnership(&_Paper.TransactOpts, newOwner)
}

// PaperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Paper contract.
type PaperApprovalIterator struct {
	Event *PaperApproval // Event containing the contract specifics and raw log

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
func (it *PaperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaperApproval)
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
		it.Event = new(PaperApproval)
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
func (it *PaperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaperApproval represents a Approval event raised by the Paper contract.
type PaperApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Paper *PaperFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*PaperApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Paper.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PaperApprovalIterator{contract: _Paper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Paper *PaperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PaperApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Paper.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaperApproval)
				if err := _Paper.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PaperNewPaperIterator is returned from FilterNewPaper and is used to iterate over the raw logs and unpacked data for NewPaper events raised by the Paper contract.
type PaperNewPaperIterator struct {
	Event *PaperNewPaper // Event containing the contract specifics and raw log

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
func (it *PaperNewPaperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaperNewPaper)
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
		it.Event = new(PaperNewPaper)
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
func (it *PaperNewPaperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaperNewPaperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaperNewPaper represents a NewPaper event raised by the Paper contract.
type PaperNewPaper struct {
	Id         *big.Int
	IssuerId   string
	Series     string
	NumberFrom *big.Int
	NumberTo   *big.Int
	Price      *big.Int
	Id         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewPaper is a free log retrieval operation binding the contract event 0x004c53bef0288e73ec569d3b97d60f916687ed9713e38c70de7c7871f3fabf78.
//
// Solidity: e NewPaper(id uint256, _issuerId string, _series string, _numberFrom uint256, _numberTo uint256, _price uint256, _id uint256)
func (_Paper *PaperFilterer) FilterNewPaper(opts *bind.FilterOpts) (*PaperNewPaperIterator, error) {

	logs, sub, err := _Paper.contract.FilterLogs(opts, "NewPaper")
	if err != nil {
		return nil, err
	}
	return &PaperNewPaperIterator{contract: _Paper.contract, event: "NewPaper", logs: logs, sub: sub}, nil
}

// WatchNewPaper is a free log subscription operation binding the contract event 0x004c53bef0288e73ec569d3b97d60f916687ed9713e38c70de7c7871f3fabf78.
//
// Solidity: e NewPaper(id uint256, _issuerId string, _series string, _numberFrom uint256, _numberTo uint256, _price uint256, _id uint256)
func (_Paper *PaperFilterer) WatchNewPaper(opts *bind.WatchOpts, sink chan<- *PaperNewPaper) (event.Subscription, error) {

	logs, sub, err := _Paper.contract.WatchLogs(opts, "NewPaper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaperNewPaper)
				if err := _Paper.contract.UnpackLog(event, "NewPaper", log); err != nil {
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

// PaperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paper contract.
type PaperOwnershipTransferredIterator struct {
	Event *PaperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaperOwnershipTransferred)
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
		it.Event = new(PaperOwnershipTransferred)
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
func (it *PaperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaperOwnershipTransferred represents a OwnershipTransferred event raised by the Paper contract.
type PaperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Paper *PaperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaperOwnershipTransferredIterator{contract: _Paper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Paper *PaperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaperOwnershipTransferred)
				if err := _Paper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PaperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Paper contract.
type PaperTransferIterator struct {
	Event *PaperTransfer // Event containing the contract specifics and raw log

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
func (it *PaperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaperTransfer)
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
		it.Event = new(PaperTransfer)
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
func (it *PaperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaperTransfer represents a Transfer event raised by the Paper contract.
type PaperTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Paper *PaperFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*PaperTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Paper.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PaperTransferIterator{contract: _Paper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Paper *PaperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PaperTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Paper.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaperTransfer)
				if err := _Paper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
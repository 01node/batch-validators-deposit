// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiDeposit

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

// MultiDepositMetaData contains all meta data concerning the MultiDeposit contract.
var MultiDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ArraysMustHaveSameLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicKeysNull\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundLeftovers\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"batchDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MultiDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiDepositMetaData.ABI instead.
var MultiDepositABI = MultiDepositMetaData.ABI

// MultiDeposit is an auto generated Go binding around an Ethereum contract.
type MultiDeposit struct {
	MultiDepositCaller     // Read-only binding to the contract
	MultiDepositTransactor // Write-only binding to the contract
	MultiDepositFilterer   // Log filterer for contract events
}

// MultiDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiDepositSession struct {
	Contract     *MultiDeposit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiDepositCallerSession struct {
	Contract *MultiDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MultiDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiDepositTransactorSession struct {
	Contract     *MultiDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MultiDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiDepositRaw struct {
	Contract *MultiDeposit // Generic contract binding to access the raw methods on
}

// MultiDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiDepositCallerRaw struct {
	Contract *MultiDepositCaller // Generic read-only contract binding to access the raw methods on
}

// MultiDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiDepositTransactorRaw struct {
	Contract *MultiDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiDeposit creates a new instance of MultiDeposit, bound to a specific deployed contract.
func NewMultiDeposit(address common.Address, backend bind.ContractBackend) (*MultiDeposit, error) {
	contract, err := bindMultiDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiDeposit{MultiDepositCaller: MultiDepositCaller{contract: contract}, MultiDepositTransactor: MultiDepositTransactor{contract: contract}, MultiDepositFilterer: MultiDepositFilterer{contract: contract}}, nil
}

// NewMultiDepositCaller creates a new read-only instance of MultiDeposit, bound to a specific deployed contract.
func NewMultiDepositCaller(address common.Address, caller bind.ContractCaller) (*MultiDepositCaller, error) {
	contract, err := bindMultiDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiDepositCaller{contract: contract}, nil
}

// NewMultiDepositTransactor creates a new write-only instance of MultiDeposit, bound to a specific deployed contract.
func NewMultiDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiDepositTransactor, error) {
	contract, err := bindMultiDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiDepositTransactor{contract: contract}, nil
}

// NewMultiDepositFilterer creates a new log filterer instance of MultiDeposit, bound to a specific deployed contract.
func NewMultiDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiDepositFilterer, error) {
	contract, err := bindMultiDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiDepositFilterer{contract: contract}, nil
}

// bindMultiDeposit binds a generic wrapper to an already deployed contract.
func bindMultiDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiDeposit *MultiDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiDeposit.Contract.MultiDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiDeposit *MultiDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiDeposit.Contract.MultiDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiDeposit *MultiDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiDeposit.Contract.MultiDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiDeposit *MultiDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiDeposit *MultiDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiDeposit *MultiDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiDeposit.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_MultiDeposit *MultiDepositCaller) DEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiDeposit.contract.Call(opts, &out, "DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_MultiDeposit *MultiDepositSession) DEPOSITAMOUNT() (*big.Int, error) {
	return _MultiDeposit.Contract.DEPOSITAMOUNT(&_MultiDeposit.CallOpts)
}

// DEPOSITAMOUNT is a free data retrieval call binding the contract method 0xec6925a7.
//
// Solidity: function DEPOSIT_AMOUNT() view returns(uint256)
func (_MultiDeposit *MultiDepositCallerSession) DEPOSITAMOUNT() (*big.Int, error) {
	return _MultiDeposit.Contract.DEPOSITAMOUNT(&_MultiDeposit.CallOpts)
}

// DEPOSITCONTRACTADDRESS is a free data retrieval call binding the contract method 0xa524679e.
//
// Solidity: function DEPOSIT_CONTRACT_ADDRESS() view returns(address)
func (_MultiDeposit *MultiDepositCaller) DEPOSITCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiDeposit.contract.Call(opts, &out, "DEPOSIT_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITCONTRACTADDRESS is a free data retrieval call binding the contract method 0xa524679e.
//
// Solidity: function DEPOSIT_CONTRACT_ADDRESS() view returns(address)
func (_MultiDeposit *MultiDepositSession) DEPOSITCONTRACTADDRESS() (common.Address, error) {
	return _MultiDeposit.Contract.DEPOSITCONTRACTADDRESS(&_MultiDeposit.CallOpts)
}

// DEPOSITCONTRACTADDRESS is a free data retrieval call binding the contract method 0xa524679e.
//
// Solidity: function DEPOSIT_CONTRACT_ADDRESS() view returns(address)
func (_MultiDeposit *MultiDepositCallerSession) DEPOSITCONTRACTADDRESS() (common.Address, error) {
	return _MultiDeposit.Contract.DEPOSITCONTRACTADDRESS(&_MultiDeposit.CallOpts)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0xca0bfcce.
//
// Solidity: function batchDeposit(bytes[] _publicKeys, bytes[] _withdrawalCredentials, bytes[] _signatures, bytes32[] _depositDataRoots) payable returns()
func (_MultiDeposit *MultiDepositTransactor) BatchDeposit(opts *bind.TransactOpts, _publicKeys [][]byte, _withdrawalCredentials [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _MultiDeposit.contract.Transact(opts, "batchDeposit", _publicKeys, _withdrawalCredentials, _signatures, _depositDataRoots)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0xca0bfcce.
//
// Solidity: function batchDeposit(bytes[] _publicKeys, bytes[] _withdrawalCredentials, bytes[] _signatures, bytes32[] _depositDataRoots) payable returns()
func (_MultiDeposit *MultiDepositSession) BatchDeposit(_publicKeys [][]byte, _withdrawalCredentials [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _MultiDeposit.Contract.BatchDeposit(&_MultiDeposit.TransactOpts, _publicKeys, _withdrawalCredentials, _signatures, _depositDataRoots)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0xca0bfcce.
//
// Solidity: function batchDeposit(bytes[] _publicKeys, bytes[] _withdrawalCredentials, bytes[] _signatures, bytes32[] _depositDataRoots) payable returns()
func (_MultiDeposit *MultiDepositTransactorSession) BatchDeposit(_publicKeys [][]byte, _withdrawalCredentials [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _MultiDeposit.Contract.BatchDeposit(&_MultiDeposit.TransactOpts, _publicKeys, _withdrawalCredentials, _signatures, _depositDataRoots)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiDeposit *MultiDepositTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiDeposit.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiDeposit *MultiDepositSession) Receive() (*types.Transaction, error) {
	return _MultiDeposit.Contract.Receive(&_MultiDeposit.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiDeposit *MultiDepositTransactorSession) Receive() (*types.Transaction, error) {
	return _MultiDeposit.Contract.Receive(&_MultiDeposit.TransactOpts)
}

// MultiDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the MultiDeposit contract.
type MultiDepositDepositedIterator struct {
	Event *MultiDepositDeposited // Event containing the contract specifics and raw log

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
func (it *MultiDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiDepositDeposited)
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
		it.Event = new(MultiDepositDeposited)
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
func (it *MultiDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiDepositDeposited represents a Deposited event raised by the MultiDeposit contract.
type MultiDepositDeposited struct {
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x3aa80c423bf036b32ea218139163f238936394aeaebf86186810f82c99c0e359.
//
// Solidity: event Deposited(bytes[] pubkeys)
func (_MultiDeposit *MultiDepositFilterer) FilterDeposited(opts *bind.FilterOpts) (*MultiDepositDepositedIterator, error) {

	logs, sub, err := _MultiDeposit.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &MultiDepositDepositedIterator{contract: _MultiDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x3aa80c423bf036b32ea218139163f238936394aeaebf86186810f82c99c0e359.
//
// Solidity: event Deposited(bytes[] pubkeys)
func (_MultiDeposit *MultiDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MultiDepositDeposited) (event.Subscription, error) {

	logs, sub, err := _MultiDeposit.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiDepositDeposited)
				if err := _MultiDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x3aa80c423bf036b32ea218139163f238936394aeaebf86186810f82c99c0e359.
//
// Solidity: event Deposited(bytes[] pubkeys)
func (_MultiDeposit *MultiDepositFilterer) ParseDeposited(log types.Log) (*MultiDepositDeposited, error) {
	event := new(MultiDepositDeposited)
	if err := _MultiDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiDepositRefundLeftoversIterator is returned from FilterRefundLeftovers and is used to iterate over the raw logs and unpacked data for RefundLeftovers events raised by the MultiDeposit contract.
type MultiDepositRefundLeftoversIterator struct {
	Event *MultiDepositRefundLeftovers // Event containing the contract specifics and raw log

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
func (it *MultiDepositRefundLeftoversIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiDepositRefundLeftovers)
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
		it.Event = new(MultiDepositRefundLeftovers)
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
func (it *MultiDepositRefundLeftoversIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiDepositRefundLeftoversIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiDepositRefundLeftovers represents a RefundLeftovers event raised by the MultiDeposit contract.
type MultiDepositRefundLeftovers struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundLeftovers is a free log retrieval operation binding the contract event 0x44b9521765e1d1e2bfca2cb7bc68db4c4abf394f49cf03ac4933887d0291c0aa.
//
// Solidity: event RefundLeftovers(address to, uint256 amount)
func (_MultiDeposit *MultiDepositFilterer) FilterRefundLeftovers(opts *bind.FilterOpts) (*MultiDepositRefundLeftoversIterator, error) {

	logs, sub, err := _MultiDeposit.contract.FilterLogs(opts, "RefundLeftovers")
	if err != nil {
		return nil, err
	}
	return &MultiDepositRefundLeftoversIterator{contract: _MultiDeposit.contract, event: "RefundLeftovers", logs: logs, sub: sub}, nil
}

// WatchRefundLeftovers is a free log subscription operation binding the contract event 0x44b9521765e1d1e2bfca2cb7bc68db4c4abf394f49cf03ac4933887d0291c0aa.
//
// Solidity: event RefundLeftovers(address to, uint256 amount)
func (_MultiDeposit *MultiDepositFilterer) WatchRefundLeftovers(opts *bind.WatchOpts, sink chan<- *MultiDepositRefundLeftovers) (event.Subscription, error) {

	logs, sub, err := _MultiDeposit.contract.WatchLogs(opts, "RefundLeftovers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiDepositRefundLeftovers)
				if err := _MultiDeposit.contract.UnpackLog(event, "RefundLeftovers", log); err != nil {
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

// ParseRefundLeftovers is a log parse operation binding the contract event 0x44b9521765e1d1e2bfca2cb7bc68db4c4abf394f49cf03ac4933887d0291c0aa.
//
// Solidity: event RefundLeftovers(address to, uint256 amount)
func (_MultiDeposit *MultiDepositFilterer) ParseRefundLeftovers(log types.Log) (*MultiDepositRefundLeftovers, error) {
	event := new(MultiDepositRefundLeftovers)
	if err := _MultiDeposit.contract.UnpackLog(event, "RefundLeftovers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

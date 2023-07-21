// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ERC5564AnnouncerMetaData contains all meta data concerning the ERC5564Announcer contract.
var ERC5564AnnouncerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"schemeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stealthAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ephemeralPubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"Announcement\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"schemeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stealthAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ephemeralPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"announce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506104128061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80634d1f95831461002d575b5f80fd5b61004760048036038101906100429190610291565b610049565b005b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f5f0eab8057630ba7676c49b4f21a0231414e79474595be8e4c432fbf6bf0f4e785856040516100a99291906103a7565b60405180910390a450505050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6100da816100c8565b81146100e4575f80fd5b50565b5f813590506100f5816100d1565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610124826100fb565b9050919050565b6101348161011a565b811461013e575f80fd5b50565b5f8135905061014f8161012b565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101a38261015d565b810181811067ffffffffffffffff821117156101c2576101c161016d565b5b80604052505050565b5f6101d46100b7565b90506101e0828261019a565b919050565b5f67ffffffffffffffff8211156101ff576101fe61016d565b5b6102088261015d565b9050602081019050919050565b828183375f83830152505050565b5f610235610230846101e5565b6101cb565b90508281526020810184848401111561025157610250610159565b5b61025c848285610215565b509392505050565b5f82601f83011261027857610277610155565b5b8135610288848260208601610223565b91505092915050565b5f805f80608085870312156102a9576102a86100c0565b5b5f6102b6878288016100e7565b94505060206102c787828801610141565b935050604085013567ffffffffffffffff8111156102e8576102e76100c4565b5b6102f487828801610264565b925050606085013567ffffffffffffffff811115610315576103146100c4565b5b61032187828801610264565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610364578082015181840152602081019050610349565b5f8484015250505050565b5f6103798261032d565b6103838185610337565b9350610393818560208601610347565b61039c8161015d565b840191505092915050565b5f6040820190508181035f8301526103bf818561036f565b905081810360208301526103d3818461036f565b9050939250505056fea2646970667358221220a3505f7ea4556fe821c38e97400a28d6448f70ab6f719fb9f84e52741f66f7ae64736f6c63430008140033",
}

// ERC5564AnnouncerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC5564AnnouncerMetaData.ABI instead.
var ERC5564AnnouncerABI = ERC5564AnnouncerMetaData.ABI

// ERC5564AnnouncerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC5564AnnouncerMetaData.Bin instead.
var ERC5564AnnouncerBin = ERC5564AnnouncerMetaData.Bin

// DeployERC5564Announcer deploys a new Ethereum contract, binding an instance of ERC5564Announcer to it.
func DeployERC5564Announcer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC5564Announcer, error) {
	parsed, err := ERC5564AnnouncerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC5564AnnouncerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC5564Announcer{ERC5564AnnouncerCaller: ERC5564AnnouncerCaller{contract: contract}, ERC5564AnnouncerTransactor: ERC5564AnnouncerTransactor{contract: contract}, ERC5564AnnouncerFilterer: ERC5564AnnouncerFilterer{contract: contract}}, nil
}

// ERC5564Announcer is an auto generated Go binding around an Ethereum contract.
type ERC5564Announcer struct {
	ERC5564AnnouncerCaller     // Read-only binding to the contract
	ERC5564AnnouncerTransactor // Write-only binding to the contract
	ERC5564AnnouncerFilterer   // Log filterer for contract events
}

// ERC5564AnnouncerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC5564AnnouncerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC5564AnnouncerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC5564AnnouncerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC5564AnnouncerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC5564AnnouncerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC5564AnnouncerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC5564AnnouncerSession struct {
	Contract     *ERC5564Announcer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC5564AnnouncerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC5564AnnouncerCallerSession struct {
	Contract *ERC5564AnnouncerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC5564AnnouncerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC5564AnnouncerTransactorSession struct {
	Contract     *ERC5564AnnouncerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC5564AnnouncerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC5564AnnouncerRaw struct {
	Contract *ERC5564Announcer // Generic contract binding to access the raw methods on
}

// ERC5564AnnouncerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC5564AnnouncerCallerRaw struct {
	Contract *ERC5564AnnouncerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC5564AnnouncerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC5564AnnouncerTransactorRaw struct {
	Contract *ERC5564AnnouncerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC5564Announcer creates a new instance of ERC5564Announcer, bound to a specific deployed contract.
func NewERC5564Announcer(address common.Address, backend bind.ContractBackend) (*ERC5564Announcer, error) {
	contract, err := bindERC5564Announcer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC5564Announcer{ERC5564AnnouncerCaller: ERC5564AnnouncerCaller{contract: contract}, ERC5564AnnouncerTransactor: ERC5564AnnouncerTransactor{contract: contract}, ERC5564AnnouncerFilterer: ERC5564AnnouncerFilterer{contract: contract}}, nil
}

// NewERC5564AnnouncerCaller creates a new read-only instance of ERC5564Announcer, bound to a specific deployed contract.
func NewERC5564AnnouncerCaller(address common.Address, caller bind.ContractCaller) (*ERC5564AnnouncerCaller, error) {
	contract, err := bindERC5564Announcer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC5564AnnouncerCaller{contract: contract}, nil
}

// NewERC5564AnnouncerTransactor creates a new write-only instance of ERC5564Announcer, bound to a specific deployed contract.
func NewERC5564AnnouncerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC5564AnnouncerTransactor, error) {
	contract, err := bindERC5564Announcer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC5564AnnouncerTransactor{contract: contract}, nil
}

// NewERC5564AnnouncerFilterer creates a new log filterer instance of ERC5564Announcer, bound to a specific deployed contract.
func NewERC5564AnnouncerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC5564AnnouncerFilterer, error) {
	contract, err := bindERC5564Announcer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC5564AnnouncerFilterer{contract: contract}, nil
}

// bindERC5564Announcer binds a generic wrapper to an already deployed contract.
func bindERC5564Announcer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC5564AnnouncerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC5564Announcer *ERC5564AnnouncerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC5564Announcer.Contract.ERC5564AnnouncerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC5564Announcer *ERC5564AnnouncerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.ERC5564AnnouncerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC5564Announcer *ERC5564AnnouncerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.ERC5564AnnouncerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC5564Announcer *ERC5564AnnouncerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC5564Announcer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC5564Announcer *ERC5564AnnouncerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC5564Announcer *ERC5564AnnouncerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.contract.Transact(opts, method, params...)
}

// Announce is a paid mutator transaction binding the contract method 0x4d1f9583.
//
// Solidity: function announce(uint256 schemeId, address stealthAddress, bytes ephemeralPubKey, bytes metadata) returns()
func (_ERC5564Announcer *ERC5564AnnouncerTransactor) Announce(opts *bind.TransactOpts, schemeId *big.Int, stealthAddress common.Address, ephemeralPubKey []byte, metadata []byte) (*types.Transaction, error) {
	return _ERC5564Announcer.contract.Transact(opts, "announce", schemeId, stealthAddress, ephemeralPubKey, metadata)
}

// Announce is a paid mutator transaction binding the contract method 0x4d1f9583.
//
// Solidity: function announce(uint256 schemeId, address stealthAddress, bytes ephemeralPubKey, bytes metadata) returns()
func (_ERC5564Announcer *ERC5564AnnouncerSession) Announce(schemeId *big.Int, stealthAddress common.Address, ephemeralPubKey []byte, metadata []byte) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.Announce(&_ERC5564Announcer.TransactOpts, schemeId, stealthAddress, ephemeralPubKey, metadata)
}

// Announce is a paid mutator transaction binding the contract method 0x4d1f9583.
//
// Solidity: function announce(uint256 schemeId, address stealthAddress, bytes ephemeralPubKey, bytes metadata) returns()
func (_ERC5564Announcer *ERC5564AnnouncerTransactorSession) Announce(schemeId *big.Int, stealthAddress common.Address, ephemeralPubKey []byte, metadata []byte) (*types.Transaction, error) {
	return _ERC5564Announcer.Contract.Announce(&_ERC5564Announcer.TransactOpts, schemeId, stealthAddress, ephemeralPubKey, metadata)
}

// ERC5564AnnouncerAnnouncementIterator is returned from FilterAnnouncement and is used to iterate over the raw logs and unpacked data for Announcement events raised by the ERC5564Announcer contract.
type ERC5564AnnouncerAnnouncementIterator struct {
	Event *ERC5564AnnouncerAnnouncement // Event containing the contract specifics and raw log

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
func (it *ERC5564AnnouncerAnnouncementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC5564AnnouncerAnnouncement)
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
		it.Event = new(ERC5564AnnouncerAnnouncement)
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
func (it *ERC5564AnnouncerAnnouncementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC5564AnnouncerAnnouncementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC5564AnnouncerAnnouncement represents a Announcement event raised by the ERC5564Announcer contract.
type ERC5564AnnouncerAnnouncement struct {
	SchemeId        *big.Int
	StealthAddress  common.Address
	Caller          common.Address
	EphemeralPubKey []byte
	Metadata        []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAnnouncement is a free log retrieval operation binding the contract event 0x5f0eab8057630ba7676c49b4f21a0231414e79474595be8e4c432fbf6bf0f4e7.
//
// Solidity: event Announcement(uint256 indexed schemeId, address indexed stealthAddress, address indexed caller, bytes ephemeralPubKey, bytes metadata)
func (_ERC5564Announcer *ERC5564AnnouncerFilterer) FilterAnnouncement(opts *bind.FilterOpts, schemeId []*big.Int, stealthAddress []common.Address, caller []common.Address) (*ERC5564AnnouncerAnnouncementIterator, error) {

	var schemeIdRule []interface{}
	for _, schemeIdItem := range schemeId {
		schemeIdRule = append(schemeIdRule, schemeIdItem)
	}
	var stealthAddressRule []interface{}
	for _, stealthAddressItem := range stealthAddress {
		stealthAddressRule = append(stealthAddressRule, stealthAddressItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC5564Announcer.contract.FilterLogs(opts, "Announcement", schemeIdRule, stealthAddressRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC5564AnnouncerAnnouncementIterator{contract: _ERC5564Announcer.contract, event: "Announcement", logs: logs, sub: sub}, nil
}

// WatchAnnouncement is a free log subscription operation binding the contract event 0x5f0eab8057630ba7676c49b4f21a0231414e79474595be8e4c432fbf6bf0f4e7.
//
// Solidity: event Announcement(uint256 indexed schemeId, address indexed stealthAddress, address indexed caller, bytes ephemeralPubKey, bytes metadata)
func (_ERC5564Announcer *ERC5564AnnouncerFilterer) WatchAnnouncement(opts *bind.WatchOpts, sink chan<- *ERC5564AnnouncerAnnouncement, schemeId []*big.Int, stealthAddress []common.Address, caller []common.Address) (event.Subscription, error) {

	var schemeIdRule []interface{}
	for _, schemeIdItem := range schemeId {
		schemeIdRule = append(schemeIdRule, schemeIdItem)
	}
	var stealthAddressRule []interface{}
	for _, stealthAddressItem := range stealthAddress {
		stealthAddressRule = append(stealthAddressRule, stealthAddressItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC5564Announcer.contract.WatchLogs(opts, "Announcement", schemeIdRule, stealthAddressRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC5564AnnouncerAnnouncement)
				if err := _ERC5564Announcer.contract.UnpackLog(event, "Announcement", log); err != nil {
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

// ParseAnnouncement is a log parse operation binding the contract event 0x5f0eab8057630ba7676c49b4f21a0231414e79474595be8e4c432fbf6bf0f4e7.
//
// Solidity: event Announcement(uint256 indexed schemeId, address indexed stealthAddress, address indexed caller, bytes ephemeralPubKey, bytes metadata)
func (_ERC5564Announcer *ERC5564AnnouncerFilterer) ParseAnnouncement(log types.Log) (*ERC5564AnnouncerAnnouncement, error) {
	event := new(ERC5564AnnouncerAnnouncement)
	if err := _ERC5564Announcer.contract.UnpackLog(event, "Announcement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

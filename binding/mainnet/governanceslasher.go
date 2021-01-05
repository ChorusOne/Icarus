// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainnet

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceSlasherABI is the input ABI used to generate the binding from.
const GovernanceSlasherABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"approveSlashing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"electionLessers\",\"type\":\"address[]\"},{\"name\":\"electionGreaters\",\"type\":\"address[]\"},{\"name\":\"electionIndices\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getApprovedSlashing\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashingApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GovernanceSlashPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// GovernanceSlasherBin is the compiled bytecode used for deploying new contracts.
var GovernanceSlasherBin = "0x60806040526000610017640100000000610066810204565b60008054600160a060020a031916600160a060020a0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b610bad806100796000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480638f32d59b1161008e5780638f32d59b14610149578063a91ee0dc14610151578063c4d66de814610177578063d6cbae021461019d578063da853aed146102c1578063f2fde38b146102f9576100c6565b8063158ef93e146100cb5780631c75f5c2146100e7578063715018a6146101155780637b1039991461011d5780638da5cb5b14610141575b600080fd5b6100d361031f565b604080519115158252519081900360200190f35b610113600480360360408110156100fd57600080fd5b50600160a060020a038135169060200135610340565b005b61011361040b565b6101256104ae565b60408051600160a060020a039092168252519081900360200190f35b6101256104bd565b6100d36104cc565b6101136004803603602081101561016757600080fd5b5035600160a060020a03166104f0565b6101136004803603602081101561018d57600080fd5b5035600160a060020a03166105f3565b6100d3600480360360808110156101b357600080fd5b600160a060020a0382351691908101906040810160208201356401000000008111156101de57600080fd5b8201836020820111156101f057600080fd5b8035906020019184602083028401116401000000008311171561021257600080fd5b91939092909160208101903564010000000081111561023057600080fd5b82018360208201111561024257600080fd5b8035906020019184602083028401116401000000008311171561026457600080fd5b91939092909160208101903564010000000081111561028257600080fd5b82018360208201111561029457600080fd5b803590602001918460208302840111640100000000831117156102b657600080fd5b5090925090506106b0565b6102e7600480360360208110156102d757600080fd5b5035600160a060020a03166108dc565b60408051918252519081900360200190f35b6101136004803603602081101561030f57600080fd5b5035600160a060020a03166108f7565b60005474010000000000000000000000000000000000000000900460ff1681565b6103486104cc565b151561038c576040805160e560020a62461bcd0281526020600482018190526024820152600080516020610b62833981519152604482015290519081900360640190fd5b600160a060020a0382166000908152600260205260409020546103b5908263ffffffff61094c16565b600160a060020a038316600081815260026020908152604091829020939093558051848152905191927f031b4696d0f74746e384670f56dacc22f2565b5637b797d8cf861dd5dfe73ed592918290030190a25050565b6104136104cc565b1515610457576040805160e560020a62461bcd0281526020600482018190526024820152600080516020610b62833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600154600160a060020a031681565b600054600160a060020a031690565b60008054600160a060020a03166104e16109b0565b600160a060020a031614905090565b6104f86104cc565b151561053c576040805160e560020a62461bcd0281526020600482018190526024820152600080516020610b62833981519152604482015290519081900360640190fd5b600160a060020a038116151561059c576040805160e560020a62461bcd02815260206004820181905260248201527f43616e6e6f7420726567697374657220746865206e756c6c2061646472657373604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383169081179091556040517f27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b90600090a250565b60005474010000000000000000000000000000000000000000900460ff1615610666576040805160e560020a62461bcd02815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a656400000000604482015290519081900360640190fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556106a4336109b4565b6106ad816104f0565b50565b600160a060020a038716600090815260026020526040812054818111610720576040805160e560020a62461bcd02815260206004820152601e60248201527f4e6f2070656e616c747920676976656e20627920676f7665726e616e63650000604482015290519081900360640190fd5b600160a060020a038916600090815260026020526040812055610741610a66565b600160a060020a03166331993fc98a836000808d8d8d8d8d8d6040518b63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808b600160a060020a0316600160a060020a031681526020018a815260200189600160a060020a0316600160a060020a0316815260200188815260200180602001806020018060200184810384528a8a82818152602001925060200280828437600083820152601f01601f19169091018581038452888152602090810191508990890280828437600083820152601f01601f19169091018581038352868152602090810191508790870280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b15801561087657600080fd5b505af115801561088a573d6000803e3d6000fd5b5050604080518481529051600160a060020a038d1693507fd2b041bb62d3ac9e704aadbea1d3a21b6f5b4677d0766e204c2d30dfc1a022f992509081900360200190a250600198975050505050505050565b600160a060020a031660009081526002602052604090205490565b6108ff6104cc565b1515610943576040805160e560020a62461bcd0281526020600482018190526024820152600080516020610b62833981519152604482015290519081900360640190fd5b6106ad816109b4565b6000828201838110156109a9576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b600160a060020a03811615156109fe5760405160e560020a62461bcd028152600401808060200182810382526026815260200180610b3c6026913960400191505060405180910390fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154604080517f4c6f636b6564476f6c64000000000000000000000000000000000000000000006020808301919091528251808303600a018152602a83018085528151918301919091207fdcf0aaed00000000000000000000000000000000000000000000000000000000909152602e8301529151600093600160a060020a03169263dcf0aaed92604e8082019391829003018186803b158015610b0a57600080fd5b505afa158015610b1e573d6000803e3d6000fd5b505050506040513d6020811015610b3457600080fd5b505190509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a165627a7a723058209796ee1b16ec1fc419aaa48c4b94f88868a78c99ba70521e0ae3ad802f02ecbc0029"

// DeployGovernanceSlasher deploys a new Ethereum contract, binding an instance of GovernanceSlasher to it.
func DeployGovernanceSlasher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernanceSlasher, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceSlasherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceSlasherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceSlasher{GovernanceSlasherCaller: GovernanceSlasherCaller{contract: contract}, GovernanceSlasherTransactor: GovernanceSlasherTransactor{contract: contract}, GovernanceSlasherFilterer: GovernanceSlasherFilterer{contract: contract}}, nil
}

// GovernanceSlasher is an auto generated Go binding around an Ethereum contract.
type GovernanceSlasher struct {
	GovernanceSlasherCaller     // Read-only binding to the contract
	GovernanceSlasherTransactor // Write-only binding to the contract
	GovernanceSlasherFilterer   // Log filterer for contract events
}

// GovernanceSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSlasherSession struct {
	Contract     *GovernanceSlasher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GovernanceSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceSlasherCallerSession struct {
	Contract *GovernanceSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GovernanceSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceSlasherTransactorSession struct {
	Contract     *GovernanceSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GovernanceSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceSlasherRaw struct {
	Contract *GovernanceSlasher // Generic contract binding to access the raw methods on
}

// GovernanceSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceSlasherCallerRaw struct {
	Contract *GovernanceSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceSlasherTransactorRaw struct {
	Contract *GovernanceSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceSlasher creates a new instance of GovernanceSlasher, bound to a specific deployed contract.
func NewGovernanceSlasher(address common.Address, backend bind.ContractBackend) (*GovernanceSlasher, error) {
	contract, err := bindGovernanceSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasher{GovernanceSlasherCaller: GovernanceSlasherCaller{contract: contract}, GovernanceSlasherTransactor: GovernanceSlasherTransactor{contract: contract}, GovernanceSlasherFilterer: GovernanceSlasherFilterer{contract: contract}}, nil
}

// NewGovernanceSlasherCaller creates a new read-only instance of GovernanceSlasher, bound to a specific deployed contract.
func NewGovernanceSlasherCaller(address common.Address, caller bind.ContractCaller) (*GovernanceSlasherCaller, error) {
	contract, err := bindGovernanceSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherCaller{contract: contract}, nil
}

// NewGovernanceSlasherTransactor creates a new write-only instance of GovernanceSlasher, bound to a specific deployed contract.
func NewGovernanceSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceSlasherTransactor, error) {
	contract, err := bindGovernanceSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherTransactor{contract: contract}, nil
}

// NewGovernanceSlasherFilterer creates a new log filterer instance of GovernanceSlasher, bound to a specific deployed contract.
func NewGovernanceSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceSlasherFilterer, error) {
	contract, err := bindGovernanceSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherFilterer{contract: contract}, nil
}

// bindGovernanceSlasher binds a generic wrapper to an already deployed contract.
func bindGovernanceSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceSlasherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGovernanceSlasherABI parses the ABI
func ParseGovernanceSlasherABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceSlasherABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceSlasher *GovernanceSlasherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceSlasher.Contract.GovernanceSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceSlasher *GovernanceSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.GovernanceSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceSlasher *GovernanceSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.GovernanceSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceSlasher *GovernanceSlasherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceSlasher *GovernanceSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceSlasher *GovernanceSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.contract.Transact(opts, method, params...)
}

// GetApprovedSlashing is a free data retrieval call binding the contract method 0xda853aed.
//
// Solidity: function getApprovedSlashing(address account) constant returns(uint256)
func (_GovernanceSlasher *GovernanceSlasherCaller) GetApprovedSlashing(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernanceSlasher.contract.Call(opts, out, "getApprovedSlashing", account)
	return *ret0, err
}

// GetApprovedSlashing is a free data retrieval call binding the contract method 0xda853aed.
//
// Solidity: function getApprovedSlashing(address account) constant returns(uint256)
func (_GovernanceSlasher *GovernanceSlasherSession) GetApprovedSlashing(account common.Address) (*big.Int, error) {
	return _GovernanceSlasher.Contract.GetApprovedSlashing(&_GovernanceSlasher.CallOpts, account)
}

// GetApprovedSlashing is a free data retrieval call binding the contract method 0xda853aed.
//
// Solidity: function getApprovedSlashing(address account) constant returns(uint256)
func (_GovernanceSlasher *GovernanceSlasherCallerSession) GetApprovedSlashing(account common.Address) (*big.Int, error) {
	return _GovernanceSlasher.Contract.GetApprovedSlashing(&_GovernanceSlasher.CallOpts, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GovernanceSlasher.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherSession) Initialized() (bool, error) {
	return _GovernanceSlasher.Contract.Initialized(&_GovernanceSlasher.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherCallerSession) Initialized() (bool, error) {
	return _GovernanceSlasher.Contract.Initialized(&_GovernanceSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GovernanceSlasher.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherSession) IsOwner() (bool, error) {
	return _GovernanceSlasher.Contract.IsOwner(&_GovernanceSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GovernanceSlasher *GovernanceSlasherCallerSession) IsOwner() (bool, error) {
	return _GovernanceSlasher.Contract.IsOwner(&_GovernanceSlasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernanceSlasher.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherSession) Owner() (common.Address, error) {
	return _GovernanceSlasher.Contract.Owner(&_GovernanceSlasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherCallerSession) Owner() (common.Address, error) {
	return _GovernanceSlasher.Contract.Owner(&_GovernanceSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovernanceSlasher.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherSession) Registry() (common.Address, error) {
	return _GovernanceSlasher.Contract.Registry(&_GovernanceSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GovernanceSlasher *GovernanceSlasherCallerSession) Registry() (common.Address, error) {
	return _GovernanceSlasher.Contract.Registry(&_GovernanceSlasher.CallOpts)
}

// ApproveSlashing is a paid mutator transaction binding the contract method 0x1c75f5c2.
//
// Solidity: function approveSlashing(address account, uint256 penalty) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactor) ApproveSlashing(opts *bind.TransactOpts, account common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "approveSlashing", account, penalty)
}

// ApproveSlashing is a paid mutator transaction binding the contract method 0x1c75f5c2.
//
// Solidity: function approveSlashing(address account, uint256 penalty) returns()
func (_GovernanceSlasher *GovernanceSlasherSession) ApproveSlashing(account common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.ApproveSlashing(&_GovernanceSlasher.TransactOpts, account, penalty)
}

// ApproveSlashing is a paid mutator transaction binding the contract method 0x1c75f5c2.
//
// Solidity: function approveSlashing(address account, uint256 penalty) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) ApproveSlashing(account common.Address, penalty *big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.ApproveSlashing(&_GovernanceSlasher.TransactOpts, account, penalty)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "initialize", registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.Initialize(&_GovernanceSlasher.TransactOpts, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.Initialize(&_GovernanceSlasher.TransactOpts, registryAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovernanceSlasher *GovernanceSlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovernanceSlasher *GovernanceSlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.RenounceOwnership(&_GovernanceSlasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.RenounceOwnership(&_GovernanceSlasher.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.SetRegistry(&_GovernanceSlasher.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.SetRegistry(&_GovernanceSlasher.TransactOpts, registryAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xd6cbae02.
//
// Solidity: function slash(address account, address[] electionLessers, address[] electionGreaters, uint256[] electionIndices) returns(bool)
func (_GovernanceSlasher *GovernanceSlasherTransactor) Slash(opts *bind.TransactOpts, account common.Address, electionLessers []common.Address, electionGreaters []common.Address, electionIndices []*big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "slash", account, electionLessers, electionGreaters, electionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0xd6cbae02.
//
// Solidity: function slash(address account, address[] electionLessers, address[] electionGreaters, uint256[] electionIndices) returns(bool)
func (_GovernanceSlasher *GovernanceSlasherSession) Slash(account common.Address, electionLessers []common.Address, electionGreaters []common.Address, electionIndices []*big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.Slash(&_GovernanceSlasher.TransactOpts, account, electionLessers, electionGreaters, electionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0xd6cbae02.
//
// Solidity: function slash(address account, address[] electionLessers, address[] electionGreaters, uint256[] electionIndices) returns(bool)
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) Slash(account common.Address, electionLessers []common.Address, electionGreaters []common.Address, electionIndices []*big.Int) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.Slash(&_GovernanceSlasher.TransactOpts, account, electionLessers, electionGreaters, electionIndices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovernanceSlasher *GovernanceSlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.TransferOwnership(&_GovernanceSlasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovernanceSlasher *GovernanceSlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceSlasher.Contract.TransferOwnership(&_GovernanceSlasher.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_GovernanceSlasher *GovernanceSlasherFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _GovernanceSlasher.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "GovernanceSlashPerformed":
		event, err = _GovernanceSlasher.ParseGovernanceSlashPerformed(log)
	case "OwnershipTransferred":
		event, err = _GovernanceSlasher.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _GovernanceSlasher.ParseRegistrySet(log)
	case "SlashingApproved":
		event, err = _GovernanceSlasher.ParseSlashingApproved(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// GovernanceSlasherGovernanceSlashPerformedIterator is returned from FilterGovernanceSlashPerformed and is used to iterate over the raw logs and unpacked data for GovernanceSlashPerformed events raised by the GovernanceSlasher contract.
type GovernanceSlasherGovernanceSlashPerformedIterator struct {
	Event *GovernanceSlasherGovernanceSlashPerformed // Event containing the contract specifics and raw log

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
func (it *GovernanceSlasherGovernanceSlashPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceSlasherGovernanceSlashPerformed)
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
		it.Event = new(GovernanceSlasherGovernanceSlashPerformed)
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
func (it *GovernanceSlasherGovernanceSlashPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceSlasherGovernanceSlashPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceSlasherGovernanceSlashPerformed represents a GovernanceSlashPerformed event raised by the GovernanceSlasher contract.
type GovernanceSlasherGovernanceSlashPerformed struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSlashPerformed is a free log retrieval operation binding the contract event 0xd2b041bb62d3ac9e704aadbea1d3a21b6f5b4677d0766e204c2d30dfc1a022f9.
//
// Solidity: event GovernanceSlashPerformed(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) FilterGovernanceSlashPerformed(opts *bind.FilterOpts, account []common.Address) (*GovernanceSlasherGovernanceSlashPerformedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.FilterLogs(opts, "GovernanceSlashPerformed", accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherGovernanceSlashPerformedIterator{contract: _GovernanceSlasher.contract, event: "GovernanceSlashPerformed", logs: logs, sub: sub}, nil
}

// WatchGovernanceSlashPerformed is a free log subscription operation binding the contract event 0xd2b041bb62d3ac9e704aadbea1d3a21b6f5b4677d0766e204c2d30dfc1a022f9.
//
// Solidity: event GovernanceSlashPerformed(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) WatchGovernanceSlashPerformed(opts *bind.WatchOpts, sink chan<- *GovernanceSlasherGovernanceSlashPerformed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.WatchLogs(opts, "GovernanceSlashPerformed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceSlasherGovernanceSlashPerformed)
				if err := _GovernanceSlasher.contract.UnpackLog(event, "GovernanceSlashPerformed", log); err != nil {
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

// ParseGovernanceSlashPerformed is a log parse operation binding the contract event 0xd2b041bb62d3ac9e704aadbea1d3a21b6f5b4677d0766e204c2d30dfc1a022f9.
//
// Solidity: event GovernanceSlashPerformed(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) ParseGovernanceSlashPerformed(log types.Log) (*GovernanceSlasherGovernanceSlashPerformed, error) {
	event := new(GovernanceSlasherGovernanceSlashPerformed)
	if err := _GovernanceSlasher.contract.UnpackLog(event, "GovernanceSlashPerformed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceSlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GovernanceSlasher contract.
type GovernanceSlasherOwnershipTransferredIterator struct {
	Event *GovernanceSlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovernanceSlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceSlasherOwnershipTransferred)
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
		it.Event = new(GovernanceSlasherOwnershipTransferred)
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
func (it *GovernanceSlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceSlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceSlasherOwnershipTransferred represents a OwnershipTransferred event raised by the GovernanceSlasher contract.
type GovernanceSlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovernanceSlasher *GovernanceSlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernanceSlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherOwnershipTransferredIterator{contract: _GovernanceSlasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovernanceSlasher *GovernanceSlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceSlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceSlasherOwnershipTransferred)
				if err := _GovernanceSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GovernanceSlasher *GovernanceSlasherFilterer) ParseOwnershipTransferred(log types.Log) (*GovernanceSlasherOwnershipTransferred, error) {
	event := new(GovernanceSlasherOwnershipTransferred)
	if err := _GovernanceSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceSlasherRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the GovernanceSlasher contract.
type GovernanceSlasherRegistrySetIterator struct {
	Event *GovernanceSlasherRegistrySet // Event containing the contract specifics and raw log

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
func (it *GovernanceSlasherRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceSlasherRegistrySet)
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
		it.Event = new(GovernanceSlasherRegistrySet)
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
func (it *GovernanceSlasherRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceSlasherRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceSlasherRegistrySet represents a RegistrySet event raised by the GovernanceSlasher contract.
type GovernanceSlasherRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GovernanceSlasher *GovernanceSlasherFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*GovernanceSlasherRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherRegistrySetIterator{contract: _GovernanceSlasher.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GovernanceSlasher *GovernanceSlasherFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *GovernanceSlasherRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceSlasherRegistrySet)
				if err := _GovernanceSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GovernanceSlasher *GovernanceSlasherFilterer) ParseRegistrySet(log types.Log) (*GovernanceSlasherRegistrySet, error) {
	event := new(GovernanceSlasherRegistrySet)
	if err := _GovernanceSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceSlasherSlashingApprovedIterator is returned from FilterSlashingApproved and is used to iterate over the raw logs and unpacked data for SlashingApproved events raised by the GovernanceSlasher contract.
type GovernanceSlasherSlashingApprovedIterator struct {
	Event *GovernanceSlasherSlashingApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceSlasherSlashingApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceSlasherSlashingApproved)
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
		it.Event = new(GovernanceSlasherSlashingApproved)
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
func (it *GovernanceSlasherSlashingApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceSlasherSlashingApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceSlasherSlashingApproved represents a SlashingApproved event raised by the GovernanceSlasher contract.
type GovernanceSlasherSlashingApproved struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashingApproved is a free log retrieval operation binding the contract event 0x031b4696d0f74746e384670f56dacc22f2565b5637b797d8cf861dd5dfe73ed5.
//
// Solidity: event SlashingApproved(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) FilterSlashingApproved(opts *bind.FilterOpts, account []common.Address) (*GovernanceSlasherSlashingApprovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.FilterLogs(opts, "SlashingApproved", accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceSlasherSlashingApprovedIterator{contract: _GovernanceSlasher.contract, event: "SlashingApproved", logs: logs, sub: sub}, nil
}

// WatchSlashingApproved is a free log subscription operation binding the contract event 0x031b4696d0f74746e384670f56dacc22f2565b5637b797d8cf861dd5dfe73ed5.
//
// Solidity: event SlashingApproved(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) WatchSlashingApproved(opts *bind.WatchOpts, sink chan<- *GovernanceSlasherSlashingApproved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GovernanceSlasher.contract.WatchLogs(opts, "SlashingApproved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceSlasherSlashingApproved)
				if err := _GovernanceSlasher.contract.UnpackLog(event, "SlashingApproved", log); err != nil {
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

// ParseSlashingApproved is a log parse operation binding the contract event 0x031b4696d0f74746e384670f56dacc22f2565b5637b797d8cf861dd5dfe73ed5.
//
// Solidity: event SlashingApproved(address indexed account, uint256 amount)
func (_GovernanceSlasher *GovernanceSlasherFilterer) ParseSlashingApproved(log types.Log) (*GovernanceSlasherSlashingApproved, error) {
	event := new(GovernanceSlasherSlashingApproved)
	if err := _GovernanceSlasher.contract.UnpackLog(event, "SlashingApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

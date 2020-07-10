// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baklava

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

// GoldTokenABI is the input ABI used to generate the binding from.
const GoldTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"transferWithComment\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"TransferComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// GoldTokenBin is the compiled bytecode used for deploying new contracts.
var GoldTokenBin = "0x6080604052600061001764010000000061006e810204565b6000805461010060a860020a031916610100600160a060020a038416908102919091178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610072565b3390565b611832806100816000396000f3fe608060405234801561001057600080fd5b50600436106101465760003560e060020a900480638da5cb5b116100bc578063a91ee0dc11610080578063a91ee0dc1461039a578063b921e163146103c0578063c4d66de8146103dd578063dd62ed3e14610403578063e1d6aceb14610431578063f2fde38b146104b657610146565b80638da5cb5b1461032a5780638f32d59b1461033257806395d89b411461033a578063a457c2d714610342578063a9059cbb1461036e57610146565b8063313ce5671161010e578063313ce56714610260578063395093511461027e57806340c10f19146102aa57806370a08231146102d6578063715018a6146102fc5780637b1039991461030657610146565b806306fdde031461014b578063095ea7b3146101c8578063158ef93e1461020857806318160ddd1461021057806323b872dd1461022a575b600080fd5b6101536104dc565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018d578181015183820152602001610175565b50505050905090810190601f1680156101ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f4600480360360408110156101de57600080fd5b50600160a060020a038135169060200135610513565b604080519115158252519081900360200190f35b6101f46105dc565b6102186105e5565b60408051918252519081900360200190f35b6101f46004803603606081101561024057600080fd5b50600160a060020a038135811691602081013590911690604001356105eb565b61026861099f565b6040805160ff9092168252519081900360200190f35b6101f46004803603604081101561029457600080fd5b50600160a060020a0381351690602001356109a4565b6101f4600480360360408110156102c057600080fd5b50600160a060020a038135169060200135610aa5565b610218600480360360208110156102ec57600080fd5b5035600160a060020a0316610ce8565b610304610cf5565b005b61030e610daf565b60408051600160a060020a039092168252519081900360200190f35b61030e610dbe565b6101f4610dd2565b610153610dfb565b6101f46004803603604081101561035857600080fd5b50600160a060020a038135169060200135610e32565b6101f46004803603604081101561038457600080fd5b50600160a060020a038135169060200135610e67565b610304600480360360208110156103b057600080fd5b5035600160a060020a0316610e7a565b610304600480360360208110156103d657600080fd5b5035610f8f565b610304600480360360208110156103f357600080fd5b5035600160a060020a0316610ffe565b6102186004803603604081101561041957600080fd5b50600160a060020a038135811691602001351661107e565b6101f46004803603606081101561044757600080fd5b600160a060020a038235169160208101359181019060608101604082013564010000000081111561047757600080fd5b82018360208201111561048957600080fd5b803590602001918460018302840111640100000000831117156104ab57600080fd5b5090925090506110a9565b610304600480360360208110156104cc57600080fd5b5035600160a060020a0316611121565b60408051808201909152600981527f43656c6f20476f6c640000000000000000000000000000000000000000000000602082015290565b6000600160a060020a0383161515610575576040805160e560020a62461bcd02815260206004820152601a60248201527f63616e6e6f742073657420616c6c6f77616e636520666f722030000000000000604482015290519081900360640190fd5b336000818152600360209081526040808320600160a060020a03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005460ff1681565b60025490565b60006105f5611188565b600160a060020a031663e5839836306040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561064d57600080fd5b505afa158015610661573d6000803e3d6000fd5b505050506040513d602081101561067757600080fd5b5051156106b85760405160e560020a62461bcd02815260040180806020018281038252602281526020018061175a6022913960400191505060405180910390fd5b600160a060020a03831615156107025760405160e560020a62461bcd02815260040180806020018281038252602a8152602001806117dd602a913960400191505060405180910390fd5b61070b84610ce8565b82111561074c5760405160e560020a62461bcd02815260040180806020018281038252602981526020018061177c6029913960400191505060405180910390fd5b600160a060020a03841660009081526003602090815260408083203384529091529020548211156107b15760405160e560020a62461bcd0281526004018080602001828103825260388152602001806117a56038913960400191505060405180910390fd5b600060fd815a908787876040516020018084600160a060020a0316600160a060020a0316815260200183600160a060020a0316600160a060020a0316815260200182815260200193505050506040516020818303038152906040526040518082805190602001908083835b6020831061083b5780518252601f19909201916020918201910161081c565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381858888f193505050503d806000811461089e576040519150601f19603f3d011682016040523d82523d6000602084013e6108a3565b606091505b50909150508015156108ff576040805160e560020a62461bcd02815260206004820152601960248201527f43656c6f20476f6c64207472616e73666572206661696c656400000000000000604482015290519081900360640190fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054610933908463ffffffff61125d16565b600160a060020a03808716600081815260036020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b601290565b6000600160a060020a0383161515610a06576040805160e560020a62461bcd02815260206004820152601a60248201527f63616e6e6f742073657420616c6c6f77616e636520666f722030000000000000604482015290519081900360640190fd5b336000908152600360209081526040808320600160a060020a038716845290915281205490610a3b828563ffffffff61129f16565b336000818152600360209081526040808320600160a060020a038b16808552908352928190208590558051858152905194955091937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3506001949350505050565b60003315610afd576040805160e560020a62461bcd02815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c00000000000000000000000000000000604482015290519081900360640190fd5b60008211610b55576040805160e560020a62461bcd02815260206004820152601660248201527f6d696e742076616c7565206d757374206265203e203000000000000000000000604482015290519081900360640190fd5b600254610b68908363ffffffff61129f16565b600255600060fd815a604080516000602080830191909152600160a060020a038a168284015260608083018a905283518084039091018152608090920192839052815193949391929182918401908083835b60208310610bd95780518252601f199092019160209182019101610bba565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381858888f193505050503d8060008114610c3c576040519150601f19603f3d011682016040523d82523d6000602084013e610c41565b606091505b5090915050801515610c9d576040805160e560020a62461bcd02815260206004820152601960248201527f43656c6f20476f6c64207472616e73666572206661696c656400000000000000604482015290519081900360640190fd5b604080518481529051600160a060020a038616916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019392505050565b600160a060020a03163190565b610cfd610dd2565b1515610d53576040805160e560020a62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60008054604051610100909104600160a060020a0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805474ffffffffffffffffffffffffffffffffffffffff0019169055565b600154600160a060020a031681565b6000546101009004600160a060020a031690565b600080546101009004600160a060020a0316610dec6112fc565b600160a060020a031614905090565b60408051808201909152600481527f63474c4400000000000000000000000000000000000000000000000000000000602082015290565b336000908152600360209081526040808320600160a060020a038616845290915281205481610a3b828563ffffffff61125d16565b6000610e738383611300565b9392505050565b610e82610dd2565b1515610ed8576040805160e560020a62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600160a060020a0381161515610f38576040805160e560020a62461bcd02815260206004820181905260248201527f43616e6e6f7420726567697374657220746865206e756c6c2061646472657373604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383169081179091556040517f27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b90600090a250565b3315610fe5576040805160e560020a62461bcd02815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c00000000000000000000000000000000604482015290519081900360640190fd5b600254610ff8908263ffffffff61129f16565b60025550565b60005460ff1615611059576040805160e560020a62461bcd02815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a656400000000604482015290519081900360640190fd5b6000805460ff19166001178155600255611072336115dc565b61107b81610e7a565b50565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b6000806110b68686611300565b90507fe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc848460405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a195945050505050565b611129610dd2565b151561117f576040805160e560020a62461bcd02815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61107b816115dc565b600154604080517f467265657a65720000000000000000000000000000000000000000000000000060208083019190915282518083036007018152602783018085528151918301919091207fdcf0aaed00000000000000000000000000000000000000000000000000000000909152602b8301529151600093600160a060020a03169263dcf0aaed92604b8082019391829003018186803b15801561122c57600080fd5b505afa158015611240573d6000803e3d6000fd5b505050506040513d602081101561125657600080fd5b5051905090565b6000610e7383836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611699565b600082820183811015610e73576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b3390565b600061130a611188565b600160a060020a031663e5839836306040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b15801561136257600080fd5b505afa158015611376573d6000803e3d6000fd5b505050506040513d602081101561138c57600080fd5b5051156113cd5760405160e560020a62461bcd02815260040180806020018281038252602281526020018061175a6022913960400191505060405180910390fd5b600160a060020a03831615156114175760405160e560020a62461bcd02815260040180806020018281038252602a8152602001806117dd602a913960400191505060405180910390fd5b61142033610ce8565b8211156114615760405160e560020a62461bcd02815260040180806020018281038252602981526020018061177c6029913960400191505060405180910390fd5b600060fd815a6040805133602080830191909152600160a060020a038a168284015260608083018a905283518084039091018152608090920192839052815193949391929182918401908083835b602083106114ce5780518252601f1990920191602091820191016114af565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381858888f193505050503d8060008114611531576040519150601f19603f3d011682016040523d82523d6000602084013e611536565b606091505b5090915050801515611592576040805160e560020a62461bcd02815260206004820152601960248201527f43656c6f20476f6c64207472616e73666572206661696c656400000000000000604482015290519081900360640190fd5b604080518481529051600160a060020a0386169133917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019392505050565b600160a060020a03811615156116265760405160e560020a62461bcd0281526004018080602001828103825260268152602001806117346026913960400191505060405180910390fd5b60008054604051600160a060020a038085169361010090930416917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b6000818484111561172b5760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156116f05781810151838201526020016116d8565b50505050905090810190601f16801561171d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616e27742063616c6c207768656e20636f6e74726163742069732066726f7a656e7472616e736665722076616c75652065786365656465642062616c616e6365206f662073656e6465727472616e736665722076616c75652065786365656465642073656e646572277320616c6c6f77616e636520666f7220726563697069656e747472616e7366657220617474656d7074656420746f207265736572766564206164647265737320307830a165627a7a7230582055aac024660d1fe4df6b48dbbb972f0d6c9676d07ab5910086ae2fdd476b1c520029"

// DeployGoldToken deploys a new Ethereum contract, binding an instance of GoldToken to it.
func DeployGoldToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoldToken, error) {
	parsed, err := abi.JSON(strings.NewReader(GoldTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GoldTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoldToken{GoldTokenCaller: GoldTokenCaller{contract: contract}, GoldTokenTransactor: GoldTokenTransactor{contract: contract}, GoldTokenFilterer: GoldTokenFilterer{contract: contract}}, nil
}

// GoldToken is an auto generated Go binding around an Ethereum contract.
type GoldToken struct {
	GoldTokenCaller     // Read-only binding to the contract
	GoldTokenTransactor // Write-only binding to the contract
	GoldTokenFilterer   // Log filterer for contract events
}

// GoldTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoldTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoldTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoldTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoldTokenSession struct {
	Contract     *GoldToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoldTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoldTokenCallerSession struct {
	Contract *GoldTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GoldTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoldTokenTransactorSession struct {
	Contract     *GoldTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GoldTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoldTokenRaw struct {
	Contract *GoldToken // Generic contract binding to access the raw methods on
}

// GoldTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoldTokenCallerRaw struct {
	Contract *GoldTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GoldTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoldTokenTransactorRaw struct {
	Contract *GoldTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoldToken creates a new instance of GoldToken, bound to a specific deployed contract.
func NewGoldToken(address common.Address, backend bind.ContractBackend) (*GoldToken, error) {
	contract, err := bindGoldToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoldToken{GoldTokenCaller: GoldTokenCaller{contract: contract}, GoldTokenTransactor: GoldTokenTransactor{contract: contract}, GoldTokenFilterer: GoldTokenFilterer{contract: contract}}, nil
}

// NewGoldTokenCaller creates a new read-only instance of GoldToken, bound to a specific deployed contract.
func NewGoldTokenCaller(address common.Address, caller bind.ContractCaller) (*GoldTokenCaller, error) {
	contract, err := bindGoldToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoldTokenCaller{contract: contract}, nil
}

// NewGoldTokenTransactor creates a new write-only instance of GoldToken, bound to a specific deployed contract.
func NewGoldTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GoldTokenTransactor, error) {
	contract, err := bindGoldToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoldTokenTransactor{contract: contract}, nil
}

// NewGoldTokenFilterer creates a new log filterer instance of GoldToken, bound to a specific deployed contract.
func NewGoldTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GoldTokenFilterer, error) {
	contract, err := bindGoldToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoldTokenFilterer{contract: contract}, nil
}

// bindGoldToken binds a generic wrapper to an already deployed contract.
func bindGoldToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoldTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoldToken *GoldTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GoldToken.Contract.GoldTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoldToken *GoldTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoldToken.Contract.GoldTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoldToken *GoldTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoldToken.Contract.GoldTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoldToken *GoldTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GoldToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoldToken *GoldTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoldToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoldToken *GoldTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoldToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_GoldToken *GoldTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_GoldToken *GoldTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GoldToken.Contract.Allowance(&_GoldToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_GoldToken *GoldTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GoldToken.Contract.Allowance(&_GoldToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_GoldToken *GoldTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_GoldToken *GoldTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GoldToken.Contract.BalanceOf(&_GoldToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_GoldToken *GoldTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GoldToken.Contract.BalanceOf(&_GoldToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GoldToken *GoldTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GoldToken *GoldTokenSession) Decimals() (uint8, error) {
	return _GoldToken.Contract.Decimals(&_GoldToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GoldToken *GoldTokenCallerSession) Decimals() (uint8, error) {
	return _GoldToken.Contract.Decimals(&_GoldToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GoldToken *GoldTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GoldToken *GoldTokenSession) Initialized() (bool, error) {
	return _GoldToken.Contract.Initialized(&_GoldToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GoldToken *GoldTokenCallerSession) Initialized() (bool, error) {
	return _GoldToken.Contract.Initialized(&_GoldToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GoldToken *GoldTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GoldToken *GoldTokenSession) IsOwner() (bool, error) {
	return _GoldToken.Contract.IsOwner(&_GoldToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GoldToken *GoldTokenCallerSession) IsOwner() (bool, error) {
	return _GoldToken.Contract.IsOwner(&_GoldToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GoldToken *GoldTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GoldToken *GoldTokenSession) Name() (string, error) {
	return _GoldToken.Contract.Name(&_GoldToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GoldToken *GoldTokenCallerSession) Name() (string, error) {
	return _GoldToken.Contract.Name(&_GoldToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GoldToken *GoldTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GoldToken *GoldTokenSession) Owner() (common.Address, error) {
	return _GoldToken.Contract.Owner(&_GoldToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GoldToken *GoldTokenCallerSession) Owner() (common.Address, error) {
	return _GoldToken.Contract.Owner(&_GoldToken.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GoldToken *GoldTokenCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GoldToken *GoldTokenSession) Registry() (common.Address, error) {
	return _GoldToken.Contract.Registry(&_GoldToken.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GoldToken *GoldTokenCallerSession) Registry() (common.Address, error) {
	return _GoldToken.Contract.Registry(&_GoldToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GoldToken *GoldTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GoldToken *GoldTokenSession) Symbol() (string, error) {
	return _GoldToken.Contract.Symbol(&_GoldToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GoldToken *GoldTokenCallerSession) Symbol() (string, error) {
	return _GoldToken.Contract.Symbol(&_GoldToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GoldToken *GoldTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GoldToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GoldToken *GoldTokenSession) TotalSupply() (*big.Int, error) {
	return _GoldToken.Contract.TotalSupply(&_GoldToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_GoldToken *GoldTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GoldToken.Contract.TotalSupply(&_GoldToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Approve(&_GoldToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Approve(&_GoldToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "decreaseAllowance", spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.DecreaseAllowance(&_GoldToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.DecreaseAllowance(&_GoldToken.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "increaseAllowance", spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.IncreaseAllowance(&_GoldToken.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.IncreaseAllowance(&_GoldToken.TransactOpts, spender, value)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(uint256 amount) returns()
func (_GoldToken *GoldTokenTransactor) IncreaseSupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "increaseSupply", amount)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(uint256 amount) returns()
func (_GoldToken *GoldTokenSession) IncreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.IncreaseSupply(&_GoldToken.TransactOpts, amount)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(uint256 amount) returns()
func (_GoldToken *GoldTokenTransactorSession) IncreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.IncreaseSupply(&_GoldToken.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GoldToken *GoldTokenTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "initialize", registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GoldToken *GoldTokenSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.Initialize(&_GoldToken.TransactOpts, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_GoldToken *GoldTokenTransactorSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.Initialize(&_GoldToken.TransactOpts, registryAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Mint(&_GoldToken.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Mint(&_GoldToken.TransactOpts, to, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GoldToken *GoldTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GoldToken *GoldTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _GoldToken.Contract.RenounceOwnership(&_GoldToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GoldToken *GoldTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GoldToken.Contract.RenounceOwnership(&_GoldToken.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GoldToken *GoldTokenTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GoldToken *GoldTokenSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.SetRegistry(&_GoldToken.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GoldToken *GoldTokenTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.SetRegistry(&_GoldToken.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Transfer(&_GoldToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.Transfer(&_GoldToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferFrom(&_GoldToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferFrom(&_GoldToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoldToken *GoldTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoldToken *GoldTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferOwnership(&_GoldToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GoldToken *GoldTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferOwnership(&_GoldToken.TransactOpts, newOwner)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_GoldToken *GoldTokenTransactor) TransferWithComment(opts *bind.TransactOpts, to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _GoldToken.contract.Transact(opts, "transferWithComment", to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_GoldToken *GoldTokenSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferWithComment(&_GoldToken.TransactOpts, to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_GoldToken *GoldTokenTransactorSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _GoldToken.Contract.TransferWithComment(&_GoldToken.TransactOpts, to, value, comment)
}

// GoldTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GoldToken contract.
type GoldTokenApprovalIterator struct {
	Event *GoldTokenApproval // Event containing the contract specifics and raw log

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
func (it *GoldTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTokenApproval)
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
		it.Event = new(GoldTokenApproval)
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
func (it *GoldTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTokenApproval represents a Approval event raised by the GoldToken contract.
type GoldTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GoldToken *GoldTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GoldTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GoldToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GoldTokenApprovalIterator{contract: _GoldToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GoldToken *GoldTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GoldTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GoldToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTokenApproval)
				if err := _GoldToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GoldToken *GoldTokenFilterer) ParseApproval(log types.Log) (*GoldTokenApproval, error) {
	event := new(GoldTokenApproval)
	if err := _GoldToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoldTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GoldToken contract.
type GoldTokenOwnershipTransferredIterator struct {
	Event *GoldTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GoldTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTokenOwnershipTransferred)
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
		it.Event = new(GoldTokenOwnershipTransferred)
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
func (it *GoldTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTokenOwnershipTransferred represents a OwnershipTransferred event raised by the GoldToken contract.
type GoldTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GoldToken *GoldTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GoldTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GoldToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GoldTokenOwnershipTransferredIterator{contract: _GoldToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GoldToken *GoldTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GoldTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GoldToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTokenOwnershipTransferred)
				if err := _GoldToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GoldToken *GoldTokenFilterer) ParseOwnershipTransferred(log types.Log) (*GoldTokenOwnershipTransferred, error) {
	event := new(GoldTokenOwnershipTransferred)
	if err := _GoldToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoldTokenRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the GoldToken contract.
type GoldTokenRegistrySetIterator struct {
	Event *GoldTokenRegistrySet // Event containing the contract specifics and raw log

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
func (it *GoldTokenRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTokenRegistrySet)
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
		it.Event = new(GoldTokenRegistrySet)
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
func (it *GoldTokenRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTokenRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTokenRegistrySet represents a RegistrySet event raised by the GoldToken contract.
type GoldTokenRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GoldToken *GoldTokenFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*GoldTokenRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GoldToken.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &GoldTokenRegistrySetIterator{contract: _GoldToken.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GoldToken *GoldTokenFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *GoldTokenRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GoldToken.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTokenRegistrySet)
				if err := _GoldToken.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_GoldToken *GoldTokenFilterer) ParseRegistrySet(log types.Log) (*GoldTokenRegistrySet, error) {
	event := new(GoldTokenRegistrySet)
	if err := _GoldToken.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoldTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GoldToken contract.
type GoldTokenTransferIterator struct {
	Event *GoldTokenTransfer // Event containing the contract specifics and raw log

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
func (it *GoldTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTokenTransfer)
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
		it.Event = new(GoldTokenTransfer)
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
func (it *GoldTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTokenTransfer represents a Transfer event raised by the GoldToken contract.
type GoldTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GoldToken *GoldTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GoldTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GoldToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoldTokenTransferIterator{contract: _GoldToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GoldToken *GoldTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GoldTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GoldToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTokenTransfer)
				if err := _GoldToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GoldToken *GoldTokenFilterer) ParseTransfer(log types.Log) (*GoldTokenTransfer, error) {
	event := new(GoldTokenTransfer)
	if err := _GoldToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoldTokenTransferCommentIterator is returned from FilterTransferComment and is used to iterate over the raw logs and unpacked data for TransferComment events raised by the GoldToken contract.
type GoldTokenTransferCommentIterator struct {
	Event *GoldTokenTransferComment // Event containing the contract specifics and raw log

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
func (it *GoldTokenTransferCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTokenTransferComment)
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
		it.Event = new(GoldTokenTransferComment)
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
func (it *GoldTokenTransferCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTokenTransferCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTokenTransferComment represents a TransferComment event raised by the GoldToken contract.
type GoldTokenTransferComment struct {
	Comment string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferComment is a free log retrieval operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_GoldToken *GoldTokenFilterer) FilterTransferComment(opts *bind.FilterOpts) (*GoldTokenTransferCommentIterator, error) {

	logs, sub, err := _GoldToken.contract.FilterLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return &GoldTokenTransferCommentIterator{contract: _GoldToken.contract, event: "TransferComment", logs: logs, sub: sub}, nil
}

// WatchTransferComment is a free log subscription operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_GoldToken *GoldTokenFilterer) WatchTransferComment(opts *bind.WatchOpts, sink chan<- *GoldTokenTransferComment) (event.Subscription, error) {

	logs, sub, err := _GoldToken.contract.WatchLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTokenTransferComment)
				if err := _GoldToken.contract.UnpackLog(event, "TransferComment", log); err != nil {
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

// ParseTransferComment is a log parse operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_GoldToken *GoldTokenFilterer) ParseTransferComment(log types.Log) (*GoldTokenTransferComment, error) {
	event := new(GoldTokenTransferComment)
	if err := _GoldToken.contract.UnpackLog(event, "TransferComment", log); err != nil {
		return nil, err
	}
	return event, nil
}

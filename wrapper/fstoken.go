// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapper

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

// WrapperABI is the input ABI used to generate the binding from.
const WrapperABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// WrapperBin is the compiled bytecode used for deploying new contracts.
const WrapperBin = `60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040805190810160405280600581526020017f4649584544000000000000000000000000000000000000000000000000000000815250600290805190602001906200009f92919062000221565b506040805190810160405280601a81526020017f4578616d706c6520466978656420537570706c7920546f6b656e00000000000081525060039080519060200190620000ed92919062000221565b506012600460006101000a81548160ff021916908360ff160217905550600460009054906101000a900460ff1660ff16600a0a620f424002600581905550600554600660008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6005546040518082815260200191505060405180910390a3620002d0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026457805160ff191683800117855562000295565b8280016001018555821562000295579182015b828111156200029457825182559160200191906001019062000277565b5b509050620002a49190620002a8565b5090565b620002cd91905b80821115620002c9576000816000905550600101620002af565b5090565b90565b6114e380620002e06000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100e0578063095ea7b31461017057806318160ddd146101d557806323b872dd14610200578063313ce5671461028557806370a08231146102b657806379ba50971461030d5780638da5cb5b1461032457806395d89b411461037b578063a9059cbb1461040b578063cae9ca5114610470578063d4ee1d901461051b578063dc39d06d14610572578063dd62ed3e146105d7578063f2fde38b1461064e575b600080fd5b3480156100ec57600080fd5b506100f5610691565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013557808201518184015260208101905061011a565b50505050905090810190601f1680156101625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017c57600080fd5b506101bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061072f565b604051808215151515815260200191505060405180910390f35b3480156101e157600080fd5b506101ea610821565b6040518082815260200191505060405180910390f35b34801561020c57600080fd5b5061026b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061087c565b604051808215151515815260200191505060405180910390f35b34801561029157600080fd5b5061029a610b27565b604051808260ff1660ff16815260200191505060405180910390f35b3480156102c257600080fd5b506102f7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b3a565b6040518082815260200191505060405180910390f35b34801561031957600080fd5b50610322610b83565b005b34801561033057600080fd5b50610339610d22565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561038757600080fd5b50610390610d47565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103d05780820151818401526020810190506103b5565b50505050905090810190601f1680156103fd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561041757600080fd5b50610456600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610de5565b604051808215151515815260200191505060405180910390f35b34801561047c57600080fd5b50610501600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f80565b604051808215151515815260200191505060405180910390f35b34801561052757600080fd5b506105306111cf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561057e57600080fd5b506105bd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506111f5565b604051808215151515815260200191505060405180910390f35b3480156105e357600080fd5b50610638600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611359565b6040518082815260200191505060405180910390f35b34801561065a57600080fd5b5061068f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113e0565b005b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107275780601f106106fc57610100808354040283529160200191610727565b820191906000526020600020905b81548152906001019060200180831161070a57829003601f168201915b505050505081565b600081600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000610877600660008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460055461147f90919063ffffffff16565b905090565b60006108d082600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461147f90919063ffffffff16565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506109a282600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461147f90919063ffffffff16565b600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610a7482600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461149b90919063ffffffff16565b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600460009054906101000a900460ff1681565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bdf57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ddd5780601f10610db257610100808354040283529160200191610ddd565b820191906000526020600020905b815481529060010190602001808311610dc057829003601f168201915b505050505081565b6000610e3982600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461147f90919063ffffffff16565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ece82600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461149b90919063ffffffff16565b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b600082600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040518082815260200191505060405180910390a38373ffffffffffffffffffffffffffffffffffffffff16638f4ffcb1338530866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561115d578082015181840152602081019050611142565b50505050905090810190601f16801561118a5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156111ac57600080fd5b505af11580156111c0573d6000803e3d6000fd5b50505050600190509392505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561125257600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561131657600080fd5b505af115801561132a573d6000803e3d6000fd5b505050506040513d602081101561134057600080fd5b8101908080519060200190929190505050905092915050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561143b57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600082821115151561149057600080fd5b818303905092915050565b600081830190508281101515156114b157600080fd5b929150505600a165627a7a72305820d30c103f89fb94e558b42ee26a3ac2bd2499f2bb8864e4fa6919b389475e87270029`

// DeployWrapper deploys a new Ethereum contract, binding an instance of Wrapper to it.
func DeployWrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wrapper{WrapperCaller: WrapperCaller{contract: contract}, WrapperTransactor: WrapperTransactor{contract: contract}, WrapperFilterer: WrapperFilterer{contract: contract}}, nil
}

// Wrapper is an auto generated Go binding around an Ethereum contract.
type Wrapper struct {
	WrapperCaller     // Read-only binding to the contract
	WrapperTransactor // Write-only binding to the contract
	WrapperFilterer   // Log filterer for contract events
}

// WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapperSession struct {
	Contract     *Wrapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapperCallerSession struct {
	Contract *WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapperTransactorSession struct {
	Contract     *WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapperRaw struct {
	Contract *Wrapper // Generic contract binding to access the raw methods on
}

// WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapperCallerRaw struct {
	Contract *WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapperTransactorRaw struct {
	Contract *WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapper creates a new instance of Wrapper, bound to a specific deployed contract.
func NewWrapper(address common.Address, backend bind.ContractBackend) (*Wrapper, error) {
	contract, err := bindWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wrapper{WrapperCaller: WrapperCaller{contract: contract}, WrapperTransactor: WrapperTransactor{contract: contract}, WrapperFilterer: WrapperFilterer{contract: contract}}, nil
}

// NewWrapperCaller creates a new read-only instance of Wrapper, bound to a specific deployed contract.
func NewWrapperCaller(address common.Address, caller bind.ContractCaller) (*WrapperCaller, error) {
	contract, err := bindWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperCaller{contract: contract}, nil
}

// NewWrapperTransactor creates a new write-only instance of Wrapper, bound to a specific deployed contract.
func NewWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapperTransactor, error) {
	contract, err := bindWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperTransactor{contract: contract}, nil
}

// NewWrapperFilterer creates a new log filterer instance of Wrapper, bound to a specific deployed contract.
func NewWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapperFilterer, error) {
	contract, err := bindWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapperFilterer{contract: contract}, nil
}

// bindWrapper binds a generic wrapper to an already deployed contract.
func bindWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrapper *WrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wrapper.Contract.WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrapper *WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrapper.Contract.WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrapper *WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrapper.Contract.WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrapper *WrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrapper *WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrapper *WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrapper.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Wrapper *WrapperCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Wrapper *WrapperSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Wrapper.Contract.Allowance(&_Wrapper.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Wrapper *WrapperCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Wrapper.Contract.Allowance(&_Wrapper.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Wrapper *WrapperCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Wrapper *WrapperSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Wrapper.Contract.BalanceOf(&_Wrapper.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Wrapper *WrapperCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Wrapper.Contract.BalanceOf(&_Wrapper.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wrapper *WrapperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wrapper *WrapperSession) Decimals() (uint8, error) {
	return _Wrapper.Contract.Decimals(&_Wrapper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wrapper *WrapperCallerSession) Decimals() (uint8, error) {
	return _Wrapper.Contract.Decimals(&_Wrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wrapper *WrapperCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wrapper *WrapperSession) Name() (string, error) {
	return _Wrapper.Contract.Name(&_Wrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wrapper *WrapperCallerSession) Name() (string, error) {
	return _Wrapper.Contract.Name(&_Wrapper.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wrapper *WrapperCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wrapper *WrapperSession) NewOwner() (common.Address, error) {
	return _Wrapper.Contract.NewOwner(&_Wrapper.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Wrapper *WrapperCallerSession) NewOwner() (common.Address, error) {
	return _Wrapper.Contract.NewOwner(&_Wrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wrapper *WrapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wrapper *WrapperSession) Owner() (common.Address, error) {
	return _Wrapper.Contract.Owner(&_Wrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wrapper *WrapperCallerSession) Owner() (common.Address, error) {
	return _Wrapper.Contract.Owner(&_Wrapper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wrapper *WrapperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wrapper *WrapperSession) Symbol() (string, error) {
	return _Wrapper.Contract.Symbol(&_Wrapper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wrapper *WrapperCallerSession) Symbol() (string, error) {
	return _Wrapper.Contract.Symbol(&_Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wrapper *WrapperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wrapper.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wrapper *WrapperSession) TotalSupply() (*big.Int, error) {
	return _Wrapper.Contract.TotalSupply(&_Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wrapper *WrapperCallerSession) TotalSupply() (*big.Int, error) {
	return _Wrapper.Contract.TotalSupply(&_Wrapper.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wrapper *WrapperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wrapper *WrapperSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wrapper.Contract.AcceptOwnership(&_Wrapper.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wrapper *WrapperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wrapper.Contract.AcceptOwnership(&_Wrapper.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.Approve(&_Wrapper.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.Approve(&_Wrapper.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Wrapper *WrapperTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Wrapper *WrapperSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Wrapper.Contract.ApproveAndCall(&_Wrapper.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_Wrapper *WrapperTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Wrapper.Contract.ApproveAndCall(&_Wrapper.TransactOpts, spender, tokens, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.Transfer(&_Wrapper.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.Transfer(&_Wrapper.TransactOpts, to, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferAnyERC20Token(&_Wrapper.TransactOpts, tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferAnyERC20Token(&_Wrapper.TransactOpts, tokenAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferFrom(&_Wrapper.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Wrapper *WrapperTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferFrom(&_Wrapper.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Wrapper *WrapperTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Wrapper *WrapperSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferOwnership(&_Wrapper.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Wrapper *WrapperTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wrapper.Contract.TransferOwnership(&_Wrapper.TransactOpts, _newOwner)
}

// WrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wrapper contract.
type WrapperApprovalIterator struct {
	Event *WrapperApproval // Event containing the contract specifics and raw log

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
func (it *WrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperApproval)
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
		it.Event = new(WrapperApproval)
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
func (it *WrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperApproval represents a Approval event raised by the Wrapper contract.
type WrapperApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_Wrapper *WrapperFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*WrapperApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wrapper.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrapperApprovalIterator{contract: _Wrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_Wrapper *WrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrapperApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wrapper.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperApproval)
				if err := _Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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

// WrapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wrapper contract.
type WrapperOwnershipTransferredIterator struct {
	Event *WrapperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WrapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperOwnershipTransferred)
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
		it.Event = new(WrapperOwnershipTransferred)
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
func (it *WrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperOwnershipTransferred represents a OwnershipTransferred event raised by the Wrapper contract.
type WrapperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Wrapper *WrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*WrapperOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Wrapper.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperOwnershipTransferredIterator{contract: _Wrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_from indexed address, _to indexed address)
func (_Wrapper *WrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrapperOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Wrapper.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperOwnershipTransferred)
				if err := _Wrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// WrapperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wrapper contract.
type WrapperTransferIterator struct {
	Event *WrapperTransfer // Event containing the contract specifics and raw log

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
func (it *WrapperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperTransfer)
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
		it.Event = new(WrapperTransfer)
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
func (it *WrapperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrapperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrapperTransfer represents a Transfer event raised by the Wrapper contract.
type WrapperTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_Wrapper *WrapperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrapperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wrapper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperTransferIterator{contract: _Wrapper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokens uint256)
func (_Wrapper *WrapperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrapperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wrapper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrapperTransfer)
				if err := _Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
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

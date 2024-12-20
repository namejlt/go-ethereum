// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_banker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_workmaxcoin\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"banker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"costcoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"workearncoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161118938038061118983398181016040528101906100319190610150565b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050670de0b6b3a7640000816100ad91906101cd565b90508060c0818152505050505061020e565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100ec826100c3565b9050919050565b6100fc816100e2565b8114610106575f5ffd5b50565b5f81519050610117816100f3565b92915050565b5f819050919050565b61012f8161011d565b8114610139575f5ffd5b50565b5f8151905061014a81610126565b92915050565b5f5f5f60608486031215610167576101666100bf565b5b5f61017486828701610109565b935050602061018586828701610109565b92505060406101968682870161013c565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101d78261011d565b91506101e28361011d565b92508282026101f08161011d565b91508282048414831517610207576102066101a0565b5b5092915050565b60805160a05160c051610f2e61025b5f395f6106d001525f81816101d601528181610690015261076f01525f81816102ad015281816104f20152818161073301526108840152610f2e5ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c8063c9116b6911610064578063c9116b6914610130578063cb13cddb1461014e578063f1e1da601461017e578063f4770b881461019a578063fc0c546a146101b65761009c565b80630ab9db5b146100a05780632e1a7d4d146100be5780634c738909146100da578063a9059cbb146100f8578063b6b55f2514610114575b5f5ffd5b6100a86101d4565b6040516100b591906109ff565b60405180910390f35b6100d860048036038101906100d39190610a4f565b6101f8565b005b6100e261032a565b6040516100ef9190610a89565b60405180910390f35b610112600480360381019061010d9190610acc565b610380565b005b61012e60048036038101906101299190610a4f565b6104da565b005b610138610621565b6040516101459190610a89565b60405180910390f35b61016860048036038101906101639190610b0a565b610677565b6040516101759190610a89565b60405180910390f35b61019860048036038101906101939190610a4f565b61068b565b005b6101b460048036038101906101af9190610a4f565b6106b8565b005b6101be610882565b6040516101cb91906109ff565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b80670de0b6b3a76400008161020d9190610b62565b90505f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905080821115610292576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028990610c23565b60405180910390fd5b670de0b6b3a7640000836102a69190610b62565b92506102d37f000000000000000000000000000000000000000000000000000000000000000033856108a6565b825f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461031e9190610c41565b92505081905550505050565b5f670de0b6b3a76400005f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461037b9190610ca1565b905090565b80670de0b6b3a7640000816103959190610b62565b90505f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508082111561041a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041190610c23565b60405180910390fd5b670de0b6b3a76400008361042e9190610b62565b9250825f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461047b9190610c41565b92505081905550825f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104cd9190610cd1565b9250508190555050505050565b670de0b6b3a7640000816104ee9190610b62565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161054d93929190610d04565b6020604051808303815f875af1158015610569573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058d9190610d6e565b6105cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c390610de3565b60405180910390fd5b805f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106179190610cd1565b9250508190555050565b5f670de0b6b3a76400005f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546106729190610ca1565b905090565b5f602052805f5260405f205f915090505481565b6106b57f000000000000000000000000000000000000000000000000000000000000000082610380565b50565b670de0b6b3a7640000816106cc9190610b62565b90507f0000000000000000000000000000000000000000000000000000000000000000811115610731576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072890610e4b565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd7f000000000000000000000000000000000000000000000000000000000000000030846040518463ffffffff1660e01b81526004016107ae93929190610d04565b6020604051808303815f875af11580156107ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107ee9190610d6e565b61082d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082490610eb3565b60405180910390fd5b805f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108789190610cd1565b9250508190555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b610920838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040516024016108d9929190610ed1565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610925565b505050565b5f5f60205f8451602086015f885af180610944576040513d5f823e3d81fd5b3d92505f519150505f821461095d576001811415610978565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b156109ba57836040517f5274afe70000000000000000000000000000000000000000000000000000000081526004016109b191906109ff565b60405180910390fd5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109e9826109c0565b9050919050565b6109f9816109df565b82525050565b5f602082019050610a125f8301846109f0565b92915050565b5f5ffd5b5f819050919050565b610a2e81610a1c565b8114610a38575f5ffd5b50565b5f81359050610a4981610a25565b92915050565b5f60208284031215610a6457610a63610a18565b5b5f610a7184828501610a3b565b91505092915050565b610a8381610a1c565b82525050565b5f602082019050610a9c5f830184610a7a565b92915050565b610aab816109df565b8114610ab5575f5ffd5b50565b5f81359050610ac681610aa2565b92915050565b5f5f60408385031215610ae257610ae1610a18565b5b5f610aef85828601610ab8565b9250506020610b0085828601610a3b565b9150509250929050565b5f60208284031215610b1f57610b1e610a18565b5b5f610b2c84828501610ab8565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610b6c82610a1c565b9150610b7783610a1c565b9250828202610b8581610a1c565b91508282048414831517610b9c57610b9b610b35565b5b5092915050565b5f82825260208201905092915050565b7f74686520616d6f756e74206d6f7265207468616e2062616e6b206f662062616c5f8201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b5f610c0d602483610ba3565b9150610c1882610bb3565b604082019050919050565b5f6020820190508181035f830152610c3a81610c01565b9050919050565b5f610c4b82610a1c565b9150610c5683610a1c565b9250828203905081811115610c6e57610c6d610b35565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610cab82610a1c565b9150610cb683610a1c565b925082610cc657610cc5610c74565b5b828204905092915050565b5f610cdb82610a1c565b9150610ce683610a1c565b9250828201905080821115610cfe57610cfd610b35565b5b92915050565b5f606082019050610d175f8301866109f0565b610d2460208301856109f0565b610d316040830184610a7a565b949350505050565b5f8115159050919050565b610d4d81610d39565b8114610d57575f5ffd5b50565b5f81519050610d6881610d44565b92915050565b5f60208284031215610d8357610d82610a18565b5b5f610d9084828501610d5a565b91505092915050565b7f7472616e73666572206465706f736974206572726f72000000000000000000005f82015250565b5f610dcd601683610ba3565b9150610dd882610d99565b602082019050919050565b5f6020820190508181035f830152610dfa81610dc1565b9050919050565b7f74686520616d6f756e74206d6f7265207468616e20776f726b6d6178636f696e5f82015250565b5f610e35602083610ba3565b9150610e4082610e01565b602082019050919050565b5f6020820190508181035f830152610e6281610e29565b9050919050565b7f7472616e7366657220776f726b6561726e636f696e206572726f7200000000005f82015250565b5f610e9d601b83610ba3565b9150610ea882610e69565b602082019050919050565b5f6020820190508181035f830152610eca81610e91565b9050919050565b5f604082019050610ee45f8301856109f0565b610ef16020830184610a7a565b939250505056fea264697066735822122009cd07d6c40211dade2086e68f4bc48dc95845a4f7dcb789e23b316152d26ec564736f6c634300081c0033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _banker common.Address, _workmaxcoin *big.Int) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend, _token, _banker, _workmaxcoin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Bank *BankCaller) Banker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "banker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Bank *BankSession) Banker() (common.Address, error) {
	return _Bank.Contract.Banker(&_Bank.CallOpts)
}

// Banker is a free data retrieval call binding the contract method 0x0ab9db5b.
//
// Solidity: function banker() view returns(address)
func (_Bank *BankCallerSession) Banker() (common.Address, error) {
	return _Bank.Contract.Banker(&_Bank.CallOpts)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address ) view returns(uint256)
func (_Bank *BankCaller) Deposited(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "deposited", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address ) view returns(uint256)
func (_Bank *BankSession) Deposited(arg0 common.Address) (*big.Int, error) {
	return _Bank.Contract.Deposited(&_Bank.CallOpts, arg0)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address ) view returns(uint256)
func (_Bank *BankCallerSession) Deposited(arg0 common.Address) (*big.Int, error) {
	return _Bank.Contract.Deposited(&_Bank.CallOpts, arg0)
}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256 balance)
func (_Bank *BankCaller) GetMyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "getMyBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256 balance)
func (_Bank *BankSession) GetMyBalance() (*big.Int, error) {
	return _Bank.Contract.GetMyBalance(&_Bank.CallOpts)
}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256 balance)
func (_Bank *BankCallerSession) GetMyBalance() (*big.Int, error) {
	return _Bank.Contract.GetMyBalance(&_Bank.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Bank *BankCaller) MyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "myBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Bank *BankSession) MyBalance() (*big.Int, error) {
	return _Bank.Contract.MyBalance(&_Bank.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_Bank *BankCallerSession) MyBalance() (*big.Int, error) {
	return _Bank.Contract.MyBalance(&_Bank.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bank *BankCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bank *BankSession) Token() (common.Address, error) {
	return _Bank.Contract.Token(&_Bank.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bank *BankCallerSession) Token() (common.Address, error) {
	return _Bank.Contract.Token(&_Bank.CallOpts)
}

// Costcoin is a paid mutator transaction binding the contract method 0xf1e1da60.
//
// Solidity: function costcoin(uint256 amount) returns()
func (_Bank *BankTransactor) Costcoin(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "costcoin", amount)
}

// Costcoin is a paid mutator transaction binding the contract method 0xf1e1da60.
//
// Solidity: function costcoin(uint256 amount) returns()
func (_Bank *BankSession) Costcoin(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Costcoin(&_Bank.TransactOpts, amount)
}

// Costcoin is a paid mutator transaction binding the contract method 0xf1e1da60.
//
// Solidity: function costcoin(uint256 amount) returns()
func (_Bank *BankTransactorSession) Costcoin(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Costcoin(&_Bank.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bank *BankSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bank *BankTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Bank *BankTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Bank *BankSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Transfer(&_Bank.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Bank *BankTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Transfer(&_Bank.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bank *BankSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bank *BankTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts, amount)
}

// Workearncoin is a paid mutator transaction binding the contract method 0xf4770b88.
//
// Solidity: function workearncoin(uint256 amount) returns()
func (_Bank *BankTransactor) Workearncoin(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "workearncoin", amount)
}

// Workearncoin is a paid mutator transaction binding the contract method 0xf4770b88.
//
// Solidity: function workearncoin(uint256 amount) returns()
func (_Bank *BankSession) Workearncoin(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Workearncoin(&_Bank.TransactOpts, amount)
}

// Workearncoin is a paid mutator transaction binding the contract method 0xf4770b88.
//
// Solidity: function workearncoin(uint256 amount) returns()
func (_Bank *BankTransactorSession) Workearncoin(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Workearncoin(&_Bank.TransactOpts, amount)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

import (
	"math/big"
	"strings"

	"github.com/go-slifeio/accounts/abi"
	"github.com/go-slifeio/common"
	"github.com/go-slifeio/core/types"
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"records\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"age\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"updateAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_age\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TestBin is the compiled bytecode used for deploying new contracts.
const TestBin = `6060604052341561000f57600080fd5b604051610497380380610497833981016040528080519060200190919080518201919060200180519060200190919050506060604051908101604052808381526020018281526020018473ffffffffffffffffffffffffffffffffffffffff168152506000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906100cc929190610129565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050506101ce565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016a57805160ff1916838001178555610198565b82800160010185558215610198579182015b8281111561019757825182559160200191906001019061017c565b5b5090506101a591906101a9565b5090565b6101cb91905b808211156101c75760008160009055506001016101af565b5090565b90565b6102ba806101dd6000396000f30060606040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063469e906714610051578063a3ed51061461015a575b600080fd5b341561005c57600080fd5b610088600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061019c565b60405180806020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252858181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156101495780601f1061011e57610100808354040283529160200191610149565b820191906000526020600020905b81548152906001019060200180831161012c57829003601f168201915b505094505050505060405180910390f35b341561016557600080fd5b61019a600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506101e5565b005b600060205280600052604060002060009150905080600001908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff168160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610289578181600101819055505b5050505600a165627a7a72305820833cbae3cd045dab86e690c8f584a92354c770a86c0ce9ae8bcc14821e7cc9070029`

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address, _name string, _age *big.Int) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestBin), backend, _addr, _name, _age)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records( address) constant returns(name string, age uint256, addr address)
func (_Test *TestCaller) Records(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name string
	Age  *big.Int
	Addr common.Address
}, error) {
	ret := new(struct {
		Name string
		Age  *big.Int
		Addr common.Address
	})
	out := ret
	err := _Test.contract.Call(opts, out, "records", arg0)
	return *ret, err
}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records( address) constant returns(name string, age uint256, addr address)
func (_Test *TestSession) Records(arg0 common.Address) (struct {
	Name string
	Age  *big.Int
	Addr common.Address
}, error) {
	return _Test.Contract.Records(&_Test.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records( address) constant returns(name string, age uint256, addr address)
func (_Test *TestCallerSession) Records(arg0 common.Address) (struct {
	Name string
	Age  *big.Int
	Addr common.Address
}, error) {
	return _Test.Contract.Records(&_Test.CallOpts, arg0)
}

// UpdateAge is a paid mutator transaction binding the contract method 0xa3ed5106.
//
// Solidity: function updateAge(_addr address, _age uint256) returns()
func (_Test *TestTransactor) UpdateAge(opts *bind.TransactOpts, _addr common.Address, _age *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "updateAge", _addr, _age)
}

// UpdateAge is a paid mutator transaction binding the contract method 0xa3ed5106.
//
// Solidity: function updateAge(_addr address, _age uint256) returns()
func (_Test *TestSession) UpdateAge(_addr common.Address, _age *big.Int) (*types.Transaction, error) {
	return _Test.Contract.UpdateAge(&_Test.TransactOpts, _addr, _age)
}

// UpdateAge is a paid mutator transaction binding the contract method 0xa3ed5106.
//
// Solidity: function updateAge(_addr address, _age uint256) returns()
func (_Test *TestTransactorSession) UpdateAge(_addr common.Address, _age *big.Int) (*types.Transaction, error) {
	return _Test.Contract.UpdateAge(&_Test.TransactOpts, _addr, _age)
}

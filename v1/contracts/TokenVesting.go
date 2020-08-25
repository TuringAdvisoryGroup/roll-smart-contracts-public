// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TokenVestingABI is the input ABI used to generate the binding from.
const TokenVestingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WEEKS_IN_SECONDS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"vestedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"releaseableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_vestingBeneficiary\",\"type\":\"address\"},{\"name\":\"_vestingPeriodInWeeks\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseAddr\",\"outputs\":[{\"name\":\"parsed\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vestingInfo\",\"outputs\":[{\"name\":\"vestingBeneficiary\",\"type\":\"address\"},{\"name\":\"releasedSupply\",\"type\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vestingBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vestingBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vestingPeriodInWeeks\",\"type\":\"uint256\"}],\"name\":\"LogTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenVestingBin is the compiled bytecode used for deploying new contracts.
const TokenVestingBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111b8806100536000396000f3fe6080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806319165587146100b45780632fa1b76b146100dd578063384711cc1461010857806342e16161146101455780636daa985014610182578063893d20e8146101ab5780638da5cb5b146101d6578063b718f83a14610201578063f2fde38b1461023e578063f78e633d14610267578063fb897ce4146102a7575b600080fd5b3480156100c057600080fd5b506100db60048036036100d69190810190610df8565b6102e4565b005b3480156100e957600080fd5b506100f26105a3565b6040516100ff9190611097565b60405180910390f35b34801561011457600080fd5b5061012f600480360361012a9190810190610df8565b6105aa565b60405161013c9190611097565b60405180910390f35b34801561015157600080fd5b5061016c60048036036101679190810190610df8565b6107b1565b6040516101799190611097565b60405180910390f35b34801561018e57600080fd5b506101a960048036036101a49190810190610e21565b610817565b005b3480156101b757600080fd5b506101c06109bc565b6040516101cd9190610fd3565b60405180910390f35b3480156101e257600080fd5b506101eb6109e5565b6040516101f89190610fd3565b60405180910390f35b34801561020d57600080fd5b5061022860048036036102239190810190610e99565b610a0a565b6040516102359190610fd3565b60405180910390f35b34801561024a57600080fd5b5061026560048036036102609190810190610df8565b610a18565b005b34801561027357600080fd5b5061028e60048036036102899190810190610df8565b610b6d565b60405161029e9493929190611017565b60405180910390f35b3480156102b357600080fd5b506102ce60048036036102c99190810190610df8565b610bbd565b6040516102db919061107c565b60405180910390f35b60006102ef826107b1565b905060008111151561030057600080fd5b61035581600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154610c8b90919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060008273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610456929190610fee565b602060405180830381600087803b15801561047057600080fd5b505af1158015610484573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104a89190810190610e70565b90508015156104ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e39061105c565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff167f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051610596929190610fee565b60405180910390a2505050565b62093a8081565b60006105b4610d13565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020608060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905060008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016106cc9190610fd3565b60206040518083038186803b1580156106e457600080fd5b505afa1580156106f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061071c9190810190610eda565b90506000610737836020015183610c8b90919063ffffffff16565b905061075483606001518460400151610c8b90919063ffffffff16565b42101515610767578093505050506107ac565b6107a68360600151610798610789866040015142610ca990919063ffffffff16565b84610cc290919063ffffffff16565b610cfd90919063ffffffff16565b93505050505b919050565b6000610810600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154610802846105aa565b610ca990919063ffffffff16565b9050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561087257600080fd5b6080604051908101604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600081526020014281526020016108bd62093a8084610cc290919063ffffffff16565b815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301559050508273ffffffffffffffffffffffffffffffffffffffff167f83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a549083836040516109af929190610fee565b60405180910390a2505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060208201519050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a7357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610aaf57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b610bc5610d13565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020608060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820154815250509050919050565b6000808284019050838110151515610c9f57fe5b8091505092915050565b6000828211151515610cb757fe5b818303905092915050565b600080831415610cd55760009050610cf7565b60008284029050828482811515610ce857fe5b04141515610cf257fe5b809150505b92915050565b60008183811515610d0a57fe5b04905092915050565b608060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6000610d5e8235611147565b905092915050565b6000610d728251611159565b905092915050565b600082601f8301121515610d8d57600080fd5b8135610da0610d9b826110df565b6110b2565b91508082526020830160208301858383011115610dbc57600080fd5b610dc783828461116f565b50505092915050565b6000610ddc8235611165565b905092915050565b6000610df08251611165565b905092915050565b600060208284031215610e0a57600080fd5b6000610e1884828501610d52565b91505092915050565b600080600060608486031215610e3657600080fd5b6000610e4486828701610d52565b9350506020610e5586828701610d52565b9250506040610e6686828701610dd0565b9150509250925092565b600060208284031215610e8257600080fd5b6000610e9084828501610d66565b91505092915050565b600060208284031215610eab57600080fd5b600082013567ffffffffffffffff811115610ec557600080fd5b610ed184828501610d7a565b91505092915050565b600060208284031215610eec57600080fd5b6000610efa84828501610de4565b91505092915050565b610f0c8161110b565b82525050565b6000603382527f7472616e736665722066726f6d2076657374696e6720746f2062656e6566696360208301527f696172792068617320746f2073756363656564000000000000000000000000006040830152606082019050919050565b608082016000820151610f856000850182610f03565b506020820151610f986020850182610fc4565b506040820151610fab6040850182610fc4565b506060820151610fbe6060850182610fc4565b50505050565b610fcd8161113d565b82525050565b6000602082019050610fe86000830184610f03565b92915050565b60006040820190506110036000830185610f03565b6110106020830184610fc4565b9392505050565b600060808201905061102c6000830187610f03565b6110396020830186610fc4565b6110466040830185610fc4565b6110536060830184610fc4565b95945050505050565b6000602082019050818103600083015261107581610f12565b9050919050565b60006080820190506110916000830184610f6f565b92915050565b60006020820190506110ac6000830184610fc4565b92915050565b6000604051905081810181811067ffffffffffffffff821117156110d557600080fd5b8060405250919050565b600067ffffffffffffffff8211156110f657600080fd5b601f19601f8301169050602081019050919050565b60006111168261111d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006111528261111d565b9050919050565b60008115159050919050565b6000819050919050565b8281833760008383015250505056fea265627a7a72305820acd58e419e4ecec10457adf6ec6b37e7605dcbc80767a0a950b3df34a3199ec56c6578706572696d656e74616cf50037`

// DeployTokenVesting deploys a new Ethereum contract, binding an instance of TokenVesting to it.
func DeployTokenVesting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenVesting, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenVestingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// TokenVesting is an auto generated Go binding around an Ethereum contract.
type TokenVesting struct {
	TokenVestingCaller     // Read-only binding to the contract
	TokenVestingTransactor // Write-only binding to the contract
	TokenVestingFilterer   // Log filterer for contract events
}

// TokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVestingSession struct {
	Contract     *TokenVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVestingCallerSession struct {
	Contract *TokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVestingTransactorSession struct {
	Contract     *TokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVestingRaw struct {
	Contract *TokenVesting // Generic contract binding to access the raw methods on
}

// TokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVestingCallerRaw struct {
	Contract *TokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVestingTransactorRaw struct {
	Contract *TokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVesting creates a new instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVesting(address common.Address, backend bind.ContractBackend) (*TokenVesting, error) {
	contract, err := bindTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// NewTokenVestingCaller creates a new read-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*TokenVestingCaller, error) {
	contract, err := bindTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingCaller{contract: contract}, nil
}

// NewTokenVestingTransactor creates a new write-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVestingTransactor, error) {
	contract, err := bindTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingTransactor{contract: contract}, nil
}

// NewTokenVestingFilterer creates a new log filterer instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVestingFilterer, error) {
	contract, err := bindTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVestingFilterer{contract: contract}, nil
}

// bindTokenVesting binds a generic wrapper to an already deployed contract.
func bindTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.TokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transact(opts, method, params...)
}

// WEEKSINSECONDS is a free data retrieval call binding the contract method 0x2fa1b76b.
//
// Solidity: function WEEKS_IN_SECONDS() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) WEEKSINSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "WEEKS_IN_SECONDS")
	return *ret0, err
}

// WEEKSINSECONDS is a free data retrieval call binding the contract method 0x2fa1b76b.
//
// Solidity: function WEEKS_IN_SECONDS() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) WEEKSINSECONDS() (*big.Int, error) {
	return _TokenVesting.Contract.WEEKSINSECONDS(&_TokenVesting.CallOpts)
}

// WEEKSINSECONDS is a free data retrieval call binding the contract method 0x2fa1b76b.
//
// Solidity: function WEEKS_IN_SECONDS() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) WEEKSINSECONDS() (*big.Int, error) {
	return _TokenVesting.Contract.WEEKSINSECONDS(&_TokenVesting.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TokenVesting *TokenVestingCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TokenVesting *TokenVestingSession) GetOwner() (common.Address, error) {
	return _TokenVesting.Contract.GetOwner(&_TokenVesting.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) GetOwner() (common.Address, error) {
	return _TokenVesting.Contract.GetOwner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_TokenVesting *TokenVestingCaller) ParseAddr(opts *bind.CallOpts, data []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "parseAddr", data)
	return *ret0, err
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_TokenVesting *TokenVestingSession) ParseAddr(data []byte) (common.Address, error) {
	return _TokenVesting.Contract.ParseAddr(&_TokenVesting.CallOpts, data)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_TokenVesting *TokenVestingCallerSession) ParseAddr(data []byte) (common.Address, error) {
	return _TokenVesting.Contract.ParseAddr(&_TokenVesting.CallOpts, data)
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) ReleaseableAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "releaseableAmount", _token)
	return *ret0, err
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmount(&_TokenVesting.CallOpts, _token)
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmount(&_TokenVesting.CallOpts, _token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) VestedAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "vestedAmount", _token)
	return *ret0, err
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, _token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(_token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, _token)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo( address) constant returns(vestingBeneficiary address, releasedSupply uint256, start uint256, duration uint256)
func (_TokenVesting *TokenVestingCaller) VestingInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	ReleasedSupply     *big.Int
	Start              *big.Int
	Duration           *big.Int
}, error) {
	ret := new(struct {
		VestingBeneficiary common.Address
		ReleasedSupply     *big.Int
		Start              *big.Int
		Duration           *big.Int
	})
	out := ret
	err := _TokenVesting.contract.Call(opts, out, "vestingInfo", arg0)
	return *ret, err
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo( address) constant returns(vestingBeneficiary address, releasedSupply uint256, start uint256, duration uint256)
func (_TokenVesting *TokenVestingSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	ReleasedSupply     *big.Int
	Start              *big.Int
	Duration           *big.Int
}, error) {
	return _TokenVesting.Contract.VestingInfo(&_TokenVesting.CallOpts, arg0)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo( address) constant returns(vestingBeneficiary address, releasedSupply uint256, start uint256, duration uint256)
func (_TokenVesting *TokenVestingCallerSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	ReleasedSupply     *big.Int
	Start              *big.Int
	Duration           *big.Int
}, error) {
	return _TokenVesting.Contract.VestingInfo(&_TokenVesting.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x6daa9850.
//
// Solidity: function addToken(_token address, _vestingBeneficiary address, _vestingPeriodInWeeks uint256) returns()
func (_TokenVesting *TokenVestingTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _vestingBeneficiary common.Address, _vestingPeriodInWeeks *big.Int) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "addToken", _token, _vestingBeneficiary, _vestingPeriodInWeeks)
}

// AddToken is a paid mutator transaction binding the contract method 0x6daa9850.
//
// Solidity: function addToken(_token address, _vestingBeneficiary address, _vestingPeriodInWeeks uint256) returns()
func (_TokenVesting *TokenVestingSession) AddToken(_token common.Address, _vestingBeneficiary common.Address, _vestingPeriodInWeeks *big.Int) (*types.Transaction, error) {
	return _TokenVesting.Contract.AddToken(&_TokenVesting.TransactOpts, _token, _vestingBeneficiary, _vestingPeriodInWeeks)
}

// AddToken is a paid mutator transaction binding the contract method 0x6daa9850.
//
// Solidity: function addToken(_token address, _vestingBeneficiary address, _vestingPeriodInWeeks uint256) returns()
func (_TokenVesting *TokenVestingTransactorSession) AddToken(_token common.Address, _vestingBeneficiary common.Address, _vestingPeriodInWeeks *big.Int) (*types.Transaction, error) {
	return _TokenVesting.Contract.AddToken(&_TokenVesting.TransactOpts, _token, _vestingBeneficiary, _vestingPeriodInWeeks)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_token address) returns()
func (_TokenVesting *TokenVestingTransactor) Release(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "release", _token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_token address) returns()
func (_TokenVesting *TokenVestingSession) Release(_token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, _token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_token address) returns()
func (_TokenVesting *TokenVestingTransactorSession) Release(_token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, _token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TokenVestingLogTokenAddedIterator is returned from FilterLogTokenAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdded events raised by the TokenVesting contract.
type TokenVestingLogTokenAddedIterator struct {
	Event *TokenVestingLogTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenVestingLogTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingLogTokenAdded)
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
		it.Event = new(TokenVestingLogTokenAdded)
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
func (it *TokenVestingLogTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingLogTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingLogTokenAdded represents a LogTokenAdded event raised by the TokenVesting contract.
type TokenVestingLogTokenAdded struct {
	Token                common.Address
	VestingBeneficiary   common.Address
	VestingPeriodInWeeks *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdded is a free log retrieval operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: e LogTokenAdded(token indexed address, vestingBeneficiary address, vestingPeriodInWeeks uint256)
func (_TokenVesting *TokenVestingFilterer) FilterLogTokenAdded(opts *bind.FilterOpts, token []common.Address) (*TokenVestingLogTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingLogTokenAddedIterator{contract: _TokenVesting.contract, event: "LogTokenAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdded is a free log subscription operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: e LogTokenAdded(token indexed address, vestingBeneficiary address, vestingPeriodInWeeks uint256)
func (_TokenVesting *TokenVestingFilterer) WatchLogTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenVestingLogTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingLogTokenAdded)
				if err := _TokenVesting.contract.UnpackLog(event, "LogTokenAdded", log); err != nil {
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

// TokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenVesting contract.
type TokenVestingOwnershipTransferredIterator struct {
	Event *TokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipTransferred)
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
		it.Event = new(TokenVestingOwnershipTransferred)
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
func (it *TokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokenVesting contract.
type TokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipTransferredIterator{contract: _TokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipTransferred)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenVestingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the TokenVesting contract.
type TokenVestingReleasedIterator struct {
	Event *TokenVestingReleased // Event containing the contract specifics and raw log

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
func (it *TokenVestingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingReleased)
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
		it.Event = new(TokenVestingReleased)
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
func (it *TokenVestingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingReleased represents a Released event raised by the TokenVesting contract.
type TokenVestingReleased struct {
	Token              common.Address
	VestingBeneficiary common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: e Released(token indexed address, vestingBeneficiary address, amount uint256)
func (_TokenVesting *TokenVestingFilterer) FilterReleased(opts *bind.FilterOpts, token []common.Address) (*TokenVestingReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingReleasedIterator{contract: _TokenVesting.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: e Released(token indexed address, vestingBeneficiary address, amount uint256)
func (_TokenVesting *TokenVestingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *TokenVestingReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingReleased)
				if err := _TokenVesting.contract.UnpackLog(event, "Released", log); err != nil {
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

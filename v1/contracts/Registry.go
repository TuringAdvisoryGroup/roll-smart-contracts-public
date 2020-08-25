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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_hashIndex\",\"type\":\"bytes32\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"approveProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"rolodex\",\"outputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"vestingBeneficiary\",\"type\":\"address\"},{\"name\":\"initialPercentage\",\"type\":\"uint8\"},{\"name\":\"vestingPeriodInWeeks\",\"type\":\"uint256\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getIndexSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseAddr\",\"outputs\":[{\"name\":\"parsed\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getIndexByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_initialPercentage\",\"type\":\"uint8\"},{\"name\":\"_vestingPeriodInWeeks\",\"type\":\"uint256\"},{\"name\":\"_vestingBeneficiary\",\"type\":\"address\"},{\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"submitProposal\",\"outputs\":[{\"name\":\"hashIndex\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"hashIndex\",\"type\":\"bytes32\"}],\"name\":\"LogProposalSubmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"LogProposalApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611ed3806100536000396000f3fe6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806326c52914146100a9578063469dfdc1146100e6578063867954cf1461012c578063893d20e8146101695780638da5cb5b14610194578063969033a4146101bf578063b718f83a146101fc578063c0af858214610239578063d1a1ca0214610276578063f2fde38b146102b3575b600080fd5b3480156100b557600080fd5b506100d060048036036100cb919081019061176a565b6102dc565b6040516100dd9190611be4565b60405180910390f35b3480156100f257600080fd5b5061010d60048036036101089190810190611741565b6107f4565b6040516101239a99989796959493929190611b3a565b60405180910390f35b34801561013857600080fd5b50610153600480360361014e91908101906117e7565b6109ff565b6040516101609190611bff565b60405180910390f35b34801561017557600080fd5b5061017e610a74565b60405161018b9190611b1f565b60405180910390f35b3480156101a057600080fd5b506101a9610a9d565b6040516101b69190611b1f565b60405180910390f35b3480156101cb57600080fd5b506101e660048036036101e19190810190611741565b610ac2565b6040516101f39190611cc1565b60405180910390f35b34801561020857600080fd5b50610223600480360361021e91908101906117a6565b610da0565b6040516102309190611b1f565b60405180910390f35b34801561024557600080fd5b50610260600480360361025b91908101906117e7565b610dae565b60405161026d9190611bff565b60405180910390f35b34801561028257600080fd5b5061029d60048036036102989190810190611828565b610e23565b6040516102aa9190611bff565b60405180910390f35b3480156102bf57600080fd5b506102da60048036036102d59190810190611718565b611218565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561033957600080fd5b6103416114d9565b6001600085815260200190815260200160002061014060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104505780601f1061042557610100808354040283529160200191610450565b820191906000526020600020905b81548152906001019060200180831161043357829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104f25780601f106104c7576101008083540402835291602001916104f2565b820191906000526020600020905b8154815290600101906020018083116104d557829003601f168201915b505050505081526020016003820160009054906101000a900460ff1660ff1660ff168152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160149054906101000a900460ff1660ff1660ff168152602001600782015481526020016008820160009054906101000a900460ff161515151581525050905061061f816020015161136d565b61062c8160400151611423565b826001600086815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180600086815260200190815260200160002060080160006101000a81548160ff02191690831515021790555083600282602001516040518082805190602001908083835b6020831015156106ec57805182526020820191506020810190506020830392506106c7565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208190555083600382604001516040518082805190602001908083835b602083101515610760578051825260208201915060208101905060208303925061073b565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055508273ffffffffffffffffffffffffffffffffffffffff167f19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa36882602001516040516107e19190611c1a565b60405180910390a2600191505092915050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108c65780601f1061089b576101008083540402835291602001916108c6565b820191906000526020600020905b8154815290600101906020018083116108a957829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109645780601f1061093957610100808354040283529160200191610964565b820191906000526020600020905b81548152906001019060200180831161094757829003601f168201915b5050505050908060030160009054906101000a900460ff16908060040154908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160149054906101000a900460ff16908060070154908060080160009054906101000a900460ff1690508a565b60006003826040518082805190602001908083835b602083101515610a395780518252602082019150602081019050602083039250610a14565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610aca6114d9565b6001600083815260200190815260200160002061014060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bd95780601f10610bae57610100808354040283529160200191610bd9565b820191906000526020600020905b815481529060010190602001808311610bbc57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c7b5780601f10610c5057610100808354040283529160200191610c7b565b820191906000526020600020905b815481529060010190602001808311610c5e57829003601f168201915b505050505081526020016003820160009054906101000a900460ff1660ff1660ff168152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160149054906101000a900460ff1660ff1660ff168152602001600782015481526020016008820160009054906101000a900460ff1615151515815250509050919050565b600060208201519050919050565b60006002826040518082805190602001908083835b602083101515610de85780518252602082019150602081019050602083039250610dc3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e8057600080fd5b610e898961136d565b610e9288611423565b8888836040516020018084805190602001908083835b602083101515610ecd5780518252602082019150602081019050602083039250610ea8565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b602083101515610f205780518252602082019150602081019050602083039250610efb565b6001836020036101000a0380198251168184511680821785525050505050509050018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401935050505060405160208183030381529060405280519060200120905061014060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018981526020018860ff1681526020018781526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018660ff168152602001858152602001600015158152506001600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190805190602001906110a9929190611577565b5060408201518160020190805190602001906110c6929190611577565b5060608201518160030160006101000a81548160ff021916908360ff1602179055506080820151816004015560a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e08201518160060160146101000a81548160ff021916908360ff16021790555061010082015181600701556101208201518160080160006101000a81548160ff021916908315150217905550905050807fc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c18a8a3360405161120493929190611c3c565b60405180910390a298975050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561127357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156112af57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006001026002826040518082805190602001908083835b6020831015156113aa5780518252602082019150602081019050602083039250611385565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054141515611420576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141790611c81565b60405180910390fd5b50565b60006001026003826040518082805190602001908083835b602083101515611460578051825260208201915060208101905060208303925061143b565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020541415156114d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114cd90611ca1565b60405180910390fd5b50565b61014060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600060ff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115b857805160ff19168380011785556115e6565b828001600101855582156115e6579182015b828111156115e55782518255916020019190600101906115ca565b5b5090506115f391906115f7565b5090565b61161991905b808211156116155760008160009055506001016115fd565b5090565b90565b60006116288235611ddd565b905092915050565b600061163c8235611def565b905092915050565b600082601f830112151561165757600080fd5b813561166a61166582611d10565b611ce3565b9150808252602083016020830185838301111561168657600080fd5b611691838284611e46565b50505092915050565b600082601f83011215156116ad57600080fd5b81356116c06116bb82611d3c565b611ce3565b915080825260208301602083018583830111156116dc57600080fd5b6116e7838284611e46565b50505092915050565b60006116fc8235611df9565b905092915050565b60006117108235611e03565b905092915050565b60006020828403121561172a57600080fd5b60006117388482850161161c565b91505092915050565b60006020828403121561175357600080fd5b600061176184828501611630565b91505092915050565b6000806040838503121561177d57600080fd5b600061178b85828601611630565b925050602061179c8582860161161c565b9150509250929050565b6000602082840312156117b857600080fd5b600082013567ffffffffffffffff8111156117d257600080fd5b6117de84828501611644565b91505092915050565b6000602082840312156117f957600080fd5b600082013567ffffffffffffffff81111561181357600080fd5b61181f8482850161169a565b91505092915050565b600080600080600080600080610100898b03121561184557600080fd5b600089013567ffffffffffffffff81111561185f57600080fd5b61186b8b828c0161169a565b985050602089013567ffffffffffffffff81111561188857600080fd5b6118948b828c0161169a565b97505060406118a58b828c01611704565b96505060606118b68b828c016116f0565b95505060806118c78b828c01611704565b94505060a06118d88b828c016116f0565b93505060c06118e98b828c0161161c565b92505060e06118fa8b828c0161161c565b9150509295985092959890939650565b61191381611e10565b82525050565b61192281611d7e565b82525050565b61193181611d90565b82525050565b61194081611d9c565b82525050565b600061195182611d73565b808452611965816020860160208601611e55565b61196e81611e88565b602085010191505092915050565b600061198782611d68565b80845261199b816020860160208601611e55565b6119a481611e88565b602085010191505092915050565b6000601382527f4e616d6520616c726561647920657869737473000000000000000000000000006020830152604082019050919050565b6000601582527f53796d626f6c20616c72656164792065786973747300000000000000000000006020830152604082019050919050565b600061014083016000830151611a396000860182611919565b5060208301518482036020860152611a51828261197c565b91505060408301518482036040860152611a6b828261197c565b9150506060830151611a806060860182611b10565b506080830151611a936080860182611b01565b5060a0830151611aa660a0860182611919565b5060c0830151611ab960c0860182611919565b5060e0830151611acc60e0860182611b10565b50610100830151611ae1610100860182611b01565b50610120830151611af6610120860182611928565b508091505092915050565b611b0a81611dc6565b82525050565b611b1981611dd0565b82525050565b6000602082019050611b346000830184611919565b92915050565b600061014082019050611b50600083018d611919565b8181036020830152611b62818c61197c565b90508181036040830152611b76818b61197c565b9050611b85606083018a611b10565b611b926080830189611b01565b611b9f60a0830188611919565b611bac60c0830187611919565b611bb960e0830186611b10565b611bc7610100830185611b01565b611bd5610120830184611928565b9b9a5050505050505050505050565b6000602082019050611bf96000830184611928565b92915050565b6000602082019050611c146000830184611937565b92915050565b60006020820190508181036000830152611c34818461197c565b905092915050565b60006060820190508181036000830152611c568186611946565b90508181036020830152611c6a8185611946565b9050611c79604083018461190a565b949350505050565b60006020820190508181036000830152611c9a816119b2565b9050919050565b60006020820190508181036000830152611cba816119e9565b9050919050565b60006020820190508181036000830152611cdb8184611a20565b905092915050565b6000604051905081810181811067ffffffffffffffff82111715611d0657600080fd5b8060405250919050565b600067ffffffffffffffff821115611d2757600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115611d5357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b6000611d8982611da6565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611de882611da6565b9050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b6000611e1b82611e22565b9050919050565b6000611e2d82611e34565b9050919050565b6000611e3f82611da6565b9050919050565b82818337600083830152505050565b60005b83811015611e73578082015181840152602081019050611e58565b83811115611e82576000848401525b50505050565b6000601f19601f830116905091905056fea265627a7a72305820a6c2b8cd86e97e0256830c72bedafd465ef6a9a9e08f64ef248f8615f0aeb45c6c6578706572696d656e74616cf50037`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(_name string) constant returns(bytes32)
func (_Registry *RegistryCaller) GetIndexByName(opts *bind.CallOpts, _name string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getIndexByName", _name)
	return *ret0, err
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(_name string) constant returns(bytes32)
func (_Registry *RegistrySession) GetIndexByName(_name string) ([32]byte, error) {
	return _Registry.Contract.GetIndexByName(&_Registry.CallOpts, _name)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(_name string) constant returns(bytes32)
func (_Registry *RegistryCallerSession) GetIndexByName(_name string) ([32]byte, error) {
	return _Registry.Contract.GetIndexByName(&_Registry.CallOpts, _name)
}

// GetIndexSymbol is a free data retrieval call binding the contract method 0x867954cf.
//
// Solidity: function getIndexSymbol(_symbol string) constant returns(bytes32)
func (_Registry *RegistryCaller) GetIndexSymbol(opts *bind.CallOpts, _symbol string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getIndexSymbol", _symbol)
	return *ret0, err
}

// GetIndexSymbol is a free data retrieval call binding the contract method 0x867954cf.
//
// Solidity: function getIndexSymbol(_symbol string) constant returns(bytes32)
func (_Registry *RegistrySession) GetIndexSymbol(_symbol string) ([32]byte, error) {
	return _Registry.Contract.GetIndexSymbol(&_Registry.CallOpts, _symbol)
}

// GetIndexSymbol is a free data retrieval call binding the contract method 0x867954cf.
//
// Solidity: function getIndexSymbol(_symbol string) constant returns(bytes32)
func (_Registry *RegistryCallerSession) GetIndexSymbol(_symbol string) ([32]byte, error) {
	return _Registry.Contract.GetIndexSymbol(&_Registry.CallOpts, _symbol)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Registry *RegistryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Registry *RegistrySession) GetOwner() (common.Address, error) {
	return _Registry.Contract.GetOwner(&_Registry.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Registry *RegistryCallerSession) GetOwner() (common.Address, error) {
	return _Registry.Contract.GetOwner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Registry *RegistryCaller) ParseAddr(opts *bind.CallOpts, data []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "parseAddr", data)
	return *ret0, err
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Registry *RegistrySession) ParseAddr(data []byte) (common.Address, error) {
	return _Registry.Contract.ParseAddr(&_Registry.CallOpts, data)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Registry *RegistryCallerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Registry.Contract.ParseAddr(&_Registry.CallOpts, data)
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex( bytes32) constant returns(token address, name string, symbol string, decimals uint8, totalSupply uint256, proposer address, vestingBeneficiary address, initialPercentage uint8, vestingPeriodInWeeks uint256, approved bool)
func (_Registry *RegistryCaller) Rolodex(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Token                common.Address
	Name                 string
	Symbol               string
	Decimals             uint8
	TotalSupply          *big.Int
	Proposer             common.Address
	VestingBeneficiary   common.Address
	InitialPercentage    uint8
	VestingPeriodInWeeks *big.Int
	Approved             bool
}, error) {
	ret := new(struct {
		Token                common.Address
		Name                 string
		Symbol               string
		Decimals             uint8
		TotalSupply          *big.Int
		Proposer             common.Address
		VestingBeneficiary   common.Address
		InitialPercentage    uint8
		VestingPeriodInWeeks *big.Int
		Approved             bool
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "rolodex", arg0)
	return *ret, err
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex( bytes32) constant returns(token address, name string, symbol string, decimals uint8, totalSupply uint256, proposer address, vestingBeneficiary address, initialPercentage uint8, vestingPeriodInWeeks uint256, approved bool)
func (_Registry *RegistrySession) Rolodex(arg0 [32]byte) (struct {
	Token                common.Address
	Name                 string
	Symbol               string
	Decimals             uint8
	TotalSupply          *big.Int
	Proposer             common.Address
	VestingBeneficiary   common.Address
	InitialPercentage    uint8
	VestingPeriodInWeeks *big.Int
	Approved             bool
}, error) {
	return _Registry.Contract.Rolodex(&_Registry.CallOpts, arg0)
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex( bytes32) constant returns(token address, name string, symbol string, decimals uint8, totalSupply uint256, proposer address, vestingBeneficiary address, initialPercentage uint8, vestingPeriodInWeeks uint256, approved bool)
func (_Registry *RegistryCallerSession) Rolodex(arg0 [32]byte) (struct {
	Token                common.Address
	Name                 string
	Symbol               string
	Decimals             uint8
	TotalSupply          *big.Int
	Proposer             common.Address
	VestingBeneficiary   common.Address
	InitialPercentage    uint8
	VestingPeriodInWeeks *big.Int
	Approved             bool
}, error) {
	return _Registry.Contract.Rolodex(&_Registry.CallOpts, arg0)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(_hashIndex bytes32, _token address) returns(bool)
func (_Registry *RegistryTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "approveProposal", _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(_hashIndex bytes32, _token address) returns(bool)
func (_Registry *RegistrySession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ApproveProposal(&_Registry.TransactOpts, _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(_hashIndex bytes32, _token address) returns(bool)
func (_Registry *RegistryTransactorSession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ApproveProposal(&_Registry.TransactOpts, _hashIndex, _token)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd1a1ca02.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address, _proposer address) returns(hashIndex bytes32)
func (_Registry *RegistryTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address, _proposer common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary, _proposer)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd1a1ca02.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address, _proposer address) returns(hashIndex bytes32)
func (_Registry *RegistrySession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address, _proposer common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposal(&_Registry.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary, _proposer)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd1a1ca02.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address, _proposer address) returns(hashIndex bytes32)
func (_Registry *RegistryTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address, _proposer common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposal(&_Registry.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary, _proposer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryLogProposalApproveIterator is returned from FilterLogProposalApprove and is used to iterate over the raw logs and unpacked data for LogProposalApprove events raised by the Registry contract.
type RegistryLogProposalApproveIterator struct {
	Event *RegistryLogProposalApprove // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalApprove)
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
		it.Event = new(RegistryLogProposalApprove)
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
func (it *RegistryLogProposalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalApprove represents a LogProposalApprove event raised by the Registry contract.
type RegistryLogProposalApprove struct {
	Name         string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogProposalApprove is a free log retrieval operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: e LogProposalApprove(name string, tokenAddress indexed address)
func (_Registry *RegistryFilterer) FilterLogProposalApprove(opts *bind.FilterOpts, tokenAddress []common.Address) (*RegistryLogProposalApproveIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalApproveIterator{contract: _Registry.contract, event: "LogProposalApprove", logs: logs, sub: sub}, nil
}

// WatchLogProposalApprove is a free log subscription operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: e LogProposalApprove(name string, tokenAddress indexed address)
func (_Registry *RegistryFilterer) WatchLogProposalApprove(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalApprove, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalApprove)
				if err := _Registry.contract.UnpackLog(event, "LogProposalApprove", log); err != nil {
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

// RegistryLogProposalSubmitIterator is returned from FilterLogProposalSubmit and is used to iterate over the raw logs and unpacked data for LogProposalSubmit events raised by the Registry contract.
type RegistryLogProposalSubmitIterator struct {
	Event *RegistryLogProposalSubmit // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalSubmit)
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
		it.Event = new(RegistryLogProposalSubmit)
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
func (it *RegistryLogProposalSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalSubmit represents a LogProposalSubmit event raised by the Registry contract.
type RegistryLogProposalSubmit struct {
	Name      string
	Symbol    string
	Proposer  common.Address
	HashIndex [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProposalSubmit is a free log retrieval operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: e LogProposalSubmit(name string, symbol string, proposer address, hashIndex indexed bytes32)
func (_Registry *RegistryFilterer) FilterLogProposalSubmit(opts *bind.FilterOpts, hashIndex [][32]byte) (*RegistryLogProposalSubmitIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalSubmitIterator{contract: _Registry.contract, event: "LogProposalSubmit", logs: logs, sub: sub}, nil
}

// WatchLogProposalSubmit is a free log subscription operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: e LogProposalSubmit(name string, symbol string, proposer address, hashIndex indexed bytes32)
func (_Registry *RegistryFilterer) WatchLogProposalSubmit(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalSubmit, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalSubmit)
				if err := _Registry.contract.UnpackLog(event, "LogProposalSubmit", log); err != nil {
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

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

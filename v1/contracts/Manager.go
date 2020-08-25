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

// ManagerABI is the input ABI used to generate the binding from.
const ManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTokenFactory\",\"type\":\"address\"}],\"name\":\"setTokenFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTokenVesting\",\"type\":\"address\"}],\"name\":\"setTokenVesting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPlatformWallet\",\"type\":\"address\"}],\"name\":\"setPlatformWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_initialPercentage\",\"type\":\"uint8\"},{\"name\":\"_vestingPeriodInWeeks\",\"type\":\"uint256\"},{\"name\":\"_vestingBeneficiary\",\"type\":\"address\"}],\"name\":\"submitProposal\",\"outputs\":[{\"name\":\"hashIndex\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRegistry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"parseAddr\",\"outputs\":[{\"name\":\"parsed\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RegistryInstance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPercentage\",\"type\":\"uint8\"}],\"name\":\"setTokenFactoryPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenFactoryInstance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"migrateManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashIndex\",\"type\":\"bytes32\"}],\"name\":\"approveProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldTF\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newTF\",\"type\":\"address\"}],\"name\":\"LogTokenFactoryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldR\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newR\",\"type\":\"address\"}],\"name\":\"LogRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"LogManagerMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ManagerBin is the compiled bytecode used for deploying new contracts.
const ManagerBin = `60806040523480156200001157600080fd5b5060405160408062002f6f83398101806040526200003391908101906200018a565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015620000de5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b1515620000ea57600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001ff565b6000620001828251620001eb565b905092915050565b600080604083850312156200019e57600080fd5b6000620001ae8582860162000174565b9250506020620001c18582860162000174565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001f882620001cb565b9050919050565b612d60806200020f6000396000f3fe6080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146100e05780632f73a9f81461010b57806343e34696146101345780638831e9cf1461015d578063893d20e8146101865780638da5cb5b146101b1578063a210853c146101dc578063a91ee0dc14610219578063b718f83a14610242578063b75ceb711461027f578063b9ae5d2c146102aa578063bbba6e92146102d3578063c7a8db19146102fe578063f20b5b2c14610327578063f2fde38b14610350575b600080fd5b3480156100ec57600080fd5b506100f5610379565b60405161010291906127c5565b60405180910390f35b34801561011757600080fd5b50610132600480360361012d91908101906120ed565b610709565b005b34801561014057600080fd5b5061015b600480360361015691908101906120ed565b610ad7565b005b34801561016957600080fd5b50610184600480360361017f91908101906120ed565b610bde565b005b34801561019257600080fd5b5061019b610d2e565b6040516101a89190612781565b60405180910390f35b3480156101bd57600080fd5b506101c6610d57565b6040516101d39190612781565b60405180910390f35b3480156101e857600080fd5b5061020360048036036101fe9190810190612224565b610d7c565b60405161021091906127e0565b60405180910390f35b34801561022557600080fd5b50610240600480360361023b91908101906120ed565b610fa4565b005b34801561024e57600080fd5b50610269600480360361026491908101906121ba565b611116565b6040516102769190612781565b60405180910390f35b34801561028b57600080fd5b50610294611124565b6040516102a19190612824565b60405180910390f35b3480156102b657600080fd5b506102d160048036036102cc9190810190612333565b61114a565b005b3480156102df57600080fd5b506102e86112f3565b6040516102f5919061283f565b60405180910390f35b34801561030a57600080fd5b50610325600480360361032091908101906120ed565b611319565b005b34801561033357600080fd5b5061034e60048036036103499190810190612168565b611555565b005b34801561035c57600080fd5b50610377600480360361037291908101906120ed565b6118dc565b005b60003073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561041657600080fd5b505afa15801561042a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061044e9190810190612116565b73ffffffffffffffffffffffffffffffffffffffff1614801561055757503073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561050757600080fd5b505afa15801561051b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061053f9190810190612116565b73ffffffffffffffffffffffffffffffffffffffff16145b80156107045750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663151ad6106040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561061b57600080fd5b505afa15801561062f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061065391908101906121fb565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156106b457600080fd5b505afa1580156106c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506106ec9190810190612116565b73ffffffffffffffffffffffffffffffffffffffff16145b905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561076457600080fd5b3073ffffffffffffffffffffffffffffffffffffffff1661078482611a31565b73ffffffffffffffffffffffffffffffffffffffff161415156107dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d390612a03565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663151ad6106040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561086057600080fd5b505afa158015610874573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061089891908101906121fb565b73ffffffffffffffffffffffffffffffffffffffff166108b782611bca565b73ffffffffffffffffffffffffffffffffffffffff1614151561090f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090690612aa3565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166357796797826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016109869190612781565b600060405180830381600087803b1580156109a057600080fd5b505af11580156109b4573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff166109e06109db83611bca565b611a31565b73ffffffffffffffffffffffffffffffffffffffff16141515610a38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2f90612a83565b60405180910390fd5b7f4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051610a8b92919061279c565b60405180910390a180600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b3257600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166343e34696826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610ba99190612781565b600060405180830381600087803b158015610bc357600080fd5b505af1158015610bd7573d6000803e3d6000fd5b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c3957600080fd5b610c41610379565b1515610c82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c79906129e3565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638831e9cf826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610cf99190612781565b600060405180830381600087803b158015610d1357600080fd5b505af1158015610d27573d6000803e3d6000fd5b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008360008160ff16118015610d95575060648160ff16105b1515610dd6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcd90612a23565b60405180910390fd5b8660008160ff16118015610ded575060128160ff16105b1515610e2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e25906129a3565b60405180910390fd5b84600181111515610e74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6b90612a63565b60405180910390fd5b610e7c610379565b1515610ebd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb4906129e3565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1a1ca028c8c8c8c8c8c8c336040518963ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610f4298979695949392919061285a565b602060405180830381600087803b158015610f5c57600080fd5b505af1158015610f70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f949190810190612191565b9350505050979650505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fff57600080fd5b3073ffffffffffffffffffffffffffffffffffffffff1661101f82611a31565b73ffffffffffffffffffffffffffffffffffffffff16141515611077576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106e906129c3565b60405180910390fd5b7fec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826040516110ca92919061279c565b60405180910390a180600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060208201519050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156111a557600080fd5b8060008160ff161180156111bc575060648160ff16105b15156111fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f490612a23565b60405180910390fd5b611205610379565b1515611246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123d906129e3565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663961f3c01836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016112bd9190612ac3565b600060405180830381600087803b1580156112d757600080fd5b505af11580156112eb573d6000803e3d6000fd5b505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561137457600080fd5b61137c610379565b15156113bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b4906129e3565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2fde38b826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016114349190612781565b600060405180830381600087803b15801561144e57600080fd5b505af1158015611462573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2fde38b826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016114dd9190612781565b600060405180830381600087803b1580156114f757600080fd5b505af115801561150b573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167f829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be17460405160405180910390a250565b61155d610379565b151561159e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611595906129e3565b60405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115f957600080fd5b611601611d63565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663969033a4836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161167891906127e0565b60006040518083038186803b15801561169057600080fd5b505afa1580156116a4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506116cd91908101906122f2565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663747ef42583602001518460400151856060015186608001518760e001518861010001518960c001516040518863ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161177197969594939291906128e6565b602060405180830381600087803b15801561178b57600080fd5b505af115801561179f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117c39190810190612116565b90506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166326c5291485846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016118409291906127fb565b602060405180830381600087803b15801561185a57600080fd5b505af115801561186e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611892919081019061213f565b90508015156118d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118cd90612a43565b60405180910390fd5b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561193757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561197357600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060606040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600060608473ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083101515611b0b5780518252602082019150602081019050602083039250611ae6565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114611b6b576040519150601f19603f3d011682016040523d82523d6000602084013e611b70565b606091505b5091509150811515611bb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bae90612963565b60405180910390fd5b611bc081611116565b9350505050919050565b600060606040516024016040516020818303038152906040527f72172364000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600060608473ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083101515611ca45780518252602082019150602081019050602083039250611c7f565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114611d04576040519150601f19603f3d011682016040523d82523d6000602084013e611d09565b606091505b5091509150811515611d50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d4790612983565b60405180910390fd5b611d5981611116565b9350505050919050565b61014060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600060ff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600081526020016000151581525090565b6000611e0d8235612c04565b905092915050565b6000611e218251612c04565b905092915050565b6000611e358251612c16565b905092915050565b6000611e498235612c22565b905092915050565b6000611e5d8251612c22565b905092915050565b600082601f8301121515611e7857600080fd5b8135611e8b611e8682612b0b565b612ade565b91508082526020830160208301858383011115611ea757600080fd5b611eb2838284612cd3565b50505092915050565b6000611ec78251612c2c565b905092915050565b600082601f8301121515611ee257600080fd5b8151611ef5611ef082612b37565b612ade565b91508082526020830160208301858383011115611f1157600080fd5b611f1c838284612ce2565b50505092915050565b600082601f8301121515611f3857600080fd5b8135611f4b611f4682612b63565b612ade565b91508082526020830160208301858383011115611f6757600080fd5b611f72838284612cd3565b50505092915050565b60006101408284031215611f8e57600080fd5b611f99610140612ade565b90506000611fa984828501611e15565b600083015250602082015167ffffffffffffffff811115611fc957600080fd5b611fd584828501611ecf565b602083015250604082015167ffffffffffffffff811115611ff557600080fd5b61200184828501611ecf565b6040830152506060612015848285016120d9565b6060830152506080612029848285016120b1565b60808301525060a061203d84828501611e15565b60a08301525060c061205184828501611e15565b60c08301525060e0612065848285016120d9565b60e08301525061010061207a848285016120b1565b6101008301525061012061209084828501611e29565b6101208301525092915050565b60006120a98235612c3e565b905092915050565b60006120bd8251612c3e565b905092915050565b60006120d18235612c48565b905092915050565b60006120e58251612c48565b905092915050565b6000602082840312156120ff57600080fd5b600061210d84828501611e01565b91505092915050565b60006020828403121561212857600080fd5b600061213684828501611e15565b91505092915050565b60006020828403121561215157600080fd5b600061215f84828501611e29565b91505092915050565b60006020828403121561217a57600080fd5b600061218884828501611e3d565b91505092915050565b6000602082840312156121a357600080fd5b60006121b184828501611e51565b91505092915050565b6000602082840312156121cc57600080fd5b600082013567ffffffffffffffff8111156121e657600080fd5b6121f284828501611e65565b91505092915050565b60006020828403121561220d57600080fd5b600061221b84828501611ebb565b91505092915050565b600080600080600080600060e0888a03121561223f57600080fd5b600088013567ffffffffffffffff81111561225957600080fd5b6122658a828b01611f25565b975050602088013567ffffffffffffffff81111561228257600080fd5b61228e8a828b01611f25565b965050604061229f8a828b016120c5565b95505060606122b08a828b0161209d565b94505060806122c18a828b016120c5565b93505060a06122d28a828b0161209d565b92505060c06122e38a828b01611e01565b91505092959891949750929550565b60006020828403121561230457600080fd5b600082015167ffffffffffffffff81111561231e57600080fd5b61232a84828501611f7b565b91505092915050565b60006020828403121561234557600080fd5b6000612353848285016120c5565b91505092915050565b61236581612c55565b82525050565b61237481612ba5565b82525050565b61238381612bb7565b82525050565b61239281612bc3565b82525050565b6123a181612c67565b82525050565b6123b081612c8b565b82525050565b60006123c182612b9a565b8084526123d5816020860160208601612ce2565b6123de81612d15565b602085010191505092915050565b60006123f782612b8f565b80845261240b816020860160208601612ce2565b61241481612d15565b602085010191505092915050565b6000603082527f696e70757420616464726573732068617320746f20626520612076616c69642060208301527f6f776e61626c6520636f6e7472616374000000000000000000000000000000006040830152606082019050919050565b6000603582527f696e70757420616464726573732068617320746f20626520612076616c69642060208301527f546f6b656e466163746f727920636f6e747261637400000000000000000000006040830152606082019050919050565b6000601e82527f68617320746f2062652061626f7665203020616e642062656c6f7720313800006020830152604082019050919050565b6000602782527f6e65772052656769737472792068617320746f206265206f776e65642062792060208301527f4d616e61676572000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000601782527f6d616e61676572206e6f7420696e697469616c697a65640000000000000000006020830152604082019050919050565b6000602b82527f6e657720546f6b656e466163746f72792068617320746f206265206f776e656460208301527f206279204d616e616765720000000000000000000000000000000000000000006040830152606082019050919050565b6000601f82527f68617320746f2062652061626f7665203020616e642062656c6f7720313030006020830152604082019050919050565b6000602882527f526567697374727920617070726f76652070726f706f73616c2068617320746f60208301527f20737563636565640000000000000000000000000000000000000000000000006040830152606082019050919050565b6000601182527f68617320746f2062652061626f766520310000000000000000000000000000006020830152604082019050919050565b6000602682527f546f6b656e466163746f727920646f6573206e6f74206f776e20546f6b656e5660208301527f657374696e6700000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000601f82527f546f6b656e56657374696e672068617320746f206265207468652073616d65006020830152604082019050919050565b61276c81612bed565b82525050565b61277b81612bf7565b82525050565b6000602082019050612796600083018461236b565b92915050565b60006040820190506127b1600083018561236b565b6127be602083018461236b565b9392505050565b60006020820190506127da600083018461237a565b92915050565b60006020820190506127f56000830184612389565b92915050565b60006040820190506128106000830185612389565b61281d602083018461236b565b9392505050565b60006020820190506128396000830184612398565b92915050565b600060208201905061285460008301846123a7565b92915050565b6000610100820190508181036000830152612875818b6123b6565b90508181036020830152612889818a6123b6565b90506128986040830189612772565b6128a56060830188612763565b6128b26080830187612772565b6128bf60a0830186612763565b6128cc60c083018561236b565b6128d960e083018461235c565b9998505050505050505050565b600060e0820190508181036000830152612900818a6123ec565b9050818103602083015261291481896123ec565b90506129236040830188612772565b6129306060830187612763565b61293d6080830186612772565b61294a60a0830185612763565b61295760c083018461236b565b98975050505050505050565b6000602082019050818103600083015261297c81612422565b9050919050565b6000602082019050818103600083015261299c8161247f565b9050919050565b600060208201905081810360008301526129bc816124dc565b9050919050565b600060208201905081810360008301526129dc81612513565b9050919050565b600060208201905081810360008301526129fc81612570565b9050919050565b60006020820190508181036000830152612a1c816125a7565b9050919050565b60006020820190508181036000830152612a3c81612604565b9050919050565b60006020820190508181036000830152612a5c8161263b565b9050919050565b60006020820190508181036000830152612a7c81612698565b9050919050565b60006020820190508181036000830152612a9c816126cf565b9050919050565b60006020820190508181036000830152612abc8161272c565b9050919050565b6000602082019050612ad86000830184612772565b92915050565b6000604051905081810181811067ffffffffffffffff82111715612b0157600080fd5b8060405250919050565b600067ffffffffffffffff821115612b2257600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115612b4e57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115612b7a57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b6000612bb082612bcd565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000612c0f82612bcd565b9050919050565b60008115159050919050565b6000819050919050565b6000612c3782612ba5565b9050919050565b6000819050919050565b600060ff82169050919050565b6000612c6082612caf565b9050919050565b6000612c7282612c79565b9050919050565b6000612c8482612bcd565b9050919050565b6000612c9682612c9d565b9050919050565b6000612ca882612bcd565b9050919050565b6000612cba82612cc1565b9050919050565b6000612ccc82612bcd565b9050919050565b82818337600083830152505050565b60005b83811015612d00578082015181840152602081019050612ce5565b83811115612d0f576000848401525b50505050565b6000601f19601f830116905091905056fea265627a7a723058204914cf314f1790362a9a4b9df2709a808b4ec283dd6c7ede3d2c9b871f2cded16c6578706572696d656e74616cf50037`

// DeployManager deploys a new Ethereum contract, binding an instance of Manager to it.
func DeployManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address, _tokenFactory common.Address) (common.Address, *types.Transaction, *Manager, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ManagerBin), backend, _registry, _tokenFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() constant returns(address)
func (_Manager *ManagerCaller) RegistryInstance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "RegistryInstance")
	return *ret0, err
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() constant returns(address)
func (_Manager *ManagerSession) RegistryInstance() (common.Address, error) {
	return _Manager.Contract.RegistryInstance(&_Manager.CallOpts)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() constant returns(address)
func (_Manager *ManagerCallerSession) RegistryInstance() (common.Address, error) {
	return _Manager.Contract.RegistryInstance(&_Manager.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() constant returns(address)
func (_Manager *ManagerCaller) TokenFactoryInstance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "TokenFactoryInstance")
	return *ret0, err
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() constant returns(address)
func (_Manager *ManagerSession) TokenFactoryInstance() (common.Address, error) {
	return _Manager.Contract.TokenFactoryInstance(&_Manager.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() constant returns(address)
func (_Manager *ManagerCallerSession) TokenFactoryInstance() (common.Address, error) {
	return _Manager.Contract.TokenFactoryInstance(&_Manager.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Manager *ManagerCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Manager *ManagerSession) GetOwner() (common.Address, error) {
	return _Manager.Contract.GetOwner(&_Manager.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Manager *ManagerCallerSession) GetOwner() (common.Address, error) {
	return _Manager.Contract.GetOwner(&_Manager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Manager *ManagerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Manager *ManagerSession) Initialized() (bool, error) {
	return _Manager.Contract.Initialized(&_Manager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Manager *ManagerCallerSession) Initialized() (bool, error) {
	return _Manager.Contract.Initialized(&_Manager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerCallerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Manager *ManagerCaller) ParseAddr(opts *bind.CallOpts, data []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "parseAddr", data)
	return *ret0, err
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Manager *ManagerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Manager.Contract.ParseAddr(&_Manager.CallOpts, data)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(data bytes) constant returns(parsed address)
func (_Manager *ManagerCallerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Manager.Contract.ParseAddr(&_Manager.CallOpts, data)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(_hashIndex bytes32) returns()
func (_Manager *ManagerTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "approveProposal", _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(_hashIndex bytes32) returns()
func (_Manager *ManagerSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveProposal(&_Manager.TransactOpts, _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(_hashIndex bytes32) returns()
func (_Manager *ManagerTransactorSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveProposal(&_Manager.TransactOpts, _hashIndex)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(_newManager address) returns()
func (_Manager *ManagerTransactor) MigrateManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "migrateManager", _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(_newManager address) returns()
func (_Manager *ManagerSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Manager.Contract.MigrateManager(&_Manager.TransactOpts, _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(_newManager address) returns()
func (_Manager *ManagerTransactorSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Manager.Contract.MigrateManager(&_Manager.TransactOpts, _newManager)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(_newPlatformWallet address) returns()
func (_Manager *ManagerTransactor) SetPlatformWallet(opts *bind.TransactOpts, _newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setPlatformWallet", _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(_newPlatformWallet address) returns()
func (_Manager *ManagerSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetPlatformWallet(&_Manager.TransactOpts, _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(_newPlatformWallet address) returns()
func (_Manager *ManagerTransactorSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetPlatformWallet(&_Manager.TransactOpts, _newPlatformWallet)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_newRegistry address) returns()
func (_Manager *ManagerTransactor) SetRegistry(opts *bind.TransactOpts, _newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setRegistry", _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_newRegistry address) returns()
func (_Manager *ManagerSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetRegistry(&_Manager.TransactOpts, _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(_newRegistry address) returns()
func (_Manager *ManagerTransactorSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetRegistry(&_Manager.TransactOpts, _newRegistry)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(_newTokenFactory address) returns()
func (_Manager *ManagerTransactor) SetTokenFactory(opts *bind.TransactOpts, _newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setTokenFactory", _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(_newTokenFactory address) returns()
func (_Manager *ManagerSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactory(&_Manager.TransactOpts, _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(_newTokenFactory address) returns()
func (_Manager *ManagerTransactorSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactory(&_Manager.TransactOpts, _newTokenFactory)
}

// SetTokenFactoryPercentage is a paid mutator transaction binding the contract method 0xb9ae5d2c.
//
// Solidity: function setTokenFactoryPercentage(_newPercentage uint8) returns()
func (_Manager *ManagerTransactor) SetTokenFactoryPercentage(opts *bind.TransactOpts, _newPercentage uint8) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setTokenFactoryPercentage", _newPercentage)
}

// SetTokenFactoryPercentage is a paid mutator transaction binding the contract method 0xb9ae5d2c.
//
// Solidity: function setTokenFactoryPercentage(_newPercentage uint8) returns()
func (_Manager *ManagerSession) SetTokenFactoryPercentage(_newPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactoryPercentage(&_Manager.TransactOpts, _newPercentage)
}

// SetTokenFactoryPercentage is a paid mutator transaction binding the contract method 0xb9ae5d2c.
//
// Solidity: function setTokenFactoryPercentage(_newPercentage uint8) returns()
func (_Manager *ManagerTransactorSession) SetTokenFactoryPercentage(_newPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactoryPercentage(&_Manager.TransactOpts, _newPercentage)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(_newTokenVesting address) returns()
func (_Manager *ManagerTransactor) SetTokenVesting(opts *bind.TransactOpts, _newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setTokenVesting", _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(_newTokenVesting address) returns()
func (_Manager *ManagerSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenVesting(&_Manager.TransactOpts, _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(_newTokenVesting address) returns()
func (_Manager *ManagerTransactorSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenVesting(&_Manager.TransactOpts, _newTokenVesting)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xa210853c.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address) returns(hashIndex bytes32)
func (_Manager *ManagerTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xa210853c.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address) returns(hashIndex bytes32)
func (_Manager *ManagerSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SubmitProposal(&_Manager.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xa210853c.
//
// Solidity: function submitProposal(_name string, _symbol string, _decimals uint8, _totalSupply uint256, _initialPercentage uint8, _vestingPeriodInWeeks uint256, _vestingBeneficiary address) returns(hashIndex bytes32)
func (_Manager *ManagerTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInWeeks *big.Int, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SubmitProposal(&_Manager.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInWeeks, _vestingBeneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Manager *ManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Manager *ManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Manager *ManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// ManagerLogManagerMigratedIterator is returned from FilterLogManagerMigrated and is used to iterate over the raw logs and unpacked data for LogManagerMigrated events raised by the Manager contract.
type ManagerLogManagerMigratedIterator struct {
	Event *ManagerLogManagerMigrated // Event containing the contract specifics and raw log

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
func (it *ManagerLogManagerMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogManagerMigrated)
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
		it.Event = new(ManagerLogManagerMigrated)
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
func (it *ManagerLogManagerMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogManagerMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogManagerMigrated represents a LogManagerMigrated event raised by the Manager contract.
type ManagerLogManagerMigrated struct {
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogManagerMigrated is a free log retrieval operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: e LogManagerMigrated(newManager indexed address)
func (_Manager *ManagerFilterer) FilterLogManagerMigrated(opts *bind.FilterOpts, newManager []common.Address) (*ManagerLogManagerMigratedIterator, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerLogManagerMigratedIterator{contract: _Manager.contract, event: "LogManagerMigrated", logs: logs, sub: sub}, nil
}

// WatchLogManagerMigrated is a free log subscription operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: e LogManagerMigrated(newManager indexed address)
func (_Manager *ManagerFilterer) WatchLogManagerMigrated(opts *bind.WatchOpts, sink chan<- *ManagerLogManagerMigrated, newManager []common.Address) (event.Subscription, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogManagerMigrated)
				if err := _Manager.contract.UnpackLog(event, "LogManagerMigrated", log); err != nil {
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

// ManagerLogRegistryChangedIterator is returned from FilterLogRegistryChanged and is used to iterate over the raw logs and unpacked data for LogRegistryChanged events raised by the Manager contract.
type ManagerLogRegistryChangedIterator struct {
	Event *ManagerLogRegistryChanged // Event containing the contract specifics and raw log

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
func (it *ManagerLogRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogRegistryChanged)
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
		it.Event = new(ManagerLogRegistryChanged)
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
func (it *ManagerLogRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogRegistryChanged represents a LogRegistryChanged event raised by the Manager contract.
type ManagerLogRegistryChanged struct {
	OldR common.Address
	NewR common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRegistryChanged is a free log retrieval operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: e LogRegistryChanged(oldR address, newR address)
func (_Manager *ManagerFilterer) FilterLogRegistryChanged(opts *bind.FilterOpts) (*ManagerLogRegistryChangedIterator, error) {

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &ManagerLogRegistryChangedIterator{contract: _Manager.contract, event: "LogRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchLogRegistryChanged is a free log subscription operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: e LogRegistryChanged(oldR address, newR address)
func (_Manager *ManagerFilterer) WatchLogRegistryChanged(opts *bind.WatchOpts, sink chan<- *ManagerLogRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogRegistryChanged)
				if err := _Manager.contract.UnpackLog(event, "LogRegistryChanged", log); err != nil {
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

// ManagerLogTokenFactoryChangedIterator is returned from FilterLogTokenFactoryChanged and is used to iterate over the raw logs and unpacked data for LogTokenFactoryChanged events raised by the Manager contract.
type ManagerLogTokenFactoryChangedIterator struct {
	Event *ManagerLogTokenFactoryChanged // Event containing the contract specifics and raw log

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
func (it *ManagerLogTokenFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogTokenFactoryChanged)
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
		it.Event = new(ManagerLogTokenFactoryChanged)
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
func (it *ManagerLogTokenFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogTokenFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogTokenFactoryChanged represents a LogTokenFactoryChanged event raised by the Manager contract.
type ManagerLogTokenFactoryChanged struct {
	OldTF common.Address
	NewTF common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogTokenFactoryChanged is a free log retrieval operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: e LogTokenFactoryChanged(oldTF address, newTF address)
func (_Manager *ManagerFilterer) FilterLogTokenFactoryChanged(opts *bind.FilterOpts) (*ManagerLogTokenFactoryChangedIterator, error) {

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return &ManagerLogTokenFactoryChangedIterator{contract: _Manager.contract, event: "LogTokenFactoryChanged", logs: logs, sub: sub}, nil
}

// WatchLogTokenFactoryChanged is a free log subscription operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: e LogTokenFactoryChanged(oldTF address, newTF address)
func (_Manager *ManagerFilterer) WatchLogTokenFactoryChanged(opts *bind.WatchOpts, sink chan<- *ManagerLogTokenFactoryChanged) (event.Subscription, error) {

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogTokenFactoryChanged)
				if err := _Manager.contract.UnpackLog(event, "LogTokenFactoryChanged", log); err != nil {
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

// ManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Manager contract.
type ManagerOwnershipTransferredIterator struct {
	Event *ManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerOwnershipTransferred)
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
		it.Event = new(ManagerOwnershipTransferred)
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
func (it *ManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerOwnershipTransferred represents a OwnershipTransferred event raised by the Manager contract.
type ManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Manager *ManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerOwnershipTransferredIterator{contract: _Manager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Manager *ManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerOwnershipTransferred)
				if err := _Manager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type Index string

const (
	IndexMultiSig   Index = "multisig"
	IndexManager    Index = "manager"
	IndexRegistry   Index = "registry"
	IndexVesting    Index = "vesting"
	IndexFactory    Index = "factory"
	IndexColdWallet Index = "coldWallet"
	IndexHotWallet  Index = "hotWallet"
)

type SmartContracts struct {
	MultiSig *MultiSigWallet
	Manager  *Manager
	Registry *Registry
	Factory  *TokenFactory
	Vesting  *TokenVesting
	NewToken func(common.Address) (*ERC20, error)
	Rolodex  map[Index]common.Address
}

func (s *SmartContracts) GetAddress(index Index) *common.Address {
	addr := s.Rolodex[index]
	return &addr
}

//LoadManager loads the manager contract on the addr address by the stipulated backend.
// Throws panic if it fails.
func LoadManager(addr common.Address, backend bind.ContractBackend) *Manager {
	manager, err := NewManager(addr, backend)
	log.Infof("loading manager contract on address %s", addr.Hex())
	if err != nil {
		log.Panicf("err occured at manager instance with : %v", err)
	}
	return manager
}

//LoadRegistry loads the registry contract on the addr address by the stipulated backend.
// Throws panic if fails.
func LoadRegistry(addr common.Address, backend bind.ContractBackend) *Registry {
	registry, err := NewRegistry(addr, backend)
	log.Infof("loading registry contract on address %s", addr.Hex())
	if err != nil {
		log.Panicf("err occured at registry instance with : %v", err)
	}
	return registry
}

//LoadVesting loads the vesting contract on the addr address by the stipulated backend.
// Throws panic if fails.
func LoadVesting(addr common.Address, backend bind.ContractBackend) *TokenVesting {
	vesting, err := NewTokenVesting(addr, backend)
	log.Infof("loading vesting contract on address %s", addr.Hex())
	if err != nil {
		log.Panicf("err occured at vesting instance with : %v", err)
	}
	return vesting
}

//LoadMultiSig loads the multisig contract on the addr address by the stipulated backend.
// Throws panic if fails.
func LoadMultiSig(addr common.Address, backend bind.ContractBackend) *MultiSigWallet {
	ms, err := NewMultiSigWallet(addr, backend)
	log.Infof("loading multisig contract on address %s", addr.Hex())
	if err != nil {
		log.Panicf("err occured at multisig instance with : %v", err)
	}
	return ms
}

//LoadFactory loads the factory contract on the addr address by the stipulated backend.
// Throws panic if fails.
func LoadFactory(addr common.Address, backend bind.ContractBackend) *TokenFactory {
	tf, err := NewTokenFactory(addr, backend)
	log.Infof("loading token factory contract on address %s", addr.Hex())
	if err != nil {
		log.Panicf("err occured at token factory instance with : %v", err)
	}
	return tf
}

func NewSmartContracts(addrs map[Index]common.Address, backend bind.ContractBackend) *SmartContracts {
	return &SmartContracts{
		Manager:  LoadManager(addrs[IndexManager], backend),
		Registry: LoadRegistry(addrs[IndexRegistry], backend),
		Vesting:  LoadVesting(addrs[IndexVesting], backend),
		MultiSig: LoadMultiSig(addrs[IndexMultiSig], backend),
		Factory:  LoadFactory(addrs[IndexFactory], backend),
		Rolodex:  addrs,
	}
}

func NewBlankSmartContracts() *SmartContracts {
	return &SmartContracts{
		Rolodex: make(map[Index]common.Address),
	}
}

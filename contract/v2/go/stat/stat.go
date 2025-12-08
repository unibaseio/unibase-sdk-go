// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stat

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

// StatMetaData contains all meta data concerning the Stat contract.
var StatMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_piece\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rsproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"piece\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPiece\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRSProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516110986100f95f395f8181610883015281816108ac01526109f001526110985ff3fe6080604052600436106100bf575f3560e01c80638da5cb5b1161007c578063ad3cb1cc11610057578063ad3cb1cc14610201578063c2e8e7171461023e578063f2fde38b1461026d578063ffa1ad741461028c575f80fd5b80638da5cb5b14610188578063900cf0cf146101c457806395ad21dc146101e2575f80fd5b80631459457a146100c35780634f1ef286146100e457806352d1902d146100f7578063715018a61461011e57806379ca7e0f1461013257806381cc0c7a14610169575b5f80fd5b3480156100ce575f80fd5b506100e26100dd366004610d01565b6102bc565b005b6100e26100f2366004610e0a565b610400565b348015610102575f80fd5b5061010b61041f565b6040519081526020015b60405180910390f35b348015610129575f80fd5b506100e261043a565b34801561013d575f80fd5b50600254610151906001600160a01b031681565b6040516001600160a01b039091168152602001610115565b348015610174575f80fd5b50600354610151906001600160a01b031681565b348015610193575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610151565b3480156101cf575f80fd5b505f54610151906001600160a01b031681565b3480156101ed575f80fd5b50600154610151906001600160a01b031681565b34801561020c575f80fd5b50610231604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516101159190610e99565b348015610249575f80fd5b5061025d610258366004610eb2565b61044d565b6040519015158152602001610115565b348015610278575f80fd5b506100e2610287366004610eeb565b6107fb565b348015610297575f80fd5b50610231604051806040016040528060058152602001640322e302e360dc1b81525081565b5f6102c561083d565b805490915060ff600160401b82041615906001600160401b03165f811580156102eb5750825b90505f826001600160401b031660011480156103065750303b155b905081158015610314575080155b156103325760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561035c57845460ff60401b1916600160401b1785555b61036586610867565b5f80546001600160a01b03808d166001600160a01b031992831617909255600180548c8416908316179055600280548b841690831617905560038054928a169290911691909117905583156103f457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b610408610878565b6104118261091c565b61041b8282610924565b5050565b5f6104286109e5565b505f8051602061104383398151915290565b610442610a2e565b61044b5f610a89565b565b5f805460408051639fa6a6e360e01b8152905183926001600160a01b031691639fa6a6e39160048083019260209291908290030181865afa158015610494573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b89190610f21565b60015460405163073a179f60e21b81529192505f916001600160a01b0390911690631ce85e7c906104ed908790600401610e99565b602060405180830381865afa158015610508573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052c9190610f21565b9050806001600160401b03165f0361054757505f9392505050565b600154604051636aadfac760e01b81526001600160401b03831660048201525f91829182916001600160a01b031690636aadfac790602401606060405180830381865afa15801561059a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105be9190610f4a565b925092509250846001600160401b0316816001600160401b0316116105e957505f9695505050505050565b6105f4600286610f9e565b94505f805b8460ff168160ff1610156107d15760015460405163ee91365b60e01b81526001600160401b038816600482015260ff831660248201525f9182916001600160a01b039091169063ee91365b906044016040805180830381865afa158015610662573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106869190610fbe565b91509150816001600160401b03165f036106a15750506107c9565b6002546040516383f0362760e01b81526001600160401b038a16600482015260ff851660248201525f916001600160a01b0316906383f0362790604401602060405180830381865afa1580156106f9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061071d9190610ff3565b9050801561072d575050506107c9565b600354604051636fb1c3fd60e11b81526001600160a01b0384811660048301526001600160401b038d1660248301529091169063df6387fa90604401602060405180830381865afa158015610784573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a89190610ff3565b905080156107b8575050506107c9565b6107c3600186611012565b94505050505b6001016105f9565b508260ff168160ff1610156107ed57505f979650505050505050565b506001979650505050505050565b610803610a2e565b6001600160a01b03811661083157604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61083a81610a89565b50565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b61086f610af9565b61083a81610b1e565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806108fe57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108f25f80516020611043833981519152546001600160a01b031690565b6001600160a01b031614155b1561044b5760405163703e46dd60e11b815260040160405180910390fd5b61083a610a2e565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561097e575060408051601f3d908101601f1916820190925261097b9181019061102b565b60015b6109a657604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610828565b5f8051602061104383398151915281146109d657604051632a87526960e21b815260048101829052602401610828565b6109e08383610b26565b505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461044b5760405163703e46dd60e11b815260040160405180910390fd5b33610a607f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461044b5760405163118cdaa760e01b8152336004820152602401610828565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610b01610b7b565b61044b57604051631afcd79f60e31b815260040160405180910390fd5b610803610af9565b610b2f82610b94565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610b73576109e08282610bf7565b61041b610c97565b5f610b8461083d565b54600160401b900460ff16919050565b806001600160a01b03163b5f03610bc957604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610828565b5f8051602061104383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f610c048484610cb6565b9050808015610c2557505f3d1180610c2557505f846001600160a01b03163b115b15610c3a57610c32610cc9565b915050610861565b8015610c6457604051639996b31560e01b81526001600160a01b0385166004820152602401610828565b3d15610c7757610c72610ce2565b610c90565b60405163d6bda27560e01b815260040160405180910390fd5b5092915050565b341561044b5760405163b398979f60e01b815260040160405180910390fd5b5f805f835160208501865af49392505050565b6040513d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b6001600160a01b038116811461083a575f80fd5b5f805f805f60a08688031215610d15575f80fd5b8535610d2081610ced565b94506020860135610d3081610ced565b93506040860135610d4081610ced565b92506060860135610d5081610ced565b91506080860135610d6081610ced565b809150509295509295909350565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610d91575f80fd5b81356001600160401b0380821115610dab57610dab610d6e565b604051601f8301601f19908116603f01168101908282118183101715610dd357610dd3610d6e565b81604052838152866020858801011115610deb575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215610e1b575f80fd5b8235610e2681610ced565b915060208301356001600160401b03811115610e40575f80fd5b610e4c85828601610d82565b9150509250929050565b5f81518084525f5b81811015610e7a57602081850181015186830182015201610e5e565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610eab6020830184610e56565b9392505050565b5f60208284031215610ec2575f80fd5b81356001600160401b03811115610ed7575f80fd5b610ee384828501610d82565b949350505050565b5f60208284031215610efb575f80fd5b8135610eab81610ced565b80516001600160401b0381168114610f1c575f80fd5b919050565b5f60208284031215610f31575f80fd5b610eab82610f06565b805160ff81168114610f1c575f80fd5b5f805f60608486031215610f5c575f80fd5b610f6584610f3a565b9250610f7360208501610f3a565b9150610f8160408501610f06565b90509250925092565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03828116828216039080821115610c9057610c90610f8a565b5f8060408385031215610fcf575f80fd5b610fd883610f06565b91506020830151610fe881610ced565b809150509250929050565b5f60208284031215611003575f80fd5b81518015158114610eab575f80fd5b60ff818116838216019081111561086157610861610f8a565b5f6020828403121561103b575f80fd5b505191905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220e9a98b7e4c911000a2a4a03d69605325b835a4666e09d9033683309ff83534be64736f6c63430008160033",
}

// StatABI is the input ABI used to generate the binding from.
// Deprecated: Use StatMetaData.ABI instead.
var StatABI = StatMetaData.ABI

// StatBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StatMetaData.Bin instead.
var StatBin = StatMetaData.Bin

// DeployStat deploys a new Ethereum contract, binding an instance of Stat to it.
func DeployStat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stat, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// Stat is an auto generated Go binding around an Ethereum contract.
type Stat struct {
	StatCaller     // Read-only binding to the contract
	StatTransactor // Write-only binding to the contract
	StatFilterer   // Log filterer for contract events
}

// StatCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatSession struct {
	Contract     *Stat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatCallerSession struct {
	Contract *StatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatTransactorSession struct {
	Contract     *StatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatRaw struct {
	Contract *Stat // Generic contract binding to access the raw methods on
}

// StatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatCallerRaw struct {
	Contract *StatCaller // Generic read-only contract binding to access the raw methods on
}

// StatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatTransactorRaw struct {
	Contract *StatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStat creates a new instance of Stat, bound to a specific deployed contract.
func NewStat(address common.Address, backend bind.ContractBackend) (*Stat, error) {
	contract, err := bindStat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// NewStatCaller creates a new read-only instance of Stat, bound to a specific deployed contract.
func NewStatCaller(address common.Address, caller bind.ContractCaller) (*StatCaller, error) {
	contract, err := bindStat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatCaller{contract: contract}, nil
}

// NewStatTransactor creates a new write-only instance of Stat, bound to a specific deployed contract.
func NewStatTransactor(address common.Address, transactor bind.ContractTransactor) (*StatTransactor, error) {
	contract, err := bindStat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatTransactor{contract: contract}, nil
}

// NewStatFilterer creates a new log filterer instance of Stat, bound to a specific deployed contract.
func NewStatFilterer(address common.Address, filterer bind.ContractFilterer) (*StatFilterer, error) {
	contract, err := bindStat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatFilterer{contract: contract}, nil
}

// bindStat binds a generic wrapper to an already deployed contract.
func bindStat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.StatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCallerSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCaller) CheckPiece(opts *bind.CallOpts, _pn []byte) (bool, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "checkPiece", _pn)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCallerSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCallerSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCaller) Eproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "eproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCallerSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatSession) Owner() (common.Address, error) {
	return _Stat.Contract.Owner(&_Stat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatCallerSession) Owner() (common.Address, error) {
	return _Stat.Contract.Owner(&_Stat.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCaller) Piece(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "piece")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCallerSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCaller) Rsproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "rsproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCallerSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactor) Initialize(opts *bind.TransactOpts, _epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "initialize", _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactorSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stat.Contract.RenounceOwnership(&_Stat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stat.Contract.RenounceOwnership(&_Stat.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.TransferOwnership(&_Stat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.TransferOwnership(&_Stat.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// StatInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stat contract.
type StatInitializedIterator struct {
	Event *StatInitialized // Event containing the contract specifics and raw log

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
func (it *StatInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatInitialized)
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
		it.Event = new(StatInitialized)
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
func (it *StatInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatInitialized represents a Initialized event raised by the Stat contract.
type StatInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatInitializedIterator, error) {

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatInitializedIterator{contract: _Stat.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatInitialized) (event.Subscription, error) {

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatInitialized)
				if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) ParseInitialized(log types.Log) (*StatInitialized, error) {
	event := new(StatInitialized)
	if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stat contract.
type StatOwnershipTransferredIterator struct {
	Event *StatOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StatOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatOwnershipTransferred)
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
		it.Event = new(StatOwnershipTransferred)
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
func (it *StatOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatOwnershipTransferred represents a OwnershipTransferred event raised by the Stat contract.
type StatOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stat *StatFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stat.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatOwnershipTransferredIterator{contract: _Stat.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stat *StatFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StatOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stat.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatOwnershipTransferred)
				if err := _Stat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stat *StatFilterer) ParseOwnershipTransferred(log types.Log) (*StatOwnershipTransferred, error) {
	event := new(StatOwnershipTransferred)
	if err := _Stat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Stat contract.
type StatUpgradedIterator struct {
	Event *StatUpgraded // Event containing the contract specifics and raw log

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
func (it *StatUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatUpgraded)
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
		it.Event = new(StatUpgraded)
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
func (it *StatUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatUpgraded represents a Upgraded event raised by the Stat contract.
type StatUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StatUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StatUpgradedIterator{contract: _Stat.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StatUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatUpgraded)
				if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) ParseUpgraded(log types.Log) (*StatUpgraded, error) {
	event := new(StatUpgraded)
	if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

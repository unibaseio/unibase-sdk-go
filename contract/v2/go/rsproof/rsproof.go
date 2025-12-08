// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rsproof

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

// IRSProofInitParams is an auto generated low-level Go binding around an user-defined struct.
type IRSProofInitParams struct {
	Token        common.Address
	Piece        common.Address
	Node         common.Address
	Base         common.Address
	Rsone        common.Address
	MinProveTime *big.Int
}

// IRSProofProofInfo is an auto generated low-level Go binding around an user-defined struct.
type IRSProofProofInfo struct {
	Fake     bool
	Chaler   common.Address
	ChalTime *big.Int
}

// RSProofMetaData contains all meta data concerning the RSProof contract.
var RSProofMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"KZG_MAX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KZG_VK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"base\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basePenalty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProof\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRSProof.ProofInfo\",\"components\":[{\"name\":\"fake\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chaler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chalTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStat\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVKRoot\",\"inputs\":[{\"name\":\"_rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIRSProof.InitParams\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"piece\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rsone\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minProveTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minProveTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"node\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"piece\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPiece\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prove\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsone\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMinProveTime\",\"inputs\":[{\"name\":\"_t\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVKRoot\",\"inputs\":[{\"name\":\"_rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_vkroot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Challenge\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Forge\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Prove\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523060805234801562000014575f80fd5b506200001f62000025565b620000d9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000765760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161250a620001005f395f8181611608015281816116310152611775015261250a5ff3fe608060405260043610610161575f3560e01c80638da5cb5b116100cd578063b29cc00911610087578063f2fde38b11610062578063f2fde38b146104e9578063f354cd5f14610508578063fc0c546a14610527578063ffa1ad7414610545575f80fd5b8063b29cc0091461046d578063d70754ec146104ab578063dd489315146104ca575f80fd5b80638da5cb5b146102e257806395ad21dc1461031e578063a617efec1461033d578063a7e029b814610407578063abe881d614610426578063ad3cb1cc1461043d575f80fd5b806352d1902d1161011e57806352d1902d1461023857806352f61ba81461024c5780635712e98c1461026b5780636c05dff814610280578063715018a61461029f57806383f03627146102b3575f80fd5b80630bd26cb51461016557806323178d221461018d5780633c530b08146101ae5780634f1ef286146101cf5780634f38ef98146101e25780635001f3b514610219575b5f80fd5b348015610170575f80fd5b5061017a60055481565b6040519081526020015b60405180910390f35b348015610198575f80fd5b506101a1610575565b6040516101849190611e49565b3480156101b9575f80fd5b506101cd6101c8366004611f34565b610591565b005b6101cd6101dd366004611fee565b610c25565b3480156101ed575f80fd5b50600454610201906001600160a01b031681565b6040516001600160a01b039091168152602001610184565b348015610224575f80fd5b50600354610201906001600160a01b031681565b348015610243575f80fd5b5061017a610c44565b348015610257575f80fd5b506101cd61026636600461203a565b610c5f565b348015610276575f80fd5b5061017a60065481565b34801561028b575f80fd5b506101cd61029a3660046120dc565b610ddb565b3480156102aa575f80fd5b506101cd6110eb565b3480156102be575f80fd5b506102d26102cd366004612162565b6110fe565b6040519015158152602001610184565b3480156102ed575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610201565b348015610329575f80fd5b50600154610201906001600160a01b031681565b348015610348575f80fd5b506103da610357366004612162565b60408051606080820183525f8083526020808401829052928401819052835191820184528082528183018181528285018281526001600160401b039790971682526007845284822060ff9687168084528186529583208054978816151585526101009097046001600160a01b031690915293905291905260019091015490915290565b604080518251151581526020808401516001600160a01b0316908201529181015190820152606001610184565b348015610412575f80fd5b506101cd610421366004612199565b61112d565b348015610431575f80fd5b5061017a630200000081565b348015610448575f80fd5b506101a1604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610478575f80fd5b5061017a6104873660046121d7565b60ff9182165f90815260086020908152604080832093909416825291909152205490565b3480156104b6575f80fd5b50600254610201906001600160a01b031681565b3480156104d5575f80fd5b506101cd6104e43660046121f3565b6111d8565b3480156104f4575f80fd5b506101cd61050336600461220a565b6111e5565b348015610513575f80fd5b506101cd610522366004612225565b611222565b348015610532575f80fd5b505f54610201906001600160a01b031681565b348015610550575f80fd5b506101a1604051806040016040528060058152602001640322e302e360dc1b81525081565b6040518060600160405280603081526020016124a56030913981565b60015460405163073a179f60e21b81525f916001600160a01b031690631ce85e7c906105c1908890600401611e49565b602060405180830381865afa1580156105dc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106009190612262565b60015460405163175164c360e31b81529192505f916001600160a01b039091169063ba8b261890610635908890600401611e49565b602060405180830381865afa158015610650573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106749190612262565b600154604051632d942d4160e21b81526001600160401b038086166004830152831660248201529192505f9182918291829182916001600160a01b039091169063b650b5049060440160a060405180830381865afa1580156106d8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106fc919061227d565b9450945094509450945060075f886001600160401b03166001600160401b031681526020019081526020015f205f8a60ff1660ff1681526020019081526020015f205f015f9054906101000a900460ff16156107735760405162461bcd60e51b815260040161076a906122e4565b60405180910390fd5b6001600160401b0387165f90815260076020908152604080832060ff8d1684529091529020600101546107d25760405162461bcd60e51b81526020600482015260076024820152661b9bc818da185b60ca1b604482015260640161076a565b6005546001600160401b0388165f90815260076020908152604080832060ff8e1684529091529020600101546108089190612304565b43106108425760405162461bcd60e51b8152602060048201526009602482015268195e18d95959081c1d60ba1b604482015260640161076a565b60ff8086165f908152600860209081526040808320938816835292905290812054900361089d5760405162461bcd60e51b81526020600482015260096024820152681b9bc81d9adc9bdbdd60ba1b604482015260640161076a565b604080516003808252608082019092525f916020820160608036833750505060ff8781165f908152600860209081526040808320938a168352929052908120548251929350918391906108f2576108f2612323565b60209081029190910101526bffffffffffffffffffffffff19606084901b1661091f600160591b82612304565b905061093967ffffffff00000000602087901b1682612304565b905061094860ff8c1682612304565b90508c8c610955836114ae565b6040518060600160405280603081526020016124a5603091396040516020016109819493929190612337565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018106905080826001815181106109d4576109d4612323565b60209081029190910101525080517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001830690819083906002908110610a1b57610a1b612323565b602090810291909101015260048054604051633f27bd4560e11b81525f926001600160a01b0390921691637e4f7a8a91610a59918f9188910161238d565b6020604051808303815f875af1158015610a75573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9991906123e2565b905080610ad15760405162461bcd60e51b815260206004820152600660248201526534b73b10383360d11b604482015260640161076a565b60015460405163ee91365b60e01b81526001600160401b038c16600482015260ff8e1660248201526001600160a01b039091169063ee91365b906044016040805180830381865afa158015610b28573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4c9190612401565b90965094506001600160401b0389811690871614610b7c5760405162461bcd60e51b815260040161076a9061242e565b610b8585611554565b5f60075f8c6001600160401b03166001600160401b031681526020019081526020015f205f8e60ff1660ff1681526020019081526020015f2060010181905550846001600160a01b03167f4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd16598a604051610c0d91906001600160401b0391909116815260200190565b60405180910390a25050505050505050505050505050565b610c2d6115fd565b610c36826116a1565b610c4082826116a9565b5050565b5f610c4d61176a565b505f8051602061248583398151915290565b5f610c686117b3565b805490915060ff600160401b82041615906001600160401b03165f81158015610c8e5750825b90505f826001600160401b03166001148015610ca95750303b155b905081158015610cb7575080155b15610cd55760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610cff57845460ff60401b1916600160401b1785555b610d08866117db565b86515f80546001600160a01b03199081166001600160a01b0393841617909155602089015160018054831691841691909117905560408901516002805483169184169190911790556060890151600380548316918416919091179055608089015160048054909216921691909117905560a0870151600555670de0b6b3a76400006006558315610dd257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60015460405163073a179f60e21b81525f916001600160a01b031690631ce85e7c90610e0b908790600401611e49565b602060405180830381865afa158015610e26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4a9190612262565b60015460405163175164c360e31b81529192505f916001600160a01b039091169063ba8b261890610e7f908790600401611e49565b602060405180830381865afa158015610e9a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebe9190612262565b60015460405163ee91365b60e01b81526001600160401b038516600482015260ff861660248201529192505f9182916001600160a01b03169063ee91365b906044016040805180830381865afa158015610f1a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f3e9190612401565b91509150816001600160401b0316836001600160401b031614610f735760405162461bcd60e51b815260040161076a9061242e565b6001600160a01b038116610fb45760405162461bcd60e51b81526020600482015260086024820152676e6f2073746f726560c01b604482015260640161076a565b6001600160401b0384165f90815260076020908152604080832060ff808a1685529252909120541615610ff95760405162461bcd60e51b815260040161076a906122e4565b6001600160401b0384165f90815260076020908152604080832060ff89168452909152902060010154156110595760405162461bcd60e51b81526020600482015260076024820152661a5b8818da185b60ca1b604482015260640161076a565b61106381336117ec565b6001600160401b038481165f90815260076020908152604080832060ff8a1684528252918290204360018201558054610100600160a81b0319163361010002179055905191851682526001600160a01b038316917f40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823910160405180910390a250505050505050565b6110f361186c565b6110fc5f6118c7565b565b6001600160401b0382165f90815260076020908152604080832060ff8086168552925290912054165b92915050565b61113561186c565b60ff8084165f908152600860209081526040808320938616835292905220541561118e5760405162461bcd60e51b815260206004820152600a6024820152691a185cc81d9adc9bdbdd60b21b604482015260640161076a565b60ff9283165f908152600860209081526040808320949095168252929092529190207f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019091069055565b6111e061186c565b600555565b6111ed61186c565b6001600160a01b03811661121657604051631e4fbdf760e01b81525f600482015260240161076a565b61121f816118c7565b50565b6001600160401b0383165f90815260076020908152604080832060ff808616855292529091205416156112675760405162461bcd60e51b815260040161076a906122e4565b6001600160401b0383165f90815260076020908152604080832060ff851684529091529020600101546112c65760405162461bcd60e51b81526020600482015260076024820152661b9bc818da185b60ca1b604482015260640161076a565b60015460405163ee91365b60e01b81526001600160401b038516600482015260ff831660248201525f9182916001600160a01b039091169063ee91365b906044016040805180830381865afa158015611321573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113459190612401565b91509150816001600160401b0316846001600160401b03161461137a5760405162461bcd60e51b815260040161076a9061242e565b6001600160a01b0381166113bb5760405162461bcd60e51b81526020600482015260086024820152676e6f2073746f726560c01b604482015260640161076a565b6005546001600160401b0386165f90815260076020908152604080832060ff881684529091529020600101546113f19190612304565b4311156114a7576001600160401b0385165f90815260076020908152604080832060ff8716845290915290205461143790829061010090046001600160a01b0316611937565b6001600160401b038581165f90815260076020908152604080832060ff88168452825291829020805460ff19166001179055905191861682526001600160a01b038316917f1cb9b342b5a7799871a1b1069c31f4be8311a3f1580fbaf8ecfd35297ce1244c910160405180910390a25b5050505050565b60606fffffffffffffffff00000000000000006001600160c01b031960c084901b1667ffffffffffffffff60801b604085901b166114ec8183612304565b604093841b93909250858416901c90506115068183612304565b60409390931b9291505083821660c01c6115208183612304565b604080516030808252606082019092529193505f919060208201818036833750505060208101939093525090949350505050565b600254600654604051637eee288d60e01b81526001600160a01b0384811660048301526024820192909252911690637eee288d906044015f604051808303815f87803b1580156115a2575f80fd5b505af11580156115b4573d5f803e3d5ffd5b505050506115dd8160026006546115cb919061244e565b5f546001600160a01b03169190611ac8565b60035460065461121f916001600160a01b0316906115cb9060029061244e565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061168357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166116775f80516020612485833981519152546001600160a01b031690565b6001600160a01b031614155b156110fc5760405163703e46dd60e11b815260040160405180910390fd5b61121f61186c565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611703575060408051601f3d908101601f191682019092526117009181019061246d565b60015b61172b57604051634c9c8ce360e01b81526001600160a01b038316600482015260240161076a565b5f80516020612485833981519152811461175b57604051632a87526960e21b81526004810182905260240161076a565b6117658383611afd565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146110fc5760405163703e46dd60e11b815260040160405180910390fd5b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00611127565b6117e3611b52565b61121f81611b77565b60025460065460405163282d3fdf60e01b81526001600160a01b038581166004830152602482019290925291169063282d3fdf906044015f604051808303815f87803b15801561183a575f80fd5b505af115801561184c573d5f803e3d5ffd5b50506006545f54610c4093506001600160a01b0316915083903090611b7f565b3361189e7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146110fc5760405163118cdaa760e01b815233600482015260240161076a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b600254600654604051637eee288d60e01b81526001600160a01b0385811660048301526024820192909252911690637eee288d906044015f604051808303815f87803b158015611985575f80fd5b505af1158015611997573d5f803e3d5ffd5b5050600280546006546001600160a01b039091169350639748dcdc9250859185916119c2919061244e565b6040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064015f604051808303815f87803b158015611a0e575f80fd5b505af1158015611a20573d5f803e3d5ffd5b5050600280546003546006546001600160a01b039283169550639748dcdc945087939290911691611a509161244e565b6040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064015f604051808303815f87803b158015611a9c575f80fd5b505af1158015611aae573d5f803e3d5ffd5b50506006545f54610c4093506001600160a01b0316915083905b611ad58383836001611bbb565b61176557604051635274afe760e01b81526001600160a01b038416600482015260240161076a565b611b0682611c1d565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611b4a576117658282611c80565b610c40611d20565b611b5a611d3f565b6110fc57604051631afcd79f60e31b815260040160405180910390fd5b6111ed611b52565b611b8d848484846001611d58565b611bb557604051635274afe760e01b81526001600160a01b038516600482015260240161076a565b50505050565b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316611c11578383151615611c05573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b806001600160a01b03163b5f03611c5257604051634c9c8ce360e01b81526001600160a01b038216600482015260240161076a565b5f8051602061248583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f611c8d8484611dc5565b9050808015611cae57505f3d1180611cae57505f846001600160a01b03163b115b15611cc357611cbb611dd8565b915050611127565b8015611ced57604051639996b31560e01b81526001600160a01b038516600482015260240161076a565b3d15611d0057611cfb611df1565b611d19565b60405163d6bda27560e01b815260040160405180910390fd5b5092915050565b34156110fc5760405163b398979f60e01b815260040160405180910390fd5b5f611d486117b3565b54600160401b900460ff16919050565b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f51148316611db4578383151615611da8573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b5f805f835160208501865af49392505050565b6040513d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b5f5b83811015611e16578181015183820152602001611dfe565b50505f910152565b5f8151808452611e35816020860160208601611dfc565b601f01601f19169290920160200192915050565b602081525f611e5b6020830184611e1e565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715611e9857611e98611e62565b60405290565b5f82601f830112611ead575f80fd5b81356001600160401b0380821115611ec757611ec7611e62565b604051601f8301601f19908116603f01168101908282118183101715611eef57611eef611e62565b81604052838152866020858801011115611f07575f80fd5b836020870160208301375f602085830101528094505050505092915050565b60ff8116811461121f575f80fd5b5f805f8060808587031215611f47575f80fd5b84356001600160401b0380821115611f5d575f80fd5b611f6988838901611e9e565b95506020870135915080821115611f7e575f80fd5b611f8a88838901611e9e565b945060408701359150611f9c82611f26565b90925060608601359080821115611fb1575f80fd5b50611fbe87828801611e9e565b91505092959194509250565b6001600160a01b038116811461121f575f80fd5b8035611fe981611fca565b919050565b5f8060408385031215611fff575f80fd5b823561200a81611fca565b915060208301356001600160401b03811115612024575f80fd5b61203085828601611e9e565b9150509250929050565b5f8082840360e081121561204c575f80fd5b60c0811215612059575f80fd5b50612062611e76565b833561206d81611fca565b8152602084013561207d81611fca565b6020820152604084013561209081611fca565b604082015260608401356120a381611fca565b606082015260808401356120b681611fca565b608082015260a0848101359082015291506120d360c08401611fde565b90509250929050565b5f805f606084860312156120ee575f80fd5b83356001600160401b0380821115612104575f80fd5b61211087838801611e9e565b94506020860135915080821115612125575f80fd5b5061213286828701611e9e565b925050604084013561214381611f26565b809150509250925092565b6001600160401b038116811461121f575f80fd5b5f8060408385031215612173575f80fd5b823561217e8161214e565b9150602083013561218e81611f26565b809150509250929050565b5f805f606084860312156121ab575f80fd5b83356121b681611f26565b925060208401356121c681611f26565b929592945050506040919091013590565b5f80604083850312156121e8575f80fd5b823561217e81611f26565b5f60208284031215612203575f80fd5b5035919050565b5f6020828403121561221a575f80fd5b8135611e5b81611fca565b5f805f60608486031215612237575f80fd5b83356122428161214e565b925060208401356122528161214e565b9150604084013561214381611f26565b5f60208284031215612272575f80fd5b8151611e5b8161214e565b5f805f805f60a08688031215612291575f80fd5b855161229c81611f26565b60208701519095506122ad81611f26565b60408701519094506122be8161214e565b60608701519093506122cf81611fca565b80925050608086015190509295509295909350565b60208082526006908201526519985a5b195960d21b604082015260600190565b8082018082111561112757634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8551612348818460208a01611dfc565b85519083019061235c818360208a01611dfc565b855191019061236f818360208901611dfc565b8451910190612382818360208801611dfc565b019695505050505050565b604081525f61239f6040830185611e1e565b8281036020848101919091528451808352858201928201905f5b818110156123d5578451835293830193918301916001016123b9565b5090979650505050505050565b5f602082840312156123f2575f80fd5b81518015158114611e5b575f80fd5b5f8060408385031215612412575f80fd5b825161241d8161214e565b602084015190925061218e81611fca565b602080825260069082015265696e7620726960d01b604082015260600190565b5f8261246857634e487b7160e01b5f52601260045260245ffd5b500490565b5f6020828403121561247d575f80fd5b505191905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52a26469706673582212205663d73e64b8d89fcd3a564bb783da41a304045a0290bcf086cc9eb7d10d92f964736f6c63430008160033",
}

// RSProofABI is the input ABI used to generate the binding from.
// Deprecated: Use RSProofMetaData.ABI instead.
var RSProofABI = RSProofMetaData.ABI

// RSProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RSProofMetaData.Bin instead.
var RSProofBin = RSProofMetaData.Bin

// DeployRSProof deploys a new Ethereum contract, binding an instance of RSProof to it.
func DeployRSProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RSProof, error) {
	parsed, err := RSProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RSProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RSProof{RSProofCaller: RSProofCaller{contract: contract}, RSProofTransactor: RSProofTransactor{contract: contract}, RSProofFilterer: RSProofFilterer{contract: contract}}, nil
}

// RSProof is an auto generated Go binding around an Ethereum contract.
type RSProof struct {
	RSProofCaller     // Read-only binding to the contract
	RSProofTransactor // Write-only binding to the contract
	RSProofFilterer   // Log filterer for contract events
}

// RSProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type RSProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RSProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RSProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RSProofSession struct {
	Contract     *RSProof          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RSProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RSProofCallerSession struct {
	Contract *RSProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RSProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RSProofTransactorSession struct {
	Contract     *RSProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RSProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type RSProofRaw struct {
	Contract *RSProof // Generic contract binding to access the raw methods on
}

// RSProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RSProofCallerRaw struct {
	Contract *RSProofCaller // Generic read-only contract binding to access the raw methods on
}

// RSProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RSProofTransactorRaw struct {
	Contract *RSProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRSProof creates a new instance of RSProof, bound to a specific deployed contract.
func NewRSProof(address common.Address, backend bind.ContractBackend) (*RSProof, error) {
	contract, err := bindRSProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RSProof{RSProofCaller: RSProofCaller{contract: contract}, RSProofTransactor: RSProofTransactor{contract: contract}, RSProofFilterer: RSProofFilterer{contract: contract}}, nil
}

// NewRSProofCaller creates a new read-only instance of RSProof, bound to a specific deployed contract.
func NewRSProofCaller(address common.Address, caller bind.ContractCaller) (*RSProofCaller, error) {
	contract, err := bindRSProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RSProofCaller{contract: contract}, nil
}

// NewRSProofTransactor creates a new write-only instance of RSProof, bound to a specific deployed contract.
func NewRSProofTransactor(address common.Address, transactor bind.ContractTransactor) (*RSProofTransactor, error) {
	contract, err := bindRSProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RSProofTransactor{contract: contract}, nil
}

// NewRSProofFilterer creates a new log filterer instance of RSProof, bound to a specific deployed contract.
func NewRSProofFilterer(address common.Address, filterer bind.ContractFilterer) (*RSProofFilterer, error) {
	contract, err := bindRSProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RSProofFilterer{contract: contract}, nil
}

// bindRSProof binds a generic wrapper to an already deployed contract.
func bindRSProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RSProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSProof *RSProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSProof.Contract.RSProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSProof *RSProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.Contract.RSProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSProof *RSProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSProof.Contract.RSProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSProof *RSProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSProof *RSProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSProof *RSProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSProof.Contract.contract.Transact(opts, method, params...)
}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofCaller) KZGMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "KZG_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofSession) KZGMAX() (*big.Int, error) {
	return _RSProof.Contract.KZGMAX(&_RSProof.CallOpts)
}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofCallerSession) KZGMAX() (*big.Int, error) {
	return _RSProof.Contract.KZGMAX(&_RSProof.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofCaller) KZGVK(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "KZG_VK")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofSession) KZGVK() ([]byte, error) {
	return _RSProof.Contract.KZGVK(&_RSProof.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofCallerSession) KZGVK() ([]byte, error) {
	return _RSProof.Contract.KZGVK(&_RSProof.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RSProof *RSProofCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RSProof *RSProofSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RSProof.Contract.UPGRADEINTERFACEVERSION(&_RSProof.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RSProof *RSProofCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RSProof.Contract.UPGRADEINTERFACEVERSION(&_RSProof.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RSProof *RSProofCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RSProof *RSProofSession) VERSION() (string, error) {
	return _RSProof.Contract.VERSION(&_RSProof.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RSProof *RSProofCallerSession) VERSION() (string, error) {
	return _RSProof.Contract.VERSION(&_RSProof.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_RSProof *RSProofCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_RSProof *RSProofSession) Base() (common.Address, error) {
	return _RSProof.Contract.Base(&_RSProof.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_RSProof *RSProofCallerSession) Base() (common.Address, error) {
	return _RSProof.Contract.Base(&_RSProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofCaller) BasePenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "basePenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofSession) BasePenalty() (*big.Int, error) {
	return _RSProof.Contract.BasePenalty(&_RSProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofCallerSession) BasePenalty() (*big.Int, error) {
	return _RSProof.Contract.BasePenalty(&_RSProof.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofCaller) GetProof(opts *bind.CallOpts, _pi uint64, _pri uint8) (IRSProofProofInfo, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "getProof", _pi, _pri)

	if err != nil {
		return *new(IRSProofProofInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRSProofProofInfo)).(*IRSProofProofInfo)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofSession) GetProof(_pi uint64, _pri uint8) (IRSProofProofInfo, error) {
	return _RSProof.Contract.GetProof(&_RSProof.CallOpts, _pi, _pri)
}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofCallerSession) GetProof(_pi uint64, _pri uint8) (IRSProofProofInfo, error) {
	return _RSProof.Contract.GetProof(&_RSProof.CallOpts, _pi, _pri)
}

// GetStat is a free data retrieval call binding the contract method 0x83f03627.
//
// Solidity: function getStat(uint64 _pi, uint8 _pri) view returns(bool)
func (_RSProof *RSProofCaller) GetStat(opts *bind.CallOpts, _pi uint64, _pri uint8) (bool, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "getStat", _pi, _pri)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStat is a free data retrieval call binding the contract method 0x83f03627.
//
// Solidity: function getStat(uint64 _pi, uint8 _pri) view returns(bool)
func (_RSProof *RSProofSession) GetStat(_pi uint64, _pri uint8) (bool, error) {
	return _RSProof.Contract.GetStat(&_RSProof.CallOpts, _pi, _pri)
}

// GetStat is a free data retrieval call binding the contract method 0x83f03627.
//
// Solidity: function getStat(uint64 _pi, uint8 _pri) view returns(bool)
func (_RSProof *RSProofCallerSession) GetStat(_pi uint64, _pri uint8) (bool, error) {
	return _RSProof.Contract.GetStat(&_RSProof.CallOpts, _pi, _pri)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofCaller) GetVKRoot(opts *bind.CallOpts, _rsn uint8, _rsk uint8) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "getVKRoot", _rsn, _rsk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofSession) GetVKRoot(_rsn uint8, _rsk uint8) (*big.Int, error) {
	return _RSProof.Contract.GetVKRoot(&_RSProof.CallOpts, _rsn, _rsk)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofCallerSession) GetVKRoot(_rsn uint8, _rsk uint8) (*big.Int, error) {
	return _RSProof.Contract.GetVKRoot(&_RSProof.CallOpts, _rsn, _rsk)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofCaller) MinProveTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "minProveTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofSession) MinProveTime() (*big.Int, error) {
	return _RSProof.Contract.MinProveTime(&_RSProof.CallOpts)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofCallerSession) MinProveTime() (*big.Int, error) {
	return _RSProof.Contract.MinProveTime(&_RSProof.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_RSProof *RSProofCaller) Node(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "node")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_RSProof *RSProofSession) Node() (common.Address, error) {
	return _RSProof.Contract.Node(&_RSProof.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_RSProof *RSProofCallerSession) Node() (common.Address, error) {
	return _RSProof.Contract.Node(&_RSProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofSession) Owner() (common.Address, error) {
	return _RSProof.Contract.Owner(&_RSProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofCallerSession) Owner() (common.Address, error) {
	return _RSProof.Contract.Owner(&_RSProof.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_RSProof *RSProofCaller) Piece(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "piece")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_RSProof *RSProofSession) Piece() (common.Address, error) {
	return _RSProof.Contract.Piece(&_RSProof.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_RSProof *RSProofCallerSession) Piece() (common.Address, error) {
	return _RSProof.Contract.Piece(&_RSProof.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RSProof *RSProofCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RSProof *RSProofSession) ProxiableUUID() ([32]byte, error) {
	return _RSProof.Contract.ProxiableUUID(&_RSProof.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RSProof *RSProofCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RSProof.Contract.ProxiableUUID(&_RSProof.CallOpts)
}

// Rsone is a free data retrieval call binding the contract method 0x4f38ef98.
//
// Solidity: function rsone() view returns(address)
func (_RSProof *RSProofCaller) Rsone(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "rsone")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsone is a free data retrieval call binding the contract method 0x4f38ef98.
//
// Solidity: function rsone() view returns(address)
func (_RSProof *RSProofSession) Rsone() (common.Address, error) {
	return _RSProof.Contract.Rsone(&_RSProof.CallOpts)
}

// Rsone is a free data retrieval call binding the contract method 0x4f38ef98.
//
// Solidity: function rsone() view returns(address)
func (_RSProof *RSProofCallerSession) Rsone() (common.Address, error) {
	return _RSProof.Contract.Rsone(&_RSProof.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RSProof *RSProofCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RSProof *RSProofSession) Token() (common.Address, error) {
	return _RSProof.Contract.Token(&_RSProof.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RSProof *RSProofCallerSession) Token() (common.Address, error) {
	return _RSProof.Contract.Token(&_RSProof.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofTransactor) Challenge(opts *bind.TransactOpts, _pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "challenge", _pn, _rn, _pri)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofSession) Challenge(_pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Challenge(&_RSProof.TransactOpts, _pn, _rn, _pri)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofTransactorSession) Challenge(_pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Challenge(&_RSProof.TransactOpts, _pn, _rn, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofTransactor) Check(opts *bind.TransactOpts, _pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "check", _pi, _ri, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofSession) Check(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Check(&_RSProof.TransactOpts, _pi, _ri, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofTransactorSession) Check(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Check(&_RSProof.TransactOpts, _pi, _ri, _pri)
}

// Initialize is a paid mutator transaction binding the contract method 0x52f61ba8.
//
// Solidity: function initialize((address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_RSProof *RSProofTransactor) Initialize(opts *bind.TransactOpts, params IRSProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "initialize", params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x52f61ba8.
//
// Solidity: function initialize((address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_RSProof *RSProofSession) Initialize(params IRSProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.Initialize(&_RSProof.TransactOpts, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x52f61ba8.
//
// Solidity: function initialize((address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_RSProof *RSProofTransactorSession) Initialize(params IRSProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.Initialize(&_RSProof.TransactOpts, params, initialOwner)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofTransactor) Prove(opts *bind.TransactOpts, _pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "prove", _pn, _rn, _pri, _proof)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofSession) Prove(_pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.Contract.Prove(&_RSProof.TransactOpts, _pn, _rn, _pri, _proof)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofTransactorSession) Prove(_pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.Contract.Prove(&_RSProof.TransactOpts, _pn, _rn, _pri, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofSession) RenounceOwnership() (*types.Transaction, error) {
	return _RSProof.Contract.RenounceOwnership(&_RSProof.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RSProof.Contract.RenounceOwnership(&_RSProof.TransactOpts)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_RSProof *RSProofTransactor) SetMinProveTime(opts *bind.TransactOpts, _t *big.Int) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "setMinProveTime", _t)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_RSProof *RSProofSession) SetMinProveTime(_t *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetMinProveTime(&_RSProof.TransactOpts, _t)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_RSProof *RSProofTransactorSession) SetMinProveTime(_t *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetMinProveTime(&_RSProof.TransactOpts, _t)
}

// SetVKRoot is a paid mutator transaction binding the contract method 0xa7e029b8.
//
// Solidity: function setVKRoot(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofTransactor) SetVKRoot(opts *bind.TransactOpts, _rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "setVKRoot", _rsn, _rsk, _vkroot)
}

// SetVKRoot is a paid mutator transaction binding the contract method 0xa7e029b8.
//
// Solidity: function setVKRoot(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofSession) SetVKRoot(_rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetVKRoot(&_RSProof.TransactOpts, _rsn, _rsk, _vkroot)
}

// SetVKRoot is a paid mutator transaction binding the contract method 0xa7e029b8.
//
// Solidity: function setVKRoot(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofTransactorSession) SetVKRoot(_rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetVKRoot(&_RSProof.TransactOpts, _rsn, _rsk, _vkroot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.TransferOwnership(&_RSProof.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.TransferOwnership(&_RSProof.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RSProof *RSProofTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RSProof *RSProofSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RSProof.Contract.UpgradeToAndCall(&_RSProof.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RSProof *RSProofTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RSProof.Contract.UpgradeToAndCall(&_RSProof.TransactOpts, newImplementation, data)
}

// RSProofChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the RSProof contract.
type RSProofChallengeIterator struct {
	Event *RSProofChallenge // Event containing the contract specifics and raw log

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
func (it *RSProofChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofChallenge)
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
		it.Event = new(RSProofChallenge)
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
func (it *RSProofChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofChallenge represents a Challenge event raised by the RSProof contract.
type RSProofChallenge struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterChallenge(opts *bind.FilterOpts, _s []common.Address) (*RSProofChallengeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofChallengeIterator{contract: _RSProof.contract, event: "Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *RSProofChallenge, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofChallenge)
				if err := _RSProof.contract.UnpackLog(event, "Challenge", log); err != nil {
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

// ParseChallenge is a log parse operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseChallenge(log types.Log) (*RSProofChallenge, error) {
	event := new(RSProofChallenge)
	if err := _RSProof.contract.UnpackLog(event, "Challenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofForgeIterator is returned from FilterForge and is used to iterate over the raw logs and unpacked data for Forge events raised by the RSProof contract.
type RSProofForgeIterator struct {
	Event *RSProofForge // Event containing the contract specifics and raw log

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
func (it *RSProofForgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofForge)
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
		it.Event = new(RSProofForge)
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
func (it *RSProofForgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofForgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofForge represents a Forge event raised by the RSProof contract.
type RSProofForge struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForge is a free log retrieval operation binding the contract event 0x1cb9b342b5a7799871a1b1069c31f4be8311a3f1580fbaf8ecfd35297ce1244c.
//
// Solidity: event Forge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterForge(opts *bind.FilterOpts, _s []common.Address) (*RSProofForgeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Forge", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofForgeIterator{contract: _RSProof.contract, event: "Forge", logs: logs, sub: sub}, nil
}

// WatchForge is a free log subscription operation binding the contract event 0x1cb9b342b5a7799871a1b1069c31f4be8311a3f1580fbaf8ecfd35297ce1244c.
//
// Solidity: event Forge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchForge(opts *bind.WatchOpts, sink chan<- *RSProofForge, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Forge", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofForge)
				if err := _RSProof.contract.UnpackLog(event, "Forge", log); err != nil {
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

// ParseForge is a log parse operation binding the contract event 0x1cb9b342b5a7799871a1b1069c31f4be8311a3f1580fbaf8ecfd35297ce1244c.
//
// Solidity: event Forge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseForge(log types.Log) (*RSProofForge, error) {
	event := new(RSProofForge)
	if err := _RSProof.contract.UnpackLog(event, "Forge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RSProof contract.
type RSProofInitializedIterator struct {
	Event *RSProofInitialized // Event containing the contract specifics and raw log

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
func (it *RSProofInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofInitialized)
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
		it.Event = new(RSProofInitialized)
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
func (it *RSProofInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofInitialized represents a Initialized event raised by the RSProof contract.
type RSProofInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RSProof *RSProofFilterer) FilterInitialized(opts *bind.FilterOpts) (*RSProofInitializedIterator, error) {

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RSProofInitializedIterator{contract: _RSProof.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RSProof *RSProofFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RSProofInitialized) (event.Subscription, error) {

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofInitialized)
				if err := _RSProof.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RSProof *RSProofFilterer) ParseInitialized(log types.Log) (*RSProofInitialized, error) {
	event := new(RSProofInitialized)
	if err := _RSProof.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RSProof contract.
type RSProofOwnershipTransferredIterator struct {
	Event *RSProofOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RSProofOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofOwnershipTransferred)
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
		it.Event = new(RSProofOwnershipTransferred)
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
func (it *RSProofOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofOwnershipTransferred represents a OwnershipTransferred event raised by the RSProof contract.
type RSProofOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RSProof *RSProofFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RSProofOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RSProofOwnershipTransferredIterator{contract: _RSProof.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RSProof *RSProofFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RSProofOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofOwnershipTransferred)
				if err := _RSProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RSProof *RSProofFilterer) ParseOwnershipTransferred(log types.Log) (*RSProofOwnershipTransferred, error) {
	event := new(RSProofOwnershipTransferred)
	if err := _RSProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofProveIterator is returned from FilterProve and is used to iterate over the raw logs and unpacked data for Prove events raised by the RSProof contract.
type RSProofProveIterator struct {
	Event *RSProofProve // Event containing the contract specifics and raw log

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
func (it *RSProofProveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofProve)
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
		it.Event = new(RSProofProve)
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
func (it *RSProofProveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofProveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofProve represents a Prove event raised by the RSProof contract.
type RSProofProve struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProve is a free log retrieval operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterProve(opts *bind.FilterOpts, _s []common.Address) (*RSProofProveIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofProveIterator{contract: _RSProof.contract, event: "Prove", logs: logs, sub: sub}, nil
}

// WatchProve is a free log subscription operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchProve(opts *bind.WatchOpts, sink chan<- *RSProofProve, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofProve)
				if err := _RSProof.contract.UnpackLog(event, "Prove", log); err != nil {
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

// ParseProve is a log parse operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseProve(log types.Log) (*RSProofProve, error) {
	event := new(RSProofProve)
	if err := _RSProof.contract.UnpackLog(event, "Prove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RSProof contract.
type RSProofUpgradedIterator struct {
	Event *RSProofUpgraded // Event containing the contract specifics and raw log

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
func (it *RSProofUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofUpgraded)
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
		it.Event = new(RSProofUpgraded)
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
func (it *RSProofUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofUpgraded represents a Upgraded event raised by the RSProof contract.
type RSProofUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RSProof *RSProofFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RSProofUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RSProofUpgradedIterator{contract: _RSProof.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RSProof *RSProofFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RSProofUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofUpgraded)
				if err := _RSProof.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_RSProof *RSProofFilterer) ParseUpgraded(log types.Log) (*RSProofUpgraded, error) {
	event := new(RSProofUpgraded)
	if err := _RSProof.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

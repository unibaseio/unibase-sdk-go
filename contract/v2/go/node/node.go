// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

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

// INodeNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type INodeNodeInfo struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeOf\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeInfoOf\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structINode.NodeInfo\",\"components\":[{\"name\":\"nodeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nodeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punish\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"money\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"_eproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rsproof\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"terminate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlock\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Punish\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_typ\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_money\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Set\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Terminated\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051611a4b6100f95f395f81816111400152818161116901526112a80152611a4b5ff3fe608060405260043610610161575f3560e01c80637eee288d116100cd578063c0c53b8b11610087578063dd752e5511610062578063dd752e55146105b0578063f2fde38b146105cf578063fc0c546a146105ee578063ffa1ad741461060c575f80fd5b8063c0c53b8b14610538578063d3748dc314610557578063d8c5115514610585575f80fd5b80637eee288d1461044357806381cc0c7a146104625780638da5cb5b14610481578063900cf0cf146104bd5780639748dcdc146104dc578063ad3cb1cc146104fb575f80fd5b806348ab5e6c1161011e57806348ab5e6c1461036e5780634f1ef2861461038d57806352d1902d146103a057806361e728b1146103c2578063715018a6146103f857806379ca7e0f1461040c575f80fd5b80630c08bf88146101655780630fdefd361461017b578063189a5a1714610281578063282d3fdf146103115780632e1a7d4d146103305780633b58524d1461034f575b5f80fd5b348015610170575f80fd5b5061017961063c565b005b348015610186575f80fd5b506102296101953660046116ab565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160a01b03165f90815260056020908152604091829020825160a081018452815460ff8082168352610100820416151593820193909352620100009092046001600160401b03169282019290925260018201546060820152600290910154608082015290565b60405161027891905f60a08201905060ff83511682526020830151151560208301526001600160401b036040840151166040830152606083015160608301526080830151608083015292915050565b60405180910390f35b34801561028c575f80fd5b506102d761029b3660046116ab565b60056020525f908152604090208054600182015460029092015460ff808316936101008404909116926201000090046001600160401b03169185565b6040805160ff909616865293151560208601526001600160401b03909216928401929092526060830191909152608082015260a001610278565b34801561031c575f80fd5b5061017961032b3660046116cb565b6107bb565b34801561033b575f80fd5b5061017961034a3660046116f3565b610876565b34801561035a575f80fd5b5061017961036936600461170a565b610abf565b348015610379575f80fd5b5061017961038836600461174b565b610af5565b61017961039b366004611779565b610b4e565b3480156103ab575f80fd5b506103b4610b6d565b604051908152602001610278565b3480156103cd575f80fd5b506103e16103dc366004611834565b610b88565b604080519215158352602083019190915201610278565b348015610403575f80fd5b50610179610be4565b348015610417575f80fd5b5060035461042b906001600160a01b031681565b6040516001600160a01b039091168152602001610278565b34801561044e575f80fd5b5061017961045d3660046116cb565b610bf7565b34801561046d575f80fd5b5060025461042b906001600160a01b031681565b34801561048c575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661042b565b3480156104c8575f80fd5b5060015461042b906001600160a01b031681565b3480156104e7575f80fd5b506101796104f636600461185c565b610c6c565b348015610506575f80fd5b5061052b604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516102789190611895565b348015610543575f80fd5b506101796105523660046118e1565b610e62565b348015610562575f80fd5b506103b4610571366004611921565b60ff165f9081526004602052604090205490565b348015610590575f80fd5b506103b461059f366004611921565b60046020525f908152604090205481565b3480156105bb575f80fd5b506101796105ca36600461174b565b610f8a565b3480156105da575f80fd5b506101796105e93660046116ab565b611068565b3480156105f9575f80fd5b505f5461042b906001600160a01b031681565b348015610617575f80fd5b5061052b604051806040016040528060058152602001640322e302e360dc1b81525081565b335f9081526005602052604090208054610100900460ff166106925760405162461bcd60e51b815260206004820152600a6024820152696e6f742061637469766560b01b60448201526064015b60405180910390fd5b805461ff00191681556001546040805163919840ad60e01b815290516001600160a01b039092169163919840ad916004808201925f9290919082900301818387803b1580156106df575f80fd5b505af11580156106f1573d5f803e3d5ffd5b5050505060015f9054906101000a90046001600160a01b03166001600160a01b0316639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610745573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610769919061193a565b81546001600160401b0391909116620100000269ffffffffffffffff00001990911617815560405133907f98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be905f90a250565b6002546001600160a01b03163314806107de57506003546001600160a01b031633145b6107fa5760405162461bcd60e51b815260040161068990611960565b6001600160a01b0382165f90815260056020526040812060028101805491928492610826908490611998565b90915550506001810154600282015411156108715760405162461bcd60e51b815260206004820152600b60248201526a1b1bd8dac8195e18d9595960aa1b6044820152606401610689565b505050565b335f90815260056020526040808220600154825163919840ad60e01b8152925191936001600160a01b039091169263919840ad926004808301939282900301818387803b1580156108c5575f80fd5b505af11580156108d7573d5f803e3d5ffd5b50508254610100900460ff1691508190506109935750805461090a906007906201000090046001600160401b03166119ab565b6001600160401b031660015f9054906101000a90046001600160a01b03166001600160a01b0316639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610963573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610987919061193a565b6001600160401b031611155b15610a08576002810154815460ff165f908152600460205260409020546109ba9084611998565b6109c49190611998565b81600101541015610a035760405162461bcd60e51b8152602060048201526009602482015268696e73756620706c6560b81b6044820152606401610689565b610a58565b6002810154610a179083611998565b81600101541015610a585760405162461bcd60e51b815260206004820152600b60248201526a1a5b9cdd59881c1b19481d60aa1b6044820152606401610689565b81816001015f828254610a6b91906119cb565b90915550505f54610a86906001600160a01b031633846110a5565b60405182815233907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d59060200160405180910390a25050565b610ac76110da565b600280546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055565b610afd6110da565b60ff82165f81815260046020908152604091829020849055815192835282018390527fc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881910160405180910390a15050565b610b56611135565b610b5f826111d9565b610b6982826111e1565b5050565b5f610b7661129d565b505f805160206119f683398151915290565b6001600160a01b0382165f9081526005602052604081208054829190610100900460ff161580610bbf5750805460ff858116911614155b15610bd0575f809250925050610bdd565b6001816001015492509250505b9250929050565b610bec6110da565b610bf55f6112e6565b565b6002546001600160a01b0316331480610c1a57506003546001600160a01b031633145b610c365760405162461bcd60e51b815260040161068990611960565b6001600160a01b0382165f90815260056020526040812060028101805491928492610c629084906119cb565b9091555050505050565b6002546001600160a01b0316331480610c8f57506003546001600160a01b031633145b610cab5760405162461bcd60e51b815260040161068990611960565b6001600160a01b0383165f90815260056020526040812060018101805491928492610cd79084906119cb565b90915550505f54610cf2906001600160a01b031684846110a5565b805460ff165f9081526004602052604090205460018201541015610e0d57805461ff00191681556001546040805163919840ad60e01b815290516001600160a01b039092169163919840ad916004808201925f9290919082900301818387803b158015610d5d575f80fd5b505af1158015610d6f573d5f803e3d5ffd5b5050505060015f9054906101000a90046001600160a01b03166001600160a01b0316639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de7919061193a565b81546001600160401b0391909116620100000269ffffffffffffffff0000199091161781555b80546040805160ff9092168252602082018490526001600160a01b0385811692908716917fc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32910160405180910390a350505050565b5f610e6b611356565b805490915060ff600160401b82041615906001600160401b03165f81158015610e915750825b90505f826001600160401b03166001148015610eac5750303b155b905081158015610eba575080155b15610ed85760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610f0257845460ff60401b1916600160401b1785555b610f0b86611380565b5f80546001600160a01b03808b166001600160a01b03199283161790925560018054928a16929091169190911790558315610f8057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f54610fa1906001600160a01b0316333084611391565b335f9081526005602052604090208054610100900460ff16158015610fe6575060ff83165f908152600460205260409020546001820154610fe3908490611998565b10155b1561100d57805469ffffffffffffffff00001961ffff1990911660ff851617610100171681555b81816001015f8282546110209190611998565b90915550506040805160ff851681526020810184905233917f3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf910160405180910390a2505050565b6110706110da565b6001600160a01b03811661109957604051631e4fbdf760e01b81525f6004820152602401610689565b6110a2816112e6565b50565b6110b283838360016113cd565b61087157604051635274afe760e01b81526001600160a01b0384166004820152602401610689565b3361110c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610bf55760405163118cdaa760e01b8152336004820152602401610689565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806111bb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166111af5f805160206119f6833981519152546001600160a01b031690565b6001600160a01b031614155b15610bf55760405163703e46dd60e11b815260040160405180910390fd5b6110a26110da565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561123b575060408051601f3d908101601f19168201909252611238918101906119de565b60015b61126357604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610689565b5f805160206119f6833981519152811461129357604051632a87526960e21b815260048101829052602401610689565b610871838361142f565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610bf55760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b611388611484565b6110a2816114a9565b61139f8484848460016114b1565b6113c757604051635274afe760e01b81526001600160a01b0385166004820152602401610689565b50505050565b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316611423578383151615611417573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b6114388261151e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561147c576108718282611581565b610b69611621565b61148c611640565b610bf557604051631afcd79f60e31b815260040160405180910390fd5b611070611484565b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f5114831661150d578383151615611501573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b806001600160a01b03163b5f0361155357604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610689565b5f805160206119f683398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f61158e8484611659565b90508080156115af57505f3d11806115af57505f846001600160a01b03163b115b156115c4576115bc61166c565b91505061137a565b80156115ee57604051639996b31560e01b81526001600160a01b0385166004820152602401610689565b3d15611601576115fc611685565b61161a565b60405163d6bda27560e01b815260040160405180910390fd5b5092915050565b3415610bf55760405163b398979f60e01b815260040160405180910390fd5b5f611649611356565b54600160401b900460ff16919050565b5f805f835160208501865af49392505050565b6040513d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b80356001600160a01b03811681146116a6575f80fd5b919050565b5f602082840312156116bb575f80fd5b6116c482611690565b9392505050565b5f80604083850312156116dc575f80fd5b6116e583611690565b946020939093013593505050565b5f60208284031215611703575f80fd5b5035919050565b5f806040838503121561171b575f80fd5b61172483611690565b915061173260208401611690565b90509250929050565b803560ff811681146116a6575f80fd5b5f806040838503121561175c575f80fd5b6116e58361173b565b634e487b7160e01b5f52604160045260245ffd5b5f806040838503121561178a575f80fd5b61179383611690565b915060208301356001600160401b03808211156117ae575f80fd5b818501915085601f8301126117c1575f80fd5b8135818111156117d3576117d3611765565b604051601f8201601f19908116603f011681019083821181831017156117fb576117fb611765565b81604052828152886020848701011115611813575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f8060408385031215611845575f80fd5b61184e83611690565b91506117326020840161173b565b5f805f6060848603121561186e575f80fd5b61187784611690565b925061188560208501611690565b9150604084013590509250925092565b5f602080835283518060208501525f5b818110156118c1578581018301518582016040015282016118a5565b505f604082860101526040601f19601f8301168501019250505092915050565b5f805f606084860312156118f3575f80fd5b6118fc84611690565b925061190a60208501611690565b915061191860408501611690565b90509250925092565b5f60208284031215611931575f80fd5b6116c48261173b565b5f6020828403121561194a575f80fd5b81516001600160401b03811681146116c4575f80fd5b6020808252600a908201526934b73b1031b0b63632b960b11b604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561137a5761137a611984565b6001600160401b0381811683821601908082111561161a5761161a611984565b8181038181111561137a5761137a611984565b5f602082840312156119ee575f80fd5b505191905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122084d3dd5a0bd0712dc9ffdb66791e093fcd9568a9092535ba0e38757e06418b9464736f6c63430008160033",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// NodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeMetaData.Bin instead.
var NodeBin = NodeMetaData.Bin

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Node.Contract.UPGRADEINTERFACEVERSION(&_Node.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Node.Contract.UPGRADEINTERFACEVERSION(&_Node.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeSession) VERSION() (string, error) {
	return _Node.Contract.VERSION(&_Node.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeCallerSession) VERSION() (string, error) {
	return _Node.Contract.VERSION(&_Node.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) (bool, *big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeSession) Check(_a common.Address, _type uint8) (bool, *big.Int, error) {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeCallerSession) Check(_a common.Address, _type uint8) (bool, *big.Int, error) {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeSession) Epoch() (common.Address, error) {
	return _Node.Contract.Epoch(&_Node.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeCallerSession) Epoch() (common.Address, error) {
	return _Node.Contract.Epoch(&_Node.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeCaller) Eproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "eproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeSession) Eproof() (common.Address, error) {
	return _Node.Contract.Eproof(&_Node.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeCallerSession) Eproof() (common.Address, error) {
	return _Node.Contract.Eproof(&_Node.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeCaller) MinStake(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "minStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeSession) MinStake(arg0 uint8) (*big.Int, error) {
	return _Node.Contract.MinStake(&_Node.CallOpts, arg0)
}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeCallerSession) MinStake(arg0 uint8) (*big.Int, error) {
	return _Node.Contract.MinStake(&_Node.CallOpts, arg0)
}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeCaller) MinStakeOf(opts *bind.CallOpts, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "minStakeOf", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeSession) MinStakeOf(_type uint8) (*big.Int, error) {
	return _Node.Contract.MinStakeOf(&_Node.CallOpts, _type)
}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) MinStakeOf(_type uint8) (*big.Int, error) {
	return _Node.Contract.MinStakeOf(&_Node.CallOpts, _type)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeCaller) NodeInfoOf(opts *bind.CallOpts, a common.Address) (INodeNodeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "nodeInfoOf", a)

	if err != nil {
		return *new(INodeNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodeNodeInfo)).(*INodeNodeInfo)

	return out0, err

}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeSession) NodeInfoOf(a common.Address) (INodeNodeInfo, error) {
	return _Node.Contract.NodeInfoOf(&_Node.CallOpts, a)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeCallerSession) NodeInfoOf(a common.Address) (INodeNodeInfo, error) {
	return _Node.Contract.NodeInfoOf(&_Node.CallOpts, a)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		NodeType     uint8
		IsActive     bool
		ExitEpoch    uint64
		StakedAmount *big.Int
		LockedAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ExitEpoch = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.StakedAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeSession) Nodes(arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	return _Node.Contract.Nodes(&_Node.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeCallerSession) Nodes(arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	return _Node.Contract.Nodes(&_Node.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCallerSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeSession) ProxiableUUID() ([32]byte, error) {
	return _Node.Contract.ProxiableUUID(&_Node.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Node.Contract.ProxiableUUID(&_Node.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeCaller) Rsproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "rsproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeSession) Rsproof() (common.Address, error) {
	return _Node.Contract.Rsproof(&_Node.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeCallerSession) Rsproof() (common.Address, error) {
	return _Node.Contract.Rsproof(&_Node.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeSession) Token() (common.Address, error) {
	return _Node.Contract.Token(&_Node.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeCallerSession) Token() (common.Address, error) {
	return _Node.Contract.Token(&_Node.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "initialize", _token, _epoch, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeSession) Initialize(_token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _token, _epoch, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeTransactorSession) Initialize(_token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _token, _epoch, initialOwner)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeTransactor) Lock(opts *bind.TransactOpts, _from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "lock", _from, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeSession) Lock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Lock(&_Node.TransactOpts, _from, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeTransactorSession) Lock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Lock(&_Node.TransactOpts, _from, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "punish", _from, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeSession) Punish(_from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeTransactorSession) Punish(_from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _to, _m)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeTransactor) Set(opts *bind.TransactOpts, _type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "set", _type, money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeSession) Set(_type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeTransactorSession) Set(_type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, money)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeTransactor) SetAddress(opts *bind.TransactOpts, _eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setAddress", _eproof, _rsproof)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeSession) SetAddress(_eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetAddress(&_Node.TransactOpts, _eproof, _rsproof)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeTransactorSession) SetAddress(_eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetAddress(&_Node.TransactOpts, _eproof, _rsproof)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeTransactor) Stake(opts *bind.TransactOpts, _type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "stake", _type, m)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeSession) Stake(_type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Stake(&_Node.TransactOpts, _type, m)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeTransactorSession) Stake(_type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Stake(&_Node.TransactOpts, _type, m)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeTransactor) Terminate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeSession) Terminate() (*types.Transaction, error) {
	return _Node.Contract.Terminate(&_Node.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeTransactorSession) Terminate() (*types.Transaction, error) {
	return _Node.Contract.Terminate(&_Node.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeTransactor) Unlock(opts *bind.TransactOpts, _from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "unlock", _from, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeSession) Unlock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Unlock(&_Node.TransactOpts, _from, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeTransactorSession) Unlock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Unlock(&_Node.TransactOpts, _from, _m)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.Contract.UpgradeToAndCall(&_Node.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.Contract.UpgradeToAndCall(&_Node.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeTransactor) Withdraw(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "withdraw", m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeSession) Withdraw(m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeTransactorSession) Withdraw(m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, m)
}

// NodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Node contract.
type NodeInitializedIterator struct {
	Event *NodeInitialized // Event containing the contract specifics and raw log

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
func (it *NodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeInitialized)
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
		it.Event = new(NodeInitialized)
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
func (it *NodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeInitialized represents a Initialized event raised by the Node contract.
type NodeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Node *NodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeInitializedIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeInitializedIterator{contract: _Node.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Node *NodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeInitialized) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeInitialized)
				if err := _Node.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Node *NodeFilterer) ParseInitialized(log types.Log) (*NodeInitialized, error) {
	event := new(NodeInitialized)
	if err := _Node.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Node contract.
type NodeOwnershipTransferredIterator struct {
	Event *NodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnershipTransferred)
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
		it.Event = new(NodeOwnershipTransferred)
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
func (it *NodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnershipTransferred represents a OwnershipTransferred event raised by the Node contract.
type NodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnershipTransferredIterator{contract: _Node.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnershipTransferred)
				if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Node *NodeFilterer) ParseOwnershipTransferred(log types.Log) (*NodeOwnershipTransferred, error) {
	event := new(NodeOwnershipTransferred)
	if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Node contract.
type NodePunishIterator struct {
	Event *NodePunish // Event containing the contract specifics and raw log

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
func (it *NodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePunish)
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
		it.Event = new(NodePunish)
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
func (it *NodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePunish represents a Punish event raised by the Node contract.
type NodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*NodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &NodePunishIterator{contract: _Node.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *NodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePunish)
				if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) ParsePunish(log types.Log) (*NodePunish, error) {
	event := new(NodePunish)
	if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Node contract.
type NodeSetIterator struct {
	Event *NodeSet // Event containing the contract specifics and raw log

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
func (it *NodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeSet)
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
		it.Event = new(NodeSet)
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
func (it *NodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeSet represents a Set event raised by the Node contract.
type NodeSet struct {
	Type uint8
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) FilterSet(opts *bind.FilterOpts) (*NodeSetIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &NodeSetIterator{contract: _Node.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *NodeSet) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeSet)
				if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) ParseSet(log types.Log) (*NodeSet, error) {
	event := new(NodeSet)
	if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Node contract.
type NodeStakedIterator struct {
	Event *NodeStaked // Event containing the contract specifics and raw log

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
func (it *NodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStaked)
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
		it.Event = new(NodeStaked)
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
func (it *NodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStaked represents a Staked event raised by the Node contract.
type NodeStaked struct {
	A    common.Address
	Type uint8
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) FilterStaked(opts *bind.FilterOpts, a []common.Address) (*NodeStakedIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Staked", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakedIterator{contract: _Node.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *NodeStaked, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Staked", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStaked)
				if err := _Node.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) ParseStaked(log types.Log) (*NodeStaked, error) {
	event := new(NodeStaked)
	if err := _Node.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeTerminatedIterator is returned from FilterTerminated and is used to iterate over the raw logs and unpacked data for Terminated events raised by the Node contract.
type NodeTerminatedIterator struct {
	Event *NodeTerminated // Event containing the contract specifics and raw log

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
func (it *NodeTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeTerminated)
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
		it.Event = new(NodeTerminated)
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
func (it *NodeTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeTerminated represents a Terminated event raised by the Node contract.
type NodeTerminated struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTerminated is a free log retrieval operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) FilterTerminated(opts *bind.FilterOpts, a []common.Address) (*NodeTerminatedIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Terminated", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeTerminatedIterator{contract: _Node.contract, event: "Terminated", logs: logs, sub: sub}, nil
}

// WatchTerminated is a free log subscription operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) WatchTerminated(opts *bind.WatchOpts, sink chan<- *NodeTerminated, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Terminated", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeTerminated)
				if err := _Node.contract.UnpackLog(event, "Terminated", log); err != nil {
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

// ParseTerminated is a log parse operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) ParseTerminated(log types.Log) (*NodeTerminated, error) {
	event := new(NodeTerminated)
	if err := _Node.contract.UnpackLog(event, "Terminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Node contract.
type NodeUpgradedIterator struct {
	Event *NodeUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeUpgraded)
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
		it.Event = new(NodeUpgraded)
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
func (it *NodeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeUpgraded represents a Upgraded event raised by the Node contract.
type NodeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Node *NodeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodeUpgradedIterator{contract: _Node.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Node *NodeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeUpgraded)
				if err := _Node.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Node *NodeFilterer) ParseUpgraded(log types.Log) (*NodeUpgraded, error) {
	event := new(NodeUpgraded)
	if err := _Node.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Node contract.
type NodeWithdrawnIterator struct {
	Event *NodeWithdrawn // Event containing the contract specifics and raw log

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
func (it *NodeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeWithdrawn)
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
		it.Event = new(NodeWithdrawn)
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
func (it *NodeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeWithdrawn represents a Withdrawn event raised by the Node contract.
type NodeWithdrawn struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) FilterWithdrawn(opts *bind.FilterOpts, a []common.Address) (*NodeWithdrawnIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Withdrawn", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeWithdrawnIterator{contract: _Node.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NodeWithdrawn, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Withdrawn", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeWithdrawn)
				if err := _Node.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) ParseWithdrawn(log types.Log) (*NodeWithdrawn, error) {
	event := new(NodeWithdrawn)
	if err := _Node.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

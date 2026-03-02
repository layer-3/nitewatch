// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custody

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

// SimpleCustodyMetaData contains all meta data concerning the SimpleCustody contract.
var SimpleCustodyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neodax\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nitewatch\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NEODAX_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NITEWATCH_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rejectWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFinalized\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawStarted\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ETHTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x6080346100b057601f61166138819003918201601f19168301916001600160401b038311848410176100b4578084926060946040528339810103126100b0578061009a61004e6100a0936100c8565b9161009461006a6040610063602085016100c8565b93016100c8565b9360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00556100dc565b50610152565b506101e5565b5060405161136890816102798239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100b057565b6001600160a01b0381165f9081525f5160206116415f395f51905f52602052604090205460ff1661014d576001600160a01b03165f8181525f5160206116415f395f51905f5260205260408120805460ff191660011790553391905f5160206115e15f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f5160206116015f395f51905f52602052604090205460ff1661014d576001600160a01b03165f8181525f5160206116015f395f51905f5260205260408120805460ff191660011790553391907f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f3905f5160206115e15f395f51905f529080a4600190565b6001600160a01b0381165f9081525f5160206116215f395f51905f52602052604090205460ff1661014d576001600160a01b03165f8181525f5160206116215f395f51905f5260205260408120805460ff191660011790553391907ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa905f5160206115e15f395f51905f529080a460019056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714610ef05750806305e95be714610b9b57806311edc78f14610a6b578063248a9ca314610a1b5780632f2ff15d146109c057806336568abe1461093857806347e7ef24146106815780635a98c2231461062957806391d14854146105b5578063a217fddf1461057d578063d547741f1461051b578063d87e1f41146101ba578063da86f315146101625763efbf64a7146100b6575f80fd5b3461015e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e576004355f52600160205260a060405f2060ff73ffffffffffffffffffffffffffffffffffffffff8254169173ffffffffffffffffffffffffffffffffffffffff600182015416906003600282015491015491604051948552602085015260408401528181161515606084015260081c1615156080820152f35b5f80fd5b3461015e575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5760206040517f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f38152f35b3461015e5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e576101f1610fcf565b6101f9610fac565b335f9081527fb62d236122444fe854e95565ed2ab7440b5f2c079b533c05c79c4ef91291a45860205260409020546044359260643592909160ff16156104cb57610241611121565b83156104a3576040519173ffffffffffffffffffffffffffffffffffffffff806020850192468452306040870152169283606086015216928360808201528560a08201528460c082015260c0815261029a60e082610ff2565b51902092835f52600160205260ff600360405f2001541661047b5760405160a0810181811067ffffffffffffffff82111761044e5760209686937f669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb98993604093845286815289810188815260038583019285845260608101936001855273ffffffffffffffffffffffffffffffffffffffff60808301945f86528b5f528f6001905281808b5f20955116167fffffffffffffffffffffffff0000000000000000000000000000000000000000855416178455511673ffffffffffffffffffffffffffffffffffffffff6001840191167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055516002820155019151151560ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084541691161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff61ff0083549260081b169116179055825191825288820152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055604051908152f35b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f157c65e1000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f360245260445ffd5b3461015e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5761057b600435610558610fac565b90610576610571825f525f602052600160405f20015490565b6110bb565b61126a565b005b3461015e575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5760206040515f8152f35b3461015e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e576105ec610fac565b6004355f525f60205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b3461015e575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5760206040517ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa8152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e576106b3610fcf565b602435906106bf611121565b81156104a35773ffffffffffffffffffffffffffffffffffffffff1690808261076357340361073b575b6040519081527f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a760203392a360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7f1841b4e1000000000000000000000000000000000000000000000000000000005f5260045ffd5b90503461073b576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481865afa9081156108b0575f91610906575b50604051917f23b872dd000000000000000000000000000000000000000000000000000000005f52336004523060245260445260205f60648180875af160015f51148116156108e7575b826040525f606052156108bb577f70a08231000000000000000000000000000000000000000000000000000000008252306004830152602082602481865afa9182156108b0575f9261087c575b508103908111156106e9577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9091506020813d6020116108a8575b8161089860209383610ff2565b8101031261015e57519083610844565b3d915061088b565b6040513d5f823e3d90fd5b827f5274afe7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60018115166108fd57833b15153d1516166107f7565b823d5f823e3d90fd5b90506020813d602011610930575b8161092160209383610ff2565b8101031261015e5751836107ad565b3d9150610914565b3461015e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5761096f610fac565b3373ffffffffffffffffffffffffffffffffffffffff8216036109985761057b9060043561126a565b7f6697b232000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461015e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e5761057b6004356109fd610fac565b90610a16610571825f525f602052600160405f20015490565b611198565b3461015e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e576020610a636004355f525f602052600160405f20015490565b604051908152f35b3461015e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e57600435610aa5611033565b610aad611121565b805f526001602052600360405f2001805460ff811615610b735760ff8160081c16610b4b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff61010091161790557f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c60206040515f8152a260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7fae899454000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f8d0fc1dd000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461015e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e57600435610bd5611033565b610bdd611121565b805f52600160205260405f2060038101805460ff811615610b735760ff8160081c16610b4b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff610100911617905573ffffffffffffffffffffffffffffffffffffffff815416600182015f73ffffffffffffffffffffffffffffffffffffffff8254169160028501908154957fffffffffffffffffffffffff000000000000000000000000000000000000000081541690557fffffffffffffffffffffffff000000000000000000000000000000000000000081541690555580155f14610dd15750814710610da9575f80809381935af13d15610da4573d67ffffffffffffffff811161044e5760405190610d1b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160183610ff2565b81525f60203d92013e5b15610d7c575b7f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c602060405160018152a260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7fb12d13eb000000000000000000000000000000000000000000000000000000005f5260045ffd5b610d25565b7fbb55fd27000000000000000000000000000000000000000000000000000000005f5260045ffd5b916040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481875afa80156108b05782915f91610ebb575b5010610da957604051917fa9059cbb000000000000000000000000000000000000000000000000000000005f5260045260245260205f60448180865af19060015f5114821615610e9a575b60405215610e6f5750610d2b565b7f5274afe7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b906001811516610eb257823b15153d15161690610e61565b503d5f823e3d90fd5b9150506020813d602011610ee8575b81610ed760209383610ff2565b8101031261015e5781905186610e16565b3d9150610eca565b3461015e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261015e57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361015e57817f7965db0b0000000000000000000000000000000000000000000000000000000060209314908115610f82575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483610f7b565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361015e57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361015e57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761044e57604052565b335f9081527fbebdf9b9a881bc7cd335909e0de2b4f6f0388553854a22ac02a758c661d9c440602052604090205460ff161561106b57565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f52336004527ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa60245260445ffd5b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f205416156110f25750565b7fe2517d3f000000000000000000000000000000000000000000000000000000005f523360045260245260445ffd5b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0054146111705760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f3ee5aeb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f205416155f1461126457805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260ff60405f2054165f1461126457805f525f60205260405f2073ffffffffffffffffffffffffffffffffffffffff83165f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a460019056fea264697066735822122094bb890eca2b27a1a5c58441aec7efb21cc1949cbbd46788a98ba57bc28cf6ee64736f6c634300081e00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0db62d236122444fe854e95565ed2ab7440b5f2c079b533c05c79c4ef91291a458bebdf9b9a881bc7cd335909e0de2b4f6f0388553854a22ac02a758c661d9c440ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// SimpleCustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleCustodyMetaData.ABI instead.
var SimpleCustodyABI = SimpleCustodyMetaData.ABI

// SimpleCustodyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleCustodyMetaData.Bin instead.
var SimpleCustodyBin = SimpleCustodyMetaData.Bin

// DeploySimpleCustody deploys a new Ethereum contract, binding an instance of SimpleCustody to it.
func DeploySimpleCustody(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, neodax common.Address, nitewatch common.Address) (common.Address, *types.Transaction, *SimpleCustody, error) {
	parsed, err := SimpleCustodyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleCustodyBin), backend, admin, neodax, nitewatch)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleCustody{SimpleCustodyCaller: SimpleCustodyCaller{contract: contract}, SimpleCustodyTransactor: SimpleCustodyTransactor{contract: contract}, SimpleCustodyFilterer: SimpleCustodyFilterer{contract: contract}}, nil
}

// SimpleCustody is an auto generated Go binding around an Ethereum contract.
type SimpleCustody struct {
	SimpleCustodyCaller     // Read-only binding to the contract
	SimpleCustodyTransactor // Write-only binding to the contract
	SimpleCustodyFilterer   // Log filterer for contract events
}

// SimpleCustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleCustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleCustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleCustodySession struct {
	Contract     *SimpleCustody    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCustodyCallerSession struct {
	Contract *SimpleCustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleCustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleCustodyTransactorSession struct {
	Contract     *SimpleCustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleCustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleCustodyRaw struct {
	Contract *SimpleCustody // Generic contract binding to access the raw methods on
}

// SimpleCustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCustodyCallerRaw struct {
	Contract *SimpleCustodyCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleCustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleCustodyTransactorRaw struct {
	Contract *SimpleCustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleCustody creates a new instance of SimpleCustody, bound to a specific deployed contract.
func NewSimpleCustody(address common.Address, backend bind.ContractBackend) (*SimpleCustody, error) {
	contract, err := bindSimpleCustody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleCustody{SimpleCustodyCaller: SimpleCustodyCaller{contract: contract}, SimpleCustodyTransactor: SimpleCustodyTransactor{contract: contract}, SimpleCustodyFilterer: SimpleCustodyFilterer{contract: contract}}, nil
}

// NewSimpleCustodyCaller creates a new read-only instance of SimpleCustody, bound to a specific deployed contract.
func NewSimpleCustodyCaller(address common.Address, caller bind.ContractCaller) (*SimpleCustodyCaller, error) {
	contract, err := bindSimpleCustody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyCaller{contract: contract}, nil
}

// NewSimpleCustodyTransactor creates a new write-only instance of SimpleCustody, bound to a specific deployed contract.
func NewSimpleCustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleCustodyTransactor, error) {
	contract, err := bindSimpleCustody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyTransactor{contract: contract}, nil
}

// NewSimpleCustodyFilterer creates a new log filterer instance of SimpleCustody, bound to a specific deployed contract.
func NewSimpleCustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleCustodyFilterer, error) {
	contract, err := bindSimpleCustody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyFilterer{contract: contract}, nil
}

// bindSimpleCustody binds a generic wrapper to an already deployed contract.
func bindSimpleCustody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleCustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCustody *SimpleCustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCustody.Contract.SimpleCustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCustody *SimpleCustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCustody.Contract.SimpleCustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCustody *SimpleCustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCustody.Contract.SimpleCustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCustody *SimpleCustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCustody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCustody *SimpleCustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCustody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCustody *SimpleCustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCustody.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.DEFAULTADMINROLE(&_SimpleCustody.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.DEFAULTADMINROLE(&_SimpleCustody.CallOpts)
}

// NEODAXROLE is a free data retrieval call binding the contract method 0xda86f315.
//
// Solidity: function NEODAX_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCaller) NEODAXROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "NEODAX_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NEODAXROLE is a free data retrieval call binding the contract method 0xda86f315.
//
// Solidity: function NEODAX_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodySession) NEODAXROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.NEODAXROLE(&_SimpleCustody.CallOpts)
}

// NEODAXROLE is a free data retrieval call binding the contract method 0xda86f315.
//
// Solidity: function NEODAX_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCallerSession) NEODAXROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.NEODAXROLE(&_SimpleCustody.CallOpts)
}

// NITEWATCHROLE is a free data retrieval call binding the contract method 0x5a98c223.
//
// Solidity: function NITEWATCH_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCaller) NITEWATCHROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "NITEWATCH_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NITEWATCHROLE is a free data retrieval call binding the contract method 0x5a98c223.
//
// Solidity: function NITEWATCH_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodySession) NITEWATCHROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.NITEWATCHROLE(&_SimpleCustody.CallOpts)
}

// NITEWATCHROLE is a free data retrieval call binding the contract method 0x5a98c223.
//
// Solidity: function NITEWATCH_ROLE() view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCallerSession) NITEWATCHROLE() ([32]byte, error) {
	return _SimpleCustody.Contract.NITEWATCHROLE(&_SimpleCustody.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleCustody *SimpleCustodySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SimpleCustody.Contract.GetRoleAdmin(&_SimpleCustody.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleCustody *SimpleCustodyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SimpleCustody.Contract.GetRoleAdmin(&_SimpleCustody.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleCustody *SimpleCustodyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleCustody *SimpleCustodySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SimpleCustody.Contract.HasRole(&_SimpleCustody.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleCustody *SimpleCustodyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SimpleCustody.Contract.HasRole(&_SimpleCustody.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleCustody *SimpleCustodyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleCustody *SimpleCustodySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleCustody.Contract.SupportsInterface(&_SimpleCustody.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleCustody *SimpleCustodyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleCustody.Contract.SupportsInterface(&_SimpleCustody.CallOpts, interfaceId)
}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 id) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodyCaller) Withdrawals(opts *bind.CallOpts, id [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "withdrawals", id)

	outstruct := new(struct {
		User      common.Address
		Token     common.Address
		Amount    *big.Int
		Exists    bool
		Finalized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Finalized = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 id) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodySession) Withdrawals(id [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	return _SimpleCustody.Contract.Withdrawals(&_SimpleCustody.CallOpts, id)
}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 id) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodyCallerSession) Withdrawals(id [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	return _SimpleCustody.Contract.Withdrawals(&_SimpleCustody.CallOpts, id)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_SimpleCustody *SimpleCustodyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_SimpleCustody *SimpleCustodySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.Contract.Deposit(&_SimpleCustody.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.Contract.Deposit(&_SimpleCustody.TransactOpts, token, amount)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodyTransactor) FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "finalizeWithdraw", withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodySession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.Contract.FinalizeWithdraw(&_SimpleCustody.TransactOpts, withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.Contract.FinalizeWithdraw(&_SimpleCustody.TransactOpts, withdrawalId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.GrantRole(&_SimpleCustody.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.GrantRole(&_SimpleCustody.TransactOpts, role, account)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodyTransactor) RejectWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "rejectWithdraw", withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodySession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RejectWithdraw(&_SimpleCustody.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RejectWithdraw(&_SimpleCustody.TransactOpts, withdrawalId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SimpleCustody *SimpleCustodyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SimpleCustody *SimpleCustodySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RenounceRole(&_SimpleCustody.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RenounceRole(&_SimpleCustody.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RevokeRole(&_SimpleCustody.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleCustody *SimpleCustodyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleCustody.Contract.RevokeRole(&_SimpleCustody.TransactOpts, role, account)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_SimpleCustody *SimpleCustodyTransactor) StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.contract.Transact(opts, "startWithdraw", user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_SimpleCustody *SimpleCustodySession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.Contract.StartWithdraw(&_SimpleCustody.TransactOpts, user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32 withdrawalId)
func (_SimpleCustody *SimpleCustodyTransactorSession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _SimpleCustody.Contract.StartWithdraw(&_SimpleCustody.TransactOpts, user, token, amount, nonce)
}

// SimpleCustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the SimpleCustody contract.
type SimpleCustodyDepositedIterator struct {
	Event *SimpleCustodyDeposited // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyDeposited)
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
		it.Event = new(SimpleCustodyDeposited)
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
func (it *SimpleCustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyDeposited represents a Deposited event raised by the SimpleCustody contract.
type SimpleCustodyDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_SimpleCustody *SimpleCustodyFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*SimpleCustodyDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyDepositedIterator{contract: _SimpleCustody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_SimpleCustody *SimpleCustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *SimpleCustodyDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyDeposited)
				if err := _SimpleCustody.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_SimpleCustody *SimpleCustodyFilterer) ParseDeposited(log types.Log) (*SimpleCustodyDeposited, error) {
	event := new(SimpleCustodyDeposited)
	if err := _SimpleCustody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCustodyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SimpleCustody contract.
type SimpleCustodyRoleAdminChangedIterator struct {
	Event *SimpleCustodyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyRoleAdminChanged)
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
		it.Event = new(SimpleCustodyRoleAdminChanged)
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
func (it *SimpleCustodyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyRoleAdminChanged represents a RoleAdminChanged event raised by the SimpleCustody contract.
type SimpleCustodyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleCustody *SimpleCustodyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SimpleCustodyRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyRoleAdminChangedIterator{contract: _SimpleCustody.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleCustody *SimpleCustodyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SimpleCustodyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyRoleAdminChanged)
				if err := _SimpleCustody.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleCustody *SimpleCustodyFilterer) ParseRoleAdminChanged(log types.Log) (*SimpleCustodyRoleAdminChanged, error) {
	event := new(SimpleCustodyRoleAdminChanged)
	if err := _SimpleCustody.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCustodyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SimpleCustody contract.
type SimpleCustodyRoleGrantedIterator struct {
	Event *SimpleCustodyRoleGranted // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyRoleGranted)
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
		it.Event = new(SimpleCustodyRoleGranted)
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
func (it *SimpleCustodyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyRoleGranted represents a RoleGranted event raised by the SimpleCustody contract.
type SimpleCustodyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SimpleCustodyRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyRoleGrantedIterator{contract: _SimpleCustody.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SimpleCustodyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyRoleGranted)
				if err := _SimpleCustody.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) ParseRoleGranted(log types.Log) (*SimpleCustodyRoleGranted, error) {
	event := new(SimpleCustodyRoleGranted)
	if err := _SimpleCustody.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCustodyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SimpleCustody contract.
type SimpleCustodyRoleRevokedIterator struct {
	Event *SimpleCustodyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyRoleRevoked)
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
		it.Event = new(SimpleCustodyRoleRevoked)
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
func (it *SimpleCustodyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyRoleRevoked represents a RoleRevoked event raised by the SimpleCustody contract.
type SimpleCustodyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SimpleCustodyRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyRoleRevokedIterator{contract: _SimpleCustody.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SimpleCustodyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyRoleRevoked)
				if err := _SimpleCustody.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleCustody *SimpleCustodyFilterer) ParseRoleRevoked(log types.Log) (*SimpleCustodyRoleRevoked, error) {
	event := new(SimpleCustodyRoleRevoked)
	if err := _SimpleCustody.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCustodyWithdrawFinalizedIterator is returned from FilterWithdrawFinalized and is used to iterate over the raw logs and unpacked data for WithdrawFinalized events raised by the SimpleCustody contract.
type SimpleCustodyWithdrawFinalizedIterator struct {
	Event *SimpleCustodyWithdrawFinalized // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyWithdrawFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyWithdrawFinalized)
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
		it.Event = new(SimpleCustodyWithdrawFinalized)
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
func (it *SimpleCustodyWithdrawFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyWithdrawFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyWithdrawFinalized represents a WithdrawFinalized event raised by the SimpleCustody contract.
type SimpleCustodyWithdrawFinalized struct {
	WithdrawalId [32]byte
	Success      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFinalized is a free log retrieval operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_SimpleCustody *SimpleCustodyFilterer) FilterWithdrawFinalized(opts *bind.FilterOpts, withdrawalId [][32]byte) (*SimpleCustodyWithdrawFinalizedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyWithdrawFinalizedIterator{contract: _SimpleCustody.contract, event: "WithdrawFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawFinalized is a free log subscription operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_SimpleCustody *SimpleCustodyFilterer) WatchWithdrawFinalized(opts *bind.WatchOpts, sink chan<- *SimpleCustodyWithdrawFinalized, withdrawalId [][32]byte) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyWithdrawFinalized)
				if err := _SimpleCustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
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

// ParseWithdrawFinalized is a log parse operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_SimpleCustody *SimpleCustodyFilterer) ParseWithdrawFinalized(log types.Log) (*SimpleCustodyWithdrawFinalized, error) {
	event := new(SimpleCustodyWithdrawFinalized)
	if err := _SimpleCustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCustodyWithdrawStartedIterator is returned from FilterWithdrawStarted and is used to iterate over the raw logs and unpacked data for WithdrawStarted events raised by the SimpleCustody contract.
type SimpleCustodyWithdrawStartedIterator struct {
	Event *SimpleCustodyWithdrawStarted // Event containing the contract specifics and raw log

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
func (it *SimpleCustodyWithdrawStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCustodyWithdrawStarted)
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
		it.Event = new(SimpleCustodyWithdrawStarted)
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
func (it *SimpleCustodyWithdrawStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCustodyWithdrawStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCustodyWithdrawStarted represents a WithdrawStarted event raised by the SimpleCustody contract.
type SimpleCustodyWithdrawStarted struct {
	WithdrawalId [32]byte
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	Nonce        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawStarted is a free log retrieval operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_SimpleCustody *SimpleCustodyFilterer) FilterWithdrawStarted(opts *bind.FilterOpts, withdrawalId [][32]byte, user []common.Address, token []common.Address) (*SimpleCustodyWithdrawStartedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SimpleCustody.contract.FilterLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCustodyWithdrawStartedIterator{contract: _SimpleCustody.contract, event: "WithdrawStarted", logs: logs, sub: sub}, nil
}

// WatchWithdrawStarted is a free log subscription operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_SimpleCustody *SimpleCustodyFilterer) WatchWithdrawStarted(opts *bind.WatchOpts, sink chan<- *SimpleCustodyWithdrawStarted, withdrawalId [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SimpleCustody.contract.WatchLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCustodyWithdrawStarted)
				if err := _SimpleCustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
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

// ParseWithdrawStarted is a log parse operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_SimpleCustody *SimpleCustodyFilterer) ParseWithdrawStarted(log types.Log) (*SimpleCustodyWithdrawStarted, error) {
	event := new(SimpleCustodyWithdrawStarted)
	if err := _SimpleCustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// QuorumCustodyMetaData contains all meta data concerning the QuorumCustody contract.
var QuorumCustodyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialSigners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"quorum_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADD_SIGNERS_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATION_EXPIRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REMOVE_SIGNERS_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSigners\",\"inputs\":[{\"name\":\"newSigners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"newQuorum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSignerCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSigner\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rejectWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSigners\",\"inputs\":[{\"name\":\"signersToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"newQuorum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signerNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalApprovals\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"hasApproved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requiredQuorum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumChanged\",\"inputs\":[{\"name\":\"oldQuorum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newQuorum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerAdded\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newQuorum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRemoved\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newQuorum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFinalized\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawStarted\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalApproved\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"currentApprovals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotRemoveLastSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeadlineExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ETHTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptySignersArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotASigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignaturesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignerAlreadyApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignerIsCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x610160806040523461034c576129c6803803809161001d8285610370565b833981019060408183031261034c5780516001600160401b03811161034c57810182601f8201121561034c578051926001600160401b0384116102f1578360051b91604051946100706020850187610370565b855260208086019382010191821161034c57602001915b8183106103505760208401516001600160401b03811690869082900361034c576040516100b5604082610370565b600d815260208101906c51756f72756d437573746f647960981b8252604051916100e0604084610370565b600183526020830191603160f81b835260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005561011d81610393565b6101205261012a84610529565b61014052519020918260e05251902080610100524660a0526040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261019360c082610370565b5190206080523060c05280511561033d5781151580610332575b15610323578051905f5b82811061022157600780546001600160401b0319168517905560405161236490816106628239608051816120b0015260a05181612167015260c05181612081015260e051816120ff0152610100518161212501526101205181610d3d01526101405181610d660152f35b81518110156102dd57600581901b8201602001516001600160a01b031690811561031457815f52600660205260ff60405f2054166103055760055491680100000000000000008310156102f15760018301806005558310156102dd5760019260055f5260205f200181848060a01b0319825416179055805f52600660205260405f208360ff198254161790557f250e2427befb4ce93c1d04e5896abb48ce7da6c28bfc92584a87b3d1331522cf6020604051888152a2016101b7565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b633b90338360e11b5f5260045ffd5b632057875960e21b5f5260045ffd5b63d173577960e01b5f5260045ffd5b5080518211156101ad565b637a67bdeb60e01b5f5260045ffd5b5f80fd5b82516001600160a01b038116810361034c57815260209283019201610087565b601f909101601f19168101906001600160401b038211908210176102f157604052565b908151602081105f1461040d575090601f8151116103cd5760208151910151602082106103be571790565b5f198260200360031b1b161790565b604460209160405192839163305a27a960e01b83528160048401528051918291826024860152018484015e5f828201840152601f01601f19168101030190fd5b6001600160401b0381116102f1575f54600181811c9116801561051f575b602082101461050b57601f81116104d9575b50602092601f821160011461047a57928192935f9261046f575b50508160011b915f199060031b1c1916175f5560ff90565b015190505f80610457565b601f198216935f8052805f20915f5b8681106104c157508360019596106104a9575b505050811b015f5560ff90565b01515f1960f88460031b161c191690555f808061049c565b91926020600181928685015181550194019201610489565b5f8052601f60205f20910160051c810190601f830160051c015b818110610500575061043d565b5f81556001016104f3565b634e487b7160e01b5f52602260045260245ffd5b90607f169061042b565b908151602081105f14610554575090601f8151116103cd5760208151910151602082106103be571790565b6001600160401b0381116102f157600154600181811c91168015610657575b602082101461050b57601f8111610624575b50602092601f82116001146105c357928192935f926105b8575b50508160011b915f199060031b1c19161760015560ff90565b015190505f8061059f565b601f1982169360015f52805f20915f5b86811061060c57508360019596106105f4575b505050811b0160015560ff90565b01515f1960f88460031b161c191690555f80806105e6565b919260206001819286850151815501940192016105d3565b60015f52601f60205f20910160051c810190601f830160051c015b81811061064c5750610585565b5f815560010161063f565b90607f169061057356fe60806040526004361015610011575f80fd5b5f3560e01c806305e95be7146113455780630ce8d6221461130a57806311edc78f1461118b57806314f8a6be146111335780631703a018146110ee5780632079fb9a1461108257806347e7ef2414610ef957806360d6e22014610ea15780637df73e2714610e3957806384b0196e14610d07578063962f2c59146109a15780639c686ee114610967578063a30bee00146108f2578063b715be81146108b7578063d87e1f411461056b578063efbf64a7146104b25763f8105157146100d4575f80fd5b346104ae576100e236611932565b9094929391335f52600660205260ff60405f205416156104865780421161045e5782156104365767ffffffffffffffff85169586151580610420575b806103dd575b156103b55761019d926101989161013b8688611b23565b600454946040519060208201927f78a54421fea8b2916ce7b78d63bd1ff4cf90f6242fb9e8e3bfc5c355863e1fbf845260408301528b606083015286608083015260a082015260a0815261019060c0826119bb565b519020611c2a565b611a14565b6004555f5b81811061023a575050506007549167ffffffffffffffff8316908082036101c557005b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000007fb367749f7733b422a7e023d43d9eb238afcd7360c002a39af5f6fbbbef107f4c9416176007556102356040519283928390929167ffffffffffffffff60209181604085019616845216910152565b0390a1005b61024d610248828486611a41565b611a51565b9073ffffffffffffffffffffffffffffffffffffffff8216801561038d57805f52600660205260ff60405f2054166103655760055468010000000000000000811015610338576001936102a882866102d49401600555611833565b90919073ffffffffffffffffffffffffffffffffffffffff8084549260031b9316831b921b1916179055565b805f52600660205260405f20837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f250e2427befb4ce93c1d04e5896abb48ce7da6c28bfc92584a87b3d1331522cf6020604051898152a2016101a2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f77206706000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f815e1d64000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fd1735779000000000000000000000000000000000000000000000000000000005f5260045ffd5b506005548481018091116103f357871115610124565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5067ffffffffffffffff6007541687101561011e565b7f7a67bdeb000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f1ab7da6b000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fa1b035c8000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f80fd5b346104ae5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576004355f52600260205260c060405f2067ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff8254169173ffffffffffffffffffffffffffffffffffffffff6001820154169060036002820154910154916040519485526020850152604084015260ff811615156060840152818160081c16608084015260481c1660a0820152f35b346104ae5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576105a2611878565b6105aa61189b565b6044359060643592335f52600660205260ff60405f205416156104865773ffffffffffffffffffffffffffffffffffffffff906105e5611a72565b16801561088f578215610867576040519173ffffffffffffffffffffffffffffffffffffffff602084019146835230604086015283606086015216928360808201528460a08201528560c082015260c0815261064260e0826119bb565b51902092835f52600260205267ffffffffffffffff600360405f20015460081c1661083f5767ffffffffffffffff600754166040519160c083019083821067ffffffffffffffff8311176103385760209787947f669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb9899460409485528782528a820189815260038684019386855260608101945f865273ffffffffffffffffffffffffffffffffffffffff608083019467ffffffffffffffff4216865260a084019687528c5f5260206002905281808c5f20955116167fffffffffffffffffffffffff0000000000000000000000000000000000000000855416178455511673ffffffffffffffffffffffffffffffffffffffff6001840191167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055516002820155019251151560ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00855416911617835551907fffffffffffffffffffffffffffffff00000000000000000000000000000000ff68ffffffffffffffff0070ffffffffffffffff0000000000000000008554935160481b169360081b16911617179055825191825288820152a460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055604051908152f35b7f157c65e1000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b7ffd684c3b000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576020600554604051908152f35b346104ae5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae5761092961189b565b6004355f52600360205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576020604051610e108152f35b346104ae576109af36611932565b94919094939293335f52600660205260ff60405f205416156104865781421161045e578315610436576005549182851015610cdf57848303968388116103f35760075467ffffffffffffffff1688811015610cce575b67ffffffffffffffff8816988915159182610cb9575b5081610cae575b50156103b557610a8f9261019891610a3a8888611b23565b600454946040519060208201927fcba211e2945e757cb8341a650bb2a20af428a18cd64a63892e8564a7c1a7bfc2845260408301528c606083015286608083015260a082015260a0815261019060c0826119bb565b600455915f5b818110610ab5576007548567ffffffffffffffff8216888082036101c557005b73ffffffffffffffffffffffffffffffffffffffff610ad8610248838587611a41565b16805f52600660205260ff60405f20541615610c8657805f52600660205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690555f5b858110610b5f575b50906001917fdc5c8906f1af1441ef2c796f82d27e2dda1b0ed7890ee1cc29787f4832ec529260206040518a8152a201610a95565b8173ffffffffffffffffffffffffffffffffffffffff610b7e83611833565b90549060031b1c1614610b9357600101610b22565b9190947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101928184116103f357610bf4906102a873ffffffffffffffffffffffffffffffffffffffff610be687611833565b90549060031b1c1691611833565b6005548015610c59577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610c2881611833565b73ffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555156103f35790936001610b2a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7fda0357f7000000000000000000000000000000000000000000000000000000005f5260045ffd5b905088111589610a22565b67ffffffffffffffff168a101591508a610a1b565b5067ffffffffffffffff8816610a05565b7fc4c85473000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae57610ddd610d617f0000000000000000000000000000000000000000000000000000000000000000611e24565b610d8a7f0000000000000000000000000000000000000000000000000000000000000000611f9a565b6020610deb60405192610d9d83856119bb565b5f84525f3681376040519586957f0f00000000000000000000000000000000000000000000000000000000000000875260e08588015260e08701906118be565b9085820360408701526118be565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b828110610e2257505050500390f35b835185528695509381019392810192600101610e13565b346104ae5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae5773ffffffffffffffffffffffffffffffffffffffff610e85611878565b165f526006602052602060ff60405f2054166040519015158152f35b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae5760206040517fcba211e2945e757cb8341a650bb2a20af428a18cd64a63892e8564a7c1a7bfc28152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae57610f2b611878565b60243590610f37611a72565b81156108675773ffffffffffffffffffffffffffffffffffffffff169081610fdb57803403610fb3575b6040519081527f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a760203392a360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7f1841b4e1000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610fb3576040517f23b872dd000000000000000000000000000000000000000000000000000000005f5233600452306024528160445260205f60648180875af19060015f5114821615611061575b6040525f606052610f6157507f5274afe7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b90600181151661107957833b15153d1516169061102a565b503d5f823e3d90fd5b346104ae5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576004356005548110156104ae5773ffffffffffffffffffffffffffffffffffffffff6110de602092611833565b90549060031b1c16604051908152f35b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae57602067ffffffffffffffff60075416604051908152f35b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae5760206040517f78a54421fea8b2916ce7b78d63bd1ff4cf90f6242fb9e8e3bfc5c355863e1fbf8152f35b346104ae5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae57600435335f52600660205260ff60405f20541615610486576111da611a72565b805f526002602052600360405f2001805467ffffffffffffffff8160081c1680156112e25760ff82166112ba57610e1081018091116103f357421115611292577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00600191161790557f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c60206040515f8152a260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b7fe91d3e7e000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fae899454000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f8d0fc1dd000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104ae575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae576020600454604051908152f35b346104ae5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ae57600435335f52600660205260ff60405f2054161561048657611394611a72565b805f52600260205260405f2060038101805467ffffffffffffffff8160081c169081156112e25760ff166112ba57610e1081018091116103f357421161180b57825f52600360205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f2054166117e357825f52600360205260405f2073ffffffffffffffffffffffffffffffffffffffff33165f5260205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790555f6005545f5b818110611781575050604051818152847fa4cb89333cc20aa626a2b0998d8608b8ce4a77e3a18dc121eb1bf316848a034a60203393a381549067ffffffffffffffff8260481c1611156114d7575b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055005b73ffffffffffffffffffffffffffffffffffffffff83541690600184019173ffffffffffffffffffffffffffffffffffffffff835416600286019460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008754951617905580155f1461165a5750814710611632575f80809381935af13d1561162d573d61156481611ae9565b9061157260405192836119bb565b81525f60203d92013e5b15611605575f925b7fffffffffffffffffffffffff000000000000000000000000000000000000000081541690557fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055557f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c602060405160018152a2808080806114b1565b7fb12d13eb000000000000000000000000000000000000000000000000000000005f5260045ffd5b61157c565b7fbb55fd27000000000000000000000000000000000000000000000000000000005f5260045ffd5b9594916040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024818b5afa80156117765782915f91611741575b501061163257604051917fa9059cbb000000000000000000000000000000000000000000000000000000005f5260045260245260205f604481808a5af19060015f5114821615611729575b604052156116fd575f939450611584565b847f5274afe7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b90600181151661107957863b15153d151616906116ec565b9150506020813d60201161176e575b8161175d602093836119bb565b810103126104ae57819051896116a1565b3d9150611750565b6040513d5f823e3d90fd5b855f52600360205260405f2073ffffffffffffffffffffffffffffffffffffffff806117ac84611833565b90549060031b1c16165f5260205260ff60405f2054166117cf575b600101611463565b916117db600191611a14565b9290506117c7565b7f2acb3fbc000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fc140b38a000000000000000000000000000000000000000000000000000000005f5260045ffd5b60055481101561184b5760055f5260205f2001905f90565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6004359073ffffffffffffffffffffffffffffffffffffffff821682036104ae57565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036104ae57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b9181601f840112156104ae5782359167ffffffffffffffff83116104ae576020808501948460051b0101116104ae57565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126104ae5760043567ffffffffffffffff81116104ae578161197b91600401611901565b9290929160243567ffffffffffffffff811681036104ae5791604435916064359067ffffffffffffffff82116104ae576119b791600401611901565b9091565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761033857604052565b67ffffffffffffffff81116103385760051b60200190565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103f35760010190565b919081101561184b5760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036104ae5790565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005414611ac15760027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f3ee5aeb5000000000000000000000000000000000000000000000000000000005f5260045ffd5b67ffffffffffffffff811161033857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b90611b2d816119fc565b91611b3b60405193846119bb565b818352611b47826119fc565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020850193013684375f5b818110611be957505050604051908160208101918294519290925f5b818110611bd0575050611bca9250037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826119bb565b51902090565b8451835260209485019486945090920191600101611b90565b73ffffffffffffffffffffffffffffffffffffffff611c0c610248838587611a41565b1690855181101561184b5760019160208260051b8801015201611b74565b604290611c3561206a565b90604051917f190100000000000000000000000000000000000000000000000000000000000083526002830152602282015220906001905f945f955b85871015611de1578660051b8301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112156104ae5783019081359167ffffffffffffffff83116104ae57602081019083360382136104ae57611cd884611ae9565b90611ce660405192836119bb565b84825260208536920101116104ae575f602085611d1896611d0f9583860137830101528761218d565b909391936121c7565b73ffffffffffffffffffffffffffffffffffffffff8083169116811115611db957805f52600660205260ff60405f20541615611d91573314611d6957611d6060019194611a14565b96019592611c71565b7f4c7612b8000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f8baa579f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f01eba551000000000000000000000000000000000000000000000000000000005f5260045ffd5b5094505091505067ffffffffffffffff6007541611611dfc57565b7f6e49c686000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff8114611e835760ff811690601f8211611e5b5760405191611e486040846119bb565b6020808452838101919036833783525290565b7fb3512b0c000000000000000000000000000000000000000000000000000000005f5260045ffd5b506040515f5f548060011c9160018216918215611f90575b602084108314611f63578385528492908115611f265750600114611ec9575b611ec6925003826119bb565b90565b505f80805290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b818310611f0a575050906020611ec692820101611eba565b6020919350806001915483858801015201910190918392611ef2565b60209250611ec69491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101611eba565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b92607f1692611e9b565b60ff8114611fbe5760ff811690601f8211611e5b5760405191611e486040846119bb565b506040515f6001548060011c9160018216918215612060575b602084108314611f63578385528492908115611f26575060011461200157611ec6925003826119bb565b5060015f90815290917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b818310612044575050906020611ec692820101611eba565b602091935080600191548385880101520191019091839261202c565b92607f1692611fd7565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016301480612164575b156120d2577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a08152611bca60c0826119bb565b507f000000000000000000000000000000000000000000000000000000000000000046146120a9565b81519190604183036121bd576121b69250602082015190606060408401519301515f1a9061229f565b9192909190565b50505f9160029190565b600481101561227257806121d9575050565b60018103612209577ff645eedf000000000000000000000000000000000000000000000000000000005f5260045ffd5b6002810361223d57507ffce698f7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6003146122475750565b7fd78bce0c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411612323579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15611776575f5173ffffffffffffffffffffffffffffffffffffffff81161561231957905f905f90565b505f906001905f90565b5050505f916003919056fea2646970667358221220d3794a72e1519439f28dd723c96b1a386e08db13e6ff8ef43fd82868c25c66f764736f6c634300081e0033",
}

// QuorumCustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use QuorumCustodyMetaData.ABI instead.
var QuorumCustodyABI = QuorumCustodyMetaData.ABI

// QuorumCustodyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuorumCustodyMetaData.Bin instead.
var QuorumCustodyBin = QuorumCustodyMetaData.Bin

// DeployQuorumCustody deploys a new Ethereum contract, binding an instance of QuorumCustody to it.
func DeployQuorumCustody(auth *bind.TransactOpts, backend bind.ContractBackend, initialSigners []common.Address, quorum_ uint64) (common.Address, *types.Transaction, *QuorumCustody, error) {
	parsed, err := QuorumCustodyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuorumCustodyBin), backend, initialSigners, quorum_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuorumCustody{QuorumCustodyCaller: QuorumCustodyCaller{contract: contract}, QuorumCustodyTransactor: QuorumCustodyTransactor{contract: contract}, QuorumCustodyFilterer: QuorumCustodyFilterer{contract: contract}}, nil
}

// QuorumCustody is an auto generated Go binding around an Ethereum contract.
type QuorumCustody struct {
	QuorumCustodyCaller     // Read-only binding to the contract
	QuorumCustodyTransactor // Write-only binding to the contract
	QuorumCustodyFilterer   // Log filterer for contract events
}

// QuorumCustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuorumCustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuorumCustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuorumCustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuorumCustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuorumCustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuorumCustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuorumCustodySession struct {
	Contract     *QuorumCustody    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuorumCustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuorumCustodyCallerSession struct {
	Contract *QuorumCustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// QuorumCustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuorumCustodyTransactorSession struct {
	Contract     *QuorumCustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// QuorumCustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuorumCustodyRaw struct {
	Contract *QuorumCustody // Generic contract binding to access the raw methods on
}

// QuorumCustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuorumCustodyCallerRaw struct {
	Contract *QuorumCustodyCaller // Generic read-only contract binding to access the raw methods on
}

// QuorumCustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuorumCustodyTransactorRaw struct {
	Contract *QuorumCustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuorumCustody creates a new instance of QuorumCustody, bound to a specific deployed contract.
func NewQuorumCustody(address common.Address, backend bind.ContractBackend) (*QuorumCustody, error) {
	contract, err := bindQuorumCustody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuorumCustody{QuorumCustodyCaller: QuorumCustodyCaller{contract: contract}, QuorumCustodyTransactor: QuorumCustodyTransactor{contract: contract}, QuorumCustodyFilterer: QuorumCustodyFilterer{contract: contract}}, nil
}

// NewQuorumCustodyCaller creates a new read-only instance of QuorumCustody, bound to a specific deployed contract.
func NewQuorumCustodyCaller(address common.Address, caller bind.ContractCaller) (*QuorumCustodyCaller, error) {
	contract, err := bindQuorumCustody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyCaller{contract: contract}, nil
}

// NewQuorumCustodyTransactor creates a new write-only instance of QuorumCustody, bound to a specific deployed contract.
func NewQuorumCustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*QuorumCustodyTransactor, error) {
	contract, err := bindQuorumCustody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyTransactor{contract: contract}, nil
}

// NewQuorumCustodyFilterer creates a new log filterer instance of QuorumCustody, bound to a specific deployed contract.
func NewQuorumCustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*QuorumCustodyFilterer, error) {
	contract, err := bindQuorumCustody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyFilterer{contract: contract}, nil
}

// bindQuorumCustody binds a generic wrapper to an already deployed contract.
func bindQuorumCustody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuorumCustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuorumCustody *QuorumCustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuorumCustody.Contract.QuorumCustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuorumCustody *QuorumCustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuorumCustody.Contract.QuorumCustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuorumCustody *QuorumCustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuorumCustody.Contract.QuorumCustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuorumCustody *QuorumCustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuorumCustody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuorumCustody *QuorumCustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuorumCustody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuorumCustody *QuorumCustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuorumCustody.Contract.contract.Transact(opts, method, params...)
}

// ADDSIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x14f8a6be.
//
// Solidity: function ADD_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodyCaller) ADDSIGNERSTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "ADD_SIGNERS_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDSIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x14f8a6be.
//
// Solidity: function ADD_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodySession) ADDSIGNERSTYPEHASH() ([32]byte, error) {
	return _QuorumCustody.Contract.ADDSIGNERSTYPEHASH(&_QuorumCustody.CallOpts)
}

// ADDSIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x14f8a6be.
//
// Solidity: function ADD_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodyCallerSession) ADDSIGNERSTYPEHASH() ([32]byte, error) {
	return _QuorumCustody.Contract.ADDSIGNERSTYPEHASH(&_QuorumCustody.CallOpts)
}

// OPERATIONEXPIRY is a free data retrieval call binding the contract method 0x9c686ee1.
//
// Solidity: function OPERATION_EXPIRY() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCaller) OPERATIONEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "OPERATION_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPERATIONEXPIRY is a free data retrieval call binding the contract method 0x9c686ee1.
//
// Solidity: function OPERATION_EXPIRY() view returns(uint256)
func (_QuorumCustody *QuorumCustodySession) OPERATIONEXPIRY() (*big.Int, error) {
	return _QuorumCustody.Contract.OPERATIONEXPIRY(&_QuorumCustody.CallOpts)
}

// OPERATIONEXPIRY is a free data retrieval call binding the contract method 0x9c686ee1.
//
// Solidity: function OPERATION_EXPIRY() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCallerSession) OPERATIONEXPIRY() (*big.Int, error) {
	return _QuorumCustody.Contract.OPERATIONEXPIRY(&_QuorumCustody.CallOpts)
}

// REMOVESIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x60d6e220.
//
// Solidity: function REMOVE_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodyCaller) REMOVESIGNERSTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "REMOVE_SIGNERS_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REMOVESIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x60d6e220.
//
// Solidity: function REMOVE_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodySession) REMOVESIGNERSTYPEHASH() ([32]byte, error) {
	return _QuorumCustody.Contract.REMOVESIGNERSTYPEHASH(&_QuorumCustody.CallOpts)
}

// REMOVESIGNERSTYPEHASH is a free data retrieval call binding the contract method 0x60d6e220.
//
// Solidity: function REMOVE_SIGNERS_TYPEHASH() view returns(bytes32)
func (_QuorumCustody *QuorumCustodyCallerSession) REMOVESIGNERSTYPEHASH() ([32]byte, error) {
	return _QuorumCustody.Contract.REMOVESIGNERSTYPEHASH(&_QuorumCustody.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QuorumCustody *QuorumCustodyCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QuorumCustody *QuorumCustodySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _QuorumCustody.Contract.Eip712Domain(&_QuorumCustody.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_QuorumCustody *QuorumCustodyCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _QuorumCustody.Contract.Eip712Domain(&_QuorumCustody.CallOpts)
}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCaller) GetSignerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "getSignerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_QuorumCustody *QuorumCustodySession) GetSignerCount() (*big.Int, error) {
	return _QuorumCustody.Contract.GetSignerCount(&_QuorumCustody.CallOpts)
}

// GetSignerCount is a free data retrieval call binding the contract method 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCallerSession) GetSignerCount() (*big.Int, error) {
	return _QuorumCustody.Contract.GetSignerCount(&_QuorumCustody.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool isSigner)
func (_QuorumCustody *QuorumCustodyCaller) IsSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "isSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool isSigner)
func (_QuorumCustody *QuorumCustodySession) IsSigner(signer common.Address) (bool, error) {
	return _QuorumCustody.Contract.IsSigner(&_QuorumCustody.CallOpts, signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address signer) view returns(bool isSigner)
func (_QuorumCustody *QuorumCustodyCallerSession) IsSigner(signer common.Address) (bool, error) {
	return _QuorumCustody.Contract.IsSigner(&_QuorumCustody.CallOpts, signer)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_QuorumCustody *QuorumCustodyCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_QuorumCustody *QuorumCustodySession) Quorum() (uint64, error) {
	return _QuorumCustody.Contract.Quorum(&_QuorumCustody.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_QuorumCustody *QuorumCustodyCallerSession) Quorum() (uint64, error) {
	return _QuorumCustody.Contract.Quorum(&_QuorumCustody.CallOpts)
}

// SignerNonce is a free data retrieval call binding the contract method 0x0ce8d622.
//
// Solidity: function signerNonce() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCaller) SignerNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "signerNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerNonce is a free data retrieval call binding the contract method 0x0ce8d622.
//
// Solidity: function signerNonce() view returns(uint256)
func (_QuorumCustody *QuorumCustodySession) SignerNonce() (*big.Int, error) {
	return _QuorumCustody.Contract.SignerNonce(&_QuorumCustody.CallOpts)
}

// SignerNonce is a free data retrieval call binding the contract method 0x0ce8d622.
//
// Solidity: function signerNonce() view returns(uint256)
func (_QuorumCustody *QuorumCustodyCallerSession) SignerNonce() (*big.Int, error) {
	return _QuorumCustody.Contract.SignerNonce(&_QuorumCustody.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_QuorumCustody *QuorumCustodyCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_QuorumCustody *QuorumCustodySession) Signers(arg0 *big.Int) (common.Address, error) {
	return _QuorumCustody.Contract.Signers(&_QuorumCustody.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_QuorumCustody *QuorumCustodyCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _QuorumCustody.Contract.Signers(&_QuorumCustody.CallOpts, arg0)
}

// WithdrawalApprovals is a free data retrieval call binding the contract method 0xa30bee00.
//
// Solidity: function withdrawalApprovals(bytes32 withdrawalId, address signer) view returns(bool hasApproved)
func (_QuorumCustody *QuorumCustodyCaller) WithdrawalApprovals(opts *bind.CallOpts, withdrawalId [32]byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "withdrawalApprovals", withdrawalId, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalApprovals is a free data retrieval call binding the contract method 0xa30bee00.
//
// Solidity: function withdrawalApprovals(bytes32 withdrawalId, address signer) view returns(bool hasApproved)
func (_QuorumCustody *QuorumCustodySession) WithdrawalApprovals(withdrawalId [32]byte, signer common.Address) (bool, error) {
	return _QuorumCustody.Contract.WithdrawalApprovals(&_QuorumCustody.CallOpts, withdrawalId, signer)
}

// WithdrawalApprovals is a free data retrieval call binding the contract method 0xa30bee00.
//
// Solidity: function withdrawalApprovals(bytes32 withdrawalId, address signer) view returns(bool hasApproved)
func (_QuorumCustody *QuorumCustodyCallerSession) WithdrawalApprovals(withdrawalId [32]byte, signer common.Address) (bool, error) {
	return _QuorumCustody.Contract.WithdrawalApprovals(&_QuorumCustody.CallOpts, withdrawalId, signer)
}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 withdrawalId) view returns(address user, address token, uint256 amount, bool finalized, uint64 createdAt, uint64 requiredQuorum)
func (_QuorumCustody *QuorumCustodyCaller) Withdrawals(opts *bind.CallOpts, withdrawalId [32]byte) (struct {
	User           common.Address
	Token          common.Address
	Amount         *big.Int
	Finalized      bool
	CreatedAt      uint64
	RequiredQuorum uint64
}, error) {
	var out []interface{}
	err := _QuorumCustody.contract.Call(opts, &out, "withdrawals", withdrawalId)

	outstruct := new(struct {
		User           common.Address
		Token          common.Address
		Amount         *big.Int
		Finalized      bool
		CreatedAt      uint64
		RequiredQuorum uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Finalized = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.RequiredQuorum = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 withdrawalId) view returns(address user, address token, uint256 amount, bool finalized, uint64 createdAt, uint64 requiredQuorum)
func (_QuorumCustody *QuorumCustodySession) Withdrawals(withdrawalId [32]byte) (struct {
	User           common.Address
	Token          common.Address
	Amount         *big.Int
	Finalized      bool
	CreatedAt      uint64
	RequiredQuorum uint64
}, error) {
	return _QuorumCustody.Contract.Withdrawals(&_QuorumCustody.CallOpts, withdrawalId)
}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 withdrawalId) view returns(address user, address token, uint256 amount, bool finalized, uint64 createdAt, uint64 requiredQuorum)
func (_QuorumCustody *QuorumCustodyCallerSession) Withdrawals(withdrawalId [32]byte) (struct {
	User           common.Address
	Token          common.Address
	Amount         *big.Int
	Finalized      bool
	CreatedAt      uint64
	RequiredQuorum uint64
}, error) {
	return _QuorumCustody.Contract.Withdrawals(&_QuorumCustody.CallOpts, withdrawalId)
}

// AddSigners is a paid mutator transaction binding the contract method 0xf8105157.
//
// Solidity: function addSigners(address[] newSigners, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodyTransactor) AddSigners(opts *bind.TransactOpts, newSigners []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "addSigners", newSigners, newQuorum, deadline, signatures)
}

// AddSigners is a paid mutator transaction binding the contract method 0xf8105157.
//
// Solidity: function addSigners(address[] newSigners, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodySession) AddSigners(newSigners []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.AddSigners(&_QuorumCustody.TransactOpts, newSigners, newQuorum, deadline, signatures)
}

// AddSigners is a paid mutator transaction binding the contract method 0xf8105157.
//
// Solidity: function addSigners(address[] newSigners, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodyTransactorSession) AddSigners(newSigners []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.AddSigners(&_QuorumCustody.TransactOpts, newSigners, newQuorum, deadline, signatures)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_QuorumCustody *QuorumCustodyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_QuorumCustody *QuorumCustodySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.Contract.Deposit(&_QuorumCustody.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_QuorumCustody *QuorumCustodyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.Contract.Deposit(&_QuorumCustody.TransactOpts, token, amount)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodyTransactor) FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "finalizeWithdraw", withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodySession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.FinalizeWithdraw(&_QuorumCustody.TransactOpts, withdrawalId)
}

// FinalizeWithdraw is a paid mutator transaction binding the contract method 0x05e95be7.
//
// Solidity: function finalizeWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodyTransactorSession) FinalizeWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.FinalizeWithdraw(&_QuorumCustody.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodyTransactor) RejectWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "rejectWithdraw", withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodySession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.RejectWithdraw(&_QuorumCustody.TransactOpts, withdrawalId)
}

// RejectWithdraw is a paid mutator transaction binding the contract method 0x11edc78f.
//
// Solidity: function rejectWithdraw(bytes32 withdrawalId) returns()
func (_QuorumCustody *QuorumCustodyTransactorSession) RejectWithdraw(withdrawalId [32]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.RejectWithdraw(&_QuorumCustody.TransactOpts, withdrawalId)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x962f2c59.
//
// Solidity: function removeSigners(address[] signersToRemove, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodyTransactor) RemoveSigners(opts *bind.TransactOpts, signersToRemove []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "removeSigners", signersToRemove, newQuorum, deadline, signatures)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x962f2c59.
//
// Solidity: function removeSigners(address[] signersToRemove, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodySession) RemoveSigners(signersToRemove []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.RemoveSigners(&_QuorumCustody.TransactOpts, signersToRemove, newQuorum, deadline, signatures)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x962f2c59.
//
// Solidity: function removeSigners(address[] signersToRemove, uint64 newQuorum, uint256 deadline, bytes[] signatures) returns()
func (_QuorumCustody *QuorumCustodyTransactorSession) RemoveSigners(signersToRemove []common.Address, newQuorum uint64, deadline *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _QuorumCustody.Contract.RemoveSigners(&_QuorumCustody.TransactOpts, signersToRemove, newQuorum, deadline, signatures)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32)
func (_QuorumCustody *QuorumCustodyTransactor) StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.contract.Transact(opts, "startWithdraw", user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32)
func (_QuorumCustody *QuorumCustodySession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.Contract.StartWithdraw(&_QuorumCustody.TransactOpts, user, token, amount, nonce)
}

// StartWithdraw is a paid mutator transaction binding the contract method 0xd87e1f41.
//
// Solidity: function startWithdraw(address user, address token, uint256 amount, uint256 nonce) returns(bytes32)
func (_QuorumCustody *QuorumCustodyTransactorSession) StartWithdraw(user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _QuorumCustody.Contract.StartWithdraw(&_QuorumCustody.TransactOpts, user, token, amount, nonce)
}

// QuorumCustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the QuorumCustody contract.
type QuorumCustodyDepositedIterator struct {
	Event *QuorumCustodyDeposited // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyDeposited)
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
		it.Event = new(QuorumCustodyDeposited)
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
func (it *QuorumCustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyDeposited represents a Deposited event raised by the QuorumCustody contract.
type QuorumCustodyDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_QuorumCustody *QuorumCustodyFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*QuorumCustodyDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyDepositedIterator{contract: _QuorumCustody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed user, address indexed token, uint256 amount)
func (_QuorumCustody *QuorumCustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *QuorumCustodyDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyDeposited)
				if err := _QuorumCustody.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_QuorumCustody *QuorumCustodyFilterer) ParseDeposited(log types.Log) (*QuorumCustodyDeposited, error) {
	event := new(QuorumCustodyDeposited)
	if err := _QuorumCustody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodyEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the QuorumCustody contract.
type QuorumCustodyEIP712DomainChangedIterator struct {
	Event *QuorumCustodyEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyEIP712DomainChanged)
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
		it.Event = new(QuorumCustodyEIP712DomainChanged)
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
func (it *QuorumCustodyEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyEIP712DomainChanged represents a EIP712DomainChanged event raised by the QuorumCustody contract.
type QuorumCustodyEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QuorumCustody *QuorumCustodyFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*QuorumCustodyEIP712DomainChangedIterator, error) {

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyEIP712DomainChangedIterator{contract: _QuorumCustody.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QuorumCustody *QuorumCustodyFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *QuorumCustodyEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyEIP712DomainChanged)
				if err := _QuorumCustody.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_QuorumCustody *QuorumCustodyFilterer) ParseEIP712DomainChanged(log types.Log) (*QuorumCustodyEIP712DomainChanged, error) {
	event := new(QuorumCustodyEIP712DomainChanged)
	if err := _QuorumCustody.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodyQuorumChangedIterator is returned from FilterQuorumChanged and is used to iterate over the raw logs and unpacked data for QuorumChanged events raised by the QuorumCustody contract.
type QuorumCustodyQuorumChangedIterator struct {
	Event *QuorumCustodyQuorumChanged // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyQuorumChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyQuorumChanged)
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
		it.Event = new(QuorumCustodyQuorumChanged)
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
func (it *QuorumCustodyQuorumChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyQuorumChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyQuorumChanged represents a QuorumChanged event raised by the QuorumCustody contract.
type QuorumCustodyQuorumChanged struct {
	OldQuorum uint64
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuorumChanged is a free log retrieval operation binding the contract event 0xb367749f7733b422a7e023d43d9eb238afcd7360c002a39af5f6fbbbef107f4c.
//
// Solidity: event QuorumChanged(uint64 oldQuorum, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) FilterQuorumChanged(opts *bind.FilterOpts) (*QuorumCustodyQuorumChangedIterator, error) {

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "QuorumChanged")
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyQuorumChangedIterator{contract: _QuorumCustody.contract, event: "QuorumChanged", logs: logs, sub: sub}, nil
}

// WatchQuorumChanged is a free log subscription operation binding the contract event 0xb367749f7733b422a7e023d43d9eb238afcd7360c002a39af5f6fbbbef107f4c.
//
// Solidity: event QuorumChanged(uint64 oldQuorum, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) WatchQuorumChanged(opts *bind.WatchOpts, sink chan<- *QuorumCustodyQuorumChanged) (event.Subscription, error) {

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "QuorumChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyQuorumChanged)
				if err := _QuorumCustody.contract.UnpackLog(event, "QuorumChanged", log); err != nil {
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

// ParseQuorumChanged is a log parse operation binding the contract event 0xb367749f7733b422a7e023d43d9eb238afcd7360c002a39af5f6fbbbef107f4c.
//
// Solidity: event QuorumChanged(uint64 oldQuorum, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) ParseQuorumChanged(log types.Log) (*QuorumCustodyQuorumChanged, error) {
	event := new(QuorumCustodyQuorumChanged)
	if err := _QuorumCustody.contract.UnpackLog(event, "QuorumChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodySignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the QuorumCustody contract.
type QuorumCustodySignerAddedIterator struct {
	Event *QuorumCustodySignerAdded // Event containing the contract specifics and raw log

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
func (it *QuorumCustodySignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodySignerAdded)
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
		it.Event = new(QuorumCustodySignerAdded)
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
func (it *QuorumCustodySignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodySignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodySignerAdded represents a SignerAdded event raised by the QuorumCustody contract.
type QuorumCustodySignerAdded struct {
	Signer    common.Address
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x250e2427befb4ce93c1d04e5896abb48ce7da6c28bfc92584a87b3d1331522cf.
//
// Solidity: event SignerAdded(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) FilterSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*QuorumCustodySignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodySignerAddedIterator{contract: _QuorumCustody.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x250e2427befb4ce93c1d04e5896abb48ce7da6c28bfc92584a87b3d1331522cf.
//
// Solidity: event SignerAdded(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *QuorumCustodySignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodySignerAdded)
				if err := _QuorumCustody.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x250e2427befb4ce93c1d04e5896abb48ce7da6c28bfc92584a87b3d1331522cf.
//
// Solidity: event SignerAdded(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) ParseSignerAdded(log types.Log) (*QuorumCustodySignerAdded, error) {
	event := new(QuorumCustodySignerAdded)
	if err := _QuorumCustody.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodySignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the QuorumCustody contract.
type QuorumCustodySignerRemovedIterator struct {
	Event *QuorumCustodySignerRemoved // Event containing the contract specifics and raw log

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
func (it *QuorumCustodySignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodySignerRemoved)
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
		it.Event = new(QuorumCustodySignerRemoved)
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
func (it *QuorumCustodySignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodySignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodySignerRemoved represents a SignerRemoved event raised by the QuorumCustody contract.
type QuorumCustodySignerRemoved struct {
	Signer    common.Address
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0xdc5c8906f1af1441ef2c796f82d27e2dda1b0ed7890ee1cc29787f4832ec5292.
//
// Solidity: event SignerRemoved(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) FilterSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*QuorumCustodySignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodySignerRemovedIterator{contract: _QuorumCustody.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0xdc5c8906f1af1441ef2c796f82d27e2dda1b0ed7890ee1cc29787f4832ec5292.
//
// Solidity: event SignerRemoved(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *QuorumCustodySignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodySignerRemoved)
				if err := _QuorumCustody.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0xdc5c8906f1af1441ef2c796f82d27e2dda1b0ed7890ee1cc29787f4832ec5292.
//
// Solidity: event SignerRemoved(address indexed signer, uint64 newQuorum)
func (_QuorumCustody *QuorumCustodyFilterer) ParseSignerRemoved(log types.Log) (*QuorumCustodySignerRemoved, error) {
	event := new(QuorumCustodySignerRemoved)
	if err := _QuorumCustody.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodyWithdrawFinalizedIterator is returned from FilterWithdrawFinalized and is used to iterate over the raw logs and unpacked data for WithdrawFinalized events raised by the QuorumCustody contract.
type QuorumCustodyWithdrawFinalizedIterator struct {
	Event *QuorumCustodyWithdrawFinalized // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyWithdrawFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyWithdrawFinalized)
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
		it.Event = new(QuorumCustodyWithdrawFinalized)
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
func (it *QuorumCustodyWithdrawFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyWithdrawFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyWithdrawFinalized represents a WithdrawFinalized event raised by the QuorumCustody contract.
type QuorumCustodyWithdrawFinalized struct {
	WithdrawalId [32]byte
	Success      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFinalized is a free log retrieval operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_QuorumCustody *QuorumCustodyFilterer) FilterWithdrawFinalized(opts *bind.FilterOpts, withdrawalId [][32]byte) (*QuorumCustodyWithdrawFinalizedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyWithdrawFinalizedIterator{contract: _QuorumCustody.contract, event: "WithdrawFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawFinalized is a free log subscription operation binding the contract event 0x150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c.
//
// Solidity: event WithdrawFinalized(bytes32 indexed withdrawalId, bool success)
func (_QuorumCustody *QuorumCustodyFilterer) WatchWithdrawFinalized(opts *bind.WatchOpts, sink chan<- *QuorumCustodyWithdrawFinalized, withdrawalId [][32]byte) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "WithdrawFinalized", withdrawalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyWithdrawFinalized)
				if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
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
func (_QuorumCustody *QuorumCustodyFilterer) ParseWithdrawFinalized(log types.Log) (*QuorumCustodyWithdrawFinalized, error) {
	event := new(QuorumCustodyWithdrawFinalized)
	if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodyWithdrawStartedIterator is returned from FilterWithdrawStarted and is used to iterate over the raw logs and unpacked data for WithdrawStarted events raised by the QuorumCustody contract.
type QuorumCustodyWithdrawStartedIterator struct {
	Event *QuorumCustodyWithdrawStarted // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyWithdrawStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyWithdrawStarted)
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
		it.Event = new(QuorumCustodyWithdrawStarted)
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
func (it *QuorumCustodyWithdrawStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyWithdrawStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyWithdrawStarted represents a WithdrawStarted event raised by the QuorumCustody contract.
type QuorumCustodyWithdrawStarted struct {
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
func (_QuorumCustody *QuorumCustodyFilterer) FilterWithdrawStarted(opts *bind.FilterOpts, withdrawalId [][32]byte, user []common.Address, token []common.Address) (*QuorumCustodyWithdrawStartedIterator, error) {

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

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyWithdrawStartedIterator{contract: _QuorumCustody.contract, event: "WithdrawStarted", logs: logs, sub: sub}, nil
}

// WatchWithdrawStarted is a free log subscription operation binding the contract event 0x669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb989.
//
// Solidity: event WithdrawStarted(bytes32 indexed withdrawalId, address indexed user, address indexed token, uint256 amount, uint256 nonce)
func (_QuorumCustody *QuorumCustodyFilterer) WatchWithdrawStarted(opts *bind.WatchOpts, sink chan<- *QuorumCustodyWithdrawStarted, withdrawalId [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "WithdrawStarted", withdrawalIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyWithdrawStarted)
				if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
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
func (_QuorumCustody *QuorumCustodyFilterer) ParseWithdrawStarted(log types.Log) (*QuorumCustodyWithdrawStarted, error) {
	event := new(QuorumCustodyWithdrawStarted)
	if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuorumCustodyWithdrawalApprovedIterator is returned from FilterWithdrawalApproved and is used to iterate over the raw logs and unpacked data for WithdrawalApproved events raised by the QuorumCustody contract.
type QuorumCustodyWithdrawalApprovedIterator struct {
	Event *QuorumCustodyWithdrawalApproved // Event containing the contract specifics and raw log

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
func (it *QuorumCustodyWithdrawalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuorumCustodyWithdrawalApproved)
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
		it.Event = new(QuorumCustodyWithdrawalApproved)
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
func (it *QuorumCustodyWithdrawalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuorumCustodyWithdrawalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuorumCustodyWithdrawalApproved represents a WithdrawalApproved event raised by the QuorumCustody contract.
type QuorumCustodyWithdrawalApproved struct {
	WithdrawalId     [32]byte
	Signer           common.Address
	CurrentApprovals *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalApproved is a free log retrieval operation binding the contract event 0xa4cb89333cc20aa626a2b0998d8608b8ce4a77e3a18dc121eb1bf316848a034a.
//
// Solidity: event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals)
func (_QuorumCustody *QuorumCustodyFilterer) FilterWithdrawalApproved(opts *bind.FilterOpts, withdrawalId [][32]byte, signer []common.Address) (*QuorumCustodyWithdrawalApprovedIterator, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.FilterLogs(opts, "WithdrawalApproved", withdrawalIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &QuorumCustodyWithdrawalApprovedIterator{contract: _QuorumCustody.contract, event: "WithdrawalApproved", logs: logs, sub: sub}, nil
}

// WatchWithdrawalApproved is a free log subscription operation binding the contract event 0xa4cb89333cc20aa626a2b0998d8608b8ce4a77e3a18dc121eb1bf316848a034a.
//
// Solidity: event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals)
func (_QuorumCustody *QuorumCustodyFilterer) WatchWithdrawalApproved(opts *bind.WatchOpts, sink chan<- *QuorumCustodyWithdrawalApproved, withdrawalId [][32]byte, signer []common.Address) (event.Subscription, error) {

	var withdrawalIdRule []interface{}
	for _, withdrawalIdItem := range withdrawalId {
		withdrawalIdRule = append(withdrawalIdRule, withdrawalIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _QuorumCustody.contract.WatchLogs(opts, "WithdrawalApproved", withdrawalIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuorumCustodyWithdrawalApproved)
				if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawalApproved", log); err != nil {
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

// ParseWithdrawalApproved is a log parse operation binding the contract event 0xa4cb89333cc20aa626a2b0998d8608b8ce4a77e3a18dc121eb1bf316848a034a.
//
// Solidity: event WithdrawalApproved(bytes32 indexed withdrawalId, address indexed signer, uint256 currentApprovals)
func (_QuorumCustody *QuorumCustodyFilterer) ParseWithdrawalApproved(log types.Log) (*QuorumCustodyWithdrawalApproved, error) {
	event := new(QuorumCustodyWithdrawalApproved)
	if err := _QuorumCustody.contract.UnpackLog(event, "WithdrawalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

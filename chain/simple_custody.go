// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neodax\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nitewatch\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NEODAX_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NITEWATCH_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rejectWithdraw\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"finalized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFinalized\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawStarted\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ETHTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MsgValueMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonZeroMsgValueForERC20\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611d9c380380611d9c833981810160405281019061003191906102c1565b600161004f6100446100d260201b60201c565b6100fb60201b60201c565b5f01819055506100675f5f1b8461010460201b60201c565b506100987f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f38361010460201b60201c565b506100c97ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa8261010460201b60201c565b50505050610311565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f61011583836101f960201b60201c565b6101ef5760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061018c61025c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506101f3565b5f90505b92915050565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61029082610267565b9050919050565b6102a081610286565b81146102aa575f5ffd5b50565b5f815190506102bb81610297565b92915050565b5f5f5f606084860312156102d8576102d7610263565b5b5f6102e5868287016102ad565b93505060206102f6868287016102ad565b9250506040610307868287016102ad565b9150509250925092565b611a7e8061031e5f395ff3fe6080604052600436106100dc575f3560e01c80635a98c2231161007e578063d547741f11610058578063d547741f146102a4578063d87e1f41146102cc578063da86f31514610308578063efbf64a714610332576100dc565b80635a98c2231461021457806391d148541461023e578063a217fddf1461027a576100dc565b8063248a9ca3116100ba578063248a9ca31461016c5780632f2ff15d146101a857806336568abe146101d057806347e7ef24146101f8576100dc565b806301ffc9a7146100e057806305e95be71461011c57806311edc78f14610144575b5f5ffd5b3480156100eb575f5ffd5b50610106600480360381019061010191906115c9565b610372565b604051610113919061160e565b60405180910390f35b348015610127575f5ffd5b50610142600480360381019061013d919061165a565b6103eb565b005b34801561014f575f5ffd5b5061016a6004803603810190610165919061165a565b6107f7565b005b348015610177575f5ffd5b50610192600480360381019061018d919061165a565b61092f565b60405161019f9190611694565b60405180910390f35b3480156101b3575f5ffd5b506101ce60048036038101906101c99190611707565b61094b565b005b3480156101db575f5ffd5b506101f660048036038101906101f19190611707565b61096d565b005b610212600480360381019061020d9190611778565b6109e8565b005b34801561021f575f5ffd5b50610228610c78565b6040516102359190611694565b60405180910390f35b348015610249575f5ffd5b50610264600480360381019061025f9190611707565b610c9c565b604051610271919061160e565b60405180910390f35b348015610285575f5ffd5b5061028e610cff565b60405161029b9190611694565b60405180910390f35b3480156102af575f5ffd5b506102ca60048036038101906102c59190611707565b610d05565b005b3480156102d7575f5ffd5b506102f260048036038101906102ed91906117b6565b610d27565b6040516102ff9190611694565b60405180910390f35b348015610313575f5ffd5b5061031c610fd6565b6040516103299190611694565b60405180910390f35b34801561033d575f5ffd5b506103586004803603810190610353919061165a565b610ffa565b604051610369959493929190611838565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103e457506103e382611083565b5b9050919050565b7ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa610415816110ec565b61041d611100565b5f60015f8481526020019081526020015f209050806003015f9054906101000a900460ff16610478576040517f8d0fc1dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060030160019054906101000a900460ff16156104c1576040517fae89945400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018160030160016101000a81548160ff0219169083151502179055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f826001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f836002015490505f845f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f846001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f84600201819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106d1578047101561062c576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8373ffffffffffffffffffffffffffffffffffffffff1682604051610651906118b6565b5f6040518083038185875af1925050503d805f811461068b576040519150601f19603f3d011682016040523d82523d5f602084013e610690565b606091505b50509050806106cb576040517fb12d13eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506107ae565b808273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161070b91906118ca565b602060405180830381865afa158015610726573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061074a91906118f7565b1015610782576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107ad83828473ffffffffffffffffffffffffffffffffffffffff166111229092919063ffffffff16565b5b857f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c60016040516107df919061160e565b60405180910390a2505050506107f3611175565b5050565b7ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa610821816110ec565b610829611100565b5f60015f8481526020019081526020015f209050806003015f9054906101000a900460ff16610884576040517f8d0fc1dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060030160019054906101000a900460ff16156108cd576040517fae89945400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018160030160016101000a81548160ff021916908315150217905550827f150e5422471a0e0b0bf81bb0c466ec4b78850d2feeea6955c7e5eb33468a9c9c5f60405161091a919061160e565b60405180910390a25061092b611175565b5050565b5f5f5f8381526020019081526020015f20600101549050919050565b6109548261092f565b61095d816110ec565b610967838361118f565b50505050565b610975611278565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109d9576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109e3828261127f565b505050565b6109f0611100565b5f8103610a29576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8190505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a9e57813414610a99576040517fbc6f88c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c06565b5f3414610ad7576040517fa57ec87300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610b1191906118ca565b602060405180830381865afa158015610b2c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5091906118f7565b9050610b7f3330858773ffffffffffffffffffffffffffffffffffffffff16611368909392919063ffffffff16565b808473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610bb991906118ca565b602060405180830381865afa158015610bd4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bf891906118f7565b610c02919061194f565b9150505b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a783604051610c639190611982565b60405180910390a350610c74611175565b5050565b7ff42609614d16e60ed8a62ea70f772fc08fb4f581d8126a6aeae13d7aee25daaa81565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f5f1b81565b610d0e8261092f565b610d17816110ec565b610d21838361127f565b50505050565b5f7f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f3610d52816110ec565b610d5a611100565b5f8403610d93576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b463087878787604051602001610dae9695949392919061199b565b60405160208183030381529060405280519060200120915060015f8381526020019081526020015f206003015f9054906101000a900460ff1615610e1e576040517f157c65e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020016001151581526020015f151581525060015f8481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555060808201518160030160016101000a81548160ff0219169083151502179055509050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16837f669c87d38156449c65caf07041b1568372d50fc03f2cc46add1d68cebc2eb9898787604051610fbd9291906119fa565b60405180910390a4610fcd611175565b50949350505050565b7f7f207140ff521d8790ff51fbcb7b65fa00c82600e052949aeb1de1aeceafd4f381565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015f9054906101000a900460ff16908060030160019054906101000a900460ff16905085565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6110fd816110f8611278565b6113bd565b50565b61110861140e565b600261111a61111561144f565b611478565b5f0181905550565b61112f8383836001611481565b61117057826040517f5274afe700000000000000000000000000000000000000000000000000000000815260040161116791906118ca565b60405180910390fd5b505050565b600161118761118261144f565b611478565b5f0181905550565b5f61119a8383610c9c565b61126e5760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061120b611278565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050611272565b5f90505b92915050565b5f33905090565b5f61128a8383610c9c565b1561135e575f5f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506112fb611278565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050611362565b5f90505b92915050565b6113768484848460016114e3565b6113b757836040517f5274afe70000000000000000000000000000000000000000000000000000000081526004016113ae91906118ca565b60405180910390fd5b50505050565b6113c78282610c9c565b61140a5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401611401929190611a21565b60405180910390fd5b5050565b611416611554565b1561144d576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f5f63a9059cbb60e01b9050604051815f525f1960601c86166004528460245260205f60445f5f8b5af1925060015f511483166114d55783831516156114c9573d5f823e3d81fd5b5f873b113d1516831692505b806040525050949350505050565b5f5f6323b872dd60e01b9050604051815f525f1960601c87166004525f1960601c86166024528460445260205f60645f5f8c5af1925060015f51148316611541578383151615611535573d5f823e3d81fd5b5f883b113d1516831692505b806040525f606052505095945050505050565b5f600261156761156261144f565b611478565b5f015414905090565b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6115a881611574565b81146115b2575f5ffd5b50565b5f813590506115c38161159f565b92915050565b5f602082840312156115de576115dd611570565b5b5f6115eb848285016115b5565b91505092915050565b5f8115159050919050565b611608816115f4565b82525050565b5f6020820190506116215f8301846115ff565b92915050565b5f819050919050565b61163981611627565b8114611643575f5ffd5b50565b5f8135905061165481611630565b92915050565b5f6020828403121561166f5761166e611570565b5b5f61167c84828501611646565b91505092915050565b61168e81611627565b82525050565b5f6020820190506116a75f830184611685565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6116d6826116ad565b9050919050565b6116e6816116cc565b81146116f0575f5ffd5b50565b5f81359050611701816116dd565b92915050565b5f5f6040838503121561171d5761171c611570565b5b5f61172a85828601611646565b925050602061173b858286016116f3565b9150509250929050565b5f819050919050565b61175781611745565b8114611761575f5ffd5b50565b5f813590506117728161174e565b92915050565b5f5f6040838503121561178e5761178d611570565b5b5f61179b858286016116f3565b92505060206117ac85828601611764565b9150509250929050565b5f5f5f5f608085870312156117ce576117cd611570565b5b5f6117db878288016116f3565b94505060206117ec878288016116f3565b93505060406117fd87828801611764565b925050606061180e87828801611764565b91505092959194509250565b611823816116cc565b82525050565b61183281611745565b82525050565b5f60a08201905061184b5f83018861181a565b611858602083018761181a565b6118656040830186611829565b61187260608301856115ff565b61187f60808301846115ff565b9695505050505050565b5f81905092915050565b50565b5f6118a15f83611889565b91506118ac82611893565b5f82019050919050565b5f6118c082611896565b9150819050919050565b5f6020820190506118dd5f83018461181a565b92915050565b5f815190506118f18161174e565b92915050565b5f6020828403121561190c5761190b611570565b5b5f611919848285016118e3565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61195982611745565b915061196483611745565b925082820390508181111561197c5761197b611922565b5b92915050565b5f6020820190506119955f830184611829565b92915050565b5f60c0820190506119ae5f830189611829565b6119bb602083018861181a565b6119c8604083018761181a565b6119d5606083018661181a565b6119e26080830185611829565b6119ef60a0830184611829565b979650505050505050565b5f604082019050611a0d5f830185611829565b611a1a6020830184611829565b9392505050565b5f604082019050611a345f83018561181a565b611a416020830184611685565b939250505056fea2646970667358221220675434c4bbddcb90fe2fd6368a14981a2f3c3f613b72899e2626a1e12ecf1b0664736f6c634300081e0033",
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
// Solidity: function withdrawals(bytes32 ) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodyCaller) Withdrawals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	var out []interface{}
	err := _SimpleCustody.contract.Call(opts, &out, "withdrawals", arg0)

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
// Solidity: function withdrawals(bytes32 ) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodySession) Withdrawals(arg0 [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	return _SimpleCustody.Contract.Withdrawals(&_SimpleCustody.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0xefbf64a7.
//
// Solidity: function withdrawals(bytes32 ) view returns(address user, address token, uint256 amount, bool exists, bool finalized)
func (_SimpleCustody *SimpleCustodyCallerSession) Withdrawals(arg0 [32]byte) (struct {
	User      common.Address
	Token     common.Address
	Amount    *big.Int
	Exists    bool
	Finalized bool
}, error) {
	return _SimpleCustody.Contract.Withdrawals(&_SimpleCustody.CallOpts, arg0)
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

package chain

//go:generate sh -c "jq .abi ../contracts/evm/out/IWithdraw.sol/IWithdraw.json > IWithdraw.abi && abigen --abi IWithdraw.abi --pkg chain --type IWithdraw --out iwithdraw.go"
//go:generate sh -c "jq .abi ../contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > SimpleCustody.abi && jq -r .bytecode.object ../contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > SimpleCustody.bin && abigen --abi SimpleCustody.abi --bin SimpleCustody.bin --pkg chain --type SimpleCustody --out simple_custody.go"

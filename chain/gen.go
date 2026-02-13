package chain

//go:generate sh -c "jq .abi ../contracts/evm/out/ICustody.sol/ICustody.json > ICustody.abi && abigen --abi ICustody.abi --pkg chain --type ICustody --out icustody.go"
//go:generate sh -c "jq .abi ../contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > SimpleCustody.abi && jq -r .bytecode.object ../contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > SimpleCustody.bin && abigen --abi SimpleCustody.abi --bin SimpleCustody.bin --pkg chain --type SimpleCustody --out simple_custody.go"

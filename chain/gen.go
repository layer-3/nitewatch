package chain

//go:generate sh -c "jq .abi ../contracts/evm/out/ICustody.sol/ICustody.json > ICustody.abi && abigen --abi ICustody.abi --pkg chain --type ICustody --out icustody.go"

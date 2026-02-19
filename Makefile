SOL_SOURCES := $(shell find contracts/evm/src -name '*.sol')
BINDINGS    := custody/iwithdraw.go custody/ideposit.go custody/simple_custody.go

.PHONY: generate
generate: $(BINDINGS)

# Sentinel tracks forge build; only re-runs when .sol sources change.
contracts/evm/out/.build-sentinel: $(SOL_SOURCES)
	cd contracts/evm && forge build
	@touch $@

custody/iwithdraw.go: contracts/evm/out/.build-sentinel
	jq .abi contracts/evm/out/IWithdraw.sol/IWithdraw.json > custody/IWithdraw.abi
	abigen --abi custody/IWithdraw.abi --pkg custody --type IWithdraw --out $@

custody/ideposit.go: contracts/evm/out/.build-sentinel
	jq .abi contracts/evm/out/IDeposit.sol/IDeposit.json > custody/IDeposit.abi
	abigen --abi custody/IDeposit.abi --pkg custody --type IDeposit --out $@

custody/simple_custody.go: contracts/evm/out/.build-sentinel
	jq .abi contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > custody/SimpleCustody.abi
	jq -r .bytecode.object contracts/evm/out/SimpleCustody.sol/SimpleCustody.json > custody/SimpleCustody.bin
	abigen --abi custody/SimpleCustody.abi --bin custody/SimpleCustody.bin --pkg custody --type SimpleCustody --out $@


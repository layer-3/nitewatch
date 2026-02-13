## Foundry

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

- **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
- **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
- **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
- **Chisel**: Fast, utilitarian, and verbose solidity REPL.

## Documentation

https://book.getfoundry.sh/

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy QuorumCustody

```shell
PRIVATE_KEY=<deployer_key> \
SIGNERS=0xAbc...,0xDef...,0x123... \
INITIAL_QUORUM=2 \
  forge script script/DeployQuorumCustody.s.sol --rpc-url $RPC_URL --broadcast
```

### Deploy SimpleCustody

```shell
PRIVATE_KEY=<deployer_key> \
ADMIN_ADDRESS=0x... \
NEODAX_ADDRESS=0x... \
NITEWATCH_ADDRESS=0x... \
  forge script script/DeploySimpleCustody.s.sol --rpc-url $RPC_URL --broadcast
```

### Add Signers

Adds signers to a QuorumCustody contract. When quorum > 1, co-signatures from
other signers are required (EIP-712 sorted ascending by signer address).

```shell
PRIVATE_KEY=<caller_key> \
CONTRACT=0x... \
NEW_SIGNERS=0xAbc...,0xDef... \
NEW_QUORUM=3 \
DEADLINE=1999999999 \
SIGNATURES=0x<sig1>,0x<sig2> \
  forge script script/AddSigners.s.sol --rpc-url $RPC_URL --broadcast
```

When quorum is 1, `SIGNATURES` can be omitted:

```shell
PRIVATE_KEY=<caller_key> \
CONTRACT=0x... \
NEW_SIGNERS=0xAbc... \
NEW_QUORUM=2 \
DEADLINE=1999999999 \
  forge script script/AddSigners.s.sol --rpc-url $RPC_URL --broadcast
```

### Remove Signers

Removes signers from a QuorumCustody contract. Same signature rules apply.

```shell
PRIVATE_KEY=<caller_key> \
CONTRACT=0x... \
SIGNERS_TO_REMOVE=0xAbc... \
NEW_QUORUM=2 \
DEADLINE=1999999999 \
SIGNATURES=0x<sig1> \
  forge script script/RemoveSigners.s.sol --rpc-url $RPC_URL --broadcast
```

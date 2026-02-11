# chain

The `chain` package provides Go bindings and a high-level listener for the `ICustody` smart contract events.

## Listener

The `Listener` struct allows you to subscribe to contract events and wait for a specified number of block confirmations before processing them. This is crucial for ensuring that your application only acts on finalized transactions, mitigating the risk of chain reorgs.

### Usage

```go
package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/layer-3/nitewatch/chain"
)

func main() {
	// 1. Connect to an Ethereum node (must support subscriptions, e.g., via WebSocket)
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 2. Create a new Listener
	// address: The deployed ICustody contract address
	// confirmations: Number of blocks to wait for confirmation (e.g., 12 for mainnet)
	contractAddress := common.HexToAddress("0x...")
	listener, err := chain.NewListener(client, contractAddress, 12)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	// 3. Subscribe to Deposited events
	ctx := context.Background()
	deposits := make(chan *chain.ICustodyDeposited)
	
	go func() {
		if err := listener.WatchDeposited(ctx, deposits); err != nil {
			log.Printf("WatchDeposited error: %v", err)
		}
	}()

	// 4. Process confirmed events
	for event := range deposits {
		log.Printf("Confirmed deposit: User=%s Token=%s Amount=%s Block=%d",
			event.User.Hex(), event.Token.Hex(), event.Amount.String(), event.Raw.BlockNumber)
	}
}
```

## Confirmations

The `confirmations` parameter in `NewListener` determines how many blocks must be mined on top of the block containing the event before it is emitted.

-   `confirmations = 0`: Events are emitted immediately upon detection. Reorgs are not handled (removed events are ignored).
-   `confirmations > 0`: Events are buffered until the chain reaches the required depth. If a reorg occurs and the event is removed from the canonical chain during the confirmation period, it is discarded and never emitted.

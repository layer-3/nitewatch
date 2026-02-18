package anviltest

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	sharedContainer testcontainers.Container
	sharedURLs      URLs
	sharedOnce      sync.Once
	sharedErr       error
)

// Anvil deterministic private keys (accounts pre-funded with 10000 ETH each).
const (
	Account0Key = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	Account1Key = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	Account2Key = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
)

// URLs holds both HTTP and WebSocket URLs for an Anvil instance.
// HTTP is used for one-shot RPC calls (deploy, transact), WS for event subscriptions.
type URLs struct {
	HTTP string
	WS   string
}

// Setup starts an Anvil testcontainer and returns HTTP and WebSocket URLs.
// The container is automatically terminated when the test finishes.
func Setup(t *testing.T) URLs {
	t.Helper()

	ctx := t.Context()

	anvilContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "ghcr.io/foundry-rs/foundry:latest",
			ExposedPorts: []string{"8545/tcp"},
			// Image ENTRYPOINT is ["/bin/sh", "-c"], so the command must be a single string.
			Cmd: []string{"anvil --host 0.0.0.0"},
			WaitingFor: wait.ForAll(
				wait.ForLog("Listening on"),
				wait.ForListeningPort("8545/tcp"),
			),
		},
		Started: true,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := anvilContainer.Terminate(context.Background()); err != nil {
			t.Log("Anvil container terminated with error:", err)
		}
	})

	host, err := anvilContainer.Host(ctx)
	require.NoError(t, err)

	port, err := anvilContainer.MappedPort(ctx, "8545")
	require.NoError(t, err)

	return URLs{
		HTTP: fmt.Sprintf("http://%s:%s", host, port.Port()),
		WS:   fmt.Sprintf("ws://%s:%s", host, port.Port()),
	}
}

// SetupShared starts a shared Anvil container (once) and returns its URLs.
// Call TerminateShared in TestMain to clean up after all tests.
func SetupShared(ctx context.Context) (URLs, error) {
	sharedOnce.Do(func() {
		var container testcontainers.Container
		container, sharedErr = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "ghcr.io/foundry-rs/foundry:latest",
				ExposedPorts: []string{"8545/tcp"},
				Cmd:          []string{"anvil --host 0.0.0.0"},
				WaitingFor: wait.ForAll(
					wait.ForLog("Listening on"),
					wait.ForListeningPort("8545/tcp"),
				),
			},
			Started: true,
		})
		if sharedErr != nil {
			return
		}
		sharedContainer = container

		host, err := container.Host(ctx)
		if err != nil {
			sharedErr = err
			return
		}
		port, err := container.MappedPort(ctx, "8545")
		if err != nil {
			sharedErr = err
			return
		}
		sharedURLs = URLs{
			HTTP: fmt.Sprintf("http://%s:%s", host, port.Port()),
			WS:   fmt.Sprintf("ws://%s:%s", host, port.Port()),
		}
	})
	return sharedURLs, sharedErr
}

// TerminateShared stops the shared Anvil container. Call from TestMain.
func TerminateShared(ctx context.Context) error {
	if sharedContainer != nil {
		return sharedContainer.Terminate(ctx)
	}
	return nil
}

// DialHTTP connects to an Anvil node via HTTP with retry for container readiness.
// Retries up to 10 times with 500ms backoff, verifying with a ChainID() call.
func DialHTTP(t *testing.T, httpURL string) *ethclient.Client {
	t.Helper()

	var client *ethclient.Client
	var err error
	for i := 0; i < 10; i++ {
		client, err = ethclient.Dial(httpURL)
		if err == nil {
			_, chainErr := client.ChainID(context.Background())
			if chainErr == nil {
				return client
			}
			client.Close()
			err = chainErr
		}
		time.Sleep(500 * time.Millisecond)
	}
	require.NoError(t, err, "failed to connect to Anvil after retries")
	return nil
}

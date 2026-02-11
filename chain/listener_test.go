package chain

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// testEvent is a simple event type for testing the generic watchWithConfirmations.
type testEvent struct {
	ID  int
	Raw types.Log
}

// mockSubscription implements ethereum.Subscription for testing.
type mockSubscription struct {
	errCh  chan error
	unsub  bool
	unsubC chan struct{}
}

func newMockSubscription() *mockSubscription {
	return &mockSubscription{
		errCh:  make(chan error, 1),
		unsubC: make(chan struct{}),
	}
}

func (m *mockSubscription) Err() <-chan error { return m.errCh }
func (m *mockSubscription) Unsubscribe() {
	if !m.unsub {
		m.unsub = true
		close(m.unsubC)
	}
}

func TestWatchWithConfirmations_ZeroConfirmations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh := make(chan *testEvent)
	headCh := make(chan *types.Header)
	sink := make(chan *testEvent, 10)

	eventSub := newMockSubscription()
	headSub := newMockSubscription()

	go func() {
		err := watchWithConfirmations(
			ctx, sink, 0,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				go func() {
					for ev := range eventCh {
						rawSink <- ev
					}
				}()
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				go func() {
					for h := range headCh {
						ch <- h
					}
				}()
				return headSub, nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
		_ = err
	}()

	// Send an event - should be delivered immediately
	eventCh <- &testEvent{
		ID: 1,
		Raw: types.Log{
			BlockNumber: 100,
			TxHash:      common.HexToHash("0x01"),
		},
	}

	select {
	case ev := <-sink:
		if ev.ID != 1 {
			t.Fatalf("expected event ID 1, got %d", ev.ID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for event")
	}

	cancel()
}

func TestWatchWithConfirmations_WithConfirmations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh := make(chan *testEvent)
	headCh := make(chan *types.Header)
	sink := make(chan *testEvent, 10)

	eventSub := newMockSubscription()
	headSub := newMockSubscription()

	go func() {
		_ = watchWithConfirmations(
			ctx, sink, 3,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				go func() {
					for ev := range eventCh {
						rawSink <- ev
					}
				}()
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				go func() {
					for h := range headCh {
						ch <- h
					}
				}()
				return headSub, nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	// Send event at block 100
	eventCh <- &testEvent{
		ID: 1,
		Raw: types.Log{
			BlockNumber: 100,
			TxHash:      common.HexToHash("0x01"),
		},
	}

	// Allow event to be buffered
	time.Sleep(50 * time.Millisecond)

	// Block 101 - not enough confirmations (depth=2, need 3)
	headCh <- &types.Header{Number: big.NewInt(101)}
	time.Sleep(50 * time.Millisecond)

	select {
	case <-sink:
		t.Fatal("event should not be delivered yet")
	default:
	}

	// Block 102 - depth=3, should confirm (102+1 >= 100+3)
	headCh <- &types.Header{Number: big.NewInt(102)}

	select {
	case ev := <-sink:
		if ev.ID != 1 {
			t.Fatalf("expected event ID 1, got %d", ev.ID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for confirmed event")
	}

	cancel()
}

func TestWatchWithConfirmations_ReorgRemoval(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh := make(chan *testEvent)
	headCh := make(chan *types.Header)
	sink := make(chan *testEvent, 10)

	eventSub := newMockSubscription()
	headSub := newMockSubscription()

	go func() {
		_ = watchWithConfirmations(
			ctx, sink, 3,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				go func() {
					for ev := range eventCh {
						rawSink <- ev
					}
				}()
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				go func() {
					for h := range headCh {
						ch <- h
					}
				}()
				return headSub, nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	txHash := common.HexToHash("0x01")

	// Send event at block 100
	eventCh <- &testEvent{
		ID:  1,
		Raw: types.Log{BlockNumber: 100, TxHash: txHash, Index: 0},
	}

	time.Sleep(50 * time.Millisecond)

	// Reorg: same event is removed
	eventCh <- &testEvent{
		ID:  1,
		Raw: types.Log{BlockNumber: 100, TxHash: txHash, Index: 0, Removed: true},
	}

	time.Sleep(50 * time.Millisecond)

	// Even with enough blocks, the removed event should not be delivered
	headCh <- &types.Header{Number: big.NewInt(105)}
	time.Sleep(50 * time.Millisecond)

	select {
	case <-sink:
		t.Fatal("removed event should not be delivered")
	default:
	}

	cancel()
}

func TestWatchWithConfirmations_MultipleEvents(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh := make(chan *testEvent)
	headCh := make(chan *types.Header)
	sink := make(chan *testEvent, 10)

	eventSub := newMockSubscription()
	headSub := newMockSubscription()

	go func() {
		_ = watchWithConfirmations(
			ctx, sink, 2,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				go func() {
					for ev := range eventCh {
						rawSink <- ev
					}
				}()
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				go func() {
					for h := range headCh {
						ch <- h
					}
				}()
				return headSub, nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	// Events at different blocks
	eventCh <- &testEvent{ID: 1, Raw: types.Log{BlockNumber: 100, TxHash: common.HexToHash("0x01")}}
	eventCh <- &testEvent{ID: 2, Raw: types.Log{BlockNumber: 101, TxHash: common.HexToHash("0x02")}}
	eventCh <- &testEvent{ID: 3, Raw: types.Log{BlockNumber: 102, TxHash: common.HexToHash("0x03")}}

	time.Sleep(50 * time.Millisecond)

	// Block 101: confirms event 1 (101+1 >= 100+2), not event 2 or 3
	headCh <- &types.Header{Number: big.NewInt(101)}

	select {
	case ev := <-sink:
		if ev.ID != 1 {
			t.Fatalf("expected event 1 first, got %d", ev.ID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out")
	}

	// Block 103: confirms events 2 and 3
	headCh <- &types.Header{Number: big.NewInt(103)}

	received := map[int]bool{}
	for i := 0; i < 2; i++ {
		select {
		case ev := <-sink:
			received[ev.ID] = true
		case <-time.After(time.Second):
			t.Fatal("timed out")
		}
	}

	if !received[2] || !received[3] {
		t.Fatalf("expected events 2 and 3, got %v", received)
	}

	cancel()
}

func TestWatchWithConfirmations_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	sink := make(chan *testEvent, 10)
	eventSub := newMockSubscription()
	headSub := newMockSubscription()

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(
			ctx, sink, 0,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return headSub, nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	cancel()

	select {
	case err := <-errCh:
		if err != context.Canceled {
			t.Fatalf("expected context.Canceled, got: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for cancellation")
	}
}

func TestWatchWithConfirmations_SubscribeEventError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sink := make(chan *testEvent, 10)

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(
			ctx, sink, 0,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				return nil, context.DeadlineExceeded
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return newMockSubscription(), nil
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("expected error")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out")
	}
}

func TestWatchWithConfirmations_SubscribeHeadError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sink := make(chan *testEvent, 10)
	eventSub := newMockSubscription()

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(
			ctx, sink, 0,
			func(rawSink chan<- *testEvent) (ethereum.Subscription, error) {
				return eventSub, nil
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return nil, context.DeadlineExceeded
			},
			func(e *testEvent) types.Log { return e.Raw },
		)
	}()

	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("expected error")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out")
	}
}

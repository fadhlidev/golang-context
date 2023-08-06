package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ProcessWithCancel(ctx context.Context, name string, delay time.Duration) {
	time.Sleep(delay)

	// checks whether the context has been cancelled
	select {
	case <-ctx.Done():
		fmt.Println(name, "has been canceled")
		return
	default:
		fmt.Println("Process of", name)
		break
	}
}

func TestContextWithCancel(T *testing.T) {
	// create cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	go ProcessWithCancel(ctx, "Process-A", 3*time.Second)
	go ProcessWithCancel(ctx, "Process-B", 2*time.Second)
	println("Total goroutine", runtime.NumGoroutine())

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
	println("Total goroutine", runtime.NumGoroutine())
}

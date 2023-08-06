package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ProcessWithTimeout(ctx context.Context, name string, delay time.Duration) {
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

func TestContextWithTimeout(T *testing.T) {
	// create context with timeout, it is like cancelable, but doesn't to call cancel function
	// however, we can still cancel the context manually
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// It is recommended to always call cancel function
	defer cancel()

	go ProcessWithTimeout(ctx, "Process-A", 3*time.Second)
	go ProcessWithTimeout(ctx, "Process-B", 2*time.Second)
	println("Total goroutine", runtime.NumGoroutine())

	time.Sleep(4 * time.Second)
	println("Total goroutine", runtime.NumGoroutine())
}

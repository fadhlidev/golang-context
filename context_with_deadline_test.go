package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ProcessWithDeadline(ctx context.Context, name string, delay time.Duration) {
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

func TestContextWithDeadline(T *testing.T) {
	// create context with deadline, it is like timeout context, but we specify when the context should cancel
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// It is recommended to always call cancel function
	defer cancel()

	go ProcessWithDeadline(ctx, "Process-A", 3*time.Second)
	go ProcessWithDeadline(ctx, "Process-B", 6*time.Second)
	go ProcessWithDeadline(ctx, "Process-C", 2*time.Second)
	println("Total goroutine", runtime.NumGoroutine())

	time.Sleep(4 * time.Second)
	println("Total goroutine", runtime.NumGoroutine())
}

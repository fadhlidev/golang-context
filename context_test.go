package golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(T *testing.T) {
	// Create an empty context using `context.Background`
	background := context.Background()
	fmt.Println(background)

	// Create an empty context using `context.TODO`
	todo := context.TODO()
	fmt.Println(todo)
}

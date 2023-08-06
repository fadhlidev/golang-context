package golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(T *testing.T) {
	parentCtx := context.Background()

	// Create a new context with parent and value (key-pair)
	childACtx := context.WithValue(parentCtx, "a", "Child-Context-A")
	childBCtx := context.WithValue(parentCtx, "b", "Child-Context-B")

	childCCtx := context.WithValue(childACtx, "c", "Sub-Child-Context-C")
	childDCtx := context.WithValue(childBCtx, "d", "Sub-Child-Context-D")

	fmt.Println(parentCtx)
	fmt.Println(childACtx)
	fmt.Println(childBCtx)
	fmt.Println(childCCtx)
	fmt.Println(childDCtx)

	// ctx.Value(key) -> Do i have value of this key ? returns the value : ask the parent
	fmt.Println(childACtx.Value("a")) // Child-Context-A
	fmt.Println(childACtx.Value("b")) // nil
	fmt.Println(childCCtx.Value("c")) // Sub-Child-Context-C
	fmt.Println(childCCtx.Value("a")) // Child-Context-A
	fmt.Println(childCCtx.Value("d")) // nil
	fmt.Println(childDCtx.Value("d")) // Sub-Child-Context-D
	fmt.Println(childDCtx.Value("b")) // Child-Context-B
}

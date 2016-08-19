package acogo

import (
	"testing"
	"fmt"
)

func TestSimpleMemory(t *testing.T) {
	var data = make([]byte, 1024)

	var memory Memory = NewSimpleMemory(true, data)

	memory.WriteStringAt(12, "你好 world.")

	fmt.Printf("%s\n", memory.ReadStringAt(12, uint64(len([]byte("Hello world.")))))

	memory.WriteWordAt(1, 12)

	fmt.Println(memory.ReadWordAt(1))
}

package acogo

import (
	"testing"
	"fmt"
)

func TestPagedMemory(t *testing.T) {
	var memory = NewPagedMemory(true)

	memory.WriteString(12, "你好 world.")

	fmt.Printf("%s\n", memory.ReadString(12, uint64(len([]byte("Hello world.")))))

	memory.WriteWord(1, 12)

	fmt.Println(memory.ReadWord(1))
}

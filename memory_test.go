package acogo

import (
	"testing"
	"fmt"
)

func TestMemory(t *testing.T) {
	var memory = NewMemory(0, true)

	memory.WriteString(12, "你好 world.")

	fmt.Printf("%s\n", memory.ReadString(12, len([]byte("Hello world."))))

	memory.WriteWord(1, 12)

	fmt.Println(memory.ReadWord(1))
}

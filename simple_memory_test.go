package acogo

import (
	"testing"
	"fmt"
)

func TestSimpleMemory(t *testing.T) {
	var data = make([]byte, 1024)

	var memory = NewSimpleMemory(true, data)

	memory.WriteString(12, "你好 world.")

	fmt.Printf("%s\n", memory.ReadString(12, uint64(len([]byte("Hello world.")))))

	memory.WriteWord(1, 12)

	fmt.Println(memory.ReadWord(1))
}

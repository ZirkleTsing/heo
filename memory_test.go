package acogo

import (
	"testing"
	"fmt"
)

func TestMemory(t *testing.T) {
	var memory = NewMemory(0, 0, true)

	memory.WriteString(0, "Hello world.")

	fmt.Printf("%s\n", memory.ReadString(0, len([]byte("Hello world."))))

	memory.WriteWord(0, 12)

	fmt.Println(memory.ReadWord(0))

}

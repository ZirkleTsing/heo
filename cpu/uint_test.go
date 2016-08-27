package cpu

import (
	"testing"
	"fmt"
)

func TestUint(t *testing.T) {
	var a = int32(-1)

	var b = uint32(a)

	var c = int(b)

	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}

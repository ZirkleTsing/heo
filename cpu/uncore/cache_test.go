package uncore

import (
	"testing"
	"github.com/mcai/acogo/cpu/mem"
	"github.com/mcai/acogo/cpu/uncore/uncoreutil"
	"fmt"
)

func TestCache(t *testing.T) {
	var geometry = mem.NewGeometry(32 * uncoreutil.KB, 16, 64)

	var cache = NewCache(geometry)

	cache.Sets[0].Lines[0].State = "test_state"

	fmt.Printf("len(cache.Sets): %d\n", len(cache.Sets))
	fmt.Printf("cache.Sets[0].Lines[0].State: %s\n", cache.Sets[0].Lines[0].State)
}

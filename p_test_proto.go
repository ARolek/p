//+build proto

package p

import (
	"testing"
)

//go:proto ignore
type L int
//go:proto ignore
var B = func(L)*L{return nil}

//go:proto B=Builtin L=B.lower
func TestB(t *testing.T) {
	var x L
	ptr := B(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}

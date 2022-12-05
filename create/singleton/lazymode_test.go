package singleton

import (
	"testing"
)

func TestLazyInstantiation(t *testing.T) {
	if LazyInstantiation() != LazyInstantiation() {
		t.Error(`Different objects appear in lazy mode`)
	}
}

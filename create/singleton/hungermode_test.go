package singleton

import (
	"testing"
)

func TestHugerInstantiation(t *testing.T) {
	if HugerInstantiation() != HugerInstantiation() {
		t.Error(`Different objects appear in Hungry Man mode`)
	}
}

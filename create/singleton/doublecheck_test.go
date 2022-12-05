package singleton

import (
	"testing"
)

func Test_dcInstantiation(t *testing.T) {
	if dcInstance != dcInstance {
		t.Error(`double check examination yields different results`)
	}
}

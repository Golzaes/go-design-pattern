package facade

import (
	"fmt"
	"testing"
)

func TestAPI(t *testing.T) {

	tests := []struct {
		name string
		api  API
		want string
	}{
		{t.Name(), NewAPI(), fmt.Sprintf(`%s\n%s`, `A model is tested`, `B model is tested`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.api.TestIt(); got != tt.want {
				t.Errorf("TestIt() = %v, want %v", got, tt.want)
			}
		})
	}

}

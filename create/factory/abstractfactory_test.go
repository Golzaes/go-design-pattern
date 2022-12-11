package factory

import "testing"

func TestSaveOrderMain(t *testing.T) {
	tests := []struct {
		name string
		d    OrderMainDAO
		want string
	}{
		{
			name: t.Name(),
			d:    &TxtMainDao{},
			want: `save main order message to txt`,
		},
		{
			name: t.Name(),
			d:    &RDBMainDao{},
			want: `save main order message to rdb`,
		},
		{
			name: t.Name(),
			d:    &JSONMainDao{},
			want: `save main order message to json`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.d
			if got := s.SaveOrderMain(); got != tt.want {
				t.Errorf("SaveOrderMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveOrderDetail(t *testing.T) {
	tests := []struct {
		name string
		d    OrderDetailDAO
		want string
	}{
		{
			name: t.Name(),
			d:    &TxtDetailDao{},
			want: `save detail order message to txt`,
		},
		{
			name: t.Name(),
			d:    &RDBDetailDao{},
			want: `save detail order message to rdb`,
		},
		{
			name: t.Name(),
			d:    &JSONDetailDao{},
			want: `save detail order message to json`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.d
			if got := s.SaveOrderDetail(); got != tt.want {
				t.Errorf("SaveOrderMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

package factory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSay(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Activity
	}{
		{t.Name(), args{name: `hello`}, &SayHello{}},
		{t.Name(), args{name: `Hello`}, &SayHello{}},
		{t.Name(), args{name: `HeLLo`}, &SayHello{}},
		{t.Name(), args{name: `HELLO`}, &SayHello{}},

		{t.Name(), args{name: `hi`}, &SayHi{}},
		{t.Name(), args{name: `hI`}, &SayHi{}},
		{t.Name(), args{name: `Hi`}, &SayHi{}},
		{t.Name(), args{name: `HI`}, &SayHi{}},

		{t.Name(), args{name: `bye`}, &SayBye{}},
		{t.Name(), args{name: `Bye`}, &SayBye{}},
		{t.Name(), args{name: `bYe`}, &SayBye{}},
		{t.Name(), args{name: `BYE`}, &SayBye{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActivity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSay_Say(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{t.Name(), args{`Payne`}, fmt.Sprintf("Hi, %s\n", `Payne`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SayHi{}
			if got := s.Say(tt.args.name); got != tt.want {
				t.Errorf("Say() = %v, want %v", got, tt.want)
			}
		})
	}
}

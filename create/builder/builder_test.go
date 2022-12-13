package builder

import "testing"

func TestBuilder1(t *testing.T) {
	builder := &Build1{}
	director := NewDirector(builder)
	director.Construct()
	if res := builder.Result(); res != `123` {
		t.Errorf(`Builder1 fail expect 123 acture %s`, res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Build2{}
	director := NewDirector(builder)
	director.Construct()
	if res := builder.Result(); res != 6 {
		t.Errorf(`Build2 fail expect 6 acture %d`, res)
	}
}

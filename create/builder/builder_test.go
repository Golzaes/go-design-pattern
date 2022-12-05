package builder

import "testing"

func TestBuilder1_Result(t *testing.T) {
	type fields struct {
		result string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: t.Name(), fields: fields{`2`}, want: `2`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder1{
				result: tt.fields.result,
			}
			if got := b.Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder2_Result(t *testing.T) {
	type fields struct {
		result int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: t.Name(), fields: fields{2}, want: 2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder2{
				result: tt.fields.result,
			}
			if got := b.Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	if res := builder.Result(); res != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	if res := builder.Result(); res != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", res)
	}
}

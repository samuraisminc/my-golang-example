package main

import (
	"testing"
)

func TestSimpleRule_Apply(t *testing.T) {
	type fields struct {
		base int
		text string
	}
	type args struct {
		value Value
	}
	fs := fields{
		base: 3,
		text: "Fizzz",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Value
	}{
		{
			name:   "1 -> 1",
			fields: fs,
			args:   args{ToValue(1)},
			want:   ToValue(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &SimpleRule{
				base: tt.fields.base,
				text: tt.fields.text,
			}
			if got := sr.Apply(tt.args.value); got.String() != tt.want.String() {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

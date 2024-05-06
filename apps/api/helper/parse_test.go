package helper

import "testing"

func TestStringToUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test StringToUint",
			args: args{s: "1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUint(tt.args.s); got != tt.want {
				t.Errorf("StringToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintToString(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test UintToString",
			args: args{i: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToString(tt.args.i); got != tt.want {
				t.Errorf("UintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

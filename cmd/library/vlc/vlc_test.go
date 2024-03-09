package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic test",
			str:  "This text will be changed, Dear",
			want: "!this text will be changed, !dear",
		},
		{
			name: "special symbols",
			str:  "?!@#$%^&*()_+./,",
			want: "?!@#$%^&*()_+./,",
		},
		{
			name: "upper case",
			str:  "HELLO",
			want: "!h!e!l!l!o",
		},
		{
			name: "numbers",
			str:  "1234567890",
			want: "1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!hello",
			want: "001000001110100100100100110001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name : "base test",
			args : args{
				bStr : "001000001110100100100100110001",
				chunkSize : 8,
			},
			want: BinaryChunks{"00100000", "11101001", "00100100", "11000100"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

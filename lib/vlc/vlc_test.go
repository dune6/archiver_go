package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "test text 1", args: "test1", want: "test1"},
		{name: "test text 2", args: "TEST2", want: "!t!e!s!t2"},
		{name: "test text 3", args: "!!!Test !3", want: "!!!!test !3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.args); got != tt.want {
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
		{name: "test encode bin 1", str: "!go", want: "001000000010010001"},
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
			name: "splitByChunks 1",
			args: args{bStr: "001000000010010001", chunkSize: 8},
			want: BinaryChunks{"00100000", "00100100", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name      string
		binChunks BinaryChunks
		want      HexChunks
	}{
		{
			name:      "encode binary chunks 1",
			binChunks: BinaryChunks{"00100000", "00100100", "01000000"},
			want:      HexChunks{"20", "24", "40"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.binChunks.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

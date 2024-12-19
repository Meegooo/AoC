package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"
)

//go:embed test.txt
var example string

//go:embed actual.txt
var actual string

func Test(t *testing.T) {
	tests := []struct {
		name string
		fun  func(io.Reader) string
		want string
	}{
		//{
		//	name: "part1",
		//	fun:  part1,
		//	want: "1928",
		//},
		{
			name: "part2",
			fun:  part2,
			want: "2858",
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("Example %v", testCase.name), func(t *testing.T) {
			if got := testCase.fun(strings.NewReader(example)); got != testCase.want {
				t.Errorf("%s() = %v, want %v", testCase.name, got, testCase.want)
			}
		})
		t.Run(fmt.Sprintf("Actual %v", testCase.name), func(t *testing.T) {
			got := testCase.fun(strings.NewReader(actual))
			t.Logf("%v", got)
		})
	}

}

func Test2(t *testing.T) {
	tests := []struct {
		name  string
		fun   func(io.Reader) string
		input string
		want  string
	}{
		{
			name:  "1",
			fun:   part1,
			input: "09091",
			want:  "0",
		},
		{
			name:  "2",
			fun:   part1,
			input: example,
			want:  "1928",
		},
		{
			name:  "3",
			fun:   part1,
			input: "12345",
			want:  "60",
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("Actual %v", testCase.name), func(t *testing.T) {
			if got := testCase.fun(strings.NewReader(testCase.input)); got != testCase.want {
				t.Errorf("%s() = %v, want %v", testCase.name, got, testCase.want)
			}
		})
	}

}

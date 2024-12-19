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
		{
			name: "part1",
			fun:  part1,
			want: "10092",
		},
		{
			name: "part2",
			fun:  part2,
			want: "9021",
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

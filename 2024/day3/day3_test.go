package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"
)

var example1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var example2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

//go:embed actual.txt
var actual string

func Test(t *testing.T) {
	tests := []struct {
		name    string
		fun     func(io.Reader) string
		want    string
		example string
	}{
		//{
		//	name: "part1",
		//	fun:  part1,
		//	want: "161",
		//	example: example1
		//},
		{
			name:    "part2",
			fun:     part2,
			want:    "48",
			example: example2,
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("Example %v", testCase.name), func(t *testing.T) {
			if got := testCase.fun(strings.NewReader(testCase.example)); got != testCase.want {
				t.Errorf("part1() = %v, want %v", got, testCase.want)
			}
		})
		t.Run(fmt.Sprintf("Actual %v", testCase.name), func(t *testing.T) {
			got := testCase.fun(strings.NewReader(actual))
			t.Logf("%v", got)
		})
	}

}

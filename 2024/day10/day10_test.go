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
			want: "36",
		},
		{
			name: "part2",
			fun:  part2,
			want: "81",
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
			name:  "0",
			fun:   part1,
			input: "0123\n1234\n8765\n9876",
			want:  "1",
		},
		{
			name:  "1",
			fun:   part1,
			input: "...0...\n...1...\n...2...\n6543456\n7.....7\n8.....8\n9.....9",
			want:  "2",
		},
		{
			name:  "2",
			fun:   part1,
			input: "..90..9\n...1.98\n...2..7\n6543456\n765.987\n876....\n987....",
			want:  "4",
		},
		{
			name:  "3",
			fun:   part1,
			input: "10..9..\n2...8..\n3...7..\n4567654\n...8..3\n...9..2\n.....01",
			want:  "3",
		},
		{
			name:  "4",
			fun:   part1,
			input: "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732",
			want:  "36",
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

package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestConvertToDigit(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "Convert 'one' to 1", input: "one", want: 1},
		{name: "Convert 'two' to 2", input: "two", want: 2},
		{name: "Convert 'three' to 3", input: "three", want: 3},
		{name: "Convert 'four' to 4", input: "four", want: 4},
		{name: "Convert 'five' to 5", input: "five", want: 5},
		{name: "Convert 'six' to 6", input: "six", want: 6},
		{name: "Convert 'seven' to 7", input: "seven", want: 7},
		{name: "Convert 'eight' to 8", input: "eight", want: 8},
		{name: "Convert 'nine' to 9", input: "nine", want: 9},
		{name: "Convert '1' to 1", input: "1", want: 1},
		{name: "Convert '2' to 2", input: "2", want: 2},
		{name: "Convert '3' to 3", input: "3", want: 3},
		{name: "Convert '4' to 4", input: "4", want: 4},
		{name: "Convert '5' to 5", input: "5", want: 5},
		{name: "Convert '6' to 6", input: "6", want: 6},
		{name: "Convert '7' to 7", input: "7", want: 7},
		{name: "Convert '8' to 8", input: "8", want: 8},
		{name: "Convert '9' to 9", input: "9", want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToDigit(tt.input)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestMatchDigits(t *testing.T) {
	line := "eighthree"
	got := matchDigits(line)
	want := []string{"eight", "three"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestExtractFirstAndLastDigitsFromLine(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantFirst int
		wantLast  int
	}{
		{
			name:      "Two digits as first and last characters in the string",
			line:      "two1nine",
			wantFirst: 2,
			wantLast:  9,
		},
		{
			name:      "",
			line:      "eightwothree",
			wantFirst: 8,
			wantLast:  3,
		},
		{
			name:      "",
			line:      "abcone2threexyz",
			wantFirst: 1,
			wantLast:  3,
		},
		{
			name:      "",
			line:      "xtwone3four",
			wantFirst: 2,
			wantLast:  4,
		},
		{
			name:      "",
			line:      "4nineeightseven2",
			wantFirst: 4,
			wantLast:  2,
		},
		{
			name:      "zoneight234",
			line:      "zoneight234",
			wantFirst: 1,
			wantLast:  4,
		},
		{
			name:      "7pqrstsixteen",
			line:      "7pqrstsixteen",
			wantFirst: 7,
			wantLast:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotLast, _ := extractFirstAndLastDigitsFromLine(tt.line)
			wantFirst := tt.wantFirst
			wantLast := tt.wantLast

			if gotFirst != wantFirst {
				t.Errorf("gotFirst %d wantFirst %d", gotFirst, wantFirst)
			}

			if gotLast != wantLast {
				t.Errorf("gotLast %d wantLast %d", gotLast, wantLast)
			}
		})
	}
}

func TestExtractCalibrationValueFromLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "Two digits as first and last characters in the string",
			line: "1abc2",
			want: 12,
		},
		{
			name: "Two digits not as the first and last characters in the string",
			line: "pqr3stu8vwx",
			want: 38,
		},
		{
			name: "Two digits where there are other digits in the string",
			line: "a1b2c3d4e5f",
			want: 15,
		},
		{
			name: "One digit",
			line: "treb7uchet",
			want: 77,
		},
		{
			name: "",
			line: "eighthree",
			want: 83,
		},
		{
			name: "",
			line: "sevenine",
			want: 79,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractCalibrationValueFromLine(tt.line)
			want := tt.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		name  string
		lines string
		want  int
	}{
		{
			name: "",
			lines: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
		{
			name: "",
			lines: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.lines))
			got := calculateSum(scanner)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

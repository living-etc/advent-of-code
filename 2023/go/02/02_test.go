package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestIdFromLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "ID From Line",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := idFromLine(tt.line)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestGameFromLine(t *testing.T) {
	tests := []struct {
		line string
		want Game
	}{
		{
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: Game{id: 1, rounds: []Round{
				{Blue: 3, Red: 4},
				{Blue: 6, Red: 1, Green: 2},
				{Green: 2},
			}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.want
			got := gameFromLine(tt.line)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestPowerOfGame(t *testing.T) {
	tests := []struct {
		name string
		game string
		want int
	}{
		{
			name: "Game 1",
			game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: 48,
		},
		{
			name: "Game 2",
			game: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: 12,
		},
		{
			name: "Game 3",
			game: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: 1560,
		},
		{
			name: "Game 4",
			game: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: 630,
		},
		{
			name: "Game 5",
			game: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerOfGame(tt.game)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestCalculateResults(t *testing.T) {
	lines := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	scanner := bufio.NewScanner(strings.NewReader(lines))
	gotSum, gotPower := calculateResults(scanner)
	wantSum := 8
	wantPower := 2286

	if gotSum != wantSum {
		t.Errorf("got sum %d, want sum %d", gotSum, wantSum)
	}

	if gotPower != wantPower {
		t.Errorf("got power %d, want power %d", gotPower, wantPower)
	}
}

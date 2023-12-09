package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var constraints = map[string]int{
	"Red":   12,
	"Green": 13,
	"Blue":  14,
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func (round *Round) SetColour(colour string, result int) *Round {
	reflect.ValueOf(round).Elem().FieldByName(colour).Set(reflect.ValueOf(result))
	return round
}

type Game struct {
	id     int
	rounds []Round
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func idFromLine(line string) int {
	re := regexp.MustCompile(`^Game (.*):`)
	matches := re.FindStringSubmatch(line)
	match := matches[1]
	id, err := strconv.Atoi(match)
	check(err)
	return id
}

func roundsFromLine(line string) []Round {
	re := regexp.MustCompile(`.*:(.*)$`)
	matches := re.FindStringSubmatch(line)
	roundsLiteral := matches[1]

	var rounds []Round

	roundSplit := strings.Split(roundsLiteral, ";")
	for _, r := range roundSplit {
		scores := strings.Split(r, ",")

		round := &Round{}
		for _, score := range scores {
			details := strings.Split(score, " ")
			result := details[1]
			colour := details[2]

			resultInt, err := strconv.Atoi(result)
			check(err)

			colourTitle := strings.Title(colour)

			round = round.SetColour(colourTitle, resultInt)
		}

		rounds = append(rounds, *round)
	}

	return rounds
}

func gameFromLine(line string) Game {
	id := idFromLine(line)
	rounds := roundsFromLine(line)
	game := Game{id: id, rounds: rounds}
	return game
}

func idOfPossibleGame(line string) (id int) {
	game := gameFromLine(line)

	id = game.id

	for _, round := range game.rounds {
		if round.Blue > constraints["Blue"] ||
			round.Green > constraints["Green"] ||
			round.Red > constraints["Red"] {
			id = 0
		}
	}
	return
}

func powerOfGame(line string) (power int) {
	minimumValue := Round{}
	game := gameFromLine(line)

	for _, round := range game.rounds {
		if minimumValue.Red == 0 || round.Red > minimumValue.Red {
			minimumValue.Red = round.Red
		}
		if minimumValue.Green == 0 || round.Green > minimumValue.Green {
			minimumValue.Green = round.Green
		}
		if minimumValue.Blue == 0 || round.Blue > minimumValue.Blue {
			minimumValue.Blue = round.Blue
		}
	}

	power = minimumValue.Red * minimumValue.Blue * minimumValue.Green

	return
}

func calculateResults(scanner *bufio.Scanner) (sum int, power int) {
	for scanner.Scan() {
		line := scanner.Text()
		sum += idOfPossibleGame(line)
		power += powerOfGame(line)
	}
	return
}

func main() {
	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	sum, power := calculateResults(scanner)

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", power)
}

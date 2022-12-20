package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pair struct {
	Key string
	Val int
}

func main() {
	f, err := os.Open("./input.txt")

	my_choices := map[string]Pair{
		"X": {Key: "Rock", Val: 1},
		"Y": {Key: "Paper", Val: 2},
		"Z": {Key: "Scissors", Val: 3},
	}

	// enemy_choices := map[string]Pair {
	//   "A": my_choices["X"],
	//   "B": my_choices["Y"],
	//   "C": my_choices["Z"],
	// }
	//
	// results := map[int]string {
	//   0:"Loss",
	//   3:"Draw",
	//   6:"Win",
	// }

	game_results := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	if err != nil {
		log.Fatal("Error: Could not open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total_score := 0

	for scanner.Scan() {

		contents := scanner.Text()
		choices := strings.Split(contents, " ")
		total_score += game_results[choices[0]][choices[1]] + my_choices[choices[1]].Val

	}
	fmt.Println("TOTAL SCORE: ", total_score)
}

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

	my_choices := map[string]map[string]int{
    "A": {
      "X": 3,
      "Y": 1,
      "Z": 2,
    },
    "B": {
      "X": 1,
      "Y": 2,
      "Z": 3,
    },
    "C": {
      "X": 2,
      "Y": 3,
      "Z": 1,
    },
	}

	// enemy_choices := map[string]Pair {
	//   "A": my_choices["X"],
	//   "B": my_choices["Y"],
	//   "C": my_choices["Z"],
	// }
	//


	result := map[string]int{
		"X": 0, 
		"Y": 3,
		"Z": 6,
	}

	// game_results := map[string]map[string]int{
	// 	"A": {
	// 		"X": 3,
	// 		"Y": 6,
	// 		"Z": 0,
	// 	},
	// 	"B": {
	// 		"X": 0,
	// 		"Y": 3,
	// 		"Z": 6,
	// 	},
	// 	"C": {
	// 		"X": 6,
	// 		"Y": 0,
	// 		"Z": 3,
	// 	},
	// }

	if err != nil {
		log.Fatal("Error: Could not open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// total_score := 0
  part2_score := 0

	for scanner.Scan() {

		contents := scanner.Text()
		choices := strings.Split(contents, " ")
		// total_score += game_results[choices[0]][choices[1]] + my_choices[choices[1]].Val

    part2_score += result[choices[1]] + my_choices[choices[0]][choices[1]]
	}
  fmt.Println("Part 2 Final Score: ", part2_score)
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// func TestRange(t *testing.T) {
//
// 	f, err := os.Open("./input_test.txt")
//
// 	if err != nil {
// 		fmt.Println("Error: Could not open file")
// 	}
//
// 	defer f.Close()
//
// 	result := parseInput(f)
//
// 	if result != 2 {
// 		t.Logf("Error: %v did not match the expected value of 2", result)
// 	}
// }

func TestOverlap(t *testing.T) {

	f, err := os.Open("./input_test.txt")

	if err != nil {
		fmt.Println("Error: Could not open file")
	}

	defer f.Close()
  total := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		contents := scanner.Text()

		result := strings.FieldsFunc(contents, func(r rune) bool { return r == ',' || r == '-' })


		intResult := make([]int, len(result))

		for k, val := range result {
			intValue, err := strconv.Atoi(val)
			if err != nil {
				panic("Error: Couldnt convert string to int")
			}
			intResult[k] = intValue
		}


    if doPairsOverlap(intResult[:2], intResult[2:]) || doPairsOverlap(intResult[2:], intResult[:2]) {
      total++
    }
    
	}

  

  if total != 4 {
    t.Errorf("Failed: expected 4, got %v", total)
  }
}

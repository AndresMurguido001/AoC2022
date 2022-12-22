package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Find how many pairs full contain each other
func main()  {

  f, err := os.Open("./input.txt")

  if err != nil {
    fmt.Println("Error: Could not open file")
  }

  defer f.Close()

  total := parseInput(f)

  fmt.Println("Total: ", total)
  
}

func parseInput(f *os.File) int  {
  
  total := 0

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {

    contents := scanner.Text()

    result := strings.FieldsFunc(contents, func(r rune) bool { return r == ',' || r == '-' })

    fmt.Println("RESULT: ", result)
    
    intResult := make([]int, len(result)) 


    for k, val := range result {

      intValue, err := strconv.Atoi(val)

      if err != nil {
        panic("Error: Couldnt convert string to int")
      }

      intResult[k] = intValue

    }


    // Part 1
    // if isPairContained(intResult[:2], intResult[2:]) || isPairContained(intResult[2:], intResult[:2]){
    //   total++
    // } 
    if doPairsOverlap(intResult[:2], intResult[2:]) {
      total++
    }

  }
  return total
}

func isPairContained(pair1, pair2 []int)  bool {
  return pair1[0] >= pair2[0] && pair1[1] <= pair2[1]
}

func doPairsOverlap(pair1, pair2 []int) bool  {
  if (pair1[0] > pair2[0]) {
    pair1, pair2 = pair2, pair1
  }

  return pair1[1] >= pair2[0]
}

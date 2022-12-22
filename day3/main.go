package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main()  {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error: Could not open file")
	}

	defer f.Close()

  // total := parseInput(f)
  total := parseInputBy3(f)

  fmt.Println("TOTAL: ", total)

}

func parseInput(f *os.File) int {

	scanner := bufio.NewScanner(f)

  total := 0

	for scanner.Scan() {

		contents := scanner.Text()

    halfway := len(contents)/2
    part1 := contents[:halfway]
    part2 := contents[halfway:]

    result := rune(intersection(part1, part2))

    priority := 0

    if (unicode.IsUpper(result)) {
      priority = int(result) - 'A' + 26 + 1
    } else {
      priority = int(result) - 'a' + 1
    }

    total += int(priority)
	}
  return total
}

func parseInputBy3(f *os.File) int {

	scanner := bufio.NewScanner(f)

  total := 0
  buffer := []string{}

	for scanner.Scan() {

    contents := scanner.Text()

    buffer = append(buffer, contents)

    if len(buffer) == 3 {
      // begin running comparisongs

      result := rune(intersection2(buffer[0], buffer[1], buffer[2]))

      priority := 0

      if (unicode.IsUpper(result)) {
        priority = int(result) - 'A' + 26 + 1
      } else {
        priority = int(result) - 'a' + 1
      }

      total += int(priority)
      buffer = []string{}
    }
  }
  return total
}

func intersection(part1 string, part2 string) byte {
  for i := range part1 {
    for j := range part2 {

      if part1[i] == part2[j] {
        return part1[i]
      }
    }
  }
  return 0
}

func intersection2(part1 string, part2 string, part3 string) byte {

  for i := range part1 {
    for j := range part2 {
      for k := range part3 {

        if part1[i] == part2[j] {
          if part3[k] == part1[i] {
            return part1[i]
          }
        }

      }
    }
  }
  return 0
}


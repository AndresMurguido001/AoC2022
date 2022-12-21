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

  fmt.Println("TOTAL: ", total)

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

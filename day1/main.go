package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main()  {

  f, err := os.Open("./test.txt")

  if err != nil {
    log.Fatal("Error: Could not open file")
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  elfs := make([]int, 0)
  current_total := 0

  for scanner.Scan() {

    contents := scanner.Text()

    if strings.Compare(contents, "") == 0 {
    
      elfs = append(elfs, current_total)
      current_total = 0
      
    } else {
    
      num, _ := strconv.Atoi(contents)
      current_total += num
    
    }

  }

  if err := scanner.Err(); err != nil {
    log.Fatal("Error: Scanner encountered an error")
  }

  maxIndex := findMax(elfs)

  fmt.Println(elfs[maxIndex])

}
// Return index of maximum value
func findMax(n []int) int {
  mi := 0
  mv := 0
  for i, x := range n {

    if x > mv {
      mv = x
      mi = i
    }

  }

  return mi
}

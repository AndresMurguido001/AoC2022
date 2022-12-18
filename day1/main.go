package main

import (
	"bufio"
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

  for scanner.Scan() {

    contents := scanner.Text()
    elfs := make([]int, 0)
    current_total := 0

    if strings.Compare(contents, "\n") == 0 {

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

}

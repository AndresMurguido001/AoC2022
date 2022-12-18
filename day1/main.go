package main

import (
	"bufio"
	"log"
	"os"
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

    if strings.Compare(contents, "\n") == 0 {
      
    }

  }


  if err := scanner.Err(); err != nil {
    log.Fatal("Error: Scanner encountered an error")
  }

}

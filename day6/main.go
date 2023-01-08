package main

import (
	"bufio"
	"log"
	"os"
)

// Looks like a sliding window problem
func main()  {

  f, err := os.Open("./input.txt")

  if err != nil {
    log.Fatalf("Error: Could not open file")
  }

  
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {

    contents := scanner.Text()



  }
  
  
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Find how many pairs full contain each other

func main()  {

  f, err := os.Open("./input.txt")

  if err != nil {
    fmt.Println("Error: Could not open file")
  }

  defer f.Close()


  scanner := bufio.NewScanner(f)

  for scanner.Scan() {

    contents := scanner.Text()

    result := strings.FieldsFunc(contents, func(r rune) bool { return r == ',' || r == '-' })
    

    fmt.Println("RESULT: ", result)


  }
  
}

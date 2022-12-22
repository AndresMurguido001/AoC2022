package main

import (
	"os"
	"testing"
)

func TestIntersection2(t *testing.T)  {

  // lowercase r appears in all 3
  input1 := "vJrwpWtwJgWrhcsFMMfFFhFp"
  input2 := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
  input3 := "PmmdzqPrVvPwwTWBwg"


  result := intersection2(input1, input2, input3)
    
  if result != byte('r') {
    t.Fatalf("Error: %v is not equal to the expected value: %d", result, byte('r'))
  }
}

func TestInputParseBy3(t *testing.T) {
  // total priority should be 70: 18 + 52
  f, err:= os.Open("./test_input.txt")
  if err != nil {
    t.Fatal("Error: Could not open file")
  }

  defer f.Close()

  result := parseInputBy3(f)

  if result != 70 {
    t.Fatalf("Error: %d was not equal to expected value 70", result)
  }
}

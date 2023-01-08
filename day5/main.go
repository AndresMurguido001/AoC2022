package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type stack struct {
  items []rune
}

// remove items from idx (exclusive)
func (s *stack) remove(last_idx int) []rune  {
  removed_items := s.items[last_idx:]
  s.items = s.items[:last_idx]
  return removed_items
}

// add items while maintaining order
func (s *stack) add(new_items []rune) []rune {
  s.items = append(s.items, new_items...)
  return s.items
}

// Pop off the end
func (s *stack) pop() rune {
  if len(s.items) == 1 {
    return s.items[0]
  }

  last_idx := len(s.items) - 1
  tmp := s.items[last_idx]
  s.items = s.items[:last_idx]
  return tmp
}

// push to the end and return new list
func (s *stack) push(x rune) []rune {
  s.items = append(s.items, x)
  return s.items
}

// Need to add to bottom in order to Initialize stacks
// in the correct order
func (s *stack) prepend(x rune)  {
  s.items = append([]rune{x}, s.items...)
}

func (s *stack) String() string {
  var str string

  for _, v := range s.items {
    str += string(v) + " "
  }

  return str

}

type movement struct {
  numItems int
  moveFrom int
  moveTo int
}

func main()  {

  f, err := os.Open("./input.txt")

  if err != nil {
    log.Fatalf("Error: Could not open file")
  }

  
  scanner := bufio.NewScanner(f)
  
  // Initialize the 9 stacks
  stacks := make([]stack, 9)

  scanner.Scan()

  for (scanner.Text() != " 1   2   3   4   5   6   7   8   9 ") {
    for i, v := range scanner.Text() {
      if unicode.IsLetter(v) {
        idx := (i-1) / 4 // index of the stack
        stacks[idx].prepend(v)
      }
    }
    scanner.Scan()
  }

  scanner.Scan() // empty line

  for scanner.Scan() {
    var numItems, from, to int

    fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &numItems, &from, &to)

    // part 1
  //   for move :=0; move < numItems; move++{
		// 	stacks[to-1].push(stacks[from-1].pop())
		// }
    // part 2
    stacks[to-1].add(stacks[from-1].remove(len(stacks[from-1].items) - numItems))


  }

  // At this point we have constructed our columns
  for x := range stacks {
    fmt.Println("STACK: ", stacks[x].String())
  }
  if scanner.Err() != nil {
    log.Fatalf("Error: Scanner encountered an error")
  }
}


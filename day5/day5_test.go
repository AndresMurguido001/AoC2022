package main

import (
	"testing"
)

var initial_stack = stack{items: []rune{'A', 'B', 'C', 'D'}}

func TestStackPush(t *testing.T) {

	expected := []rune{'A', 'B', 'C', 'D', 'E'}
	initial_stack.push('E')

	for x := range initial_stack.items {
		if expected[x] != initial_stack.items[x] {
			t.Logf("Error: %v did not match expected value: %v", initial_stack.items[x], expected[x])
		}
	}

}

func TestStackPop(t *testing.T) {

	expected := []rune{'A', 'B', 'C' }

  item := initial_stack.pop()

  if item != 'C' {
    t.Logf("Error: Expected 'C' to be returned from pop()")
  }

  if len(initial_stack.items) != len(expected) {
    t.Logf("Error: Expected length to be 3")
  }
}


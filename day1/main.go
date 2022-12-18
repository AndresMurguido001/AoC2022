package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Key int
	Val int
}

type PairList []Pair

func main() {

	f, err := os.Open("./test.txt")

	if err != nil {
		log.Fatal("Error: Could not open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfs := make(PairList, 0)
	current_total := 0
  current_elf := 0

	for scanner.Scan() {

		contents := scanner.Text()

		if strings.Compare(contents, "") == 0 {

			elfs = append(elfs, Pair{ Key: current_elf, Val: current_total })
			current_total = 0
      current_elf++

		} else {

			num, _ := strconv.Atoi(contents)
			current_total += num

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: Scanner encountered an error")
	}

	// maxIndex := findMax(elfs)
	// fmt.Println(elfs[maxIndex])

  sortPairList(&elfs)

  top3_total := 0
  for _, v := range elfs[:3] {
    top3_total += v.Val
  }

  fmt.Println(top3_total)

}
func (a PairList) Len() int           { return len(a) }
func (a PairList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }


// ByValue implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByValue struct { PairList }
func (a ByValue) Less(i, j int) bool { return a.PairList[j].Val < a.PairList[i].Val }


func sortPairList(m* PairList) {
  sort.Sort(ByValue{ *m })
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

package main

import (
	"fmt"
	"os"
	g "./Gravity"
	r "./Radix_LSD"
	util "./Utils"
)

const (
	maxNum  int = 100
	arrSize int = 100
)

var (
	sortMap = make(map[string]SortMethod)
	gravitySortSlow = SortMethod {
		Short: "gravity_slow",
		Name: "Gravity Sort Slow",
		Sort: g.SortSlow,
	}.add()
	gravitySortFast = SortMethod {
		Short: "gravity_fast",
		Name: "Gravity Sort Fast",
		Sort: g.SortFast,
	}.add()
	gravitySortFastest = SortMethod {
		Short: "gravity_fastest",
		Name: "Gravity Sort Fastest",
		Sort: g.SortFastest,
	}.add()
	radixLSD = SortMethod {
		Short: "radix_lsd",
		Name: "Radix LSD",
		Sort: r.Sort,
	}.add()
)

type SortMethod struct {
	Sort func([]int, int, int) []int
	Name string
	Short string
}

func init() {
}

func (s SortMethod) add() SortMethod {
	sortMap[s.Short] = s
	return s
}

func main() {
	var o string
	var output *os.File

	fmt.Print("Output to Stdout or /dev/null (std/null): ")
	fmt.Scanln(&o)

	switch o {
	case "std":
		output = os.Stdout
	case "null":
		output = nil
	default:
		fmt.Println("No valid output defined")
		os.Exit(1)
	}

	fmt.Println("Now choose a sorting algorithm:", func() string {
		var out string
		var count int
		for _, algo := range sortMap {
			out += algo.Short
			if count != len(sortMap) - 1 {
				out += ", "
			} 
			count++
		}
		return out
	}())

	var algo string
	fmt.Scanln(&algo)

	random := util.MakeArray(arrSize, maxNum)

	if val, ok := sortMap[algo]; ok {
		fmt.Fprintln(output, val.Sort(random, arrSize, maxNum))		
		return
	}
	fmt.Println("nothing found")
}

package main

import (
	"fmt"

	"github.com/turbo-d/go-patterns/hof"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)

	s1Doubled := hof.Map(s1, func(x int) int {
		return x * 2
	})
	fmt.Println(s1Doubled)

	s1EvenOnly := hof.Filter(s1, func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	})
	fmt.Println(s1EvenOnly)

	isPositive := hof.Every(s1, func(x int) bool {
		if x > 0 {
			return true
		}
		return false
	})
	fmt.Println(isPositive)

	isEven := hof.Every(s1, func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	})
	fmt.Println(isEven)

	someDivisibleBy4 := hof.Some(s1, func(x int) bool {
		if x%4 == 0 {
			return true
		}
		return false
	})
	fmt.Println(someDivisibleBy4)

	sum := hof.Reduce(s1, func(accumulator, x int) int {
		return accumulator + x
	})
	fmt.Println(sum)

	firstDivBy3, _ := hof.Find(s1, func(x int) bool {
		if x%3 == 0 {
			return true
		}
		return false
	})
	fmt.Println(firstDivBy3)

	firstDivBy3Idx, _ := hof.FindIndex(s1, func(x int) bool {
		if x%3 == 0 {
			return true
		}
		return false
	})
	fmt.Println(firstDivBy3Idx)
}

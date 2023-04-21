package main

import "fmt"

func main() {
	numSeq := numSequence(0, 10)

	for _, num := range numSeq {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}

func numSequence(min int, max int) []int {
	a := []int{}

	for i := min; i <= max; i++ {
		a = append(a, i)
	}

	return a

}

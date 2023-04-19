package main

import (
	"golang.org/x/exp/constraints"
)

func main() {

}

func LargestElementInSlice[T constraints.Ordered](s []T) T {
	largest := s[0]
	for _, item := range s {
		if item > largest {
			largest = item
		}
	}
	return largest
}

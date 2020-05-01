package main

import (
	"fmt"
	"sort"
)

//BinarySearch ...
func BinarySearch(d []int, s int) bool {
	sort.Ints(d)
	length := len(d)
	middleIndex := (length / 2)
	result := false

	for middleIndex != 0 {
		if s == d[middleIndex] {
			result = true
			break
		} else {
			if s > d[middleIndex] {
				d = append(d[middleIndex:])
			} else {
				d = append(d[:middleIndex])
			}
		}
		middleIndex = len(d) / 2
	}

	if len(d) == 1 && d[0] == s {
		result = true
	}

	return result
}

func main() {
	d := []int{2}
	s := 2

	fmt.Println(BinarySearch(d, s))
}

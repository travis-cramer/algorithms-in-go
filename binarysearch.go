package algorithms

import "fmt"

func Binarysearch(array []string, target string) (index int) {
	var L int
	L = 0
	var R int
	R = len(array) - 1
	var middle int
	var done bool
	for !done {
		middle = (L + R) / 2
		if array[middle] < target {
			L = middle + 1
		} else if array[middle] > target {
			R = middle - 1
		} else {
			done = true
		}
		index = middle
	}
	return
}

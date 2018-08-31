package main

import (
	"fmt"
)


func main() {
	test_binarysearch()
}


func binarysearch(array []string, target string) (index int) {
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


func test_binarysearch() {
	// tests binarysearch, output should be 0...(len(array) - 1)
	array := [10]string{"apple", "banana", "carrot", "danish", "energy", "fun", "greatness", "hairy", "igneous", "jim"}
	for i := 0; i < len(array); i++ {
		var index int
		index = binarysearch(array[:], array[i])
		if index != i {
			fmt.Println("Failed.")
			return
		}
	}
	fmt.Println("Success!")
	return
}

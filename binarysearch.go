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

func test_binarysearch() {
	// tests binarysearch
	array1 := [10]string{"apple", "banana", "carrot", "danish", "energy", "fun", "greatness", "hairy", "igneous", "jim"}
	for i := 0; i < len(array1); i++ {
		var index int
		index = Binarysearch(array1[:], array1[i])
		if index != i {
			fmt.Println("Failed.")
			return
		}
	}
	array2 := [9]string{"apple", "banana", "carrot", "danish", "energy", "fun", "greatness", "hairy", "igneous"}
	for i := 0; i < len(array2); i++ {
		var index int
		index = Binarysearch(array2[:], array2[i])
		if index != i {
			fmt.Println("Failed.")
			return
		}
	}
	fmt.Println("Success!")
	return
}

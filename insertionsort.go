package main

import "fmt"


func main() {
	array := []int{5, 4, 3, 2, 1, 13, 12, 14, 3, 4, 4, 4, 8, 7, 4, 9}
	fmt.Println(array)
	fmt.Println("---------->")
	fmt.Println(insertionsort(array[:]))
}


func insertionsort(A []int) []int {
	// we start with the second element of the array (to compare it to the first)
	for i := 1; i < len(A); i++ {
		for j := i - 1; j >=0; j-- {
			// j runs backwards through A from i
			if A[j] > A[j+1] {
				//swap
				temp := A[j+1]
				A[j+1]=A[j]
				A[j]=temp
			}
		}
	}

	return A
}
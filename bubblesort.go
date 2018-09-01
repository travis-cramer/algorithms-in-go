package algorithms

func Bubblesort(A []int) []int {
	var sorted bool  //checks whether array is sorted or not
	for !sorted {
		var i int  //index
		var swapped bool  //checks whether anything swapped on this loop
		for ; i < len(A) - 1; i++ {
			if A[i] > A[i+1] {
				//swap
				var temp int
				temp = A[i+1]
				A[i+1] = A[i]
				A[i] = temp
				swapped = true  //something swapped
			}
		}
		if !swapped {
			sorted = true  //nothing swapped, array is sorted
		}

	}
	return A
}

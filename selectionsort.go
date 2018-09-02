package algorithms

func Selectionsort(A []int) []int {
	for i := 0; i < len(A); i++ {
		// find where the lowest, not-yet-sorted element is
		lowest_index := i
		for j := i+1; j < len(A); j++ {
			if A[j] < A[lowest_index] {
				lowest_index = j
			}
		}
		//swap lowest with element at index i
		temp := A[i]
		A[i] = A[lowest_index]
		A[lowest_index] = temp
	}
	return A
}

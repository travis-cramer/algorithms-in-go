package algorithms

func Insertionsort(A []int) []int {
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
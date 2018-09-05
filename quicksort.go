package algorithms

func Quicksort(A []int) []int {
	QuicksortRecursive(A[:], 0, len(A) - 1)
	return A
}

func QuicksortRecursive(A []int, iBegin int, iEnd int) {
	if iBegin == iEnd {
		// it's sorted, length of subarray is 1
		return
	}
	iPivot := pivotsort(A[:], iBegin, iEnd)

	if iPivot > iBegin {
		QuicksortRecursive(A[:], iBegin, iPivot-1)
	}
	if iPivot < iEnd {
		QuicksortRecursive(A[:], iPivot+1, iEnd)
	}
	return
}

func pivotsort(A []int, iBegin int, iEnd int) (iPivot int) {
	/* Takes in a subarray, sorts it such that all elements greater than the pivot element are to the right of the pivot
	and all elements less than the pivot element are to the left of the pivot element.
	The last element, A[iEnd], is assigned as the pivot element.
	NOTE: iBegin and iEnd are both inclusive. */
	i := iBegin
	j := iEnd - 1 // because iEnd is the index of the pivot element, which we don't iterate through
	for i < j {
		if A[i] > A[iEnd] && A[j] < A[iEnd] {
			// swap A[i] and A[j]
			temp := A[i]
			A[i] = A[j]
			A[j] = temp
			i++
			j--
		} else if A[i] > A[iEnd] {
			j--
		} else if A[j] < A[iEnd] {
			i++
		} else {
			// if A[i] == A[j] == A[iEnd] just move on to next pair of integers
			i++
			j--
		}
	}
	if i == j {
		if A[j] > A[iEnd] {
			// swap A[j] and A[iEnd]
			temp := A[j]
			A[j] = A[iEnd]
			A[iEnd] = temp
			iPivot = j
		} else {
			// swap A[j+1] and A[iEnd]
			temp := A[j+1]
			A[j+1] = A[iEnd]
			A[iEnd] = temp
			iPivot = j+1
		}
	} else if j < i {
		// swap A[i] and A[iEnd]
		temp := A[i]
		A[i] = A[iEnd]
		A[iEnd] = temp
		iPivot = i
	}
	return iPivot
}

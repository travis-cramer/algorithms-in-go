package algorithms

func Mergesort(A []int) []int {
	n := len(A)
	B := make([]int, n)         // instantiate a work array
	copyarray(A[:], 0, n, B[:]) // copy contents of A into B
	for run_size := 1; run_size < n; run_size = run_size * 2 {
		var l int         // a "left" index
		m := l + run_size // a "middle" index
		r := m + run_size // a "right index"
		// don't run merge if m is >= n (then A[m:r] is null, i.e. nothing to merge)
		for m < n {
			// r can never be greater than n, otherwise index error
			if r > n {
				r = n
			}
			merge(A[:], l, m, r, B[:])
			// increment each index by twice the run_size
			l = l + run_size*2
			m = m + run_size*2
			r = r + run_size*2
		}
	}
	return A
}

func merge(A []int, l int, m int, r int, B []int) {
	// merges and sorts A[l:m] and A[m:r] (which are already sorted) into B[l:r]
	// l, m, and r are left, middle, and right indexes
	i := l // index running through A[l:m]
	j := m // index running through A[m:r]
	k := l // index running through B[l:r]
	for k < r {
		if i < m && j < r {
			// compare and insert
			if A[i] <= A[j] {
				B[k] = A[i]
				i++
			} else {
				B[k] = A[j]
				j++
			}
			// if either A[l:m] or A[m:r] is all inserted, then insert the remains of the other slice
		} else if i < m {
			B[k] = A[i]
			i++
		} else {
			B[k] = A[j]
			j++
		}
		k++
	}
	// copy B[l:r] (sorted) into A[l:r] (unsorted)
	copyarray(B[:], l, r, A[:])
}

func copyarray(A []int, l int, r int, B []int) {
	// copies B[l:r] into A[l:r]
	// note r is exclusive
	for i := l; i < r; i++ {
		B[i] = A[i]
	}
}

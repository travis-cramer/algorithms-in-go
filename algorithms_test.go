package algorithms_test

import (
	"testing"
	"math/rand"
	algorithms "github.com/travis-cramer/algorithms-in-go"
)

var big_array [1000]int
var bigger_array [10000]int
//var biggest_array [100000]int

var big_array_copy [1000]int
var bigger_array_copy [10000]int
//var biggest_array_copy [100000]int

var cases []struct{in, want []int}

func init() {
	populateArrayWithRandInts(big_array[:])
	populateArrayWithRandInts(bigger_array[:])
	//populateArrayWithRandInts(biggest_array[:])

	copyarrayintoarray(big_array[:], big_array_copy[:])
	copyarrayintoarray(bigger_array[:], bigger_array_copy[:])
	//copyarrayintoarray(biggest_array[:], biggest_array_copy[:])

	cases = []struct {
		in, want []int
	}{
		{big_array[:], algorithms.Quicksort(big_array_copy[:])},
		{bigger_array[:], algorithms.Quicksort(bigger_array_copy[:])},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{12, 3, -2, 8, 9, 1, 6}, []int{-2, 1, 3, 6, 8, 9, 12}},
	}
}

func populateArrayWithRandInts(A []int) {
	rand.Seed(42)
	for i := 0; i < len(A); i++ {
		A[i] = rand.Intn(len(A))
	}
}

func copyarrayintoarray(A []int, B []int) {
	for i := 0; i < len(A); i++ {
		B[i] = A[i]
	}
}

func TestIntSorters(t *testing.T) {

	// Bubblesort
	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		got := algorithms.Bubblesort(in)
		for i := 0; i < len(in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Bubblesort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}

	// Insertionsort
	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		got := algorithms.Insertionsort(in)
		for i := 0; i < len(in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Insertionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
	//Selectionsort
	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		got := algorithms.Selectionsort(in)
		for i := 0; i < len(in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Selectionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}

	//Mergesort
	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		got := algorithms.Mergesort(in)
		for i := 0; i < len(in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Mergesort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}

	//Quicksort
	for _, c := range cases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		got := algorithms.Quicksort(in)
		for i := 0; i < len(in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Quicksort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}

func TestStringSearchers(t *testing.T) {
	cases := []struct {
		array  []string
		target string
		index  int
	}{
		{[]string{"apple"}, "apple", 0},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "banana", 1},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "greatness", 6},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "fun", 5},
	}
	for _, c := range cases {
		got := algorithms.Binarysearch(c.array, c.target)
		if got != c.index {
			t.Errorf("Binarysearch(%q, %q) == %q, want %q", c.array, c.target, got, c.index)
		}
	}
}

func BenchmarkBubblesort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			in := make([]int, len(c.in))
			copy(in, c.in)
			algorithms.Bubblesort(in)
		}
	}
}

func BenchmarkSelectionsort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			in := make([]int, len(c.in))
			copy(in, c.in)
			algorithms.Selectionsort(in)
		}
	}
}

func BenchmarkInsertionsort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			in := make([]int, len(c.in))
			copy(in, c.in)
			algorithms.Insertionsort(in)
		}
	}
}

func BenchmarkMergesort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			in := make([]int, len(c.in))
			copy(in, c.in)
			algorithms.Mergesort(in)
		}
	}
}

func BenchmarkQuicksort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			in := make([]int, len(c.in))
			copy(in, c.in)
			algorithms.Quicksort(in)
		}
	}
}

package algorithms

import "testing"

func TestIntSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, -4, 3, -2, 1}, []int{-4, -2, 1, 3, 5}},
		{[]int{5, -4, 3, 6, -2, 1}, []int{-4, -2, 1, 3, 5, 6}},
		{[]int{12, 5, 8, 3, 9, 100, 10, 11}, []int{3, 5, 8, 9, 10, 11, 12, 100}},
	}
	for _, c := range cases {
		got := Bubblesort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Bubblesort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
	for _, c := range cases {
		got := Insertionsort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Insertionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
	for _, c := range cases {
		got := Selectionsort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Selectionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
	for _, c := range cases {
		got := Mergesort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Mergesort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}

func TestStringSearch(t *testing.T) {
	cases := []struct {
		array []string; target string; index int
	}{
		{[]string{"apple"}, "apple", 0},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "banana", 1},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "greatness", 6},
		{[]string{"apple", "banana", "cherry", "danish", "energy", "fun", "greatness", "hairy"}, "fun", 5},
	}
	for _, c := range cases {
		got := Binarysearch(c.array, c.target)
		if got != c.index {
			t.Errorf("Binarysearch(%q, %q) == %q, want %q", c.array, c.target, got, c.index)
		}
	}
}

package algorithms

import "testing"

func TestBubblesort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, c := range cases {
		got := Bubblesort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Bubblesort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}

func TestInsertionsort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, c := range cases {
		got := Insertionsort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Insertionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}

func TestSelectionsort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, c := range cases {
		got := Selectionsort(c.in)
		for i := 0; i < len(c.in); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Selectionsort(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}

func TestBinarysearch(t *testing.T) {
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

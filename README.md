# algorithms-in-go
This is a package of fundamental CS algorithms that I've coded in Go. No other author's code was consulted in the writing of 
these algorithms, just gifs demonstrating their procedures, such as this one for insertionsort:  
https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif

# why
I'm undertaking this project to study the typical algorithms that engineers think of when recruiters mention "Data Structures 
and Algorithms", and to learn Go along the way.

# usage
Use Golang's `go get` command to clone this repo into your local go workspace:
```
go get "github.com/travis-cramer/algorithms-in-go"
```
This will place the repository in the /src/ folder of your go workspace, as specified by your GOPATH (see: 
https://golang.org/doc/code.html).  
Now, you can import the package into your go program with:
```
import algorithms "github.com/travis-cramer/algorithms-in-go"
```
Of course, the `algorithms` package alias can be changed to whatever you prefer.

There are currently 6 algorithms in this package:
* `Bubblesort()`, which takes in an array of ints (`[]int`) and returns the same, sorted in numerical order
* `Insertionsort()`, same as `Bubblesort()`
* `Selectionsort()`, same as `Insertionsort()`
* `Mergesort()`, same as `Selectionsort()`
* `Quicksort()`, same as `Mergesort()`
* `Binarysearch()`, which takes in an array of strings (`[]string`) **sorted in alphabetical order** as the first 
argument and a `string` (that is an element of the array) as the second. Returns the index (`int`) of the given 
string within the array

# note
Note that these functions were written to accept arrays of any length, so you must input them as slices:
```
algorithms.Bubblesort(my_array[:])
```

# example
A quick example that you can copy and paste into a Go script to test that the package is imported properly:
```
package main

import (
	"fmt"
	algorithms "github.com/travis-cramer/algorithms-in-go"
)

func main() {
	array := [5]int{2, 5, 1, 3, 4}
	fmt.Println(array)
	algorithms.Mergesort(array[:])
	fmt.Println(array)
}
```
As output you should see two arrays: the first unsorted and the second sorted, in numerical order.

# testing
For a demonstration of the performance of each of these sorting algorithms, use Go's benchmark feature in its built-in testing 
framework. Run the following in this repo's directory:
```
go test -bench .
```

# help me
Feedback is appreciated. If you notice anything wrong, something that I can improve on, or any way that I can conform 
better to the conventions of Go, please let me know by opening an issue on this repo.

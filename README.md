# algorithms-in-go
This is a package of misc. algorithms that I've coded in Go. No other author's code was consulted in the writing of 
these functions, just gifs demonstrating the procedure of the algorithm, such as this one for insertionsort:  
https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif

# why
I'm undertaking this project to learn Go and study the typical algorithms that people think of when
recruiters mention "Data Structures and Algorithms".

# usage
```
go get "github.com/travis-cramer/algorithms-in-go"
```
Use Golang's `go get` command to clone this repo into your local go workspace. Now, you can import it into your own go files 
with:
```
import algorithms "github.com/travis-cramer/algorithms-in-go"
```
Of course, the `algorithms` package alias can be changed to whatever you like.

There are currently 6 algorithms in this package:
* `Bubblesort()`, which takes in an array of ints (`[]int`) and returns the same, sorted in numerical order.
* `Insertionsort()`, same as `Bubblesort()`.
* `Selectionsort()`, same as `Insertionsort()`.
* `Mergesort()`, same as `Selectionsort()`.
* `Quicksort()`, same as `Mergesort()`.
* `Binarysearch()`, which takes in an array of strings (`[]string`) **sorted in alphabetical order** as the first 
argument and a `string` (that is an element of the array) as the second. Returns the index (`int`) of the given 
string within the array.

Note that in Go since these functions accept arrays of any length, we input them as arguments by slicing them. Ex:
```
my_sorted_array = algorithms.Bubblesort(my_array[:])
```

# help me
Feedback is appreciated. If you notice anything wrong, something that I can improve on, or any way that I can conform 
better to the conventions of Golang, please let me know by opening an issue on this repo.

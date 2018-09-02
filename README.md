# algorithms-in-go
This is a package of misc. algorithms that I've coded in Go.

# why
I'm undertaking this project to learn Go and study the typical algorithms that people think of when
recruiters mention "Data Structures and Algorithms".

# Usage

```
import algorithms "github.com/travis-cramer/algorithms-in-go"
```  
Of course, the `algorithms` variable name can be changed to whatever you like.

There are currently 3 methods in this package:
* `Bubblesort()`, which takes in an array of ints (`[]int`) and returns the same, sorted in numerical order.
* `Insertionsort()`, same as `Bubblesort()`.
* `Binarysearch()`, which takes in an array of strings (`[]string`) as the first argument and a `string` as the 
second. Returns the index (`int`) of the given string within the array.

Note that in Go, we input arrays as arguments by slicing them. Ex:  
```
my_sorted_array = algorithms.Bubblesort(my_array[:])
```

# help me
Feedback is appreciated. We are all sometimes noobs in life, but me more than others.

# Golang Sorting Algorithms

 Some sorting algorithms written in [Golang](https://golang.org/)

## Installation

`$ go get github.com/indmind/sort`

## Usage

```Go
package main

import (
	"fmt"

	"github.com/indmind/sort"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}

	sort.Bubble(arr) // etc...

	fmt.Println(arr) // [1 2 3 4 5]

	/*
		!!PUT INTO VARIABLE FOR 'MERGE' sort AND 'QUICK' sort!!
		exp:
			arr := []int{5, 4, 3, 2, 1}

			sorted := sort.Quick(arr) // arr : [5 4 3 2 1] and sorted : [1 2 3 4 5]

	*/
}
```

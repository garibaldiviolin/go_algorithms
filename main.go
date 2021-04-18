package main

import (
	"fmt"
	binary_search "go_algorithms/binary_search"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binary_search.BinarySearch(list, 3))
}

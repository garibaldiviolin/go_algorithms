package main

import (
	"fmt"
	binary_search "go_algorithms/binary_search"
)

func main() {
	list := make([]int, 10_000_000)
    for i := range list {
        list[i] = i + 1
    }

	fmt.Println(binary_search.BinarySearch(list, -1))
	fmt.Println(binary_search.BinarySearch(list, 0))
	fmt.Println(binary_search.BinarySearch(list, 1))
	fmt.Println(binary_search.BinarySearch(list, 7_777_777))
	fmt.Println(binary_search.BinarySearch(list, 10_000_000))
	fmt.Println(binary_search.BinarySearch(list, 10_000_001))
	fmt.Println(binary_search.BinarySearch(list, 10_000_002))
}

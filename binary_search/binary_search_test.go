package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := make([]int, 10)
	for i := range list {
		list[i] = i + 1
	}

	parameters := []struct {
		number            int
		expected_found    bool
		expected_position int
	}{
		{-1, false, -1},
		{0, false, -1},
		{1, true, 0},
		{2, true, 1},
		{3, true, 2},
		{4, true, 3},
		{5, true, 4},
		{6, true, 5},
		{7, true, 6},
		{8, true, 7},
		{9, true, 8},
		{10, true, 9},
		{11, false, -1},
		{12, false, -1},
	}

	for i := range parameters {
		found, position := BinarySearch(list, parameters[i].number)

		if found != parameters[i].expected_found {
			t.Errorf("Number 5 expected to be found.")
		}
		if position != parameters[i].expected_position {
			t.Errorf("Position expected to be 4.")
		}
	}

}

package binary_search

func BinarySearch(list []int, number int) (bool, int) {
	list_length := len(list)
	start_position := 0
	end_position := list_length - 1

	for {
		middle_position := int((start_position + end_position) / 2)
		middle_number := list[middle_position]

		if end_position < 0 || start_position >= list_length {
			break
		}

		if middle_number == number {
			return true, middle_position
		} else if middle_number < number {
			start_position = middle_position + 1
		} else {
			end_position = middle_position - 1
		}
	}

	return false, -1
}

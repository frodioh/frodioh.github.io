package search

func binarySearch(list []int, item int) int  {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == item {
			return mid
		} else if list[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
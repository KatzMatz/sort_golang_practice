package sort

func MergeSortInt(array []int) []int {
	var size int = len(array)

	if size > 1 {
		// Split part
		var mid_idx int = size / 2
		var left_arr []int = MergeSortInt(array[:mid_idx])
		var right_arr []int = MergeSortInt(array[mid_idx:])

		// Merge part
		var sorted []int = make([]int, size)
		var left_idx int = 0
		var right_idx int = 0

		for i := 0; i < size; i++ {
			if (right_idx == size-mid_idx) || (left_idx < mid_idx && left_arr[left_idx] <= right_arr[right_idx]) {
				sorted[i] = left_arr[left_idx]
				left_idx++
			} else {
				sorted[i] = right_arr[right_idx]
				right_idx++
			}
		}
		return sorted

	} else {
		return array
	}

}

func QuickSortInt(array []int) []int {
	var size int = len(array)
	if size > 1 {
		var sorted []int = make([]int, size)

		var pivot int = array[0]
		var left []int
		var right []int

		for _, item := range array[1:] {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortInt(left)
		right = QuickSortInt(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return array
	}

}

func HeapSortInt(array []int) []int {
	var size int = len(array)
	// var i int = (size / 2) - 1
	var sorted []int = make([]int, size)
	copy(sorted, array)
	var max_idx int = 0
	for unsorted_length := size; unsorted_length > 1; unsorted_length-- {
		// build heap tree
		for i := (int(unsorted_length/2) - 1); i >= 0; i-- {
			max_idx = i
			if sorted[max_idx] < sorted[2*i+1] {
				max_idx = 2*i + 1
			}
			if (unsorted_length > 2*i+2) && sorted[max_idx] < sorted[2*i+2] {
				max_idx = 2*i + 2
			}

			if max_idx != i {
				sorted[i], sorted[max_idx] = sorted[max_idx], sorted[i]
			}
		}

		// swap root and end of slice
		sorted[0], sorted[unsorted_length-1] = sorted[unsorted_length-1], sorted[0]
	}
	return sorted
}

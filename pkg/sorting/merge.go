package sorting

// MergeSorter реализует сортировку слиянием.
type MergeSorter struct{}

func (m MergeSorter) Name() string { return "MergeSort" }

func (m MergeSorter) Sort(data []int) []int {
	arr := make([]int, len(data))
	copy(arr, data)
	mergeSort(arr, 0, len(arr)-1)
	return arr
}

func mergeSort(a []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(a, low, mid)
	mergeSort(a, mid+1, high)
	merge(a, low, mid, high)
}

func merge(a []int, low, mid, high int) {
	temp := make([]int, 0, high-low+1)
	left, right := low, mid+1
	for left <= mid && right <= high {
		if a[left] <= a[right] {
			temp = append(temp, a[left])
			left++
		} else {
			temp = append(temp, a[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, a[left])
		left++
	}
	for right <= high {
		temp = append(temp, a[right])
		right++
	}
	copy(a[low:high+1], temp)
}

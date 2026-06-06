package sorting

// QuickSorter реализует быструю сортировку.
type QuickSorter struct{}

func (q QuickSorter) Name() string { return "QuickSort" }

func (q QuickSorter) Sort(data []int) []int {
	arr := make([]int, len(data))
	copy(arr, data)
	quick(arr, 0, len(arr)-1)
	return arr
}

func quick(a []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(a, low, high)
	quick(a, low, p-1)
	quick(a, p+1, high)
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return i
}

package sorting

// BubbleSorter реализует пузырьковую сортировку.
type BubbleSorter struct{}

func (b BubbleSorter) Name() string { return "BubbleSort" }

func (b BubbleSorter) Sort(data []int) []int {
	arr := make([]int, len(data))
	copy(arr, data)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

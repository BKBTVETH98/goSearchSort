package searching

import "fmt"

// BinarySearcher реализует бинарный поиск.
type BinarySearcher struct{}

func (b BinarySearcher) Name() string { return "BinarySearch" }

func (b BinarySearcher) Search(data []int, target int) (int, error) {
	left, right := 0, len(data)-1
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == target {
			return mid, nil
		}
		if data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, fmt.Errorf("error: значение не найдено")
}

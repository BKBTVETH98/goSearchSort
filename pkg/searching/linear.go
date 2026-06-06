package searching

import "fmt"

// LinearSearcher реализует линейный поиск.
type LinearSearcher struct{}

func (l LinearSearcher) Name() string { return "LinearSearch" }

func (l LinearSearcher) Search(data []int, target int) (int, error) {
	for i, v := range data {
		if v == target {
			return i, nil
		}
	}
	return 0, fmt.Errorf("error: значение не найдено")
}

package sorting

// Sorter описывает общий интерфейс алгоритмов сортировки.
type Sorter interface {
	Name() string
	Sort([]int) []int
}

package searching

type Searcher interface {
	Name() string
	Search([]int, int) (int, error)
}

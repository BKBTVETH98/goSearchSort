package benchmark

import (
	"time"

	"sortsearch/pkg/searching"
	"sortsearch/pkg/sorting"
)

type SortResult struct {
	Duration time.Duration
	Sorted   []int
	Error    error
}

type SearchResult struct {
	Duration time.Duration
	Index    int
	Error    error
}

func MeasureSort(alg sorting.Sorter, data []int) SortResult {
	start := time.Now()
	result := alg.Sort(data)
	return SortResult{
		Duration: time.Since(start),
		Sorted:   result,
	}
}

func MeasureSearch(alg searching.Searcher, data []int, target int) SearchResult {
	start := time.Now()
	index, err := alg.Search(data, target)
	return SearchResult{
		Duration: time.Since(start),
		Index:    index,
		Error:    err,
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sortsearch/pkg/benchmark"
	"sortsearch/pkg/searching"
	"sortsearch/pkg/sorting"
)

const arrSize = 100000

func main() {

	data := make([]int, arrSize)
	dataSort := make([]int, arrSize)
	for i := range data {
		data[i] = rand.Intn(arrSize)
	}

	for i := range dataSort {
		dataSort[i] = i
	}

	fmt.Println("Demo библиотеки sortsearch")
	fmt.Println("Размер массива:", len(data))

	sorters := []sorting.Sorter{
		sorting.BubbleSorter{},
		sorting.QuickSorter{},
		sorting.MergeSorter{},
	}

	for _, alg := range sorters {
		result := benchmark.MeasureSort(alg, data)
		fmt.Printf("%s: время=%d ns\n", alg.Name(), result.Duration.Nanoseconds())
	}

	searchers := []searching.Searcher{
		searching.LinearSearcher{},
		searching.BinarySearcher{},
	}

	fmt.Println("\nПоиск элемента", 9999, "в отсортированном массиве")

	for _, alg := range searchers { // Ищем элемент 9999 в отсортированном массиве
		result := benchmark.MeasureSearch(alg, dataSort, 9999)
		if result.Error != nil {
			fmt.Printf("%s: ошибка=%v\n", alg.Name(), result.Error)
			continue
		}
		fmt.Printf("%s: индекс=%d время=%d ns\n", alg.Name(), result.Index, result.Duration.Nanoseconds())
	}

}

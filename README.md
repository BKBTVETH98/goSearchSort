# SortSearch

Учебная библиотека алгоритмов сортировки и поиска на Go. Она реализует несколько популярных алгоритмов сортировки и поиска, а также позволяет измерять время выполнения каждого алгоритма.

## Структура

- `pkg/sorting` — интерфейс `Sorter` и реализации алгоритмов сортировки:
  - `BubbleSorter`
  - `QuickSorter`
  - `MergeSorter`
- `pkg/searching` — интерфейс `Searcher` и реализации алгоритмов поиска:
  - `LinearSearcher`
  - `BinarySearcher`
- `pkg/benchmark` — функции измерения времени выполнения алгоритмов:
  - `MeasureSort`
  - `MeasureSearch`
- `cmd/demo` — пример использования библиотеки.

## Пример использования

```go
package main

import (
    "fmt"
    "math/rand"
    "sortsearch/pkg/benchmark"
    "sortsearch/pkg/searching"
    "sortsearch/pkg/sorting"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    data := []int{9, 3, 5, 1, 7}

    sorter := sorting.QuickSorter{}
    result := benchmark.MeasureSort(sorter, data)
    fmt.Println("Sorted:", result.Sorted)
    fmt.Println("Sort duration:", result.Duration)

    sorted := result.Sorted
    searcher := searching.BinarySearcher{}
    searchResult := benchmark.MeasureSearch(searcher, sorted, 5)
    fmt.Println("Index:", searchResult.Index)
    fmt.Println("Search duration:", searchResult.Duration)
}
```

## Запуск

```bash
go run ./cmd/demo
```

## Тесты

```bash
go test ./...
```

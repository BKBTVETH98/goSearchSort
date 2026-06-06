package benchmark

import (
	"reflect"
	"testing"

	"sortsearch/pkg/searching"
	"sortsearch/pkg/sorting"
)

func TestMeasureSort(t *testing.T) {
	alg := sorting.BubbleSorter{}
	data := []int{4, 2, 3, 1}
	got := MeasureSort(alg, data)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got.Sorted, want) {
		t.Fatalf("MeasureSort(%s) failed: got %v, want %v", alg.Name(), got.Sorted, want)
	}
}

func TestMeasureSearch(t *testing.T) {
	alg := searching.BinarySearcher{}
	data := []int{1, 2, 3, 4, 5}
	got := MeasureSearch(alg, data, 4)
	if got.Error != nil || got.Index != 3 {
		t.Fatalf("MeasureSearch(%s) failed: got index=%d err=%v", alg.Name(), got.Index, got.Error)
	}
}

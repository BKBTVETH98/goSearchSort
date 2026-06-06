package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSorter(t *testing.T) {
	alg := BubbleSorter{}
	data := []int{5, 1, 4, 2, 8}
	got := alg.Sort(data)
	want := []int{1, 2, 4, 5, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s failed: got %v, want %v", alg.Name(), got, want)
	}
}

func TestQuickSorter(t *testing.T) {
	alg := QuickSorter{}
	data := []int{5, 1, 4, 2, 8}
	got := alg.Sort(data)
	want := []int{1, 2, 4, 5, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s failed: got %v, want %v", alg.Name(), got, want)
	}
}

func TestMergeSorter(t *testing.T) {
	alg := MergeSorter{}
	data := []int{5, 1, 4, 2, 8}
	got := alg.Sort(data)
	want := []int{1, 2, 4, 5, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s failed: got %v, want %v", alg.Name(), got, want)
	}
}

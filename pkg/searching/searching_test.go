package searching

import "testing"

func TestLinearSearcher(t *testing.T) {
	alg := LinearSearcher{}
	data := []int{2, 3, 7, 1, 5}
	index, err := alg.Search(data, 7)
	if err != nil || index != 2 {
		t.Fatalf("%s failed: got index=%d err=%v", alg.Name(), index, err)
	}
}

func TestBinarySearcher(t *testing.T) {
	alg := BinarySearcher{}
	data := []int{1, 3, 5, 7, 9}
	index, err := alg.Search(data, 7)
	if err != nil || index != 3 {
		t.Fatalf("%s failed: got index=%d err=%v", alg.Name(), index, err)
	}
}

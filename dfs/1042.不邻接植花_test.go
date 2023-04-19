package dfs

import (
	"reflect"
	"testing"
)

func TestGardenNoAdj(t *testing.T) {
	n := 4
	paths := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}
	want := []int{1, 2, 1, 2}

	got := gardenNoAdj(n, paths)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("gardenNoAdj(%d, %v) = %v; want %v", n, paths, got, want)
	}
}

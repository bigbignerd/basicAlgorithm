package merge

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	Sort(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr is not sorted.")
	}
	// t.Logf("%v", arr)
}

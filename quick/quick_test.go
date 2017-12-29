package quick

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	Sort(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("sort fail.")
	}
	// t.Logf("%v", arr)
}
func TestQuickSortV3(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, 100)
	var arrCopy []int = make([]int, n)
	copy(arrCopy, arr)
	v1Time := util.ExecuteTime(Sort, arr, n)
	v3Time := util.ExecuteTime(SortV3, arrCopy, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("v1 sort error.")
	}
	if !util.IsSortedAsc(arrCopy, n) {
		t.Error("v3 sort error.")
	}
	t.Logf("v1 exeucute time %.6f s", v1Time)
	t.Logf("v3 exeucute time %.6f s", v3Time)

}

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
func TestMergSortBU(t *testing.T) {
	n := 10
	arr := util.RandArr(n, 1, n)
	SortBU(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("bottom to up sort arr not sorted.")
	}
}
func TestReverseOrderCount(t *testing.T) {
	n := 9
	// arr := util.RandArr(n, 1, n)
	var arr []int = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("%v", arr)
	count := ReverseSort(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("排序失败")
	}
	t.Logf("逆序对数量为：%d", count)
}

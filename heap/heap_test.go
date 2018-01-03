package heap

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

func TestMaxHeapInsert(t *testing.T) {
	var cap int = 10
	var maxheap MaxHeap = MaxHeap{Data: make([]int, cap+1)}
	for i := 0; i < 10; i++ {
		maxheap.Insert(util.RandNumber(1, 60))
	}
	t.Logf("%v", maxheap.Data)
}
func TestMaxHeapExtract(t *testing.T) {
	var cap int = 10
	var maxheap MaxHeap = MaxHeap{Data: make([]int, cap+1)}
	for i := 0; i < 10; i++ {
		maxheap.Insert(util.RandNumber(1, 60))
	}
	for !maxheap.IsEmpty() {
		t.Logf("%d,", maxheap.ExtractMax())
	}
}

//测试v1 v2版本堆排序
func TestHeapSortv1(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	// t.Logf("未排序arr：%v", arr)
	var copyArr []int = make([]int, n)
	copy(copyArr, arr)
	// t.Logf("未排序copy arr：%v", copyArr)

	v1Time := util.ExecuteTime(SortV1, arr, n)
	v2Time := util.ExecuteTime(SortV2, copyArr, n)
	// t.Logf("排序后arr：%v", arr)
	// t.Logf("排序后copy arr：%v", copyArr)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr sort fail.")
	}
	if !util.IsSortedAsc(copyArr, n) {
		t.Error("copy arr sort fail.")
	}
	t.Logf("v1执行时间:%.6f", v1Time)
	t.Logf("v2执行时间:%.6f", v2Time)
}

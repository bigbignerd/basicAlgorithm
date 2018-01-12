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

func TestHeapSortV1(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	SortV1(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr is not sorted.")
	}
}
func TestHeapSortV2(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	SortV2(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr is not sorted.")
	}
}

func TestHeapSortv3(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	SortV3(arr, n)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr is not sorted.")
	}
}

//测试对比v1 v2 v3版本堆排序
func TestHeapSortCompare(t *testing.T) {
	n := 1000000
	arr := util.RandArr(n, 1, n)
	// t.Logf("未排序arr：%v", arr)
	var copyArr []int = make([]int, n)
	copy(copyArr, arr)
	var arr3 []int = make([]int, n)
	copy(arr3, arr)
	// t.Logf("未排序copy arr：%v", copyArr)

	v1Time := util.ExecuteTime(SortV1, arr, n)
	v2Time := util.ExecuteTime(SortV2, copyArr, n)
	v3Time := util.ExecuteTime(SortV3, arr3, n)
	// t.Logf("排序后arr：%v", arr)
	// t.Logf("排序后copy arr：%v", copyArr)
	if !util.IsSortedAsc(arr, n) {
		t.Error("arr sort fail.")
	}
	if !util.IsSortedAsc(copyArr, n) {
		t.Error("copy arr sort fail.")
	}
	if !util.IsSortedAsc(arr3, n) {
		t.Error("arr3 sort fail.")
	}
	t.Logf("v1执行时间:%.6f", v1Time)
	t.Logf("v2执行时间:%.6f", v2Time)
	t.Logf("v3执行时间:%.6f", v3Time)
}

type FloatHeap float64

func (f FloatHeap) Less(i, j FloatHeap) bool {
	return i < j
}
func TestMinheap(t *testing.T) {
	data := []FloatHeap{12.1, 9.5, 8.3, 7.6, 6.1, 3.2}
	miniheap := NewMinHeap(6)
	for _, v := range data {
		miniheap.Insert(v)
	}
	t.Log(miniheap.ExtractMin())
	t.Log(miniheap.ExtractMin())
	t.Log(miniheap.ExtractMin())
	t.Log(miniheap.ExtractMin())
	t.Log(miniheap.ExtractMin())
}

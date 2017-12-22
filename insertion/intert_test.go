package insertion

import (
	"github.com/bigbignerd/basicAlgorithm/selection"
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

func TestInsert(t *testing.T) {
	num := 100
	arr := util.RandArr(num, 1, num)
	Sort(arr, num)
	if !util.IsSortedAsc(arr, num) {
		t.Errorf("%v is not sorted", arr)
	}
	// t.Logf("%v", arr)
}

func TestInsertDesc(t *testing.T) {
	num := 20
	arr := util.RandArr(num, 1, num)
	SortDesc(arr, num)
	if !util.IsSortedDesc(arr, num) {
		t.Errorf("%v is not sorted as desc.", arr)
	}
	// t.Logf("%v", arr)
}

//测试选择排序与插入排序执行时间
func TestCompareInsertAndSelect(t *testing.T) {
	num := 1000
	arr := util.RandArr(num, 1, num)
	var copyArr []int = make([]int, len(arr))
	copy(copyArr, arr)
	insertSort := util.ExecuteTime(Sort, copyArr, num)
	selectSort := util.ExecuteTime(selection.Sortv2, arr, num)

	t.Logf("select sort execute time: %.5f s", selectSort)
	t.Logf("insert sort execute time: %.5f s", insertSort)
}

//测试改进后的插入排序v2版
func TestNewInsert(t *testing.T) {
	num := 100000
	arr := util.RandArr(num, 1, num)
	var copyArr []int = make([]int, len(arr))
	copy(copyArr, arr)

	insertV1 := util.ExecuteTime(Sort, arr, num)
	insertV2 := util.ExecuteTime(SortV2, copyArr, num)
	// arrt := util.RandArr(10, 1, 10)
	// if !util.IsSortedAsc(arrt, 10) {
	// 	t.Error("arrt is not sorted")
	// }
	if !util.IsSortedAsc(arr, num) {
		t.Error("arr is not sorted.")
	}
	if !util.IsSortedAsc(copyArr, num) {
		t.Error("copyArr is not sorted.")
	}
	t.Logf("insertv1 sort execute time: %.5f s", insertV1)
	t.Logf("insertv2 sort execute time: %.5f s", insertV2)
}

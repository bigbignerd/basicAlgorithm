package selection

import (
	// "fmt"
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

//v1 使用arry
func TestSelectSortV1(t *testing.T) {
	var arr [10]int = [10]int{10, 9, 8, 6, 7, 5, 4, 3, 2, 1}
	Sortv1(&arr, 10)
	if !util.IsSortedAsc(arr[:], 10) {
		t.Errorf("%v is not sorted", arr)
	}
	// fmt.Printf("%v", arr)
}

//v2 使用生成随机slice
func TestSelectSortV2(t *testing.T) {
	num := 10000
	arr := util.RandArr(num, 1, num)
	Sortv2(arr, num)
	if !util.IsSortedAsc(arr, num) {
		t.Error("NOT sorted")
	}
	// fmt.Printf("%v", arr)
}

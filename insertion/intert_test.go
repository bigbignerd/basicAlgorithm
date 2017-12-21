package insertion

import (
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
	t.Logf("%v", arr)
}

func TestInsertDesc(t *testing.T) {
	num := 20
	arr := util.RandArr(num, 1, num)
	SortDesc(arr, num)
	if !util.IsSortedDesc(arr, num) {
		t.Errorf("%v is not sorted as desc.", arr)
	}
	t.Logf("%v", arr)
}

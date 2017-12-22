package bubble

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
)

func TestSort(t *testing.T) {
	num := 100
	arr := util.RandArr(num, 1, num)
	Sort(arr, num)
	if !util.IsSortedAsc(arr, num) {
		t.Error("arr is not sorted.")
	}
	t.Logf("%v", arr)
}

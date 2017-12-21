package insertion

import "github.com/bigbignerd/basicAlgorithm/util"

// 插入排序的思想类似抓扑克牌：每拿到一个新的元素依次跟手中的牌做对比，直到这个元素比前面的元素大停止交换
func Sort(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				util.Swap(&arr[j], &arr[j-1])
			}
		}
	}
}

// 从大到小插入排序
func SortDesc(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] > arr[j-1] {
				util.Swap(&arr[j], &arr[j-1])
			}
		}
	}
}

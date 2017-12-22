package bubble

import "github.com/bigbignerd/basicAlgorithm/util"

//冒泡排序
func Sort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				util.Swap(&arr[j], &arr[j+1])
			}
		}
	}
}

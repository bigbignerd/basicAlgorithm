package insertion

import "github.com/bigbignerd/basicAlgorithm/util"

// 插入排序的思想类似抓扑克牌：每拿到一个新的元素依次跟手中的牌做对比，直到这个元素比前面的元素大停止交换
func Sort(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			util.Swap(&arr[j], &arr[j-1])
		}
	}
}

//插入排序优化思路：
//之前的每一次比较，如果发生交换操作，会产生三次的赋值操作
//所以改为保存当前带插入元素的值，然后依次向前比较，进行位置移动的操作，直接赋值

func SortV2(arr []int, n int) {
	for i := 1; i < n; i++ {
		curEle := arr[i]
		var j int
		for j = i; j > 0 && arr[j-1] > curEle; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curEle
	}
}

// 从大到小插入排序
func SortDesc(arr []int, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] > arr[j-1] {
				util.Swap(&arr[j], &arr[j-1])
			} else {
				break
			}
		}
	}
}

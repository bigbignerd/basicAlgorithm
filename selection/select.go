package selection

import "github.com/bigbignerd/basicAlgorithm/util"

//选择排序的思想是：
//1.从前到后遍历所有未排序的元素  	0->n
//2.找到所有未排序的元素中最小的那个 i+1 -> n
//3.最小的元素与当前i位置（未排序的第一个元素）的元素交换位置(前i个元素是排好序的)

//v1 使用数组
func Sortv1(arr *[10]int, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		util.Swap(&arr[i], &arr[minIndex])
	}
}

//v2 改为使用slice
func Sortv2(arr []int, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		//获取当前最小值 后面的未排序的元素中 最小的值的索引
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		util.Swap(&arr[i], &arr[minIndex])
	}
}

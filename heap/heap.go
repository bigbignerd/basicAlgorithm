package heap

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	// "log"
)

func SortV1(arr []int, n int) {
	var maxheap MaxHeap = MaxHeap{Data: make([]int, n+1), count: 0}
	for i := 0; i < n; i++ {
		maxheap.Insert(arr[i])
	}
	for j := n - 1; j >= 0; j-- {
		arr[j] = maxheap.ExtractMax()
	}
}

//v2 版本优化：将一个二叉树变成一个二叉堆的过程，对非叶子节点开始执行shiftDown操作
func SortV2(arr []int, n int) {
	maxheap := MaxHeap{Data: make([]int, n+1)}
	copy(maxheap.Data[1:n+1], arr)
	maxheap.SetCount(n)
	//找到第一个非叶子节点
	for i := maxheap.count / 2; i >= 1; i-- {
		maxheap.shiftDown(i)
	}

	for j := n - 1; j >= 0; j-- {
		arr[j] = maxheap.ExtractMax()
	}
}

//v3 原地堆排序：对数组进行heapify 每次产生一个最大堆 然后将首元素交换到数组末尾 再对剩余的元素在进行heapify【shiftDown的过程】
func SortV3(arr []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDownArr(arr, n, i)
	}
	for j := n - 1; j > 0; j-- {
		util.Swap(&arr[0], &arr[j])
		shiftDownArr(arr, j, 0)
	}
}

//数组的shiftdown索引从零开始：
//左节点：2*i+1 右节点：2*i+2 最后一个非叶子节点:(count-1)/2 父节点：(i-1)/2
func shiftDownArr(arr []int, n int, k int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j++
		}
		if arr[j] <= arr[k] {
			break
		}
		util.Swap(&arr[j], &arr[k])
		k = j
	}
}

package heap

import (
// "github.com/bigbignerd/basicAlgorithm/util"
// "log"
)

func SortV1(arr []int, n int) {
	var maxheap MaxHeap = MaxHeap{Data: make([]int, n+1)}
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

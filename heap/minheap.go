package heap

import (
	"log"
)

//存储任意元素的最小堆实现
type MinHeap struct {
	count int
	data  []interface{}
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		data: make([]interface{}, size+1),
	}
}

//向最小堆中插入元素
func (mh *MinHeap) Insert(item interface{}) {
	//堆中元素从1开始
	mh.count++
	mh.data[mh.count] = item
	mh.shiftup(mh.count)
}
func (mh *MinHeap) shiftup(k int) {
	if k < 1 || k > len(mh.data) {
		log.Fatal("shift up i 越界")
	}
	//跟父节点进行比较
	for k/2 >= 1 && mh.less(k, k/2) {
		//交换k/2 与 k
		mh.swap(k, k/2)
		k /= 2
	}
}

//取出最小元素
func (mh *MinHeap) ExtractMin() interface{} {
	if mh.count <= 0 {
		log.Fatal("堆中元素为空")
		return nil
	}
	//交换堆顶元素与最后一个元素
	min := mh.data[1]
	mh.swap(1, mh.count)
	mh.count--
	mh.shiftdown(1)
	//删除最后一个元素
	mh.data = mh.data[0 : mh.count+1]
	return min
}
func (mh *MinHeap) shiftdown(k int) {
	for 2*k <= mh.count {
		j := 2 * k
		if j+1 <= mh.count && mh.less(j+1, j) {
			j += 1
		}
		//当前值k小于 左右两个孩子中最小的那个j
		if mh.less(k, j) {
			break
		}
		mh.swap(k, j)
		k = j
	}
}

//判断data[a] 是否小于 data[b]
func (mh *MinHeap) less(a, b int) bool {
	lessthan := false
	switch mh.data[a].(type) {
	case float64:
		lessthan = mh.data[a].(float64) < mh.data[b].(float64)
	}
	return lessthan
}
func (mh *MinHeap) swap(k1, k2 int) {
	mh.data[k1], mh.data[k2] = mh.data[k2], mh.data[k1]
}

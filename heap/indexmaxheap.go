package heap

import (
	"errors"
	"github.com/bigbignerd/basicAlgorithm/util"
)

type IndexMaxHeap struct {
	count   int
	Cap     int
	Data    []int
	Indexes []int
	Reverse []int //Indexes的反向索引
}

//插入新的元素
func (h *IndexMaxHeap) Insert(i, item int) error {
	if h.count+1 > h.cap {
		return errors.New("数量超出容量")
	}
	if i+1 < 1 || i+1 > h.cap {
		return errors.New("index 越界")
	}
	//针对用户使用数组是从0开始
	i += 1
	//保存索引
	h.Indexes[h.count+1] = i
	//保存reverse
	h.Reverse[i] = h.count + 1
	//保存数据
	h.Data[i] = item
	h.count++
	//对新添加的最后一个元素 开始执行shift up 操作
	h.shiftUp(h.count)
}

func (h *IndexMaxHeap) shiftUp(k int) {
	//比较数据 堆化索引
	for k > 1 && h.Data[h.Indexes[k/2]] < h.Data[h.Indexes[k]] {
		util.Swap(&h.Indexes[k], &h.Indexes[k/2])
		h.Reverse[h.Indexes[k]] = k
		h.Reverse[h.Indexes[k/2]] = k / 2
		k /= 2
	}
	return nil
}

//获取最大元素
func (h *IndexMaxHeap) ExtractMax() (int, error) {
	if h.count <= 0 {
		return 0, errors.New("堆中元素为空")
	}
	ret := h.Data[h.Indexes[1]]
	util.Swap(&h.Data[h.Indexes[1]], &h.Data[h.Indexes[h.count]])
	h.Reverse[h.Indexes[1]] = 1
	h.Reverse[h.Indexes[h.count]] = 0
	h.count--
	h.shiftDown(1)
	return ret, nil
}
func (h *IndexMaxHeap) shiftDown(k int) {
	for 2*k <= h.count {
		j := 2 * k
		if j+1 <= h.count && h.Data[h.Indexes[j]] < h.Data[h.Indexes[j+1]] {
			j += 1
		}
		if h.Data[h.Indexes[k]] >= h.Data[h.Indexes[j]] {
			break
		}
		util.Swap(&h.Indexes[j], &h.Indexes[k])
		h.Reverse[h.Indexes[j]] = j
		h.Reverse[h.Indexes[k]] = k
		k = j
	}
}

//获取最大值索引
func (h *IndexMaxHeap) ExtractMaxIndex() (int, error) {
	if h.count <= 0 {
		return 0, errors.New("堆为空")
	}
	//减1是因为用户使用的数组是从0开始
	index := h.Indexes[1] - 1
	util.Swap(&h.Data[h.Indexes[1]], &h.Data[h.Indexes[h.count]])
	h.Reverse[h.Indexes[1]] = 1
	h.Reverse[h.Indexes[h.count]] = 0
	h.count--
	h.shiftDown(1)
	return index, nil
}

//获取堆指定索引元素 这里因为只是对索引进行了堆化，数组本身没有发生改变
func (h *IndexMaxHeap) GetItem(i int) (int, error) {
	if !h.contain(i) {
		return 0, errors.New("堆中不存在此索引")
	}
	return h.Data[i+1], nil
}

func (h *IndexMaxHeap) Change(i int, item int) error {
	if !h.contain(i) {
		return errors.New("out of range.")
	}
	i += 1
	h.Data[i] = item
	// for j := 1; j <= h.count; j++ {
	// 	if h.Indexes[j] == i {
	// 		h.shiftDown(j)
	// 		h.shiftUp(j)
	// 	}
	// }

	//优化 通过维护的reverses数组直接找到j
	j = h.Reverse[i]
	h.shiftDown(j)
	h.shiftUp(j)
	return nil
}

func (h *IndexMaxHeap) contain(i) bool {
	if i+1 < 1 || i+1 > h.Cap P{
		return false
	}
	return h.Reverse[i+1] != 0
}

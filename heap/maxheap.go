package heap

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	// "log"
)

//最大堆定义：每个节点有两个子节点，每个节点的元素值都要小于其父节点的元素值 的一棵满二叉树
type MaxHeap struct {
	count int   //堆中元素数量
	Data  []int //堆中数据
}

func (m *MaxHeap) Size() int {
	return m.count
}
func (m *MaxHeap) SetCount(count int) {
	m.count = count
}
func (m *MaxHeap) IsEmpty() bool {
	return m.count == 0
}

//堆中插入元素
func (m *MaxHeap) Insert(item int) {
	if m.count > len(m.Data)-1 {
		// panic("insert fail.")
		return 0
	}
	m.count++
	m.Data[m.count] = item
	m.shiftUp(m.count)

}

func (m *MaxHeap) shiftUp(k int) {
	//不断的跟当前的父级元素进行比较,向上移动到合适的位置
	for k > 1 && m.Data[k/2] < m.Data[k] {
		util.Swap(&m.Data[k/2], &m.Data[k])
		k /= 2
	}
}

//取得堆中最大元素
func (m *MaxHeap) ExtractMax() int {
	if m.count < 1 {
		return -1
	}
	ret := m.Data[1]
	util.Swap(&m.Data[1], &m.Data[m.count])
	m.count--
	//shift down 操作
	m.shiftDown(1)
	return ret
}
func (m *MaxHeap) shiftDown(k int) {

	for 2*k <= m.count {
		j := 2 * k
		//比较左右节点的大小
		//之前的bug:这里的if判断写成了for 造成排序结果出错
		if j+1 <= m.count && m.Data[j] < m.Data[j+1] {
			j += 1
		}
		//和左右子树中较大的一个比较，位置不需要移动的情况
		if m.Data[k] >= m.Data[j] {
			break
		}
		util.Swap(&m.Data[k], &m.Data[j])
		k = j
	}
}

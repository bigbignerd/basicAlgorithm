package heap

import (
	"errors"
	// "log"
	"reflect"
)

//为什么要引入索引堆？
//1.堆中存储的是非int这种类型，比如字符串或其他的对象类型，在heapify过程中的交换操作浪费大量资源
//2.向堆中插入元素后想要再找到这个元素比较困难，因为k,v已经不对应了，需要遍历整个堆元素才能找到
//应用：比如k代表系统进程的id,v代表进程的优先级，如果使用普通的堆，再v入堆后与k则不再对应，查找耗时

type IndexMinHeap struct {
	data    []interface{}
	cap     int
	count   int
	Indexes []int //索引
	Reverse []int //逆向索引
}

func NewIndexMinHeap(size int) *IndexMinHeap {
	imh := &IndexMinHeap{
		data:    make([]interface{}, size+1),
		cap:     size,
		count:   0,
		Indexes: make([]int, size+1),
		Reverse: make([]int, size+1),
	}
	return imh
}

func (imh *IndexMinHeap) Insert(i int, item interface{}) error {
	if i < 0 || i >= imh.cap {
		return errors.New("Insert i out of range.")
	}
	imh.count++
	i += 1 //用户索引从0开始，堆从1开始
	imh.data[i] = item
	imh.Indexes[imh.count] = i
	imh.Reverse[i] = imh.count
	imh.shiftUp(imh.count)

	return nil
}

func (imh *IndexMinHeap) shiftUp(i int) {
	for i > 1 && imh.less(i, i/2) {
		imh.swap(i, i/2)
		imh.Reverse[imh.Indexes[i]] = i
		imh.Reverse[imh.Indexes[i/2]] = i / 2
		i /= 2
	}
}

//堆中取出最小值【最小值在堆中删除】
func (imh *IndexMinHeap) ExtractMin() (interface{}, error) {
	if imh.count <= 0 {
		return nil, errors.New("堆中元素为空")
	}
	ret := imh.data[imh.Indexes[1]]

	imh.swap(1, imh.count)
	imh.Reverse[imh.Indexes[1]] = 1
	imh.Reverse[imh.Indexes[imh.count]] = 0
	imh.count--
	imh.shiftDown(1)

	return ret, nil
}

//堆中取出最小值索引 【最小值在堆中删除】
func (imh *IndexMinHeap) ExtractMinIndex() (int, error) {
	if imh.count <= 0 {
		return 0, errors.New("堆中元素为空")
	}
	index := imh.Indexes[1] - 1 //用户索引为0
	imh.swap(1, imh.count)
	imh.Reverse[imh.Indexes[1]] = 1
	imh.Reverse[imh.Indexes[imh.count]] = 0
	imh.count--
	imh.shiftDown(1)

	return index, nil
}
func (imh *IndexMinHeap) shiftDown(i int) {
	for 2*i < imh.count {
		j := 2 * i
		if 2*i+1 <= imh.count && imh.less(2*i+1, 2*i) {
			j += 1
		}
		if imh.less(i, j) {
			break
		}
		imh.swap(j, i)
		imh.Reverse[imh.Indexes[j]] = j
		imh.Reverse[imh.Indexes[i]] = i

		i = j
	}
}

//获取堆顶数据 【不从堆中删除】
func (imh *IndexMinHeap) GetMin() interface{} {
	if imh.count <= 0 {
		return nil
	}
	return imh.data[imh.Indexes[1]]
}

//获取堆顶元素索引 【不从堆中删除】
func (imh *IndexMinHeap) GetMinIndex() int {
	if imh.count <= 0 {
		return -1
	}
	return imh.Indexes[1] - 1
}
func (imh *IndexMinHeap) Contain(i int) bool {
	return imh.Reverse[i+1] != 0
}

//根据索引获取堆中数据
func (imh *IndexMinHeap) GetItem(index int) (interface{}, error) {
	if imh.count <= 0 || !imh.Contain(index) {
		return nil, errors.New("堆中元素为空")
	}
	return imh.data[index+1], nil
}

//将最小堆中索引为i的元素改为新的值
func (imh *IndexMinHeap) Change(i int, newItem interface{}) {
	if !imh.Contain(i) {
		return
	}
	i += 1 //用户 索引0开始
	imh.data[i] = newItem
	imh.shiftUp(imh.Reverse[i])
	imh.shiftDown(imh.Reverse[i])
}

func (imh *IndexMinHeap) IsEmpty() bool {
	return imh.count == 0
}

func (imh *IndexMinHeap) swap(i, j int) {
	imh.Indexes[i], imh.Indexes[j] = imh.Indexes[j], imh.Indexes[i]
}

func (imh *IndexMinHeap) less(i, j int) bool {
	datai := imh.data[imh.Indexes[i]]
	dataj := imh.data[imh.Indexes[j]]

	f := reflect.ValueOf(datai).MethodByName("Less")
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(datai)
	params[1] = reflect.ValueOf(dataj)

	result := f.Call(params)

	return result[0].Interface().(bool)
}

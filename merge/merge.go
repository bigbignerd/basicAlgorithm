package merge

import (
	"math"
)

//归并排序应用：求数组元素中逆序对的个数
//归并排序思路：就是拆分+比较+合并 归成有序序列 时间复杂度nlogn
func Sort(arr []int, n int) {
	// mergesort(arr, 0, n-1)
	mergesortv2(arr, 0, n-1)
}

//递归拆分
func mergesort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort(arr, l, mid)
	mergesort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

//递归拆分优化v2
func mergesortv2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort(arr, l, mid)
	mergesort(arr, mid+1, r)
	//因为左右都已经进行了排序，都是从小到大的排列了，所以只有在左侧最后一个元素的值 > 右侧第一个元素的值得时候才需要进行合格并的操作
	//这个优化 对于近乎有序的序列有很大的速度提升
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

//自底向上的归并排序
func SortBU(arr []int, n int) {
	//size 从1 到 2xsize 不断循环
	for size := 1; size <= n; size += size {
		for i := 0; i+size < n; i += size + size {
			right := math.Min(float64(i+size+size-1), float64(n-1))
			merge(arr, i, i+size-1, int(right))
		}
	}
}

//执行合并
func merge(arr []int, l int, mid int, r int) {
	//当前待合并处理的元素个数[l,r]
	var aux []int = make([]int, r-l+1)
	copy(aux, arr[l:r+1])
	//i,j 分别代表待合并的左侧和右侧元素的起始位置
	var i, j int = l, mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

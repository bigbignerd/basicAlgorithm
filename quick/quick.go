package quick

import (
	"github.com/bigbignerd/basicAlgorithm/util"
)

//快速排序应用：求数组中第n大的元素
//快速排序思想：选定一个元素 比较 使得左侧的元素都小于它 右侧都是大于它的，递归执行
func Sort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}
func SortV3(arr []int, n int) {
	quickSortV3(arr, 0, n-1)
}

//v4版本针对大量重复元素的优化：
//三路快排中心思想是：将数组分为 < v; ==v; > v 三部分，然后对小于v和大于v的部分进行快速排序
func SortV4(arr []int, n int) {
	quickSortV4(arr, 0, n-1)
}
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	// 获取第一个元素所应该在的位置
	p := partition(arr, l, r)
	//优化 v2
	// p := partitionv2(arr, l, r)
	//优化 v3
	// p := partitionv3(arr, l, r)
	//分别对p左侧和右侧的元素进行快速排序
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}
func quickSortV3(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partitionv3(arr, l, r)
	//分别对p左侧和右侧的元素进行快速排序
	quickSortV3(arr, l, p-1)
	quickSortV3(arr, p+1, r)
}
func quickSortV4(arr []int, l int, r int) {
	if l >= r {
		return
	}
	util.Swap(&arr[l], &arr[util.RandNumber(l, r)])
	v := arr[l]
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < v {
			util.Swap(&arr[i], &arr[lt+1])
			lt++
			i++
		} else if arr[i] > v {
			util.Swap(&arr[i], &arr[gt-1])
			gt--
		} else { // == v
			i++
		}
	}
	util.Swap(&arr[l], &arr[lt])
	quickSortV4(arr, l, lt-1)
	quickSortV4(arr, gt, r)
}
func partition(arr []int, l int, r int) int {
	j := l      //j为分界点[l,j] 都为小于arr[l]的元素[j+1,r]为大于arr的元素
	v := arr[l] //第一个元素的值
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			util.Swap(&arr[j+1], &arr[i])
			j++
		}
	}
	util.Swap(&arr[l], &arr[j])
	return j
}

//partition v2 改进，因为快速排序最差的情况是一个完全有序的数组
//那么他的时间复杂度就变成了O(n2)，针对的改进是每次选定的v值得位置由固定的l变成随机
func partitionv2(arr []int, l int, r int) int {
	util.Swap(&arr[l], &arr[util.RandNumber(l, r)])
	j := l
	v := arr[l]
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			util.Swap(&arr[j+1], &arr[i])
			j++
		}
	}
	util.Swap(&arr[j], &arr[l])
	return j

}

//v3版本 针对数组中重复的元素较多的情况，将数组切分为 [l+1,i] [j,r]
//然后比较arr[i] 与 arr[j] 进行交换
func partitionv3(arr []int, l int, r int) int {
	v := arr[l]
	i, j := l+1, r
	for true {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		util.Swap(&arr[i], &arr[j])
		i++
		j--
	}
	util.Swap(&arr[l], &arr[j])
	return j
}

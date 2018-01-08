package binarytree

//二分查找迭代版本 LgN级别
func BinarySearch(arr []int, n int, target int) int {
	//定义查找区间为[l,r]前闭后闭
	l, r := 0, n-1
	//历史上著名的bug l与r都是int的最大值得情况下则l+r越界
	//mid := (l+r)/2
	mid := l + (r-l)/2
	for l <= r {
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	//不存在此target值
	return -1
}

//二分查找递归版本
func BinarySearchV2(arr []int, l, r, target int) int {
	if l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			BinarySearchV2(arr, mid+1, r, target)
		} else {
			BinarySearchV2(arr, l, mid-1, target)
		}
	}
	return -1
}

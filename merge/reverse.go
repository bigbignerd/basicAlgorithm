package merge

var count int = 0

//归并排序算法应用 求逆序对
func ReverseSort(arr []int, n int) int {
	revermergesort(arr, 0, n-1)
	return count
}

//递归拆分
func revermergesort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	revermergesort(arr, l, mid)
	revermergesort(arr, mid+1, r)
	mergeCount(arr, l, mid, r)
}
func mergeCount(arr []int, l int, mid int, r int) {
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
			count += (mid - i + 1)
		}
	}
}

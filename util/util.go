package util

import (
	"log"
	"math/rand"
	"time"
)

type SortFunc func(arr []int, n int)

//获取sort 排序执行的时间
func ExecuteTime(sort SortFunc, arr []int, n int) float64 {
	start := time.Now()
	sort(arr, n)
	delta := time.Now().Sub(start)
	return float64(delta) / 1e9
}

//交换两个元素值
func Swap(i *int, j *int) {
	var temp *int = i
	*i, *j = *j, *temp
}

//生成随机数组
func RandArr(num int, min int, max int) []int {
	if max <= min {
		log.Fatal("max的值小于或等于min的值")
	}
	source := rand.NewSource(time.Now().Unix())
	randWithSource := rand.New(source)

	arr := make([]int, num)
	for k, _ := range arr {
		arr[k] = randWithSource.Int()%(max-min+1) + min
	}
	return arr
}

//是否已经正确的按从小到大排序
func IsSortedAsc(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

//是否已经正确的按从大到小排序
func IsSortedDesc(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

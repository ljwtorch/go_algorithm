package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	//获取分区位置
	mid := len(arr) / 2
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])
	//排序后合并
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	//存放结果
	var result []int
	for {
		//如果任何一个区间遍历结束，则退出
		if i >= m || j >= n {
			break
		}
		//对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//如果左侧区间还没有遍历结束，将剩余数据放到结果里
	if i != m {
		for ; i < m; i++ {
			result = append(result, left[i])
		}
	}
	//如果右侧区间还没有遍历完，将剩余数组放到结果集
	if j != m {
		for ; j < n; j++ {
			result = append(result, right[j])
		}
	}
	return result
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	fmt.Println(mergeSort(nums))
}

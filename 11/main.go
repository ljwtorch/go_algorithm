package main

import "fmt"

// InsertSort 插入排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1 //下标
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
	fmt.Println("排序后的数组为 =", arr)
}

func main() {
	arr := []int{77, 55, 33, 99, 66, 11}
	fmt.Println("排序前的数组为 =", arr)
	InsertSort(arr)
}

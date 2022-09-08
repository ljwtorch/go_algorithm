package main

import "fmt"

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		maxNum := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if maxNum < arr[j] { //从大到小，或者从小到大
				maxNum = arr[j]
				maxIndex = j
			}
		}
		//交换
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
	return arr
}

func main() {
	//从大到小
	arr := []int{77, 55, 33, 99, 66, 11}
	fmt.Println("排序前的数组 =", arr)
	SelectSort(arr)
	fmt.Println("排序后的数组 =", arr)
}

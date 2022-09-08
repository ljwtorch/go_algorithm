package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	temp := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{77, 55, 33, 99, 66, 11}
	fmt.Println("排序前的数组 =", arr)
	BubbleSort(arr)
	fmt.Println("排序后的数组 =", arr)
}

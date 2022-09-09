package main

import "fmt"

/*
1.left 表示数组左边的下标
2.right 表示数组右边的下标
3.array 表示要排序的数组
*/
//快速排序
func QuickSort(left int, right int, array []int) {
	l := left
	r := right
	//pivot中轴
	pivot := array[(left+right)/2]
	temp := 0

	//for 循环的目标是将比 pivot 小的数放到左边 大的放在右边
	for l < r {
		//从pivot的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot的右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// l >= r 表明本次分解任务完成
		if l >= r {
			break
		}
		//进行交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化:找到了一侧有等于pivot的数，然后和另一侧交换，交换后在另一侧的去掉即可，不需要再交换回来
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	//如果 l==r 再移动
	if l == r {
		l++
		r--
	}
	//左递归
	if left < r {
		QuickSort(left, r, array)
	}
	//右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := []int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	fmt.Println(arr)
	QuickSort(0, 8, arr)
	fmt.Println(arr)
}

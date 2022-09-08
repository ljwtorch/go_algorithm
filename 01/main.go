package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑色棋子
	chessMap[2][3] = 2 //蓝色棋子

	//2.输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	/*
	   思路：
	   1.遍历chessmap，发现有元素不为0，创建node结构体
	   2.然后放入对应的切片中即可
	   3.标准的稀疏数组应该还有一个记录元素的二维数组的规模（行列，默认值）
	*/
	var sparseArr []ValNode

	//这里表示原始数组的规模
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode结点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				//存入切片中
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("当前稀疏数组是：")
	for k, v := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", k, v.row, v.col, v.val)
	}

	//先创建一个原始数组
	var chessMap2 [11][11]int
	//遍历sparseArr
	for k, v := range sparseArr {
		if k != 0 {
			chessMap2[v.row][v.col] = v.val
		}
	}
	//查看chessmap2
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}

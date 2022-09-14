package main

import "fmt"

// SetWay 找路函数 i j 表示对地图的哪一个点进行测试
func SetWay(maps *[8][7]int, i int, j int) bool {
	//分析出什么情况下，就找到出路
	if maps[6][5] == 2 {
		return true
	} else {
		if maps[i][j] == 0 {
			//假设这个点是可以通，但是需要探测下右上左
			maps[i][j] = 2
			if SetWay(maps, i+1, j) { //下
				return true
			} else if SetWay(maps, i, j+1) { //右
				return true
			} else if SetWay(maps, i-1, j) { //上
				return true
			} else if SetWay(maps, i, j-1) { //左
				return true
			} else {
				maps[i][j] = 3 //死路
				return false
			}
		} else { //说明这个点不能探测
			return false
		}
	}
}

func main() {
	//1.如果元素值为1，就是墙
	//2.如果元素的值为0，就是没有走过的点
	//3.如果元素的值为2，就是一个通路
	//4.如果元素的值为3，就是走过的点，但是不通
	//创建二维数组模仿迷宫
	var maps [8][7]int

	//给地图上边和下边设置为1
	for i := 0; i < 7; i++ {
		maps[0][i] = 1
		maps[7][i] = 1
	}

	//给地图左边右边设置为1
	for i := 0; i < 8; i++ {
		maps[i][0] = 1
		maps[i][6] = 1
	}

	maps[3][1] = 1
	maps[3][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(maps[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//开始测试
	SetWay(&maps, 1, 1)

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(maps[i][j], " ")
		}
		fmt.Println()
	}
}

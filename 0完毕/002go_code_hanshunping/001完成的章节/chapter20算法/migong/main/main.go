package main

import(
	"fmt"
)

func initMyMap(myMap *[8][7]int)  {
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[1][2] = 1
	myMap[2][2] = 1
	showMyMap(myMap)
}


func showMyMap(myMap *[8][7]int)  {
	fmt.Println("地图情况是：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j],"  ")
		}
		fmt.Println()
	}
}

/*
编写一个函数，完成老鼠战略
myMap *[8][7]int:地图，保证是同一个地图
i,j表示对地图哪一个点进行测试
*/
func SetWay(myMap *[8][7]int, i int, j int) bool {
	// 分析出什么情况下，就找到出路
	// myMap[6][5] == 2
	if myMap[6][5] == 2 {
		return true
	}else{
		// 说明继续找
		if myMap[i][j] == 0 {//可以探测
			// 假设这个点是通的,但是需要探测，上下左右
			// 换一种策略 下右上左
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) {//下
				return true
			}else if SetWay(myMap, i, j+1) {//右
				return true
			}else if SetWay(myMap, i-1, j){//上
				return true
			}else if SetWay(myMap, i, j-1){//左
				return true
			}else{
				myMap[i][j] = 3
				return false
			}
		}else{//b不能走，1墙，3走过不通
			return false

		}
	}
}
func main()  {
	/*
	先创建一个二维数组，模拟迷宫
	规则：
	1、如果元素的值为1，就是墙
	2、如果元素的值为0，就是没有走过的点
	3、如果元素的值为2，是一条通路
	4、如果元素的值为3，是走过的点，但是走不通
	*/
	var myMap [8][7]int
	initMyMap(&myMap)
	SetWay(&myMap,1,1)
	showMyMap(&myMap)
}

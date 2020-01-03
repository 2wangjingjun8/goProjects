package main

import(
	"fmt"
)

type ValNode struct{
	Row int
	Col int
	Val int
}

func main()  {
	var sparseArr [11][11]int
	sparseArr[1][2] = 1
	sparseArr[2][3] = 2
	for _, v := range sparseArr {
		for _, j := range v {
			fmt.Printf("%d\t",j)
		}
		fmt.Println()
	}

	// 原始数组转稀疏数组
	var valNode []ValNode
	valNode0 := ValNode{
		Row:11,
		Col:11,
		Val:0,
	}
	valNode = append(valNode,valNode0)
	for i, v := range sparseArr {
		for j, v2 := range v {
			if v2 != 0 {
				valNode0 := ValNode{
					Row:i,
					Col:j,
					Val:v2,
				}
				valNode = append(valNode,valNode0)
				
			}
		}
	}

	// 输出稀疏数组
	for _, v := range valNode {
		fmt.Printf("%d %d %d\n",v.Row,v.Col,v.Val)
	}

	// 稀疏数组转原始数组
	// 先创建一个空数组
	var sparseArr2 [11][11]int
	for i, v := range valNode {
		if i == 0 {
			continue
		}
		sparseArr2[v.Row][v.Col] = v.Val
	}

	// 输出原始数组
	for _, v := range sparseArr2 {
		for _, v2 := range v {
			fmt.Printf("%d \t",v2)
		}
		fmt.Println()
	}



}
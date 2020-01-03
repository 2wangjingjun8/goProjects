package main

import (
	"fmt"
)

func main()  {
	var scores [3][4]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%v个班第%v个学生的成绩:",i+1,j+1)
			fmt.Scanln(&scores[i][j])
			
		}
	}
	fmt.Println(scores)
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		fmt.Printf("第%v个班总成绩sum = %v, 平均成绩是%v \n",i+1,sum,sum/float64(len(scores[i])))
		totalSum += sum
	}
	fmt.Printf("各个班总成绩totalSum = %v, 平均成绩是%v \n",totalSum,totalSum/12)

}
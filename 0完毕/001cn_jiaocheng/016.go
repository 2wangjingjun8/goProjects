package main

import "fmt"

func main()  {
	// var salary float32 = balance[9]
	var n = [5][2]int{{0,0},{1,2},{2,4},{4,6},{6,8}}
	var i,j int
	for i = 0; i < 5; i++{
		for j = 0; j < 2;j++{
			fmt.Printf("n[%d][%d] = %d\n",i,j,n[i][j])
		}
	}

	


}
package main

import (
	"fmt"
	"bufio"
	"os"
)
// Count struct
type Count struct{
	EnCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main()  {
	count := Count{}
	file,err := os.Open("./aa.txt")
	if err != nil{
		fmt.Println("open file err = ",err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		for _,v := range str{
			fmt.Println(string(v))
			switch {
				case v >= 'a' && v <= 'z':
					count.EnCount++
				case v >= 'A' && v <= 'Z':
					count.EnCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				default:
					count.OtherCount++
			}
		}
		if err != nil{
			fmt.Println("read file err = ",err)
			break
		}
	}

	fmt.Printf("EnCount = %v, NumCount = %v, SpaceCount = %v, OtherCount = %v",count.EnCount,
	count.NumCount,count.SpaceCount,count.OtherCount)
}


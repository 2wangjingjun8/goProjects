package main
import (
	"fmt"
)

type AreaUtils struct{

}

func (su AreaUtils) print(m int, n int){
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main()  {
	su := AreaUtils{}
	su.print(5,15)
	// fmt.Println()
}
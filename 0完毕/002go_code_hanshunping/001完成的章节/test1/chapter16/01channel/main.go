package main
import (
	"fmt"
)

type Cat struct{
	Name string
	Age int
}

func main()  {
	var chan1 chan int
	chan1 = make(chan int, 3)
	chan1 <- 20
	chan1 <- 10
	chan1 <- 4
	num1 := <-chan1
	fmt.Println(num1)

	chan2 := make(chan interface{},10)
	chan2 <- "哈哈"
	chan2 <- 4
	cat := Cat{"小白猫",2}
	chan2 <- cat
	<-chan2
	<-chan2
	aa := <-chan2
	aa1 := aa.(Cat)
	fmt.Println(aa1.Name)


}
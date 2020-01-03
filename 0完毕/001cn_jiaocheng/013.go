package main
import "fmt"
var aa int = 10
func main(){
	var a int = 100

	if (a<20) {
		fmt.Println("a小于20\n")
	}else{
		fmt.Println("a不小于20\n")
	}
	fmt.Printf("a的值为：%d",a)
	fmt.Printf("aa的值为：%d",aa)
	
}
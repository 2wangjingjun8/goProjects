package main
import "fmt"

func main(){
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("%d + %d = %d \n",a,b,c)
	c = a - b
	fmt.Printf("%d - %d = %d \n",a,b,c)
	c = a * b
	fmt.Printf("%d * %d = %d \n",a,b,c)
	c = a / b
	fmt.Printf("%d / %d = %d \n",a,b,c)
	c = a % b
	fmt.Printf("%d 取模 %d = %d \n",a,b,c)
	a = 20
	a--
	fmt.Printf("a = %d \n",a)
	a = 20
	a++
	fmt.Printf("a = %d \n",a)
	
	
}
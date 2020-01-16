package main

import(
	"fmt"
	"os"
)
func main() {
	file,err := os.Open("./aa.go")
	if err != nil {
		fmt.Println("err = ",err)
	}
	fmt.Println("file = ",file)
	err = file.Close()
	if err != nil {
		fmt.Println("err = ",err)
	}
}
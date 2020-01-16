package main

import (
	"fmt"
	"flag"
)

func main()  {
	var username string
	var password string
	var host string
	var port int
	flag.StringVar(&username,"u","","默认为空")
	flag.StringVar(&password,"p","123456","默认为123456")
	flag.StringVar(&host,"h","localhost","默认为localhost")
	flag.IntVar(&port,"port",3306,"默认为3306")
	flag.Parse()
	fmt.Printf("username = %v password = %v host = %v port = %v",username, password, host, port)
}
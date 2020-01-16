package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn err = " , err )
		return
	}
	defer c.Close()
	fmt.Println(c)

	_,err = c.Do("dbsize")
	if err != nil {
		fmt.Println("Do err = " , err )
		return
	}
	fmt.Println("Do OK")
	
	_,err = c.Do("set","name","xiaoxiao")
	if err != nil {
		fmt.Println("Do err = " , err)
		return
	}
	fmt.Println("set OK")
	
	res,err := redis.String(c.Do("get","name"))
	if err != nil {
		fmt.Println("Do err = " , err)
		return
	}
	fmt.Println("get OK, name = ",res)

}
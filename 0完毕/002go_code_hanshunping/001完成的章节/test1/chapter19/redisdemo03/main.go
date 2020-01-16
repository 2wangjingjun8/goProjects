package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("Dial err = ",err )
		return
	}
	_,err = c.Do("hmset","user1","name","xiaoxiao","age",20)
	if err != nil {
		fmt.Println("hmset err = ",err )
		return
	}
	res,err := redis.Strings(c.Do("hmget","user1","name","age"))
	if err != nil {
		fmt.Println("hmget err = ",err )
		return
	}
	fmt.Println("hmget res = ",res )
	for i,v := range res {
		fmt.Printf("res[%d] = %v\n",i,v)
	}

}
package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("Dail err = ",err )
		return 
	}
	// fmt.Println("c = ", c)

	_,err = c.Do("hset","stu","name","xiaoxiao")
	if err != nil {
		fmt.Println("hset err = ",err )
		return 
	}
	_,err = c.Do("hset","stu","age",20)
	if err != nil {
		fmt.Println("hset err = ",err )
		return 
	}

	
	name,err := redis.String(c.Do("hget","stu","name"))
	if err != nil {
		fmt.Println("hget err = ",err )
		return 
	}
	fmt.Println("Dail name = ",name )
	
	age,err := redis.Int(c.Do("hget","stu","age"))
	if err != nil {
		fmt.Println("hget err = ",err )
		return 
	}
	fmt.Println("Dail age = ",age )
}
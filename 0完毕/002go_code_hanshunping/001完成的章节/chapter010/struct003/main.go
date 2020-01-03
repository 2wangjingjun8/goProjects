package main
import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	var monster = Monster{"笑笑",20}
	fmt.Println(monster)
	jsonstr,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json格式化错误")
	}
	fmt.Println(string(jsonstr))
}
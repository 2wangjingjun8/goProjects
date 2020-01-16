package main

import(
	"fmt"
	"encoding/json"
)

type Monkey struct {
	Name string `json:"name"`
	Age int
	BirthDay string
	Salary float64
	Skill string
}

func testStruct(){
	var monkey Monkey
	struct_str := "{\"name\":\"孙悟空\",\"age\":500,\"BirthDay\":\"2019-09-05\",\"Salary\":60.00,\"Skill\":\"fly\"}"
	err := json.Unmarshal([]byte(struct_str),&monkey)
	if err != nil {
		fmt.Println("err = ",err)
	}
	fmt.Printf("monkey = %v name=%v\n",monkey,monkey.Name)
}

func testMap(){
	var map1 map[string]interface{}
	map1_str := "{\"name\":\"孙悟空\",\"Age\":500,\"BirthDay\":\"2019-09-05\",\"Salary\":60.00,\"Skill\":[\"fly\",\"swim\"]}"	
	err := json.Unmarshal([]byte(map1_str),&map1)
	if err != nil {
		fmt.Println("err = ",err)
	}
	fmt.Println("map1 = ",map1)
}

func testSlice(){
	var slice []map[string]interface{}
	slice_str := "[{\"age\":18,\"name\":\"xiaowang~~\",\"salary\":6500,\"sex\":\"男\"},"+
	"{\"age\":18,\"name\":\"~~hah\",\"salary\":7000,\"sex\":\"男\"}]"
	err := json.Unmarshal([]byte(slice_str),&slice)
	if err != nil {
		fmt.Println("err = ",err)
	}
	fmt.Print("slice = %V",slice)

}

func main() {
	testStruct()
	testMap()
	testSlice()

}
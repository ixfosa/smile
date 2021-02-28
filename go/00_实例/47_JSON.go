package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println("bloB = ", string(bolB)) // bloB =  true

	intB, _ := json.Marshal(2)
	fmt.Println("intB = ", intB) // intB =  [50]

	fltB, _ := json.Marshal(6.66)
	fmt.Println("fltB = ", string(fltB)) // fltB =  6.66

	slcD := []string{"long", "qing", "zhong"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("slcB = ", string(slcB)) // slcB =  ["long","qing","zhong"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("mapB = ", string(mapB)) //mapB =  {"apple":5,"lettuce":7}


	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}


	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	num := dat["num"].(float64)
	fmt.Println("num = ", num) // num =  6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("str1 = ", str1) // str1 =  a


	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	if err := json.Unmarshal([]byte(str), res); err != nil {
		panic(err)
	}
	fmt.Println("res = ", res) // res =  &{1 [apple peach]}
	fmt.Println(res.Fruits) // [apple peach]

	// 也可以和 os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	encoder.Encode(d) // {"apple":5,"lettuce":7}
}

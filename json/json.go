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
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcB, _ := json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]string) //type casting
	fmt.Println(strs)
	str1 := strs[0]
	fmt.Println(str1)

	resp2JsonStr := `{"page" : 1,"fruits":["apple","peach"]} `
	var resp2d *Response2 = &Response2{}
	if err := json.Unmarshal([]byte(resp2JsonStr), resp2d); err != nil {
		panic(err)
	}
	fmt.Println(*resp2d)

	var encode *json.Encoder = json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	encode.Encode(d)

	resp1JsonStr :=
		`{
	      "Page" : 1,
	      "Fruits" : ["apple","peach"]
		}
		`
	var byte = []byte(resp1JsonStr)
	var resp1 *Response1 = &Response1{}

	if err := json.Unmarshal(byte, resp1); err != nil {
		panic(err)
	}
	fmt.Println(resp1.Fruits[0])
}

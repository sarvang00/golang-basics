package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to session on JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web", "dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeonline.in", "sar123", nil},
	}

	// package courses as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
			"coursename": "MERN Bootcamp",
			"Price": 199,
			"website": "learncodeonline.in",
			"tags": ["full-stack", "js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON failed!")
	}

	// some cases where you just wanna add data to k-v pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)

	fmt.Println(lcoCourse)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v; value type is %T\n", k, v, v)
	}
}

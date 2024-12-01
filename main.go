package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	jsonFile, _ := ioutil.ReadFile("./1.json")
	var data map[string]interface{}
	json.Unmarshal(jsonFile, &data)
	fmt.Println(data["test"])
}

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var slice01 = []string{"hello", "NYH"}
	var slice02 = make([]string, len(slice01))
	data, _ := json.Marshal(slice01)
	json.Unmarshal(data, &slice02)

	fmt.Println("slice02 = ", strings.Join(slice02, " "))
	
}
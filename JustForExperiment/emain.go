//go:generate PrintGenVar

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonString := `{"name": "John", "age": 30, "city": "New York", "isMarried": true, "height": 1.75, "weight": 75.5}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(data)
	}

	for k, v := range data {
		fmt.Printf("key: %s, value: %v, type: %v \n", k, v, reflect.TypeOf(v))
	}
}

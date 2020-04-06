package main

import (
	"encoding/json"
	"fmt"
)

var raw = `[
    {
        "type": "group"
    },
    {
        "type": "row",
        "first_name": "Stanislav"
    },
    {
        "type": "row",
        "first_name": "Chris"
    },
    {
        "type": "group" 
    },
    {
        "type": "row",
        "first_name": "Tomas"
    }
]`
var engineers []Engineer

func main() {
	err := json.Unmarshal([]byte(raw), &engineers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unfiltered: ", engineers)
	filtered := engineers[:0]
	for _, v := range engineers {
		switch v.Type {
		case "row":
			filtered = append(filtered, v)
		}
	}
	fmt.Println("filtered: ", filtered)

}

type Engineer struct {
	Type      string `json:"type"`
	FirstName string `json:"first_name"`
}

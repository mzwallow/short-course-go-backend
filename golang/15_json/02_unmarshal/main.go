package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name           string   `json:"name,omitempty"`
	Age            int      `json:"age"`
	FavoriteColors []string `json:"favorite_colors"`
}

func main() {
	jsonBlob := []byte(`
	{
		"age": 20,
		"favorite_colors": [
			"Red",
			"Green",
			"Blue"
		]
	}
	`)

	var p Person
	if err := json.Unmarshal(jsonBlob, &p); err != nil {
		log.Fatalln("json.Unmarshall failed")
	}
	fmt.Println(p)
}

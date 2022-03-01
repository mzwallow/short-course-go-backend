package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name           string   `json:"name"`
	Age            int      `json:"age"`
	FavoriteColors []string `json:"favorite_colors"`
}

func main() {
	// p := Person{"John", 20, []string{"Red", "Green", "Blue"}}
	p := Person{
		Name:           "John",
		Age:            20,
		FavoriteColors: []string{"Red", "Green", "Blue"},
	}

	d, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("json.Marshal failed with '%s'\n", err)
	}
	fmt.Println("Person in compact JSON:", string(d))
}

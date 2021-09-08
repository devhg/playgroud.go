package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City       string      `json:"city"`
	Street     string      `json:"street"`
	ZipCode    string      `json:"zip_code,omitempty"`
	Coordinate *coordinate `json:"coordinate,omitempty"`
}

type coordinate struct {
	Lat *float64 `json:"latitude,omitempty"`
	Lng *float64 `json:"longitude,omitempty"`
}

func main() {
	data := `{
        "latitude": 1.0,
        "longitude": 0.0
    }`

	c := &coordinate{}
	_ = json.Unmarshal([]byte(data), c)
	fmt.Printf("%#v\n", c)

	addressBytes, _ := json.MarshalIndent(c, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}

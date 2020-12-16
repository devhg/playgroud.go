package omitempty

import (
	"encoding/json"
	"fmt"
	"testing"
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

func TestMarshal(t *testing.T) {
	data := `{
        "latitude": 1.0,
        "longitude": 0.0
    }`
	c := &coordinate{}
	json.Unmarshal([]byte(data), c)
	fmt.Printf("%#v\n", c)

	addressBytes, _ := json.MarshalIndent(c, "", "    ")
	fmt.Printf("%s\n", string(addressBytes))
}

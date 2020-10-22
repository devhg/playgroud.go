package m_gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y int
	Name string
}

func Test() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)

	//var network bytes.Buffer
	//encoder := gob.NewEncoder(&network)
	//decoder := gob.NewDecoder(&network)
	//
	//err := encoder.Encode(P{3, 4, 5, "pp"})
	//if err != nil {
	//	log.Fatal("encode error ", err)
	//}
	//
	//var q Q
	//err = decoder.Decode(&q)
	//if err != nil {
	//	log.Fatal("decode error ", err)
	//}
	//
	//fmt.Printf("%q: {%d, %d}\n", q.Name, q.X, q.Y)
}

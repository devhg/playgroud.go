package main

import (
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string
}

type Foo struct {
	p Person
}

func main() {
	s0 := "something"
	s1 := "something"
	s2 := "something"[3:]
	spew.Dump(s0, s1, s2)

	spew.Dump(&s0)
	spew.Dump(&s1)
	spew.Dump((*reflect.StringHeader)(unsafe.Pointer(&s0)))
	spew.Dump((*reflect.StringHeader)(unsafe.Pointer(&s1)))

	f := Foo{p: Person{Name: "123"}}
	mp := map[int]int{1: 23}
	b := []byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}
	spew.Dump(f, mp, b, s0)
}

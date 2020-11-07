package gredis

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	err := Setup("")
	if err != nil {
		panic(err)
	}
}

func TestDelete(t *testing.T) {
	b, err := Delete("age")
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

func TestExists(t *testing.T) {
	exists := Exists("name")
	fmt.Println(exists)
}

func TestGet(t *testing.T) {
	get, err := Get("name")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(get))
}

func TestLikeDeletes(t *testing.T) {
	err := LikeDeletes("a")
	if err != nil {
		panic(err)
	}
}

func TestSet(t *testing.T) {
	set, err := Set("a", "1234", 200)
	if err != nil {
		panic(err)
	}
	fmt.Println(set)
}

type TestStruct struct {
	Id    int    `redis:"id" json:"id"`
	Name  string `redis:"name" json:"name"`
	Sex   string `redis:"sex" json:"sex"`
	Desc  string `redis:"desc" json:"desc"`
	Desc1 string `redis:"desc1" json:"desc1"`
	Desc2 string `redis:"desc2" json:"desc2"`
	Desc3 string `redis:"desc3" json:"desc3"`
	Desc4 string `redis:"desc4" json:"desc4"`
	Desc5 string `redis:"desc5" json:"desc5"`
	Desc6 string `redis:"desc6" json:"desc6"`
	Desc7 string `redis:"desc7" json:"desc7"`
	Desc8 string `redis:"desc8" json:"desc8"`
}

var testStruct = TestStruct{
	Id:    1,
	Name:  "devhui",
	Sex:   "男",
	Desc:  "1",
	Desc1: "2",
	Desc2: "",
}

// Test for hash store
func TestHashStore(t *testing.T) {
	_, err := DoHashStore("hash", testStruct)
	if err != nil {
		panic(err)
	}

	dest := &TestStruct{}
	err = DoHashGet("hash", dest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", *dest)
}

// Test for Gob store
func TestDoGobStore(t *testing.T) {
	_, err := DoGobStore("gob", testStruct)
	if err != nil {
		panic(err)
	}

	dest := &TestStruct{}
	err = DoGobGet("gob", dest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", *dest)
}

// Test for json store
func TestDoJsonStore(t *testing.T) {
	_, err := DoJsonStore("json", testStruct)
	if err != nil {
		panic(err)
	}

	dest := &TestStruct{}
	err = DoJsonGet("json", dest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", *dest)
}

func TestSubscribe(t *testing.T) {
	go Subscribe("channel1")
	go Publish("channel1", "this is message")
	time.Sleep(time.Second * 3)
}

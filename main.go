package main

import (
	_ "embed"
	"fmt"
	_ "net/http/pprof"
)

/*-------------------INPUT-------------------*/
//var reader = bufio.new(os.Stdin)
//var writer = bufio.NewWriter(os.Stdout)

//var writer = bufio.NewWriter(os.Stdout)
//var reader = bufio.NewScanner(os.Stdin)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

type UserIsEmpty struct {
}

type UserHasField struct {
	Age uint64 `json:"age"`
}

func FPrint(uIsEmpty UserIsEmpty, uHasField UserHasField) {
	fmt.Printf("FPrint uIsEmpty:%p uHasField:%p\n", &uIsEmpty, &uHasField)
}

func main() {
	uIsEmpty := UserIsEmpty{}
	uHasField := UserHasField{
		Age: 10,
	}
	FPrint(uIsEmpty, uHasField)
	fmt.Printf("main: uIsEmpty:%p uHasField:%p\n", &uIsEmpty, &uHasField)
}

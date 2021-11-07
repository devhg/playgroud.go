package safemap_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/devhg/playgroud.go/pkg/safemap"
)

func TestNewSafeMap(t *testing.T) {
	smap := safemap.NewSafeMap()
	smap.Add(1, 2)
	smap.Add(3, 4)
	if get, ok := smap.Get(1); ok {
		fmt.Println(get)
	}

	smap.Delete(1)
	if _, ok := smap.Get(1); ok {
		log.Fatal("err")
	}

	smap.Close()
}

package m_goroutine_channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeout(t *testing.T) {
	//timeout1()
	//timeout()
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, "statusUpdate()")
	}
}

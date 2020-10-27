package taskpool

import (
	"fmt"
	"testing"
)

func TestNewGoPool(t *testing.T) {
	pool := NewGoPool(WithMaxLimit(10))
	defer pool.Wait()
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 100; i++ {
		v := i
		group.Go(func() error {
			select {
			case <-ctx.Done():
				fmt.Println("canceled ", v)
				return nil
			default:
				if v > 90 {
					fmt.Println("ended ", v)
					return fmt.Errorf("error in %d", v)
				}
				fmt.Println("did ", v)
				return nil
			}
		})
	}
	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}

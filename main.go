package main

import "fmt"

/*-------------------INPUT-------------------*/
//var reader = bufio.new(os.Stdin)
//var writer = bufio.NewWriter(os.Stdout)

//var writer = bufio.NewWriter(os.Stdout)
//var reader = bufio.NewScanner(os.Stdin)

func main() {
	mp := map[int]interface{}{}
	mp[1] = 1
	if _, ok := mp[1]; ok {
		fmt.Println(ok)
	}
}

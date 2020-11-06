package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Respond struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  User   `json:"result,omitempty"`
}

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	//args := os.Args
	//fmt.Println(args)
	//body := `{"status":1,"message":"OK","result":{"name":"brad","age":19}}`
	//respond := &Respond{}
	//_ = json.Unmarshal([]byte(body), respond)
	//if respond.Status != 1 {
	//	panic("获取失败")
	//}
	//fmt.Printf("%+v", respond)
	s := "124"
	split := strings.Split(s, "")
	for _, s2 := range split {
		fmt.Println(strconv.Atoi(s2))
	}
}

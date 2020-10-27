package main

import (
	"fmt"
	"os"
)

//编译时在二进制文件中加入版本信息
//go编译时可以通过ldflags动态的为程序里的某些变量赋值，
//我们可以利用这个特性将go的版本信息和commit的版本信息编
//译到我们的二进制文件中
var (
	gitHash   string
	buildTime string
	goVersion string
)

/*
go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'" -o main pkg/version/goversion.go
./main -v

output:
Git Commit Hash: 75bbf78d103750248044ca7571f8984b346639cf
Build TimeStamp: Thu Oct 22 21:17:50 2020 +0800
GoLang Version: go version go1.15.3 darwin/amd64

Go 其他参数
-o 输出的二进制文件名
-v 编译时显示包名
-p n 开启并发编译，默认情况下该值为 CPU 逻辑核数
-a 强制重新构建
-n 打印编译时会用到的所有命令，但不真正执行
-x 打印编译时会用到的所有命令
-race 开启竞态检测
*/
func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "-v" || args[1] == "--version") {
		fmt.Printf("Git Commit Hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
		return
	}
}

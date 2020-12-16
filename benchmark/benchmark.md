# Go 性能测试

## 基准测试

基准测试主要是通过测试CPU和内存的效率问题，来评估被测试代码的性能，进而找到更好的解决方案。

1. 基准测试的代码文件必须以_test.go结尾
2. 基准测试的函数必须以Benchmark开头，必须是可导出的
3. 基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
4. 基准测试函数不能有返回值
5. b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
6. 最后的for循环很重要，被测试的代码要放到循环里
7. b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能



命令如下

```shell
go test -bench=.
# -bench=标记，接受一个表达式作为参数, .表示运行所有的基准测试。可以只写方法名字，挑选基准测试

go test -bench=. -run=none
# -run=匹配一个从来没有的单元测试方法，过滤掉单元测试的输出，我们这里使用none，因为我们基本上不会创建这个名字的单元测试方法。因为默认情况下 go test 会运行单元测试，为了防止单元测试的输出影响我们查看基准测试的结果.

goos: darwin
goarch: amd64
pkg: github.com/devhg/LearnGo/benchmark/strcontact
//Benchmark名字-CPU        循环次数          平均每次执行时间
BenchmarkAdd-8              5223            212483 ns/op
BenchmarkFormat-8           3456            353151 ns/op
BenchmarkBuffer-8         129862              8876 ns/op
BenchmarkBuilder-8        157110              6456 ns/op
BenchmarkAppend-8          86055             13071 ns/op
PASS
ok      github.com/devhg/LearnGo/benchmark/strcontact   7.559s

//看到函数后面的-8了吗？这个表示运行时对应的GOMAXPROCS的值。
//循环次数：运行for循环的次数也就是调用被测试代码的次数
//平均每次运行时间：执行一次操作花费的时间
```

测试时间默认是1秒，也就是1秒的时间调用次数，我们可以通过以下修改。

```bash
go test -bench=. -benchtime=3s -run=none
# 可以通过-benchtime指定时间
```



## 性能对比

`-benchmem`可以提供每次操作分配内存的次数，以及每次操作分配的字节数。

```shell
go test -bench=. -benchmem -run=none
# -benchmem可以提供每次操作分配内存的次数，以及每次操作分配的字节数。
goos: darwin
goarch: amd64
pkg: github.com/devhg/LearnGo/benchmark/strcontact
BenchmarkAdd-8              5299            248707 ns/op         1063894 B/op        999 allocs/op
BenchmarkFormat-8           3362            358401 ns/op         1096746 B/op       3000 allocs/op
BenchmarkBuffer-8         129304              8602 ns/op            6576 B/op          7 allocs/op
BenchmarkBuilder-8        162169              6431 ns/op            7416 B/op         11 allocs/op
BenchmarkAppend-8          87915             13102 ns/op            9464 B/op         12 allocs/op
PASS
ok      github.com/devhg/LearnGo/benchmark/strcontact   9.539s
```

在代码开发中，对于我们要求性能的地方，编写基准测试非常重要，这有助于我们开发出性能更好的代码。不过性能、可用性、复用性等也要有一个相对的取舍，不能为了追求性能而过度优化。





## 结合pprof性能监控

```shell
go test -bench=. -benchmem -cpuprofile profile.out
# 生成cpu profile
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
# 同时生成cpu和内存 profile
```

使用pprof工具查看

```bash
go tool pprof profile.out
go tool pprof memprofile.out
go tool pprof -http=":8081" [binary] [profile]
```



报错：Failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in %PATH%

是你电脑没有安装gvedit导致的

fq进入gvedit官网https://graphviz.gitlab.io/_pages/Download/Download_windows.html 下载稳定版

mac 安装， 安装好后就可以使用web进行展现了

`brew install graphviz`

https://my.oschina.net/solate/blog/3034188
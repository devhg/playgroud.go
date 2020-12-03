```shell script
# 逃逸分析
go build -gcflags "-m -l -N" xx.go

# trace分析
#
#https://studygolang.com/articles/12639
#https://studygolang.com/articles/9693
go tool trace trace.out


```

1. 基准测试的代码文件必须以_test.go结尾
2. 基准测试的函数必须以Benchmark开头，必须是可导出的
3. 基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
4. 基准测试函数不能有返回值
5. b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
6. 最后的for循环很重要，被测试的代码要放到循环里
7. b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
https://my.oschina.net/solate/blog/3034188
```cassandraql
go test -bench=. 
```
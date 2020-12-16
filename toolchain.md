```shell script
# 逃逸分析
go build -gcflags "-m -l -N" xx.go

# trace分析
#
#https://studygolang.com/articles/12639
#https://studygolang.com/articles/9693
go tool trace trace.out

```

```shell script
# 逃逸分析
go build -gcflags "-m -l -N" xx.go

# trace分析
#
#https://studygolang.com/articles/12639
#https://studygolang.com/articles/9693
go tool trace trace.out



# 代码检查工具
# https://studygolang.com/articles/9619
#go vet，只在一个单独的包内可用，不能使用flag 选项（来激活某些指定的检测）。
#go tool vet更加完整，它可用用于文件和目录。目录被递归遍历来找到包。go tool vet也可以按照检测的分组激活选项。

go tool vet <directory|files>

#虽然vet是不完美的，但是它仍然是一个非常有价值的工具，它应该在所有的Go项目中定期使用。
#它是那么有价值，以至于它甚至可以让我们怀疑是不是有些检测不应该被编译器检测到。为什么有人会编译一个检测到有prrintf格式错误的代码呢？
```


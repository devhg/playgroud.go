package main

import (
	"fmt"
)

// n q, n个服务，q次查询。
// 下面n行，表示第i个服务依赖的服务数量k，下面k个服务id分别x，xx。k=0表示没有依赖
// 下面q行，表示q次查询，每次查询x，y，x=1表示开启y服务，x=0表示关闭y服务
// 重复 开启、关闭同一个服务不会造成任何影响
// 当服务 a 引入了服务 b 作为依赖之后，服务 a 启动时 b 会随之启动，b 停止时 a 会随之停止

// 3 2
// 1 2
// 1 3
// 0
// 1 1
// 0 2

// 3
// 1
type service struct {
	isOpen   bool
	requires []int
}

var serviceMap = make(map[int]*service)
var serviceStatus = make(map[int]bool)
var opened = 0

func main() {
	n, q := 0, 0
	fmt.Scan(&n, &q)
	for i := 1; i <= n; i++ {
		e := &service{isOpen: false}
		serviceMap[i] = e
		serviceStatus[i] = false

		subs := 0
		fmt.Scan(&subs)
		if subs == 0 {
			continue
		}

		for j := 0; j < subs; j++ {
			v := 0
			fmt.Scan(&v)
			e.requires = append(e.requires, v)
		}
	}

	x, y := 0, 0
	for i := 0; i < q; i++ {
		fmt.Scan(&x, &y)
		if x == 1 {
			open(y)
		} else {
			closes(y)
		}
		fmt.Println(opened)
	}
}

// open by bfs
func open(id int) {
	if isOpened := serviceStatus[id]; isOpened {
		return
	}

	q := []int{id}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		e := serviceMap[front]
		if e.isOpen {
			continue
		}

		opened++
		e.isOpen = true
		serviceStatus[front] = true
		q = append(q, e.requires...)
	}
}

// closes by dfs Recursion
func closes(id int) {
	if isOpened := serviceStatus[id]; !isOpened {
		return
	}

	opened--
	serviceStatus[id] = false
	serviceMap[id].isOpen = false
	for k, v := range serviceMap {
		if k != id && v.isOpen && InArray(id, v.requires) {
			closes(k)
		}
	}
}

func InArray(x int, arr []int) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

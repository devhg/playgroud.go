package m_goroutine_channel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。

1. WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。

2. WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，
当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。

3. WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。

4. WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过
Context.Value方法访问到，后面我们会专门讲。
*/
func TestContext(t *testing.T) {
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	go watch(ctx, "监控3")

	time.Sleep(4 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出", ctx.Err())
			return
		default:
			fmt.Println(name, "goroutine监控中")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContext2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//go handle(ctx, 500*time.Millisecond)
	go handle(ctx, 1500*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestChan(t *testing.T) {
	ints := make(chan int, 1)
	ints <- 1
	close(ints)

	go func() {
		fmt.Println(<-ints)
	}()
}

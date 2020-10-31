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


5. type CancelFunc func()取消函数的类型，
该函数可以取消一个Context，以及这个节点Context下所有的所有的Context，不管有多少层级。
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

/*
Context 使用原则
1. 不要把Context放在结构体中，要以参数的方式传递
2. 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
3. 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
4. Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
5. Context是线程安全的，可以放心的在多个goroutine中传递
*/

/*************************************************/

//我们可以使用context.WithValue方法附加一对K-V的键值对，这里Key必须是等价性的，也就是具有可比性；Value值要是线程安全的。
//
//这样我们就生成了一个新的Context，这个新的Context带有这个键值对，在使用的时候，可以通过Value方法读取ctx.Value(key)。
//
//记住，使用WithValue传值，一般是必须的值，不要什么值都传递。
func TestWithValue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, "name", "监控1")
	go watchValueCtx(valueCtx)
	go watchValueCtx(valueCtx)
	go watchValueCtx(valueCtx)

	time.Sleep(2 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(1 * time.Second)
}

func watchValueCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("name"), "监控推出", ctx.Err())
			return
		default:
			fmt.Println(ctx.Value("name"), "监控中")
			time.Sleep(1 * time.Second)
		}
	}
}

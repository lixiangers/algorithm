//实现1/0轮番打印
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//method1()
	//method2()
	//method3()
	method5()
}

// 采用双channel 方式　循环打印
func method1() {
	defer fmt.Println("goroutine count:", runtime.NumGoroutine())

	ch1 := make(chan int)
	ch2 := make(chan int)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	go func(ctx context.Context) {
		//保证在发送的goroutine中close channel.否则会发送 send at closed channel panic
		defer close(ch2)
		for {
			select {
			case num, ok := <-ch1:
				if ok {
					fmt.Println(num, time.Now().UnixNano())
					ch2 <- 1
				} else {
					fmt.Println("goroutine1 end")
					return
				}
			case <-ctx.Done():
				fmt.Println("goroutine1 end")
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		//保证在发送的goroutine中close channel.否则会发送 send at closed channel panic
		defer close(ch1)
		for {
			select {
			case num, ok := <-ch2:
				if ok {
					fmt.Println(num, time.Now().UnixNano())
					ch1 <- 0
				} else {
					fmt.Println("goroutine2 end")
					return
				}

			case <-ctx.Done():
				fmt.Println("goroutine2 end")
			}
		}
	}(ctx)

	ch1 <- 0

	select {
	case <-ctx.Done():
		fmt.Println("goroutine main end")
	}

}

// 采用单channel 方式　循环打印 0/1
// 这里两个goroutine 在一次发送读取回合中只现在一个goroutine中打印。这样的话能保证是交替打印的。例如:发送goroutine中在发送后打印，
// 对应的接收goroutine不会打印。因为发送后还有打印，打印期间接收G会阻塞一直等到发送G打印完成后发送数据。这次发送G发送后不打印。接收G
// 会打印。发送G一直阻塞到　接收G打印完成后开始准备接收数据.打印是顺序的
// 如果采用发送接收过程部分奇偶依次打印.在接收/打印fmt.println的顺序无法保证。接收和发送都是实时的，他们和打印不是原子的。无法保证他们的
// 顺序
func method2() {
	defer fmt.Println("goroutine count:", runtime.NumGoroutine())

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	msgChan := make(chan int)

	var sendNumber = 1
	go func(ctx context.Context) {
		defer close(msgChan)
		for {
			select {
			case msgChan <- sendNumber:
				if sendNumber%2 == 1 {
					fmt.Println("groutine-1:", 0, ">>>>", sendNumber)
				}
				sendNumber++
			case <-ctx.Done():
				fmt.Println("goroutine1 end")
				return
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case number, ok := <-msgChan:
				if ok {
					if number%2 == 0 {
						fmt.Println("groutine-2:", 1, ">>>>", number)
					}
				} else {
					fmt.Println("goroutine2 end")
				}
			case <-ctx.Done():
				fmt.Println("goroutine2 end")
			}
		}
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("goroutine main end")
	}

}

// runtime.GOMAXPROCS(1)设置单核，runtime.Gosched()让出时间片
func method3() {
	//var group sync.WaitGroup
	runtime.GOMAXPROCS(1)

	//group.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(2*i - 1)
			runtime.Gosched()
		}
		//group.Done()
	}()

	//group.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(2 * i)
			runtime.Gosched()
		}
		//group.Done()
	}()

	//不能使用sync.WaitGroup ，因为这样main goroutine也会占用时间片。不符合预期。
	//time.Sleep会让main goroutine　不占用时间片。真正的只有２个goroutine在轮流执行
	//group.Wait()
	time.Sleep(time.Second * 5)
}

// 错误的示例
// 在接收发送是两个G.发送和接收是实时感知的。println和发送/接收不是原子的。两个打印会同时进行。顺序无法保证
func method4() {
	msg := make(chan int)
	var POOL = 100
	go func() {
		for i := 1; i <= POOL; i++ {
			select {
			case msg <- i:
				fmt.Println("groutine-1:", 0)
			}
		}
	}()
	go func() {
		for i := 1; i <= POOL; i++ {
			select {
			case <-msg:
				fmt.Println("groutine-2:", 1)
			}
		}
	}()
}

//通过lock sync.cond实现轮流打印1 0
func method5() {
	var mailbox uint8
	var lock sync.Mutex
	var sendCond = sync.NewCond(&lock)
	var recvCond = sync.NewCond(&lock)

	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	// sign 用于传递演示完成的信号。
	go func(ctx context.Context, index int) { // 用于发信。
		for {
			select {
			case <-ctx.Done():
				return
			default:
				lock.Lock()
				for mailbox == 1 {
					//fmt.Println("信箱满了")
					sendCond.Wait()
				}
				mailbox = 1
				fmt.Println(index, ":1")
				lock.Unlock()
				recvCond.Broadcast()
			}
		}
	}(ctx, 0)

	go func(ctx context.Context, index int) { // 用于收信。
		for {
			select {
			case <-ctx.Done():
				return
			default:
				lock.Lock()
				for mailbox == 0 {
					//没有信
					//fmt.Println("信箱空了")
					recvCond.Wait()
				}
				mailbox = 0
				fmt.Println(index, ":0")
				lock.Unlock()
				sendCond.Broadcast()
			}
		}
	}(ctx, 0)

	select {
	case <-ctx.Done():
		fmt.Println("goroutine main end")
	}
}

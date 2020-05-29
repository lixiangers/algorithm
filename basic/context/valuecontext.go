package main

import (
	"context"
	"fmt"
	"time"
)

// value context 的上游parent context 有cancel context的话，value context就可以监听Done()
// 但是如果valueContext的上游节点没有cancelContext的话，<-Done()会被阻塞的。因为Done()是nil
// 原因是我们调用context.Done()的是获取的parent的Done().parent 是做为内匿名类型在valueContext 中的
// type valueCtx struct {
//	Context
//	key, val interface{}
//}
func main() {
	//testValueContextHaveCancelContext()
	testValueContextWithoutCancelContext()
}

func testValueContextHaveCancelContext() {
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
	valueCtx1 := context.WithValue(timeoutCtx, "s", "d")
	valueCtx2 := context.WithValue(valueCtx1, "s", "d")

	// value context也有deline done() 也可以监听Done()
	fmt.Println(valueCtx2.Deadline())
	fmt.Println(valueCtx2.Done())

	go func(ctx context.Context) {
		for {
			select {
			case <-valueCtx2.Done():
				fmt.Println("value  context")
				return
			}
		}
	}(valueCtx2)

	<-valueCtx2.Done()
	time.Sleep(time.Second)
	fmt.Println("end")
}

// output: all goroutines are asleep - deadlock!
// valueContext的上游节点没有cancelContext的话，<-Done()会被阻塞的。因为Done()是nil
func testValueContextWithoutCancelContext() {
	timeoutCtx := context.WithValue(context.Background(), "x", "y")
	valueCtx1 := context.WithValue(timeoutCtx, "s", "d")
	valueCtx2 := context.WithValue(valueCtx1, "s", "d")

	// value context也有deline done() 也可以监听Done()
	fmt.Println(valueCtx2.Deadline())
	fmt.Println(valueCtx2.Done())

	go func(ctx context.Context) {
		for {
			select {
			// 一直被阻塞
			case <-valueCtx2.Done():
				fmt.Println("value  context")
				return
			}
		}
	}(valueCtx2)

	// 不注释的话 all goroutines are asleep - deadlock!
	//<-valueCtx2.Done()
	time.Sleep(time.Second * 10)
	fmt.Println("end")
}

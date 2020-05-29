package main

import "fmt"

func main() {
	closure()
}

//5: 1
//4 2
//2: 3
//1: 0
func closure() {
	i := 0
	// 不是闭包 i的值直接压入都defer栈中
	defer fmt.Println("1:", i)
	// 是闭包，对i的引用
	defer func() {
		i += 1
		fmt.Println("2:", i)
	}()

	// defer 返回的是一个func 不会对里面的func执行
	defer func() func() {
		return func() {
			i += 1
			fmt.Println("3", i)
		}
	}()
	// defer返回的func会被执行，同时也是闭包
	defer func() func() {
		return func() {
			i += 1
			fmt.Println("4", i)
		}
	}()()

	i += 1
	fmt.Println("5:", i)
}

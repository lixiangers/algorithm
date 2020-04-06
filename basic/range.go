package main

import (
	"fmt"
	"strings"
)

type MyInt1 int
type MyInt2 = int
type testStruct struct {
	name string
}

func (t testStruct) Method1(param string) string {
	return "A"
}

func (t *testStruct) Method2(param string) string {
	return "B"
}

func GetValue() int {
	return 1
}

var x error = nil

type person struct {
	name string
}

func hello(num ...int) {
	num[0] = 18
}

func increaseA() int {
	var i int
	defer func() {
		i++
		fmt.Println("AAA")
	}()
	return i
}
func g(x *interface{}) {
}

type S struct {
	m string
}

func f() *S {
	return &S{"x"}
}

type Math struct {
	x, y int
}

var p *int
type T struct {
	x int
	y *int
}
func test1() {
	//删除字符串右边的在cuts中的字符。直到不在里面的字符停止
	fmt.Println(strings.TrimRight("AEABBA", "AB"))
	fmt.Println(strings.TrimRight("abbdcddaaaaa", "dac"))
}
type ConfigOne struct {
	Daemon string
}



func Stop(stop chan<- bool) {}

func main() {
	test1()
}

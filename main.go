package main

import (
	list2 "container/list"
	"errors"
	"fmt"
	"github.com/lixiangers/algorithm/highlevel"
	"math"
	"reflect"
	"runtime"
	"time"
)

const cl = 100

var bl = 123

type testfunc func(string2 string) string
type test struct {
	name string
	//m    map[int]string
}

type People interface {
	Speak(string) string
}

type Stduent struct {
	//m   map[string]string
	m   interface{}
	min int
}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func productor(channel chan<- string) {
	for {
		channel <- fmt.Sprintf("1")
		time.Sleep(time.Second * time.Duration(1))
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
	}
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
		stu.Age = 10
	}
	fmt.Println(stus)
}

func main() {
	l := list2.New()
	l.Back()
	runtime.GOMAXPROCS(1)
	//channel := make(chan string) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	//go productor(channel)        // 将 productor 函数交给协程处理, 产生的结果传入信道中
	//customer(channel)            // 主线程从信道中取数据

	pase_student()
	var peo People = &Stduent{}
	// 实例方法的范围比指针的方法更广。如果一个指针实现了一个接口，那struct并不会。反之一个struct实现了一个接口。指针也实现了这个接口.
	// 对于不牵扯到指针的化。方法无论是指针方法还是struct方法，都可以互相调用
	think := "bitch"

	stduent := Stduent{}
	stduent.Speak("")
	s := &Stduent{}
	s.Speak("dxx")

	fmt.Println(peo.Speak(think))

	//i := GetValue()
	//
	//switch i.(type) {
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//default:
	//	println("unknown")
	//}

	println(DeferFunc1(1)) //4
	println(DeferFunc2(1)) //1
	println(DeferFunc3(1)) //3

	list := new([]int)
	ints := *list
	fmt.Println(len(ints))
	ints = append(ints, 1)
	fmt.Println(ints)
	fmt.Println(reflect.TypeOf(list))

	m := new(map[int]int)
	m2 := *m
	//m2[1]=2
	fmt.Println(m2)
	fmt.Println(reflect.TypeOf(m))

	s2 := new(Stduent)
	fmt.Println(s2.Speak("s"))
	fmt.Println(reflect.TypeOf(s2))

	s1 := []int{1, 2, 3}
	s3 := []int{4, 5}
	s1 = append(s1, s3...)
	fmt.Println(s1)

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	m   map[string]string
	//	age int
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	m   map[string]string
	//	age int
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}

	sm11 := struct {
		age int
		p   Stduent
	}{p: Stduent{m: make(map[int]int, 3)}}

	sm12 := struct {
		age int
		p   Stduent
	}{p: Stduent{m: make(map[int]int)}}

	if reflect.DeepEqual(sm11, sm12) {
		fmt.Println("sm11 ==sm12")
	} else {
		fmt.Println("sn1 !=sm")
	}
	//if sm11 == sm12 {
	//	fmt.Println("sm11 == sm12")
	//}

	var x *int = nil
	Foo(x)

	intmap := map[int]string{1: "a", 2: "bb", 3: "ccc",}
	fmt.Println(intmap)

	println(&bl, bl)
	//println(&cl, cl)

	//loop:
	for i := 0; i < 10; i++ {
		println(i)
	}
	goto loop1
loop1:
	fmt.Println("end")

	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

	var u1 MyUser1
	var u2 MyUser2
	u1.m1()
	u2.m2()

	var u User
	u.m2()

	my := MyStruct{}
	//my.m1() //ambiguous selector my.m1 结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示
	my.T1.m1()
	my.T2.m1()

	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))

	funs := test2_2()
	for _, f := range funs {
		f()
	}

	a, b := test3(100)
	a()
	b()
	math.Min

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println("fatal")
	//	}
	//}()
	//
	//defer func() {
	//	panic("defer panic")
	//}()
	//panic("panic")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err)
			fmt.Println(f())
			fmt.Println(reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	//panic("panic")

	testSliceRange()

	var sss string = "天气真好123xx"
	for i, j := range sss {
		fmt.Println(i, string(j))
	}

	fmt.Println(hammingDistance(1, 4))
	lengthOfLongestSubstring("你好吧abc你呀")
	ss := "abcd"
	fmt.Println(reflect.TypeOf(ss[0]))

	var c1 chan int = nil
	select {
	case <-c1:
		fmt.Println("nil channel read")
	default:
		fmt.Println("default")
	}

	highlevel.GetFullPermutation([]int{1})
	head:=&ListNode{Val:4}
	node1:=&ListNode{Val:2}
	node2:=&ListNode{Val:1}
	node3:=&ListNode{Val:3}
	head.Next=node1
	node1.Next=node2
	node2.Next=node3

    fmt.Println(sortList(head))
}

//func GetValue() interface{} { return 1 }

func DeferFunc1(i int) (t int) {
	t = i
	defer func() { t += 3 }()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() { t += 3 }()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() { t += i }()
	return 2
}
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

type User struct{}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() { fmt.Println("MyUser1.m1") }
func (i User) m2()    { fmt.Println("User.m2") }

type T1 struct{}

func (t T1) m1() { fmt.Println("T1.m1") }

type T2 = T1
type MyStruct struct {
	T1
	T2
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		// 变量作用域 因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}

	return err
}

//func DoTheThing(reallyDoIt bool) (err error) {
//	var result string
//	if reallyDoIt {
//		result, err = tryTheThing()
//		if err != nil || result != "it worked" {
//			err = ErrDidNotWork
//		}
//	}
//	return err
//}

func tryTheThing() (string, error) { return "", ErrDidNotWork }

func test2() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() { println(&i, i) })
	}
	return funs
}
func test2_2() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() { println(&x, x) })
	}
	return funs
}

func test3(x int) (func(), func()) {
	return func() {
		println(x)
		x += 10
	}, func() { println(x) }
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum, currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func testSliceRange() {
	s1 := make([]int, 5, 5)
	s1[0] = 1
	s1[1] = 2
	fmt.Println(len(s1), cap(s1)) // len=5 cap=5
	fmt.Println(s1)               //1,2,0,0,0
	change(s1)
	fmt.Println(s1) // 1,2,0,0,0   len=5再添加就会造成 change会造成扩容，对s1没有影响．扩容会产生新的slice,新的slice的有新的数组指针

	s2 := s1[0:2]
	change(s2) // len=2 cap=5
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s1) // 1,2,3,0,0

}

func change(s []int) {
	s = append(s, 3)
}

/**
 * 快慢指针 .
 * 通过快慢指针找到中点，反转后半部分链表且进行比较
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找出中点，快指针到了链表结尾，慢指针也就到了链表中点
	mid := findMid(head)
	// 翻转后半部分链表
	rev := reverse(mid)
	// 比对前后链表
	for rev != nil && head != nil {
		if head.Val != rev.Val {
			return false
		}
		rev, head = rev.Next, head.Next
	}
	return true
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	// 经过遍历，后半部分链表会变成一个头节点为 prev，最后为 nil 的链表
	var prev, curr *ListNode = nil, head
	for curr != nil {
		after := curr.Next
		curr.Next = prev

		prev = curr
		curr = after
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hammingDistance(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor > 0 {
		if xor&1 > 0 {
			distance++
		}
		xor >>= 1
	}

	return distance
}
func lengthOfLongestSubstring(s string) int {
	max, start := 0, 0
	m := make(map[string]int)

	runeSlice := []rune(s)
	for i, r := range runeSlice {
		if index, ok := m[string(r)]; ok {
			// 存在，还要判断是否在start ,end范围内．判断的标准是 i>=start
			if index >= start {
				//重复的下一个位置才算窗口的左边界
				start = index + 1
			}
		}

		m[string(r)] = i
		// 每次都要计算当前窗口的值并和最大的窗口值比较保存
		if i-start+1 > max {
			max = i - start + 1
		}
	}

	return max
}
func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	result := [][]int{}
	used := make([]int, len(nums))
	callPermute(nums, 0, used, &result)
	return result
}

var numbers []int

func callPermute(candidate []int, index int, usedIndex []int, result *[][]int) {
	if len(numbers) == len(candidate) {
		temp := make([]int, len(numbers))
		copy(temp, numbers)

		*result = append(*result, temp)
		numbers = nil
		return
	}

	for i, item := range candidate {
		if usedIndex[i] == 0 {
			numbers = append(numbers, item)
			usedIndex[i] = 1
			callPermute(candidate, index+1, usedIndex, result)
			usedIndex[i] = 0
		}
	}
}
func sortList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}

	//快慢指针找到分割点
	slow,fast:=head,head.Next
	for fast!=nil && fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}

	rightHeadNode:=slow.Next
	// 需要设置nil分割链表
	slow.Next=nil

	//对左部分进行排序
	node1:=sortList(head)
	// 对右部分进行排序
	node2:=sortList(rightHeadNode)
	// 合并
	node:=merge(node1,node2)
	return node
}

func merge(node1 *ListNode,node2 *ListNode)*ListNode{
	if node1==nil{
		return node2
	}

	if node2==nil{
		return node1
	}

	head:=new(ListNode)
	cur:=head

	for (node1!=nil && node2!=nil){
		if node1.Val<node2.Val{
			cur.Next,node1=node1,node1.Next
		}else{
			cur.Next,node2=node2,node2.Next
		}
		cur=cur.Next
	}

	// 还有某个链表没有完的情况
	if node1!=nil{
		cur.Next=node1
	}else if node2!=nil{
		cur.Next=node2
	}

	return head.Next
}


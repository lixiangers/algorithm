package main

import "fmt"

func False() bool {
	return false
}

func False2() bool {
	return true
}

func IntFunc() int {
	return 1
}

// ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
// A missing switch expression is equivalent to the boolean value true
func main() {
	// {在下方的时候,会自动加上; 等价于switch False; 这种情况下没有 Expression ,则默认Expression是true
	// 1
	switch False(); {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	// 2
	switch False(); {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	// 1,2 是等价的

	// 3. 运行出错:cannot use 100 (type untyped int) as type bool 证明是默认Expression是true
	//switch IntFunc()
	//{
	//default:
	//	fmt.Println("default")
	//case 100:
	//	fmt.Println("1")
	//}

	switch IntFunc() {
	case 2:
		fmt.Println("2")
	case 1:
		fmt.Println("1")
		// fallthrough 直接执行下面的case.无论下面的case条件是否满足
		fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}
}

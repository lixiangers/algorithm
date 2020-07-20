package string

import (
	"fmt"
	"unicode/utf8"
)

type AcNode struct {
	fail      *AcNode
	isPattern bool
	next      map[rune]*AcNode
	length    int
	content   rune
}

func newAcNode(c rune) *AcNode {
	return &AcNode{
		fail:      nil,
		isPattern: false,
		next:      map[rune]*AcNode{},
		content:   c,
	}
}

type AcAutoMachine struct {
	root *AcNode
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: newAcNode([]rune("/")[0]),
	}
}

func (ac *AcAutoMachine) AddPattern(pattern string) {
	chars := []rune(pattern)
	iter := ac.root
	for _, c := range chars {
		if _, ok := iter.next[c]; !ok {
			iter.next[c] = newAcNode(c)
		}
		iter = iter.next[c]
	}
	iter.isPattern = true
	iter.length = utf8.RuneCountInString(pattern)
}

func (node *AcNode) String() string {
	return string(node.content)
}

func (ac *AcAutoMachine) Build() {
	queue := []*AcNode{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.next {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				// fail指针设置原则是：
				// 1）查看father->failNode下有没有和自己一样的子节点，有则fail指针取该子节点
				// 2）否则，沿father->failNode->failNode继续查询下，如果一直没有，fail指针就取根节点
				failAcNode := parent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.next[char]; ok {
						child.fail = failAcNode.next[char]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					child.fail = ac.root
				}
			}
			queue = append(queue, child)
		}
	}
}

func (ac *AcAutoMachine) Query(content string) (results []string) {
	chars := []rune(content)
	p := ac.root
	fmt.Printf("start:%s\n", p.String())
	for i, c := range chars {
		fmt.Printf("loop:%s\n", string(c))
		for {
			_, ok := p.next[c]
			if !ok && p != ac.root {
				// 如果主串中的当前字符不在Tire树当前节点的children中，找到fail指针的节点
				p = p.fail
				//fmt.Printf("for set p is fail:%s\n",p.String())
			} else {
				break
			}
		}

		p = p.next[c]
		if p != nil {
			fmt.Printf("for set p is next:%s\n", p.String())
		}

		// 如果没有匹配的，从root开始重新匹配
		if p == nil {
			p = ac.root
			fmt.Printf("set p is root:%s\n", p.String())
		} else {
			tmp := p
			for tmp != ac.root {
				// 打印出可以匹配的模式串
				if tmp.isPattern == true {
					pos := i - tmp.length + 1
					results = append(results, string([]rune(content)[pos:pos+tmp.length]))
					fmt.Printf("匹配起始下标%d,长度:%d \n", pos, tmp.length)
				}
				tmp = tmp.fail
			}
		}
	}
	return
}

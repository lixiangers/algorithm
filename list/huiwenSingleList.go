// 判断是否回文单链表
package list

import "fmt"

// 快慢指针实现找到中点位置,注意奇偶.慢指针走1,快指针走2
// 慢指针在走的过程中反转前部分指针
// 找到中点后开一个新指针，从中点往回走。慢指针继续往后走，直到走到尾部.走的过程判断是否每个都一样。都一样则是回文否则不是
// 现在反转的方案会造成改变以前list的数据。后续采用Stack方式，采用新增空间的方式不改变以前list结构达到同一效果
func IsHuiWenSingListWithReverse(list *SingleList) bool {
	quickIndex := 0
	slowIndex := 0
	if nil == list || list.len == 1 {
		return false
	}

	var slowCurrentNode *SingleListNode = list.root
	var quickCurrentNode *SingleListNode = list.root

	var after *SingleListNode
	var pre *SingleListNode

	// 循环到到快指针走到尾部
	for quickIndex != list.len-1 {
		// 快指针先走否则会慢指针的反转会影响快指针
		if nil != quickCurrentNode.next {
			quickCurrentNode = quickCurrentNode.next
			quickIndex += 1
		}
		if nil != quickCurrentNode.next {
			quickCurrentNode = quickCurrentNode.next
			quickIndex += 1
		}

		// 反转前部分
		after = slowCurrentNode.next
		slowCurrentNode.next = pre
		pre = slowCurrentNode

		slowCurrentNode = after
		slowIndex++
	}

	// 慢指针一直走到tail,在中间新开一个指针往前走(前面已经反转)
	if list.len%2 == 0 {
		// 当时偶数的时候after和pre 不用调整
	} else {
		//当是奇数的时候after需要后移一位。中间这位不做比较
		after = after.next
	}
	result := true

	for nil != after && nil != pre {
		if after.value != pre.value {
			result = false
			break
		}
		slowIndex++
		after = after.next
		pre = pre.next
	}

	fmt.Println("middle index:", slowIndex, " is hui wei:", result)
	return result
}

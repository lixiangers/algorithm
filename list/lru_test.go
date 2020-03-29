package list

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回  1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)    // 返回  4

	//[null,null,null,1,null,2,null,1,3,4]

	//[null,null,null,1,null,-1,null,-1,3,4]
}

package main

import "sync"

var onceSingleDataStruct1Init sync.Once
var onceSingleDataStruct1Instance *SingleDataStruct1

var onceSingleDataStruct2Instance *SingleDataStruct2 = &SingleDataStruct2{name: "init"}

var onceSingleDataStruct3InstanceLock sync.Mutex
var onceSingleDataStruct3Instance *SingleDataStruct3

var onceSingleDataStruct4InstanceLock sync.Mutex
var onceSingleDataStruct4Instance *SingleDataStruct4

// 单例模式实现１.利用once实现.实现简单安全
func GetSingleDataStruct1Instance() *SingleDataStruct1 {
	onceSingleDataStruct1Init.Do(func() {
		onceSingleDataStruct1Instance = &SingleDataStruct1{name: "init"}
	})
	return onceSingleDataStruct1Instance
}

// 单例模式实现2.饿汉模式.如果初始化复制，会加长初始化时间
func GetSingleDataStruct2Instance() *SingleDataStruct2 {
	return onceSingleDataStruct2Instance
}

//　单例模式3.懒汉模式。双重锁机制.避免了不需要的锁，提高了效率
func GetSingleDataStruct3Instance() *SingleDataStruct3 {
	if onceSingleDataStruct3Instance == nil {
		onceSingleDataStruct3InstanceLock.Lock()
		defer onceSingleDataStruct3InstanceLock.Unlock()

		if onceSingleDataStruct3Instance == nil {
			onceSingleDataStruct3Instance = &SingleDataStruct3{name: "xxx"}
		}
	}
	return onceSingleDataStruct3Instance
}

//　单例模式4.懒汉模式。双重锁机制. 每次都加锁，效率低
func GetSingleDataStruct4Instance() *SingleDataStruct4 {
	onceSingleDataStruct4InstanceLock.Lock()
	defer onceSingleDataStruct4InstanceLock.Unlock()

	if onceSingleDataStruct4Instance == nil {
		onceSingleDataStruct4Instance = &SingleDataStruct4{name: "xxx"}
	}
	return onceSingleDataStruct4Instance
}

type SingleDataStruct1 struct {
	name string
}

type SingleDataStruct2 struct {
	name string
}

type SingleDataStruct3 struct {
	name string
}

type SingleDataStruct4 struct {
	name string
}

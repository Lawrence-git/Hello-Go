package main

import (
	"fmt"
	"sync"
)

type info struct {
	//sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区
	//RWMutex 锁：通过 RLock() 来允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作。
	rm   sync.RWMutex
	age  int
	name string
}

func mutexDemo() {
	ms := new(info)
	ms.age = 20
	ms.name = "TestDemo"
	fmt.Println(ms.name)
	updateName(ms)
	fmt.Println(ms.name)
}

func updateName(ifdemo *info) {
	ifdemo.rm.RLock()
	ifdemo.name = "Test update"
	ifdemo.rm.RUnlock()
}

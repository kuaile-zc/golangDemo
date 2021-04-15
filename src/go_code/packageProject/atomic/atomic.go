package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) add(add int) {
	a.lock.Lock()
	defer  a.lock.Unlock()
	a.value += add;
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer  a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.add(1)
	go func() {
		a.add(1)
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

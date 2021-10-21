package main

import (
	"fmt"
	"sync"
)

//控制循环次数
var count = 10

func main() {
	wg := sync.WaitGroup{}

	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)
	chanD := make(chan struct{}, 1)

	chanA <- struct{}{}

	wg.Add(4)

	go printA(&wg, chanA, chanB)
	go printB(&wg, chanB, chanC)
	go printC(&wg, chanC, chanD)
	go printD(&wg, chanD, chanA)

	wg.Wait()
}

func printA(wg *sync.WaitGroup, chanA, chanB chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanA
		fmt.Println("张三")

		chanB <- struct{}{}
	}
}

func printB(wg *sync.WaitGroup, chanB, chanC chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanB
		fmt.Println("李四")

		chanC <- struct{}{}
	}
}

func printC(wg *sync.WaitGroup, chanC, chanA chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		fmt.Println("王五")

		chanA <- struct{}{}
	}
}

func printD(wg *sync.WaitGroup, chanC, chanA chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		fmt.Println("赵六")

		chanA <- struct{}{}
	}
}

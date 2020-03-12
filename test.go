package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时执行
	interval := time.Duration(time.Second * 1)
	ticker_ := time.NewTicker(interval)
	defer ticker_.Stop()

	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	for {
		<-ticker_.C
		go f1(ch1)
		go f2(ch2)
		go f3(ch3)

		<-ch1
		<-ch2
		<-ch3
	}

}

func f1(ch chan int) {
	fmt.Println("f1")
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("f2")
	ch <- 1
}

func f2(ch chan int) {
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("f4")
	fmt.Println(time.Now().Unix())
	ch <- 2
}

func f3(ch chan int) {
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("f5")
	fmt.Println(time.Now().Unix(), "f5")
	ch <- 3
}

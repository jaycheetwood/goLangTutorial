package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3) // n will be between 0 and 3
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomethingInRoutine(msg string) {
	pause()
	fmt.Println(msg)
}

func goRoutineTutorial() {
	doSomething("sync1")

	go doSomething("async1")
	go doSomething("async2")
	go doSomething("async3")

	doSomething("sync2")

	time.Sleep(time.Second * 10)
}

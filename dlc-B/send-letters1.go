//DLC-B
package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "A (3s)\t"
	time.Sleep(1 * time.Second)
	channel <- "B (4s)\t"
	time.Sleep(1 * time.Second)
	channel <- "C (5s)\t"
	time.Sleep(1 * time.Second)
	channel <- "D (6s)\t"
}

func main() {
	fmt.Println(time.Now())
	channel := make(chan string, 1) //добавлен аргумент!
	go sendLetters(channel)
	time.Sleep(3 * time.Second)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now()) //это и предыдущее выводятся одновременно.
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(time.Now())
}

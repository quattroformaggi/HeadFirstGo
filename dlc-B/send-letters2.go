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
	channel := make(chan string, 4) //изменён аргумент!
	go sendLetters(channel)
	time.Sleep(3 * time.Second)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now()) //все выводятся одновременно
	fmt.Println(time.Now())
}

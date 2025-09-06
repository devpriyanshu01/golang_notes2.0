package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inside Main Fn.")
	go printNumbers()
	go printLetters()

	time.Sleep(time.Second*3)
}

func printNumbers() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for _,rune := range "ABCDE" {
		fmt.Println(time.Now())
		fmt.Println(string(rune))
		time.Sleep(time.Millisecond * 500)
	}
}
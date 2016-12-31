package main

import (
	"fmt"
	"strconv"
	"time"
)

type Token struct {
	recipient int
	data      string
}

var N int

func main() {
	N = 10
	channel := make(chan Token)
	token := Token{recipient: 10, data: "token"}
	for i := 1; i <= N; i++ {
		if i == 1 {
			go first(channel, token)
		} else if i == N {
			go receive(channel, token, i)
		} else {
			go receive(channel, token, i)
			go send(channel, token, i)
		}
		time.Sleep(1 * 1e9)
	}
}

var i int

func receive(channel chan<- Token, t Token, index int) {
	if index != N {
		fmt.Println("Thread  " + strconv.Itoa(index) + "  receive ..." + t.data)
		channel <- t
	} else {
		fmt.Println("Token ¹ " + strconv.Itoa(index) + " data : " + t.data + " reached the destination.")
	}
}

func send(channel <-chan Token, t Token, index int) {
	s := <-channel
	fmt.Println("Thread  " + strconv.Itoa(index) + "  send ..." + s.data)
}
func first(channel chan Token, t Token) {
	fmt.Println("Thread  " + strconv.Itoa(i) + "  send ..." + t.data)
	channel <- t
}
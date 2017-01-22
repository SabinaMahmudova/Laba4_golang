package main

import "fmt"
import "strconv"
import "strings"

type Token struct {
	channel chan Token
	recipient int
	data      string
}
var n=10;
func main() {
	var chann = make([]chan Token, n+1)
	output := make(chan string)             //сообщ с дан. токена/что токен не достиг адресата. нужен чтобы после вывода output функция main завершила работу	
	for i := range chann{
		chann[i] = make(chan Token)
	}
	for i := 0; i <= n-1; i++ {
			go run(chann[i], chann[i+1], output, i)
	}
	var token Token;
	token.recipient = 8
	token.data = "token"
	chann[0] <- token
  	fmt.Println(<-output)
}
func run(a <-chan  Token, b chan <- Token, output chan string, num int) {
	token := <-a
	fmt.Println("Thread", num, ": received Token for thread", token.recipient)
	if token.recipient == num {                                                                                      // если мы достигли адресата
		output <- strings.Join([]string{"I am the reciever ", strconv.Itoa(num), ". Data: ", token.data, }, "")
	}else if num!= n-1 {                                                                                             //иначе проверяем что мы не дошли до конца цепочки и передаем токен дальше
		fmt.Println("Thread", strconv.Itoa(num), ": sending to", strconv.Itoa(num+1))
		b <- token
	}else{                                                                                                           //мы достигли конца, выводим сообщение
		output <- "You reaced the end of the chain. There's no reciever"
	}
}
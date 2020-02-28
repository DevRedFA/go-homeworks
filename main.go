package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//В качестве задания предлагается написать самый простой аналог netcat.
//Нужно принимать ввод в stdin и отправлять в Echo Server. Ответ Echo Server писать в stdout.

func main() {

	conn, _ := net.Dial("tcp", "localhost:8000")

	go readData(conn)

	c := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(c)

	fmt.Println("Print you message to server: ")
	for s.Scan() {
		_, err := fmt.Fprintf(conn, s.Text()+"\n")
		if err != nil {
			fmt.Println("Error during sending message to server:", err)
		}
		fmt.Println("Print you message to server: ")

	}
	if err := s.Err(); err != nil {
		fmt.Println("error while reading data from console, err:", err)
	}

}

func readData(c net.Conn) {
	defer c.Close()
	s := bufio.NewScanner(c)
	for s.Scan() {
		fmt.Println("Answer from server: " + s.Text())
	}
	if err := s.Err(); err != nil {
		fmt.Println("error while reading answer from server, err:", err)
	}
}

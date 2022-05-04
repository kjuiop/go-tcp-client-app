package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var number int = 0

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Failed to Dial : ", err)
	}
	defer conn.Close()

	go func(c net.Conn) {
		for {
			number = number + 1
			send := "Hello " + strconv.Itoa(number)
			_, err = c.Write([]byte(send))
			if err != nil {
				fmt.Println("Failed to write data : ", err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}(conn)

	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				fmt.Println("Failed to Read data : ", err)
				break
			}

			fmt.Println(string(recv[:n]))
		}
	}(conn)

	fmt.Scanln()
}

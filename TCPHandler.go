package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func tcpCall() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			break
		}

		fmt.Print("-> ", string(netData))
		sarr := strings.SplitAfter(string(netData), " ")

		if strings.TrimSpace(string(sarr[0])) == "put" {

			if len(sarr) != 3 {
				fmt.Println(sarr[0])
				c.Write([]byte("Please provide key and value"))
				return
			}

			c.Write([]byte(KeyCreate(c, sarr[1], sarr[2]) + "\n"))
		} else if strings.TrimSpace(string(sarr[0])) == "get" {
			if len(sarr) != 2 {
				c.Write([]byte("Please provide key "))
				return
			}
			c.Write([]byte(GetKey(sarr[1]) + "\n"))
		} else {
			c.Write([]byte("Please provide put or get"))
		}
		fmt.Println(temp)
	}
	c.Close()
}

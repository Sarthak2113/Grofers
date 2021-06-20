package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/garyburd/redigo/redis"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		sarr := strings.SplitAfter(text, " ")
		fmt.Fprintf(c, text+"\n")

		if strings.TrimSpace(string(sarr[0])) == "watch" {
			c, _ := redis.Dial("tcp", ":6379")
			c1 := c
			defer c1.Close()
			for {
				rep, _ := c1.Do("MONITOR")
				str := fmt.Sprintf("%v", rep)
				fmt.Println(str)
			}
		}

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(sarr[0])) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

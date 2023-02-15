package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 制定目标
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
		Zone: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	// 释放
	defer c.Close()

	// 读数据
	input := bufio.NewReader(os.Stdin)
	s, _ := input.ReadString('\n')
	c.Write([]byte(s))
	var buf [1024]byte
	n, err := c.Read(buf[:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:n]))
}

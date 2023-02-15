package main

import (
	"fmt"
	"net"
)

// udp server
func main() {
	// 监听端口两种方式
	//listen, err := net.Listen("udp", "127.0.0.1:30000")
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
		Zone: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()
	// 接收数据
	for {
		var buf [1024]byte

		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err)
		}

		// 读数据
		fmt.Println("接收到的数据：", string(buf[:n]))

		// 发送数据
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println(err)
		}
	}
}

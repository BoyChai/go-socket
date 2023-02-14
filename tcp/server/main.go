package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp服务端
func main() {
	// 监听杜纳开
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 等待客户端链接
	for {
		// listen.Accept如果没有链接会一直阻塞
		conn, err := listen.Accept() //等待建立链接
		if err != nil {
			log.Fatalln(err)
			continue
		}
		// 启动一个goroutine去处理链接
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 处理完关闭链接
	//针对链接做数据接收发送操作
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
			return
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据是:", recv)
		//conn.Write(buf[:n])
		conn.Write([]byte("ok"))
	}
}

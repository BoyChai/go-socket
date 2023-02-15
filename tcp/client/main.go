package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp客户端
func main() {
	// 和服务端建立链接
	dial, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 最后关闭链接
	defer dial.Close()
	// 接收键盘数据
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		// 发送接收的数据
		dial.Write([]byte(s))
		// 接收服务端返回的数据
		var buf [1024]byte
		n, _ := dial.Read(buf[:])
		fmt.Println(string(buf[:n]))
	}

}

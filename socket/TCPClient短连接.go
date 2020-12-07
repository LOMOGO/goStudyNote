package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal("tcp地址获取失败:", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("获取tcp连接失败:", err)
	}

	result := make([]byte, 256)
	_, err = conn.Read(result)
	if err != nil {
		log.Fatal("从服务端读取数据失败：", err)
	}
	fmt.Println(string(result))
}

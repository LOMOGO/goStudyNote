package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/*
Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,
//UDP缺少了对客户端连接请求的Accept函数。其他基本几乎一模一样，只有TCP换成了UDP而已。UDP的几个主要函数如下所示：

- func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
- func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
- func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
- func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error)
- func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
*/

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient2(conn)
	}
}

func handleClient2(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
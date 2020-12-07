package main

import (
	"fmt"
	"net"
	"os"
)

//一个UDP的客户端代码如下所示,我们可以看到不同的就是TCP换成了UDP而已：

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("anything"))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
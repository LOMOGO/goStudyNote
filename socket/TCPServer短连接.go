package main

import (
	"log"
	"net"
	"strconv"
	"sync"
)

//在Go语言的net包中有一个类型TCPConn， TCPConn可以用在客户端和服务端来读写数据
/*
func (c *TCPConn) Write(b []byte) (int, error)
func (c *TCPConn) Read(b []byte) (int, error)

TCPAddr类型可以表示一个TCP的地址信息，定义如下：
type TCPAddr struct {
	IP IP
	Port int
	Zone string
}
在GO语言中通过ResolveTCPAddr获取一个TCPAddr
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
- net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)。
- addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22"。

TCP client
GO语言中通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCP类型的对象，当连接建立时服务端也创建一个同类型的对象，
此时客户端的服务端通过各自拥有的TCPConn对象来进行数据交换。一般而言，客户端通过TCPConn对象将请求信息发送到服务器端，读取服务器端响应的消息。服务器读取并解析
来自客户端的请求，并返回应答信息，这个连接只有当任一端关闭了连接之后才失效，否则连接就一直在使用。
建立连接的函数定义如下：
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
- network参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
- laddr表示本机地址，一般设置为nil
- raddr表示远程的服务地址
*/

func main() {
	server := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

type ID struct {
	count int
	mu sync.Mutex
}

var id ID

func handleClient(conn net.Conn) {
	defer conn.Close()
	id.mu.Lock()
	id.count++
	id.mu.Unlock()
	conn.Write([]byte("hello user:" + strconv.Itoa(id.count)))
}

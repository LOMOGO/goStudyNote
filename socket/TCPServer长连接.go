package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type ID1 struct {
	count int
	mu sync.Mutex
}

var id1 ID1

func main() {
	server := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		conn, err := listener.Accept()
		checkErr(err)
		go handleClient1(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//conn.SetReadDeadline（）设置了超时，当一定时间内客户端无请求发送，conn便会自动关闭，request在创建时需要指定一个最大长度以防止
//flood attack；每次读取到请求处理完毕后，需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后
func handleClient1(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	request := make([]byte, 128)
	//这个长连接对比刚刚的短连接多了一个东西：在handleClient里面多了一个for循环，for循环需要一些条件才能退出
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if readLen == 0 {
			break
		} else {
			id1.mu.Lock()
			id1.count++
			id1.mu.Unlock()
			conn.Write([]byte("用户:" + strconv.Itoa(id1.count) + "说:" + string(request)))
		}

		request = make([]byte, 128)
	}
}

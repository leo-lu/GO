// test project main.go
package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	"time"
)

var session map[string]*net.TCPConn

func main() {
	session = make(map[string]*net.TCPConn)

	service := "10.0.28.216:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	checkErr(err)
	for {
		conn, err := listener.AcceptTCP()
		if nil != err {
			fmt.Println(time.Now())
			continue
		}

		go handleClient(conn)

		session[conn.RemoteAddr().String()] = conn

		fmt.Println("session count:", len(session))
		for k, v := range session {
			fmt.Println("[session] >>", k, v)
		}
		fmt.Println("==================================")
	}

	os.Exit(0)
}

func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	fmt.Println("handleClient", time.Now().Format("2006-01-02 15:04:05"), conn.RemoteAddr())

	for {
		var buff [1024]byte
		rLen, err := conn.Read(buff[0:])

		if 0 < rLen {
			fmt.Println(string(buff[0:rLen]))
			conn.Write([]byte("hello client!!!~"))
		}
		if !checkErr(err) {
			delete(session, conn.RemoteAddr().String())
			conn.Close()
			fmt.Println("session closed... cur count:", len(session))
			break
		}
	}
}

func checkErr(err error) bool {
	if nil != err {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

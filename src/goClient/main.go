package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"os"
)

func main() {
	service := "10.0.28.216:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	go handleMsg(conn)
	_, err = conn.Write([]byte("hello server!!"))
	checkErr(err)
	select {}
}

func handleMsg(conn *net.TCPConn) {
	defer conn.Close()

	for {
		var buff [1024]byte
		len, err := conn.Read(buff[0:])
		checkErr(err)
		if 0 < len {
			fmt.Println(string(buff[0:len]))
		}
	}
}

func checkErr(err error) {
	if nil != err {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

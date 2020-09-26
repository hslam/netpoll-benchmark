package main

import (
	"flag"
	"net"
	"time"
)

var addr string
var async bool
var sleep int

func init() {
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.BoolVar(&async, "async", false, "-async=false")
	flag.IntVar(&sleep, "sleep", 0, "-sleep=0")
	flag.Parse()
}

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	var buf = make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		if sleep > 0 {
			time.Sleep(time.Millisecond * time.Duration(sleep))
		}
		conn.Write(buf[:n])
	}
	conn.Close()
}

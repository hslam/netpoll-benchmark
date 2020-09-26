package main

import (
	"flag"
	"fmt"
	"github.com/hslam/stats"
	"net"
	"sync/atomic"
)

var addr string
var clients int
var total int
var msg int
var bar bool

func init() {
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.IntVar(&total, "total", 100000, "-total=100000")
	flag.IntVar(&clients, "clients", 1, "-clients=1")
	flag.IntVar(&msg, "msg", 512, "-msg=512")
	flag.BoolVar(&bar, "bar", true, "-bar=true")
	flag.Parse()
	stats.SetBar(bar)
	fmt.Printf("./client -addr=%s -total=%d -clients=%d -msg=%d\n", addr, total, clients, msg)
}

func main() {
	if clients < 1 || total < 1 {
		return
	}

	var wrkClients []stats.Client
	for i := 0; i < clients; i++ {
		if conn, err := net.Dial("tcp", addr); err != nil {
			panic(err)
		} else {
			defer conn.Close()
			client := &WrkClient{conn, make([]byte, msg), make([]byte, msg*2), -1}
			_, _, ok := client.Call()
			if !ok {
				panic("call err")
			}
			wrkClients = append(wrkClients, client)
		}
	}
	stats.StartPrint(1, total, wrkClients)
	for i := 0; i < clients; i++ {
		fmt.Printf("%04d:%d\n", i, wrkClients[i].(*WrkClient).count)
	}
}

type WrkClient struct {
	net.Conn
	msg   []byte
	buf   []byte
	count int64
}

func (c *WrkClient) Call() (int64, int64, bool) {
	atomic.AddInt64(&c.count, 1)
	var err error
	_, err = c.Conn.Write(c.msg)
	if err != nil {
		return int64(len(c.msg)), 0, false
	}
	var n int
	n, err = c.Conn.Read(c.buf)
	if err != nil {
		return int64(len(c.msg)), 0, false
	}
	return int64(len(c.msg)), int64(n), true
}

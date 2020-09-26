package main

import (
	"flag"
	"github.com/panjf2000/gnet"
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
	var events EventServer
	if err := gnet.Serve(&events, "tcp://localhost"+addr, Option); err != nil {
		panic(err.Error())
	}
}
func Option(opts *gnet.Options) {
	opts.NumEventLoop = 20
}

type EventServer struct {
}

func (es *EventServer) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	return
}

func (es *EventServer) OnShutdown(svr gnet.Server) {
}

func (es *EventServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (es *EventServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func (es *EventServer) PreWrite() {
}

func (es *EventServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if async {
		go func(frame []byte, c gnet.Conn) {
			if sleep > 0 {
				time.Sleep(time.Millisecond * time.Duration(sleep))
			}
			c.AsyncWrite(frame)
		}(frame, c)
	} else {
		out = frame
	}
	return
}

func (es *EventServer) Tick() (delay time.Duration, action gnet.Action) {
	return
}

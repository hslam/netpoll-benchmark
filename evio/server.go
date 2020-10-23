package main

import (
	"flag"
	"github.com/tidwall/evio"
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
	var events evio.Events
	events.NumLoops = 20
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		if in != nil {
			if async {
				go func(c evio.Conn, in []byte) {
					if sleep > 0 {
						time.Sleep(time.Millisecond * time.Duration(sleep))
					}
					c.SetContext(in)
					c.Wake()
				}(c, in)
			} else {
				out = in
			}
		} else {
			out = c.Context().([]byte)
		}
		return
	}
	if err := evio.Serve(events, "tcp://localhost"+addr); err != nil {
		panic(err.Error())
	}
}

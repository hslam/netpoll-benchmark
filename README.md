# netpoll-benchmark

#### Comparison to other packages
|Package| [net](https://github.com/golang/go/tree/master/src/net "net")| [netpoll](https://github.com/hslam/netpoll "netpoll")|[gnet](https://github.com/panjf2000/gnet "gnet")|[evio](https://github.com/tidwall/evio "evio")|
|:--:|:--|:--|:--|:--|
|Low memory usage|No|Yes|Yes|Yes|
|Non-blocking I/O|No|Yes|Yes|Yes|
|Splice/sendfile|Yes|Yes|No|No|
|Rescheduling|Yes|Yes|No|No|
|Compatible with the net.Conn interface|Yes|Yes|No|No|

<img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-qps.png" width = "400" height = "300" alt="mock 0ms" align=center> <img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-mock-time-qps.png" width = "400" height = "300" alt="mock 1ms" align=center>
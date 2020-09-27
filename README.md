# netpoll-benchmark

#### Comparison to other packages
|Package| [net](https://github.com/golang/go/tree/master/src/net "net")| [netpoll](https://github.com/hslam/netpol "netpoll")|[gnet](https://github.com/panjf2000/gnet "gnet")|[evio](https://github.com/tidwall/evio "evio")|
|:--:|:--|:--|:--|:--|
|Sync Handler|Yes|Yes|Yes|Yes|
|Async Handler|Yes|Yes|Yes|No|
|Low Memory Usage|No|Yes|Yes|Yes|
|Non-Blocking I/O|No|Yes|Yes|Yes|
|Rescheduling|Yes|Yes|No|No|
|Compatible With The net.Conn|Yes|Yes|No|No|

##### Network QPS

<img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-qps.png"  alt="mock 0ms" align=center>

##### Mocking 1ms QPS

<img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-mock-time-qps.png" alt="mock 1ms" align=center>

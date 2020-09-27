# netpoll-benchmark

## Benchmark
**Network QPS**
<img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-qps.png"  alt="mock 0ms" align=center>

**Mocking 1ms QPS**
<img src="https://raw.githubusercontent.com/hslam/netpoll-benchmark/master/netpoll-mock-time-qps.png" alt="mock 1ms" align=center>

net
```
./client -addr=:9999 -total=1000000 -clients=1 -msg=512
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	57.30s
	Requests per second:	17452.50
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	4.21ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.05ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.07ms
	99%	time for request:	0.09ms
	99.9%	time for request:	0.17ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	8935681.30 Byte/s (8.94 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	8935681.30 Byte/s (8.94 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

netpoll
```
./client -addr=:9999 -total=1000000 -clients=1 -msg=512
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	45.20s
	Requests per second:	22125.73
	Fastest time for request:	0.03ms
	Average time per request:	0.04ms
	Slowest time for request:	2.75ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.03ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.04ms
	90%	time for request:	0.05ms
	95%	time for request:	0.06ms
	99%	time for request:	0.08ms
	99.9%	time for request:	0.16ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	11328375.24 Byte/s (11.33 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	11328375.24 Byte/s (11.33 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

gnet
```
./client -addr=:9999 -total=1000000 -clients=1 -msg=512
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	43.92s
	Requests per second:	22767.99
	Fastest time for request:	0.03ms
	Average time per request:	0.04ms
	Slowest time for request:	1.45ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.03ms
	5%	time for request:	0.03ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.04ms
	90%	time for request:	0.05ms
	95%	time for request:	0.05ms
	99%	time for request:	0.07ms
	99.9%	time for request:	0.16ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	11657211.76 Byte/s (11.66 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	11657211.76 Byte/s (11.66 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

evio
```
./client -addr=:9999 -total=1000000 -clients=1 -msg=512
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	46.83s
	Requests per second:	21353.80
	Fastest time for request:	0.03ms
	Average time per request:	0.04ms
	Slowest time for request:	3.83ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.04ms
	90%	time for request:	0.05ms
	95%	time for request:	0.06ms
	99%	time for request:	0.08ms
	99.9%	time for request:	0.17ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	10933147.79 Byte/s (10.93 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	10933147.79 Byte/s (10.93 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
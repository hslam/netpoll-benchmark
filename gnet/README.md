# gnet

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
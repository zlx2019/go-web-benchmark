# Go Web 性能压测

## fiber
- 第一次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.90ms   14.41ms 277.12ms   96.13%
    Req/Sec    23.01k     9.34k  145.73k    85.85%
  2264317 requests in 10.09s, 336.87MB read
Requests/sec: 224424.69
Transfer/sec:     33.39MB
```
- 第二次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.66ms   22.25ms 263.36ms   92.98%
    Req/Sec    21.41k    16.01k  140.76k    80.24%
  2063332 requests in 10.08s, 306.97MB read
Requests/sec: 204675.15
Transfer/sec:     30.45MB
```

- 第三次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.78ms   16.79ms 217.82ms   91.36%
    Req/Sec    20.36k    17.73k  142.84k    83.02%
  1971950 requests in 10.08s, 293.37MB read
Requests/sec: 195618.86
Transfer/sec:     29.10MB
```

# hertz
- 第一次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   532.92us    1.23ms  43.41ms   98.50%
    Req/Sec    21.52k     4.26k   38.78k    74.10%
  2143861 requests in 10.05s, 318.95MB read
Requests/sec: 213342.83
Transfer/sec:     31.74MB
```
- 第二次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.77ms   17.02ms 196.67ms   91.54%
    Req/Sec    21.58k    16.11k  146.93k    78.82%
  2096808 requests in 10.06s, 311.95MB read
Requests/sec: 208345.60
Transfer/sec:     31.00MB
```
- 第三次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.72ms   11.23ms 181.67ms   94.65%
    Req/Sec    23.87k    10.98k  166.30k    85.03%
  2356363 requests in 10.09s, 350.56MB read
Requests/sec: 233506.72
Transfer/sec:     34.74MB
```

## gin
- 第一次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.18ms    2.11ms  36.62ms   90.61%
    Req/Sec    17.76k     1.90k   22.00k    67.59%
  1782706 requests in 10.10s, 236.32MB read
Requests/sec: 176494.40
Transfer/sec:     23.40MB
```

- 第二次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.06ms    1.84ms  33.38ms   91.12%
    Req/Sec    18.04k     2.41k   38.47k    80.10%
  1804372 requests in 10.10s, 239.19MB read
Requests/sec: 178650.87
Transfer/sec:     23.68MB
```

- 第三次
```shell
Running 10s test @ http://127.0.0.1:8080/ping
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.63ms    3.05ms  47.14ms   90.51%
    Req/Sec    14.47k     3.24k   21.45k    61.31%
  1451917 requests in 10.10s, 192.47MB read
Requests/sec: 143716.49
Transfer/sec:     19.05MB
```
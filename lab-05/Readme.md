# Бенчмарк

```bash
ab -n 10000 -c 100 http://localhost:80/api/v1/accounts
```

## Без балансировкой

```bash
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/accounts
Document Length:        52 bytes

Concurrency Level:      100
Time taken for tests:   8.516 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      4870000 bytes
HTML transferred:       520000 bytes
Requests per second:    1174.23 [#/sec] (mean)
Time per request:       85.162 [ms] (mean)
Time per request:       0.852 [ms] (mean, across all concurrent requests)
Transfer rate:          558.45 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       8
Processing:     1   84 186.5      2    1216
Waiting:        0   83 186.5      2    1216
Total:          1   84 186.6      2    1216

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      3
  75%      6
  80%     16
  90%    419
  95%    523
  98%    649
  99%    756
 100%   1216 (longest request)
```

## С балансировкой

```bash
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/accounts
Document Length:        52 bytes

Concurrency Level:      100
Time taken for tests:   5.628 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      4870000 bytes
HTML transferred:       520000 bytes
Requests per second:    1776.68 [#/sec] (mean)
Time per request:       56.285 [ms] (mean)
Time per request:       0.563 [ms] (mean, across all concurrent requests)
Transfer rate:          844.96 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0      21
Processing:     1   55 145.5      3     845
Waiting:        0   55 145.5      3     845
Total:          1   55 145.6      3     845

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      5
  75%      7
  80%     10
  90%    321
  95%    473
  98%    546
  99%    585
 100%    845 (longest request)
```


```bash
ab -n 10000 -c 150 http://localhost:80/api/v1/accounts
```

## Без балансировкой

```bash
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/accounts
Document Length:        52 bytes

Concurrency Level:      150
Time taken for tests:   9.559 seconds
Complete requests:      10000
Failed requests:        10
   (Connect: 0, Receive: 0, Length: 10, Exceptions: 0)
Non-2xx responses:      10
Total transferred:      4870150 bytes
HTML transferred:       520080 bytes
Requests per second:    1046.11 [#/sec] (mean)
Time per request:       143.388 [ms] (mean)
Time per request:       0.956 [ms] (mean, across all concurrent requests)
Transfer rate:          497.53 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.7      0      11
Processing:     1  140 291.1      2    1694
Waiting:        0  140 291.1      2    1693
Total:          1  140 291.2      3    1694

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      7
  80%     56
  90%    701
  95%    784
  98%    906
  99%   1016
 100%   1694 (longest request)
```

## С балансировкой

```bash
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/accounts
Document Length:        52 bytes

Concurrency Level:      150
Time taken for tests:   6.948 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      4870000 bytes
HTML transferred:       520000 bytes
Requests per second:    1439.20 [#/sec] (mean)
Time per request:       104.225 [ms] (mean)
Time per request:       0.695 [ms] (mean, across all concurrent requests)
Transfer rate:          684.46 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.7      0      22
Processing:     1  101 262.6      3    1317
Waiting:        0  100 262.6      2    1314
Total:          1  101 262.7      3    1317

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      5
  80%      7
  90%    685
  95%    829
  98%    897
  99%    952
 100%   1317 (longest request)
```



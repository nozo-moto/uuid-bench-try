# Uuid-Bench-Try

check it

https://www.reddit.com/r/golang/comments/q9dewu/faster_implementation_for_uuidnewstring/

```
go test -bench . -benchmem
```

result is

```
goos: darwin
goarch: amd64
pkg: github.com/nozo-moto/uuid-bench-try
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkUseRandPool-12       	 7165122	       161.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkNotUseRandPool-12    	 1488103	       809.7 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/nozo-moto/uuid-bench-try	3.457s
```

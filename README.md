# Test 1

Set up an of size 1000x1000x1000 in stack. Iterate to create a product of 1...100 and fill each cell.

## Java21 vs go1.21.3 darwin/amd64 vs rustc 1.74.0 

```bash
javac JavaPerfTest1.java ; java JavaPerfTest1
go build -ldflags "-s -w" GoPerfTest1.go  ; ./GoPerfTest1
cd rustperf ; cargo build --release ; ./target/release/rustperf ; cd ..
```

```bash
Elapsed Time: 46210 milliseconds
Elapsed Time: 58482 milliseconds
Elapsed Time: 3740 milliseconds
```

```bash
Java
CPU Usage: 2.73486328125%
Memory Pool: CodeHeap 'non-nmethods'
  Committed Memory: 2 MB
  Used Memory: 1 MB
  Max Memory: 7 MB
  Usage Threshold: 0 MB

Go
CPU Usage: 6.20%
Heap Memory Usage:
  Alloc: 7629 MB
  TotalAlloc: 7629 MB
  Sys: 7769 MB
  NumGC: 1
  
Rust
CPU Usage: 16%
31744 bytes allocated/4804608 bytes resident
```


# Test 1

Set up an of size 1000x1000x1000 in stack. Iterate to create a product of 1...100 for each of the cells and fill it. 

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


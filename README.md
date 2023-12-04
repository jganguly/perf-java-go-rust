# Test 1

Set up an of size 1000x1000x1000x10 in stack. Iterate over it and calculate `Pi` using the Leibniz formula.

## Java21 vs go1.21.3 darwin/amd64 vs rustc 1.74.0 

```bash
javac JavaPerfTest1.java ; java JavaPerfTest1
go build -ldflags "-s -w" GoPerfTest1.go  ; ./GoPerfTest1
cd rustperf ; cargo build --release ; ./target/release/rustperf ; cd ..
```

```bash
Elapsed Time: 7051 milliseconds
Elapsed Times: 7330 milliseconds
Elapsed Time: 6692 milliseconds
```


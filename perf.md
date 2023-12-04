# Test 1

Set up an of size 1000_0000_000 in stack. Iterate over it and calculate `Pi` using the Leibniz formula.

## Output: Java21

```bash
 javac JavaPerfTest1.java
 java JavaPerfTest1
```

```bash
Array Size: 1000000000
Elapsed Time: 7283 milliseconds
```

## go1.21.3 darwin/amd64

```bash
go build -ldflags "-s -w" GoPerfTest1.go 
./GoPerfTest1
```

```bash
Array Size: 1000000000
Elapsed Times: 3387 milliseconds
```

## rustc 1.74.0 (79e9716c9 2023-11-13)`

```bash
cargo build --release
./target/release/rustperf
Array Size: 1000000000
Elapsed Time: 3906 milliseconds
```

**Java : Go : Rust**
1 : 0.47 : 0.54
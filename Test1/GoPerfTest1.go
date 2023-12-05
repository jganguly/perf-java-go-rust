package main

import (
	"fmt"
	// "runtime"
	"time"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	// "github.com/shirou/gopsutil/mem"	
)

func main() {
	startTime := time.Now()
	runYourCode()
	elapsedTime := time.Since(startTime).Milliseconds()

	msg := fmt.Sprintf("Elapsed Time: %d milliseconds", elapsedTime)
	fmt.Println(msg)
	perfEx()
}

func runYourCode() {
	// Allocate a large slice to consume memory
	n := 1000 * 1000 * 1000
	largeSlice := make([]int, n)

	for i := 0; i < n; i++ {
		largeSlice[i] = calculateProduct()
	}
}

func calculateProduct() int {
	product := 1
	for i := 1; i <= 100; i++ {
		product *= i
	}
	return product
}


func perfEx() {
	cpuUsage, err := getCPUUsage()
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	}

	memoryUsage()

	// Get and print memory usage
	// memoryUsage, err := getMemoryUsage()
	// if err != nil {
	// 	fmt.Println("Error getting memory usage:", err)
	// } else {
	// 	fmt.Printf("Memory Usage: Used %v MB, Total %v MB\n", memoryUsage.Used/1024/1024, memoryUsage.Total/1024/1024)
	// }
}

func getCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}


func memoryUsage() {
	// Allocate memory
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 1024)
	}

	// Collect and print memory statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Heap Memory Usage:\n")
	fmt.Printf("  Alloc: %v MB\n", toMB(memStats.Alloc))
	fmt.Printf("  TotalAlloc: %v MB\n", toMB(memStats.TotalAlloc))
	fmt.Printf("  Sys: %v MB\n", toMB(memStats.Sys))
	fmt.Printf("  NumGC: %v\n", memStats.NumGC)
}

func toMB(bytes uint64) uint64 {
	return bytes / 1024 / 1024
}


// func getMemoryUsage() (*mem.VirtualMemoryStat, error) {
// 	memory, err := mem.VirtualMemory()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return memory, nil
// }



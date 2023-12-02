package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	elapsedTime()

	// Print CPU Usage
	cpuUsage, err := getCPUUsage()
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	}

	// Print Memory Usage
	memoryUsage := getMemoryUsage()
	fmt.Println("Memory Usage:", formatMemoryUsage(memoryUsage))
}

func getCPUUsage() (float64, error) {
	// Unfortunately, Go does not have a standard library to directly retrieve CPU usage.
	// External libraries or system-specific calls may be needed for this purpose.
	return -1, fmt.Errorf("CPU usage not supported in Go without external libraries")
}

func getMemoryUsage() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func formatMemoryUsage(memStats runtime.MemStats) string {
	max := memStats.HeapSys / (1024 * 1024)
	used := memStats.HeapInuse / (1024 * 1024)
	committed := memStats.HeapAlloc / (1024 * 1024)

	return fmt.Sprintf("Used: %dMB, Committed: %dMB, Max: %dMB", used, committed, max)
}

func elapsedTime() {
	startTime := time.Now()

	// Allocate memory
	arraySize := 1000000
	memoryArray := make([]int, arraySize)

    for i:= 1; i < arraySize; i++ {
    	memoryArray[i] = i;
    }


	elapsedTime := time.Since(startTime)

	fmt.Println("Memory Test Array Size:", arraySize)
	fmt.Println("Memory Test Time:", elapsedTime)
}

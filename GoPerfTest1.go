package main

import (
	"fmt"
	// "runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"	
)

func main() {
	startTime := time.Now()
	runYourCode()
	elapsedTime := time.Since(startTime).Milliseconds()

	msg := fmt.Sprintf("Elapsed Times: %d milliseconds", elapsedTime)
	fmt.Println(msg)
}

func runYourCode() {
	// Allocate a large array to consume memory
	n := 1000 * 1000 * 1000
	fmt.Println("Array Size:", n)
	largeArray := make([]int, n)
	for i := 0; i < n; i++ {
		largeArray[i] = i
	}
	// CPU-intensive operation - calculating Pi using the Leibniz formula
	pi := 0.0

	for i := 0; i < n; i++ {
		pi += float64((i % 2 * 2 - 1)) * 4.0 / float64(2*i+1)
	}
}


func perfEx() {
	cpuUsage, err := getCPUUsage()
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	}

	// Get and print memory usage
	memoryUsage, err := getMemoryUsage()
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
	} else {
		fmt.Printf("Memory Usage: Used %v MB, Total %v MB\n", memoryUsage.Used/1024/1024, memoryUsage.Total/1024/1024)
	}
}

func getCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func getMemoryUsage() (*mem.VirtualMemoryStat, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return memory, nil
}



package main

import (
	"fmt"
	// "runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"	
)

func main() {
	elapsedTime()
	// perfEx()
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

func elapsedTime() {
	startTime := time.Now()

	// Allocate memory
	arraySize := 1000*1000*1000
	memoryArray := make([]int, arraySize)

    for i:= 1; i < arraySize; i++ {
    	memoryArray[i] = i;
    }


	elapsedTime := time.Since(startTime).Milliseconds()

	fmt.Println("Array Size:", arraySize)
	fmt.Println("Elapsed Time:", elapsedTime)
}

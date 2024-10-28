package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Запуск горутин для потребления памяти
	for i := 0; i < 10; i++ {
		go consumeMemory()
	}

	// Мониторинг использования памяти
	var memStats runtime.MemStats
	for {
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Allocated memory: %d bytes\n", memStats.Alloc)
		time.Sleep(2 * time.Second)
	}
}

func consumeMemory() {
	data := make([]byte, 100*1024*1024) // 100 MB
	for i := range data {
		data[i] = byte(i)
	}
	time.Sleep(5 * time.Second)
}

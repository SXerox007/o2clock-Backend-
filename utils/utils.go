package utils

import (
	"fmt"
	"runtime"
	"time"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(time.Now())
	fmt.Println("Alloc = ", bToMb(m.Alloc))
	fmt.Println("TotalAlloc = ", bToMb(m.TotalAlloc))
	fmt.Println("Sys =  ", bToMb(m.Sys))
	fmt.Println("NumGC = ", m.NumGC)
	fmt.Println("Lookups = ", bToMb(m.Lookups))
	fmt.Println("Mallocs = ", bToMb(m.Mallocs))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Job memory usage
func CreateJobMemUsage() {
	jt := NewJobMemUsage()
	for {
		<-jt.t.C
		PrintMemUsage()
		jt.updateJobMemCheck()
	}
}

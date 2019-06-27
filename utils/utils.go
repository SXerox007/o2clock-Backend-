package utils

import (
	"fmt"
	"o2clock/constants/appconstant"
	"runtime"
	"time"

	"google.golang.org/grpc"
)

func ValidateAndPrintMemUsage(srv *grpc.Server) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(time.Now())
	fmt.Println(appconstant.MEM_ALLOC, bToMb(m.Alloc))
	fmt.Println(appconstant.MEM_TOTAL_ALLOC, bToMb(m.TotalAlloc))
	fmt.Println(appconstant.MEM_SYS, bToMb(m.Sys))
	fmt.Println(appconstant.NUM_GC, m.NumGC)
	fmt.Println(appconstant.LOOKUPS, bToMb(m.Lookups))
	fmt.Println(appconstant.MALLOCS, bToMb(m.Mallocs))
	fmt.Println()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Job memory usage
func CreateJobMemUsage(srv *grpc.Server) {
	jt := NewJobMemUsage()
	for {
		<-jt.t.C
		ValidateAndPrintMemUsage(srv)
		jt.updateJobMemCheck()
	}
}

// current srv status
func CurrentMemStatus() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return appconstant.MEM_ALLOC + fmt.Sprint(bToMb(m.Alloc)) +
		"\n" + appconstant.MEM_TOTAL_ALLOC + fmt.Sprint(bToMb(m.TotalAlloc)) +
		"\n" + appconstant.MEM_SYS + fmt.Sprint(bToMb(m.Sys))
}

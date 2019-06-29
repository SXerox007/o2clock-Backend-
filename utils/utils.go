package utils

import (
	"crypto/rand"
	"fmt"
	"io"
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

// generate random string
func GenerateRandomString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// generate random string with type
func RandomStringGenerateWithType(size int, randType string) string {
	var dictionary string
	switch randType {
	case appconstant.ALPHA_NUM:
		dictionary = appconstant.DIC_ALPHA_NUM
		break
	case appconstant.ALPHA:
		dictionary = appconstant.DIC_ALPHA
		break
	case appconstant.NUM:
		dictionary = appconstant.DIC_NUM
		break
	default:
		return ""
	}
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

//

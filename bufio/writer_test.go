package bufio

import (
	"os"
	"testing"
	"time"
)

func Benchmark_Sync_Writer(b *testing.B) {
	fd, _ := os.Create("./test.log")
	bw := NewBufioWriter(4096, time.Millisecond*200)
	bw.SetOutputFile(fd)

	for n := 0; n < b.N; n++ {
		bw.WriteString("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111\n")
	}

	//time.Sleep(time.Second)
	bw.Close()
	//time.Sleep(time.Second)
}

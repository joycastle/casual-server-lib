package bufio

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type BufioWriter struct {
	writer    *bufio.Writer
	size      int
	fd        *os.File
	mu        *sync.Mutex
	fcyc      time.Duration
	closeFlag bool
}

func NewBufioWriterDefault() *BufioWriter {
	return NewBufioWriter(100, time.Millisecond*100)
}

func NewBufioWriter(size int, fcyc time.Duration) *BufioWriter {
	bw := &BufioWriter{
		size:      size,
		mu:        new(sync.Mutex),
		fcyc:      fcyc,
		closeFlag: false,
	}
	go bw.syncFlush()
	return bw
}

func (bw *BufioWriter) SetOutputFile(fd *os.File) {
	bw.mu.Lock()
	defer bw.mu.Unlock()
	bw.fd = fd
	bw.writer = bufio.NewWriterSize(fd, bw.size)
}

func (bw *BufioWriter) Close() {
	bw.mu.Lock()
	defer bw.mu.Unlock()
	bw.writer.Flush()
	bw.closeFlag = true
}

func (bw *BufioWriter) WriteString(msg string) {
	bw.mu.Lock()
	defer bw.mu.Unlock()
	bw.writer.WriteString(msg)
}

func (bw *BufioWriter) syncFlush() {
	<-time.After(bw.fcyc)
	if bw.closeFlag {
		return
	}

	if bw.writer != nil {
		bw.mu.Lock()
		bw.writer.Flush()
		bw.mu.Unlock()
	}

	bw.syncFlush()
}

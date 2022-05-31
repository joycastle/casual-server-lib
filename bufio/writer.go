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
	return NewBufioWriter(4096, time.Millisecond*200)
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
	lastFd := bw.fd
	bw.fd = fd
	bw.writer = bufio.NewWriterSize(fd, bw.size)
	lastFd.Close()
}

func (bw *BufioWriter) Close() {
	bw.mu.Lock()
	defer bw.mu.Unlock()
	if bw.writer != nil {
		bw.writer.Flush()
	}
	if bw.fd != nil {
		bw.fd.Close()
	}
	bw.closeFlag = true
}

func (bw *BufioWriter) WriteString(msg string) {
	bw.mu.Lock()
	defer bw.mu.Unlock()
	if bw.writer != nil {
		bw.writer.WriteString(msg)
	}
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

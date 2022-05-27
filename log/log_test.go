package log

import (
	"errors"
	"testing"
)

func Test_log(t *testing.T) {
	// t.Fatal("not implemented")
	log := NewLogger(LogConf{"./testlogs/test.log-*-*-*", "ALL", 1})
	log.Infof("This is Logger:%s", "hello")
	log.Info("This is Logger", errors.New("Hello"))
	log.Debugf("This is Logger:%s", "hello")
	log.Debug("This is Logger", errors.New("Hello"))

	log.DisableColor()

	log.Warnf("This is Logger:%s", "hello")
	log.Warn("This is Logger", errors.New("Hello"))
	log.Fatalf("This is Logger:%s", "hello")
	log.Fatal("This is Logger", errors.New("Hello"))
}

func Benchmark_log(b *testing.B) {
	log := NewLogger(LogConf{"./testlogs/test.log", "ALL", 1})
	for n := 0; n < b.N; n++ {
		log.Fatalf("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	}
}

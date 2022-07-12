package faketime

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFakeTimeNow(t *testing.T) {
	time.Local = time.UTC

	fk := &FakeTime{mu: new(sync.RWMutex)}
	fk.initialNano()
	if err := fk.SetTargetTimeFormatV1("2022-01-18 00:00:01"); err != nil {
		t.Fatal(err)
	}

	targets := []int64{1642464001000011000, 1642464001000011000, 1642464001999999999}
	for _, nano := range targets {
		if nano != time.Unix(nano/1e9, nano-(nano/1e9)*1e9).UnixNano() {
			t.Fatal("1e9 count is fatal")
		}
		nano++
		if nano != time.Unix(nano/1e9, nano-(nano/1e9)*1e9).UnixNano() {
			t.Fatal("1e9 count is fatal")
		}
	}

	for m := 1; m <= 12; m++ {
		for d := 1; d <= 28; d++ {
			for h := 0; h <= 23; h++ {
				for i := 1; i <= 59; i++ {
					timeDesc := fmt.Sprintf("2022-%.2d-%.2d %.2d:%.2d:%.2d", m, d, h, i, i)
					if err := fk.SetTargetTimeFormatV1(timeDesc); err != nil {
						t.Fatal(err)
					}
					fkTimeDesc := fk.Now().Format("2006-01-02 15:04:05")
					if fkTimeDesc != timeDesc {
						t.Fatal(timeDesc, "!=", fkTimeDesc)
					}
				}
			}
		}
	}
}

func TestSince(t *testing.T) {
	SetTargetTimeFormatV1("2022-06-01 12:31:56")
	a, _ := time.Parse("2006-01-02 15:04:05", "2022-06-01 12:31:56")
	offset := int64(0)
	for i := 0; i < 5; i++ {
		b := Since(a).Nanoseconds() / 1000000
		if i != 0 {
			if b-offset < 1000 || b-offset > 1100 {
				t.Fatal(b - offset)
			}
		}
		offset = b
		time.Sleep(time.Second)
	}
}

func TestDiffWithActicalNow(t *testing.T) {
	fk := NewFakeTime()
	for i := 0; i < 100; i++ {
		fktime := fk.Now()
		actTime := time.Now()

		fktSec := fktime.Unix()
		actSec := actTime.Unix()
		fktNano := fktime.UnixNano()
		actNano := actTime.UnixNano()
		if fktSec-actSec != 0 {
			t.Fatal(fktSec, actSec, fktSec-actSec, fktNano, actNano, fktNano-actNano, i)
		}
		time.Sleep(1000)
	}
}

func TestScan(t *testing.T) {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = time.UTC
	timeDesc := "2022-07-01 12:31:56"
	a, _ := time.Parse("2006-01-02 15:04:05", timeDesc)
	stamp := a.Unix()
	SetTargetTimeFormatV1(timeDesc)
	for i := 0; i < 10; i++ {
		now := Now().Unix()
		if now != stamp {
			t.Fatal(now, stamp, now-stamp)
		}
		stamp++
		time.Sleep(time.Second)
	}

	//DebubForHttp(1122)
	//time.Sleep(time.Second * 1000)
}

func BenchmarkNow5000Disable(b *testing.B) {
	Disable()
	b.ReportAllocs()
	b.ResetTimer()
	// 设置并发数
	b.SetParallelism(5000)
	// 测试多线程并发模式
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Now()
		}
	})
}

func BenchmarkNow5000Enable(b *testing.B) {
	Enable()
	b.ReportAllocs()
	b.ResetTimer()
	// 设置并发数
	b.SetParallelism(5000)
	// 测试多线程并发模式
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Now()
		}
	})
}

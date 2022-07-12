package faketime

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var faketime *FakeTime = NewFakeTime().Disable()

func Now() time.Time {
	return faketime.Now()
}

func Since(t time.Time) time.Duration {
	return faketime.Since(t)
}

func SetTargetTimeFormatV1(timeDesc string) error {
	return faketime.SetTargetTimeFormatV1(timeDesc)
}

func SetTargetTimeFormatV2(timeDesc string) error {
	return faketime.SetTargetTimeFormatV2(timeDesc)
}

func Enable() {
	faketime.Enable()
}

func Disable() {
	faketime.Disable()
}

type FakeTime struct {
	startNano  int64
	defineNano int64
	enable     bool
	mu         *sync.RWMutex
}

func (fk *FakeTime) initialNano() *FakeTime {
	fk.mu.Lock()
	defer fk.mu.Unlock()
	nano := time.Now().UnixNano()
	fk.startNano = nano
	fk.defineNano = nano
	return fk
}

func NewFakeTime() *FakeTime {
	fk := &FakeTime{mu: new(sync.RWMutex)}
	fk.initialNano()
	return fk
}

func (fk *FakeTime) SetTargetTimeFormatV1(timeDesc string) error {
	return fk.setTargetTime("2006-01-02 15:04:05", timeDesc)
}

func (fk *FakeTime) SetTargetTimeFormatV2(timeDesc string) error {
	return fk.setTargetTime("20060102150405", timeDesc)
}

func (fk *FakeTime) setTargetTime(format, timeDesc string) error {
	fk.mu.Lock()
	defer fk.mu.Unlock()

	formatTime, err := time.Parse(format, timeDesc)
	if err != nil {
		return err
	}

	fk.defineNano = formatTime.UnixNano()
	fk.startNano = time.Now().UnixNano()
	fk.enable = true

	return nil
}

func (fk *FakeTime) Enable() *FakeTime {
	fk.mu.Lock()
	defer fk.mu.Unlock()
	fk.enable = true
	return fk
}

func (fk *FakeTime) Disable() *FakeTime {
	fk.mu.Lock()
	defer fk.mu.Unlock()
	fk.enable = false
	return fk
}

func (fk *FakeTime) Now() time.Time {
	return fk.now()
}

func (fk *FakeTime) now() time.Time {
	fk.mu.RLock()
	defer fk.mu.RUnlock()

	if !fk.enable {
		return time.Now()
	}

	offset := time.Now().UnixNano() - fk.startNano
	define := fk.defineNano + offset
	return time.Unix(define/1e9, define-(define/1e9)*1e9)
}

func (fk *FakeTime) Since(t time.Time) time.Duration {
	return fk.now().Sub(t)
}

func DebubForHttp(port int) {

	nowt := "faketime: current server time: %s\n\n"
	usage := "faketime: using http://{ip:port}/faketime?v=20220701235959 to set server time as date: 2022-07-01 23:59:59"
	errset := "faketime: set error: %s"

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf(nowt, Now().Format("2006-01-02 15:04:05"))+usage)
	})

	r.HandleFunc("/faketime", func(w http.ResponseWriter, r *http.Request) {
		fakeTime := r.URL.Query().Get("v")
		if err := SetTargetTimeFormatV2(fakeTime); err != nil {
			fmt.Fprintf(w, fmt.Sprintf(errset, err.Error()))
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	go srv.ListenAndServe()
}

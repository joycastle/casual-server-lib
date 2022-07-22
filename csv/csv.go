package csv

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/joycastle/casual-server-lib/log"
	"github.com/joycastle/casual-server-lib/util"
)

type Header struct {
	Name string
	Type string
	Desc string
}

type Csv struct {
	folder   string
	files    []string
	fmd5s    map[string]string
	headers  map[string]map[int]Header
	contents map[string][][]string
	mu       *sync.RWMutex
}

func NewCsvReader(folder string) *Csv {
	return &Csv{
		folder:   folder,
		fmd5s:    make(map[string]string),
		headers:  make(map[string]map[int]Header),
		contents: make(map[string][][]string),
		mu:       new(sync.RWMutex),
	}
}
func (c *Csv) Use(files ...string) *Csv {
	c.files = append(c.files, files...)
	return c
}

func (c *Csv) read(f string) error {
	fpath := filepath.Join(c.folder, f)
	fnode := strings.TrimRight(f, filepath.Ext(f))

	if md5v, err := util.Md5FileLower32(fpath); err != nil {
		return err
	} else {
		c.mu.Lock()
		if old, ok := c.fmd5s[fnode]; ok && old == md5v {
			c.mu.Unlock()
			return fmt.Errorf("md5 %s: same at %s", fpath, md5v)
		}
		c.fmd5s[fnode] = md5v
		c.mu.Unlock()
	}

	fd, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer fd.Close()

	contents, err := csv.NewReader(fd).ReadAll()
	if err != nil {
		return err
	}

	if len(contents) <= 3 {
		return fmt.Errorf("read %s: at least three lines", fpath)
	}

	out := make(map[int]Header)
	rowsLen := len(contents[0])
	for i := 0; i < rowsLen; i++ {
		h := Header{}
		h.Name = strings.Trim(contents[0][i], " ")
		h.Type = strings.Trim(contents[1][i], " ")
		h.Desc = strings.Trim(contents[2][i], " ")
		out[i] = h
	}

	c.mu.Lock()
	c.headers[f] = out
	c.contents[fnode] = contents[3:]
	c.mu.Unlock()

	return nil
}

func (c *Csv) Load() error {
	for _, f := range c.files {
		if err := c.read(f); err != nil {
			return err
		}
	}
	return nil
}

func (c *Csv) Loading(ctx context.Context, sec int, report bool) {
	go func() {
		t := time.NewTicker(time.Duration(sec) * time.Second)
		for {
			select {
			case <-t.C:
				if err := c.Load(); err != nil {
					log.Get("error").Warn("csv: loading:", err)
				} else {
					log.Get("run").Info("csv: loading: Ok")
				}

				if report {
					if s, err := c.PrintAll(); err != nil {
						log.Get("error").Fatal("csv: loading print:", err)
					} else {
						log.Get("run").Info("csv: loading print: ", s)
					}
				}

			case <-ctx.Done():
				log.Get("run").Info("csv: loading: Exit")
				t.Stop()
				return
			}
		}
	}()
}

func (c *Csv) PrintAll() (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	b, err := json.Marshal(c.contents)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

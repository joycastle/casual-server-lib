package csv

import (
	"strconv"
	"strings"
)

const (
	FormatType_INT = iota
	FormatType_STRING
	FormatType_INT_ARRAY
	FormatType_FLOAT
)

type Format struct {
	v          string
	vt         int
	vint       int
	vfloat64   float64
	vintslice  []int
	vintslice2 [][]int
}

func NewFormat(t int, s string) *Format {
	return &Format{v: strings.Trim(s, " "), vt: t}
}

func (f *Format) ToInt() int {
	return f.vint
}

func (f *Format) ToString() string {
	return f.v
}

func (f *Format) ToIntSlice() []int {
	return f.vintslice
}

func (f *Format) ToFloat64() float64 {
	return f.vfloat64
}

func (f *Format) ToIntSlice2() [][]int {
	return f.vintslice2
}

func (f *Format) Parse() error {
	switch f.vt {
	case FormatType_INT:
		if err := f.parseInt(); err != nil {
			return err
		}
	}
	return nil
}

func (f *Format) parseInt() error {
	d, err := strconv.Atoi(f.v)
	if err != nil {
		return err
	}
	f.vint = d
	return nil
}

func (f *Format) parseString() error {
	return nil
}

func (f *Format) parseIntSlice() error {
	list := strings.Split(f.v, "|")
	out := []int{}
	for _, v := range list {
		d, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		out = append(out, d)
	}
	f.vintslice = out
	return nil
}

func (f *Format) parseFloat64() error {
	d, err := strconv.ParseFloat(f.v, 32)
	if err != nil {
		return err
	}
	f.vfloat64 = d
	return nil
}

func (f *Format) parseIntSlice2() error {
	s := strings.Replace(f.v, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)
	list := strings.Split(s, ",")
	out := [][]int{}
	for _, v := range list {
		arr := strings.Split(v, "|")
		tmp := []int{}
		for _, vv := range arr {
			d, err := strconv.Atoi(vv)
			if err != nil {
				return err
			}
			tmp = append(tmp, d)
		}
		if len(tmp) > 0 {
			out = append(out, tmp)
		}
	}
	f.vintslice2 = out
	return nil
}

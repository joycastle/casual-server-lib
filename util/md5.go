package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func Md5StringUpper32(s string) string {
	return strings.ToUpper(Md5StringLower32(s))
}

func Md5StringLower32(s string) string {
	e := md5.New()
	io.WriteString(e, s)
	return fmt.Sprintf("%x", e.Sum(nil))
}

func Md5FileUpper32(fileName string) (string, error) {
	v, err := Md5FileLower32(fileName)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(v), nil
}

func Md5FileLower32(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return Md5StringLower32(string(b)), nil
}

package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMd5String(t *testing.T) {
	if Md5StringLower32("jskldjfaskdfjsljdfdsjkfjdsjfskld") != "4f356d2b696709c69ed60ff6cc08fde8" {
		t.Fatal()
	}

	if Md5StringLower32("你好朋友！#%SDFSD") != "eb1b8d93456dd227a7f3e45070bc6666" {
		t.Fatal()
	}

	if err := ioutil.WriteFile("./test_md5.txt", []byte("jskldjfaskdfjsljdfdsjkfjdsjfskld"), 0666); err != nil {
		t.Fatal(err)
	}

	if md5, err := Md5FileLower32("./test_md5.txt"); err != nil || md5 != "4f356d2b696709c69ed60ff6cc08fde8" {
		t.Fatal()
	}

	if err := os.Remove("./test_md5.txt"); err != nil {
		t.Fatal(err)
	}
}

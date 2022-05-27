package config

import (
	"testing"
)

func Test_config(t *testing.T) {

	if err := InitConfig("./dev.yaml"); err != nil {
		t.Fatal(err)
	}

	if Logs["main"].Output != "./main.log-*-*-*" || Logs["main"].Level != "INFO" {
		t.Fatal("parse error")
	}
}

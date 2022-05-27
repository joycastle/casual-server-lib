package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type ParseFunc func(v *viper.Viper) error

var registerParseFunc []ParseFunc

func RegisterParser(f ParseFunc) {
	registerParseFunc = append(registerParseFunc, f)
}

func InitConfig(fileName string) error {
	ext := filepath.Ext(fileName)
	ext = strings.TrimLeft(ext, ".")

	isSupport := false
	for _, v := range viper.SupportedExts {
		if v == ext {
			isSupport = true
			break
		}
	}
	if !isSupport {
		return fmt.Errorf("viper not support the format \".%s\"", ext)
	}

	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType(ext)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	for _, parseHandler := range registerParseFunc {
		if err := parseHandler(v); err != nil {
			return err
		}
	}

	return nil
}

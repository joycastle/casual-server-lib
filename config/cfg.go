package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/joycastle/casual-server-lib/util"
	"github.com/spf13/viper"
)

type ParseFunc func(v *viper.Viper) error

var registerParseFunc []ParseFunc
var ErrFileNotExists error = errors.New("file not exists")

func RegisterParser(f ParseFunc) {
	registerParseFunc = append(registerParseFunc, f)
}

func InitConfig(path string) error {
	fileFullPathNames := []string{}

	defaultViper := viper.New()

	if util.IsDir(path) {
		fileFullPathNames = util.ReadDirFiles(path)
	} else {
		fileFullPathNames = append(fileFullPathNames, path)
	}

	for _, fileFullPathName := range fileFullPathNames {
		fpath := filepath.Dir(fileFullPathName)
		fileFullName := filepath.Base(fileFullPathName)
		ext := filepath.Ext(fileFullName)
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
		v.SetConfigName(fileFullName)
		v.SetConfigType(ext)
		v.AddConfigPath(fpath)

		if err := v.ReadInConfig(); err != nil {
			return err
		}

		if err := defaultViper.MergeConfigMap(v.AllSettings()); err != nil {
			return err
		}
	}

	for _, parseHandler := range registerParseFunc {
		if err := parseHandler(defaultViper); err != nil && err != ErrFileNotExists {
			return err
		}
	}

	return nil
}

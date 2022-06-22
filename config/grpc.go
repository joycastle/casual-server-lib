package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	CFG_GRPC      = "s-grpc"
	CFG_GRPC_ADDR = "addr"
)

var Grpc map[string]string = make(map[string]string)

func init() {
	RegisterParser(parseGrpc)
}

func parseGrpc(v *viper.Viper) error {
	mps := v.GetStringMap(CFG_GRPC)

	if len(mps) == 0 {
		return ErrFileNotExists
	}

	for k, v := range mps {
		vv := v.(map[string]interface{})

		var c string

		if s, ok := vv[CFG_GRPC_ADDR]; !ok {
			return fmt.Errorf("GRPC config file not contains \"%s\"", CFG_GRPC_ADDR)
		} else {
			c = s.(string)
		}

		Grpc[k] = c
	}

	return nil
}

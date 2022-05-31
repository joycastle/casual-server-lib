package main

import (
	"fmt"
	"runtime"

	"github.com/joycastle/casual-server-lib/config"
	"github.com/spf13/viper"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	//1.初始化配置

	//config包里面默认会解析log配置
	//
	//log配置文件格式(.yaml)
	// logs:
	//   error:
	//     output: ./error.log
	//     level: DEBUG
	//  main:
	//    output: ./main.log-*-*-*
	//    level: INFO
	if err := config.InitConfig("./config/dev.yaml"); err != nil {
		panic(err)
	}

	//打印log配置
	fmt.Println(config.Logs)

	//初始化log组件
	//log.InitLogs(config.Logs)

	//log.Get("error").Debug("customize error log record...")
	//log.Get("error").Info("customize error log record...")
	//log.Get("error").Warn("customize error log record...")
	//log.Get("error").Fatal("customize error log record...")

	//log.Get("main").Debugf("customize main log record...%s", "test")
	//log.Get("main").Infof("customize main log record...%s", "test")
	//log.Get("main").Warnf("customize main log record...%s", "test")
	//log.Get("main").Fatalf("customize main log record...%s", "test")

	//2.自定义日志配置 (基于viper)
	var mydefine []int

	parse_mydefine := func(v *viper.Viper) error {
		mydefine = v.GetIntSlice("mydefine")
		return nil
	}

	//注册解析函数
	config.RegisterParser(parse_mydefine)

	if err := config.InitConfig("./config/dev.yaml"); err != nil {
		panic(err)
	}

	fmt.Println(mydefine)
}

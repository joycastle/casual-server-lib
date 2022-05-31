package main

import (
	"runtime"

	"github.com/joycastle/casual-server-lib/log"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//log default -------------------------------start
	log.Debug("log debug....")
	log.Debugf("log debugf....")
	log.Info("log info....")
	log.Infof("log infof....")
	log.Warn("log warn....")
	log.Warnf("log warnf....")
	log.Fatal("log fatal....")
	log.Fatalf("log fatalf....")
	//log default -------------------------------end

	//log customize -----------------------------------------start
	logConfigs := map[string]log.LogConf{
		"error": log.LogConf{
			Output: "./logs/error.log",
			Level:  "WARN",
		},
		"user": log.LogConf{
			Output: "./logs/user.log",
			Level:  "INFO",
		},
		"access": log.LogConf{
			Output: "./logs/access.log-*-*-*", //split with time year-month-day eg: access.log-2022-01-01
			Level:  "INFO",
		},
	}

	log.InitLogs(logConfigs)

	log.Get("error").Debug("customize error log record...")
	log.Get("error").Info("customize error log record...")
	log.Get("error").Warn("customize error log record...")
	log.Get("error").Fatal("customize error log record...")

	log.Get("user").Debugf("customize user log record...%s", "test")
	log.Get("user").Infof("customize user log record...%s", "test")
	log.Get("user").Warnf("customize user log record...%s", "test")
	log.Get("user").Fatalf("customize user log record...%s", "test")

	//log customize ---------------------------------------------------end
}

package mysql

import (
	"fmt"
	"time"

	"github.com/joycastle/casual-server-lib/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConf struct {
	Addr     string
	Username string
	Password string
	Database string
	Options  string

	MaxIdle     int
	MaxOpen     int
	MaxLifeTime time.Duration

	SlowSqlTime time.Duration

	SlowLogger string
	ErrLogger  string
	StatLogger string
}

var (
	mysqlPoolMap map[string]*gorm.DB = make(map[string]*gorm.DB)
	mysqlNodes   []string
)

func InitMysql(configs map[string]MysqlConf) error {
	for node, config := range configs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Addr, config.Database)
		if config.Options != "" {
			dsn = dsn + "?" + config.Options
		}

		gormConfig := &gorm.Config{}
		if config.SlowLogger != "" && config.SlowSqlTime > 0 {
			gormConfig.Logger = logger.New(log.Get(config.SlowLogger).Logger, logger.Config{
				SlowThreshold:             config.SlowSqlTime,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				//Colorful:                  true,
			})
		}

		gdb, err := gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			log.Get(config.ErrLogger).Fatal("mysql:", err, dsn, node)
			return err
		}

		if sqlDb, err := gdb.DB(); err != nil {
			log.Get(config.ErrLogger).Fatal("mysql:", err, dsn, node)
			return err
		} else {
			sqlDb.SetMaxIdleConns(config.MaxIdle)
			sqlDb.SetMaxOpenConns(config.MaxOpen)
			sqlDb.SetConnMaxLifetime(config.MaxLifeTime)

			if config.StatLogger != "" {
				//stats monitor
				go func(node string) {
					//type DBStats struct {
					//  MaxOpenConnections int // Maximum number of open connections to the database.

					// Pool Status
					//  OpenConnections int // The number of established connections both in use and idle.
					//  InUse           int // The number of connections currently in use.
					//  Idle            int // The number of idle connections.

					// Counters
					//  WaitCount         int64         // The total number of connections waited for.
					//  WaitDuration      time.Duration // The total time blocked waiting for a new connection.
					//  MaxIdleClosed     int64         // The total number of connections closed due to SetMaxIdleConns.
					//  MaxIdleTimeClosed int64         // The total number of connections closed due to SetConnMaxIdleTime.
					//  MaxLifetimeClosed int64         // The total number of connections closed due to SetConnMaxLifetime.
					//}
					for {
						time.Sleep(time.Second * 20)
						stat := sqlDb.Stats()
						infos := fmt.Sprintf("mysql stat: Connection open:%d, inUse:%d, idle:%d, waitCount:%d, waitDuration:%v dsn:%s",
							stat.OpenConnections,
							stat.InUse,
							stat.Idle,
							stat.WaitCount,
							stat.WaitDuration,
							node)

						log.Get(config.StatLogger).Infof(infos)
					}
				}(node)
			}
		}

		mysqlPoolMap[node] = gdb
		mysqlNodes = append(mysqlNodes, node)
	}

	return nil
}

func Get(sn string) *gorm.DB {
	if g, ok := mysqlPoolMap[sn]; ok {
		return g
	}
	log.Get("error").Fatalf(fmt.Sprintf("mysql node \"%s\" not exists, choose from %v", sn, mysqlNodes))
	panic(fmt.Sprintf("mysql: not exists node:%s", sn))
	return nil
}

package flowcontrol

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/joycastle/casual-server-lib/mysql"
)

func TestMain(m *testing.M) {
	mysqlConfigs := map[string]mysql.MysqlConf{
		"default-master": mysql.MysqlConf{
			Addr:        "127.0.0.1",
			Username:    "root",
			Password:    "123456",
			Database:    "db_game",
			Options:     "charset=utf8mb4&parseTime=True",
			MaxIdle:     16,
			MaxOpen:     128,
			MaxLifeTime: time.Second * 300,
			SlowSqlTime: time.Second * 1,
			SlowLogger:  "slow",
			StatLogger:  "stat",
		},

		"default-slave": mysql.MysqlConf{
			Addr:        "127.0.0.1",
			Username:    "root",
			Password:    "123456",
			Database:    "db_game",
			Options:     "charset=utf8mb4&parseTime=True",
			MaxIdle:     16,
			MaxOpen:     128,
			MaxLifeTime: time.Second * 300,
			SlowSqlTime: time.Second * 1,
			SlowLogger:  "slow",
			StatLogger:  "stat",
		},
	}

	if err := mysql.InitMysql(mysqlConfigs); err != nil {
		panic(err)
	}
	m.Run()
}

func TestCreateTable(t *testing.T) {
	fc := NewFlowControl().SetMysqlNode("default-master", "default-slave").Use("robot-server")
	fc.Startup()

	flow, err := CreateFlow("default-master", "robot-server", "机器人服务流量控制", "levin")
	if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		t.Fatal(err)
	}
	if flow.ID == 0 {
		flow, err = GetFlowByName("default-master", "robot-server")
		if err != nil {
			t.Fatal(err)
		}
	}

	flowConfig1, err := CreateFlowConfig("default-master", flow.ID, MethodRand, "100")
	if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		t.Fatal(err)
	}
	if flowConfig1.ID == 0 {
		flowConfig1, err = GetFlowConfigByFlowIDAndStrategy("default-master", flow.ID, MethodRand)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := OpenFlowConfig("default-master", flowConfig1.ID); err != nil {
		t.Fatal(err)
	}

	flowConfig2, err := CreateFlowConfig("default-master", flow.ID, MethodRemainder10, "0|1|2|3|5")
	if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		t.Fatal(err)
	}
	if flowConfig2.ID == 0 {
		flowConfig2, err = GetFlowConfigByFlowIDAndStrategy("default-master", flow.ID, MethodRemainder10)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := OpenFlowConfig("default-master", flowConfig2.ID); err != nil {
		t.Fatal(err)
	}

	flowConfig3, err := CreateFlowConfig("default-master", flow.ID, MethodWhiteList, "use white")
	if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		t.Fatal(err)
	}
	if flowConfig3.ID == 0 {
		flowConfig3, err = GetFlowConfigByFlowIDAndStrategy("default-master", flow.ID, MethodWhiteList)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := OpenFlowConfig("default-master", flowConfig3.ID); err != nil {
		t.Fatal(err)
	}

	/*
		for i := 0; i < 1000; i++ {
			_, err := fc.CreateFlowWhiteList(flowConfig3.ID, fmt.Sprintf("%d", i))
			if err != nil && !strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
				t.Fatal(err)
			}
		}*/

	if s, hit := fc.IsHit("xxxx", "1000", 1000); hit {
		t.Fatal(1, s)
	}

	if err := CloseFlowConfig("default-master", flowConfig1.ID); err != nil {
		t.Fatal(err)
	}

	if err := CloseFlowConfig("default-master", flowConfig3.ID); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	for i := 0; i < 1000; i++ {
		y := i % 10
		if y == 0 || y == 1 || y == 2 || y == 3 || y == 5 {
			if s, hit := fc.IsHit("robot-server", fmt.Sprintf("%d", i), int64(i)); !hit {
				t.Fatal(i, s)
			}
		} else {
			if s, hit := fc.IsHit("robot-server", fmt.Sprintf("%d", i), int64(i)); hit {
				t.Fatal(y, i, s)
			}
		}
	}

	if err := CloseFlowConfig("default-master", flowConfig2.ID); err != nil {
		t.Fatal(err)
	}

	if err := OpenFlowConfig("default-master", flowConfig3.ID); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	for i := 0; i < 2000; i++ {
		if s, hit := fc.IsHit("robot-server", fmt.Sprintf("%d", i), int64(i)); !hit {
			t.Fatal(i, s)
		}
	}

	if err := CloseFlowConfig("default-master", flowConfig2.ID); err != nil {
		t.Fatal(err)
	}

	if err := CloseFlowConfig("default-master", flowConfig3.ID); err != nil {
		t.Fatal(err)
	}

	if err := OpenFlowConfig("default-master", flowConfig1.ID); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 20)

	for i := 0; i < 2000; i++ {
		if s, hit := fc.IsHit("robot-server", fmt.Sprintf("%d", i), int64(i)); !hit || s != MethodRand {
			t.Fatal(i, s)
		}
	}
}

func BenchmarkHit(b *testing.B) {
	Use("robot-server")
	SetMysqlNode("default-master", "default-slave")
	Startup()
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()
	b.ResetTimer()
	// 设置并发数
	b.SetParallelism(5000)
	// 测试多线程并发模式
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := rand.Intn(99999999)
			IsHit("robot-server", fmt.Sprintf("%d", i), int64(i))
		}
	})
}

package flowcontrol

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joycastle/casual-server-lib/log"
	"github.com/joycastle/casual-server-lib/mysql"
)

const (
	MethodRand        = "rand"
	MethodRemainder10 = "remainder10"
	MethodWhiteList   = "whitelist"
)

type FLowControl struct {
	flowMap          map[string]Flow
	flowConfigMap    map[uint64][]FlowConfig
	flowWhiteListMap map[uint64]map[string]struct{}
	dbSlave          string
	dbMaster         string
	names            []string
	mu               *sync.RWMutex
}

func NewFlowControl() *FLowControl {
	return &FLowControl{
		flowMap:          make(map[string]Flow),
		flowConfigMap:    make(map[uint64][]FlowConfig),
		flowWhiteListMap: make(map[uint64]map[string]struct{}),
		mu:               new(sync.RWMutex),
	}
}

func (fc *FLowControl) SetMysqlNode(master, slave string) *FLowControl {
	fc.dbMaster = master
	fc.dbSlave = slave
	return fc
}

func (fc *FLowControl) Use(names ...string) *FLowControl {
	for _, name := range names {
		fc.names = append(fc.names, name)
	}
	return fc
}

func (fc *FLowControl) Startup() {
	go fc.reload()
}

func (fc *FLowControl) reload() {
	for {
		start := time.Now()

		var (
			flows       []Flow
			flowConfigs []FlowConfig
			whiteLists  []FlowWhiteList
		)

		if len(fc.names) == 0 {
			log.Get("error").Fatal("flowcontrol:", "not set any items,  using function 'Use()' to set")
			time.Sleep(time.Second * 5)
			continue
		}

		//db load
		if r := mysql.Get(fc.dbSlave).Where("name IN ?", fc.names).Find(&flows); r.Error != nil {
			log.Get("error").Fatal("flowcontrol:", r.Error)
			time.Sleep(time.Second * 5)
			continue
		}

		flowIds := []uint64{}
		for _, flow := range flows {
			flowIds = append(flowIds, flow.ID)
		}

		if len(flowIds) > 0 {
			if r := mysql.Get(fc.dbSlave).Where("flow_id IN ?", flowIds).Find(&flowConfigs); r.Error != nil {
				log.Get("error").Fatal("flowcontrol:", r.Error)
				time.Sleep(time.Second * 5)
				continue
			}
		}

		configIds := []uint64{}
		for k, flowConfig := range flowConfigs {
			if flowConfig.Open == 1 {
				if flowConfig.Strategy == MethodWhiteList {
					configIds = append(configIds, flowConfig.ID)
				} else if flowConfig.Strategy == MethodRand {
					intv, err := strconv.Atoi(flowConfig.Value)
					if err != nil {
						log.Get("error").Fatal("flowcontrol:", err)
						continue
					}
					flowConfigs[k].ValueParseInt = intv
				} else if flowConfig.Strategy == MethodRemainder10 {
					vs := strings.Split(flowConfig.Value, "|")
					out := make(map[int]struct{})
					for _, v := range vs {
						intv, err := strconv.Atoi(v)
						if err != nil {
							log.Get("error").Fatal("flowcontrol:", err)
							continue
						}
						out[intv] = struct{}{}
					}
					flowConfigs[k].ValueParseMap = out
				}
			}
		}

		if len(configIds) > 0 {
			maxID := uint64(0)
			limit := 1000
			isFatal := false
			for {
				var wlist []FlowWhiteList
				if r := mysql.Get(fc.dbSlave).Where("config_id IN ? AND id > ?", configIds, maxID).Order("id ASC").Limit(limit).Find(&wlist); r.Error != nil {
					log.Get("error").Fatal("flowcontrol:", r.Error)
					isFatal = true
					break
				}

				for _, v := range wlist {
					whiteLists = append(whiteLists, v)
				}

				if len(wlist) < limit {
					break
				} else {
					maxID = wlist[limit-1].ID
				}
			}

			if isFatal {
				time.Sleep(time.Second * 5)
				continue
			}
		}

		//data merge
		fc.mu.Lock()
		fc.flowMap = make(map[string]Flow)
		for _, flow := range flows {
			fc.flowMap[flow.Name] = flow
		}

		fc.flowConfigMap = make(map[uint64][]FlowConfig)
		for _, flowConfig := range flowConfigs {
			if flowConfig.Open == 1 {
				if _, ok := fc.flowConfigMap[flowConfig.FlowID]; !ok {
					fc.flowConfigMap[flowConfig.FlowID] = []FlowConfig{}
				}
				fc.flowConfigMap[flowConfig.FlowID] = append(fc.flowConfigMap[flowConfig.FlowID], flowConfig)
			}
		}

		fc.flowWhiteListMap = make(map[uint64]map[string]struct{})
		for _, v := range whiteLists {
			if _, ok := fc.flowWhiteListMap[v.ConfigID]; !ok {
				fc.flowWhiteListMap[v.ConfigID] = make(map[string]struct{})
			}
			fc.flowWhiteListMap[v.ConfigID][v.SubID] = struct{}{}
		}

		log.Get("run").Info("flowcontrol:", "flow map", fc.flowMap)
		log.Get("run").Info("flowcontrol:", "flowconfig map", fc.flowConfigMap)
		log.Get("run").Info("flowcontrol:", "flowwhitelist map", fc.flowWhiteListMap)

		fc.mu.Unlock()

		cost := time.Since(start).Nanoseconds() / 1000000
		log.Get("run").Info("flowcontrol: reload cost:", cost, "ms")
		time.Sleep(time.Second * 20)
	}
}

func (fc *FLowControl) IsHit(ftype string, idxs string, idxi int64) (string, bool) {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	if flow, ok := fc.flowMap[ftype]; !ok {
		return "none", false
	} else {
		if flowConfigs, ok := fc.flowConfigMap[flow.ID]; !ok {
			return "none", false
		} else {
			hit := false
			hitStrategy := ""
			for _, flowConfig := range flowConfigs {
				if hit {
					break
				}

				switch flowConfig.Strategy {
				case MethodWhiteList:
					_, hit = fc.flowWhiteListMap[flowConfig.ID][idxs]
					hitStrategy = MethodWhiteList
				case MethodRand:
					if rand.Intn(100) < flowConfig.ValueParseInt {
						hit = true
					} else {
						hit = false
					}
					hitStrategy = MethodRand
				case MethodRemainder10:
					index := idxi % 10
					_, hit = flowConfig.ValueParseMap[int(index)]
					hitStrategy = MethodRemainder10
				}
			}
			return hitStrategy, hit
		}
	}
	return "none", false
}

package flowcontrol

import (
	"time"

	"github.com/joycastle/casual-server-lib/mysql"
)

func CreateFlow(node string, name string, desc string, author string) (Flow, error) {
	var flow Flow
	flow.Name = name
	flow.Desc = desc
	flow.Author = author
	flow.AddTime = time.Now().Unix()
	flow.UpdateTime = time.Now().Unix()

	if r := mysql.Get(node).Create(&flow); r.Error != nil {
		return flow, r.Error
	}
	return flow, nil
}

func GetFlowByName(node string, name string) (Flow, error) {
	var flow Flow
	if r := mysql.Get(node).Where("name = ?", name).Limit(1).Find(&flow); r.Error != nil {
		return flow, r.Error
	}
	return flow, nil
}

func CreateFlowConfig(node string, flowID uint64, strategy string, value string) (FlowConfig, error) {
	var flowc FlowConfig
	flowc.FlowID = flowID
	flowc.Strategy = strategy
	flowc.Value = value
	flowc.AddTime = time.Now().Unix()
	flowc.UpdateTime = time.Now().Unix()

	if r := mysql.Get(node).Create(&flowc); r.Error != nil {
		return flowc, r.Error
	}
	return flowc, nil
}

func GetFlowConfigByFlowIDAndStrategy(node string, flowID uint64, strategy string) (FlowConfig, error) {
	var flowc FlowConfig
	if r := mysql.Get(node).Where("flow_id = ? AND strategy = ?", flowID, strategy).Limit(1).Find(&flowc); r.Error != nil {
		return flowc, r.Error
	}
	return flowc, nil
}

func OpenFlowConfig(node string, configID uint64) error {
	if r := mysql.Get(node).Model(&FlowConfig{}).Where("id = ?", configID).Limit(1).Update("open", 1); r.Error != nil {
		return r.Error
	}
	return nil
}

func CloseFlowConfig(node string, configID uint64) error {
	if r := mysql.Get(node).Model(&FlowConfig{}).Where("id = ?", configID).Limit(1).Update("open", 0); r.Error != nil {
		return r.Error
	}
	return nil
}

func CreateFlowWhiteList(node string, configID uint64, subID string) (FlowWhiteList, error) {
	var fwl FlowWhiteList
	fwl.ConfigID = configID
	fwl.SubID = subID
	if r := mysql.Get(node).Create(&fwl); r.Error != nil {
		return fwl, r.Error
	}
	return fwl, nil
}

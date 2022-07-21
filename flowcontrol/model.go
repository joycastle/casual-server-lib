package flowcontrol

type Flow struct {
	ID         uint64 `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	Desc       string `json:"desc" gorm:"column:desc"`
	Author     string `json:"author" gorm:"column:author"`
	AddTime    int64  `json:"add_time" gorm:"column:add_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
}

func (f *Flow) TableName() string {
	return "flow"
}

type FlowConfig struct {
	ID         uint64 `json:"id" gorm:"column:id"`
	FlowID     uint64 `json:"flow_id" gorm:"column:flow_id"`
	Strategy   string `json:"strategy" gorm:"column:strategy"`
	Value      string `json:"value" gorm:"column:value"`
	Open       uint   `json:"open" gorm:"column:open"`
	Extra      string `json:"extra" gorm:"column:extra"`
	AddTime    int64  `json:"add_time" gorm:"column:add_time"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`

	ValueParseInt   int              `gorm:"-"`
	ValueParseSlice []int            `gorm:"-"`
	ValueParseMap   map[int]struct{} `gorm:"-"`
}

func (fc *FlowConfig) TableName() string {
	return "flow_config"
}

type FlowWhiteList struct {
	ID       uint64 `json:"id" gorm:"column:id"`
	ConfigID uint64 `json:"config_id" gorm:"column:config_id"`
	SubID    string `json:"sub_id" gorm:"column:sub_id"`
	AddTime  int64  `json:"add_time" gorm:"column:add_time"`
}

func (fwl *FlowWhiteList) TableName() string {
	return "flow_white_list"
}

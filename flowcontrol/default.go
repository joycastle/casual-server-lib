package flowcontrol

var defaultFlowControl *FLowControl = NewFlowControl()

func IsHit(ftype string, idxs string, idxi int64) (string, bool) {
	return defaultFlowControl.IsHit(ftype, idxs, idxi)
}

func SetMysqlNode(master, slave string) *FLowControl {
	return defaultFlowControl.SetMysqlNode(master, slave)
}

func Use(names ...string) *FLowControl {
	return defaultFlowControl.Use(names...)
}

func Startup() {
	defaultFlowControl.Startup()
}

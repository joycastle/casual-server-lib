package log

var loggers map[string]*Logger = make(map[string]*Logger)

func InitLogs(cfgs map[string]LogConf) {
	for sn, cfg := range cfgs {
		loggers[sn] = NewLogger(cfg)
	}
}

func Get(sn string) *Logger {
	if l, ok := loggers[sn]; ok {
		return l
	}
	return Default
}

func Debugf(f string, v ...any) {
	Default.Debugf(f, v...)
}

func Debug(v ...any) {
	Default.Debug(v...)
}

func Infof(f string, v ...any) {
	Default.Infof(f, v...)
}

func Info(v ...any) {
	Default.Info(v...)
}

func Warnf(f string, v ...any) {
	Default.Warnf(f, v...)
}

func Warn(v ...any) {
	Default.Warn(v...)
}

func Fatalf(f string, v ...any) {
	Default.Fatalf(f, v...)
}

func Fatal(v ...any) {
	Default.Fatal(v...)
}

func EnableColor() {
	Default.EnableColor()
}

func DisableColor() {
	Default.DisableColor()
}

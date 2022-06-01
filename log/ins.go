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
	return DefaultLogger
}

func Debugf(f string, v ...any) {
	DefaultLogger.Debugf(f, v...)
}

func Debug(v ...any) {
	DefaultLogger.Debug(v...)
}

func Infof(f string, v ...any) {
	DefaultLogger.Infof(f, v...)
}

func Info(v ...any) {
	DefaultLogger.Info(v...)
}

func Warnf(f string, v ...any) {
	DefaultLogger.Warnf(f, v...)
}

func Warn(v ...any) {
	DefaultLogger.Warn(v...)
}

func Fatalf(f string, v ...any) {
	DefaultLogger.Fatalf(f, v...)
}

func Fatal(v ...any) {
	DefaultLogger.Fatal(v...)
}

func EnableColor() {
	DefaultLogger.EnableColor()
}

func DisableColor() {
	DefaultLogger.DisableColor()
}

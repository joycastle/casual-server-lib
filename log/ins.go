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
	return defaultLogger
}

func Debugf(f string, v ...any) {
	defaultLogger.Debugf(f, v...)
}

func Debug(v ...any) {
	defaultLogger.Debug(v...)
}

func Infof(f string, v ...any) {
	defaultLogger.Infof(f, v...)
}

func Info(v ...any) {
	defaultLogger.Info(v...)
}

func Warnf(f string, v ...any) {
	defaultLogger.Warnf(f, v...)
}

func Warn(v ...any) {
	defaultLogger.Warn(v...)
}

func Fatalf(f string, v ...any) {
	defaultLogger.Fatalf(f, v...)
}

func Fatal(v ...any) {
	defaultLogger.Fatal(v...)
}

func EnableColor() {
	defaultLogger.EnableColor()
}

func DisableColor() {
	defaultLogger.DisableColor()
}

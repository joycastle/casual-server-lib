package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"

	level_all   = 0x0
	level_debug = 0x1
	level_info  = 0x2
	level_warn  = 0x4
	level_fatal = 0x8
	level_off   = 0xF
)

var (
	logLevelMapForConfig map[string]int          = make(map[string]int)
	logColorMap          map[int8]map[int]string = make(map[int8]map[int]string)
	DefaultLogger        *Logger
)

func init() {
	logLevelMapForConfig["ALL"] = level_all
	logLevelMapForConfig["DEBUG"] = level_debug
	logLevelMapForConfig["INFO"] = level_info
	logLevelMapForConfig["WARN"] = level_warn
	logLevelMapForConfig["FATAL"] = level_fatal
	logLevelMapForConfig["OFF"] = level_off

	//color print set
	logColorMap[1] = make(map[int]string)
	logColorMap[1][level_debug] = fmt.Sprintf("%sDEBUG%s ", blue, reset)
	logColorMap[1][level_info] = fmt.Sprintf("%sINFO%s ", green, reset)
	logColorMap[1][level_warn] = fmt.Sprintf("%sWARN%s ", red, reset)
	logColorMap[1][level_fatal] = fmt.Sprintf("%sFATAL%s ", yellow, reset)

	//no color print set
	logColorMap[0] = make(map[int]string)
	logColorMap[0][level_debug] = "DEBUG "
	logColorMap[0][level_info] = "INFO "
	logColorMap[0][level_warn] = "WARN "
	logColorMap[0][level_fatal] = "FATAL "

	//init default logger
	DefaultLogger = NewLogger(LogConf{"", "ALL", 0, 0}).SetTraceLevel(4).EnableColor()
}

/*
log conf demo
 Output : error.log
 Level : ALL, DEBUG, INFO, WARN, FATAL, OFF
 ExpireDays : 30
*/

type LogConf struct {
	Output      string
	Level       string
	ExpireDays  int64
	TraceOffset int
}

type Logger struct {
	*log.Logger
	cfg           LogConf
	Fptr          *os.File
	Fname         string
	isEnableColor int8
	level         int
	traceLevel    int
}

func NewLogger(cfg LogConf) *Logger {
	l := &Logger{
		cfg:           cfg,
		isEnableColor: 0,
		traceLevel:    3,
	}

	if level, ok := logLevelMapForConfig[cfg.Level]; ok {
		l.level = level
	} else {
		l.level = level_off
	}

	if cfg.TraceOffset > 0 {
		l.traceLevel = l.traceLevel + cfg.TraceOffset
		l.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	} else {
		l.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	l.setup_file()

	return l
}

func (l *Logger) EnableColor() *Logger {
	l.isEnableColor = 1
	return l
}

func (l *Logger) DisableColor() *Logger {
	l.isEnableColor = 0
	return l
}

func (l *Logger) SetTraceLevel(tl int) *Logger {
	l.traceLevel = tl
	return l
}

func (l *Logger) pf(level int, f string, v ...any) {
	if level < l.level {
		return
	}
	l.Logger.Output(l.traceLevel, logColorMap[l.isEnableColor][level]+fmt.Sprintf(f, v...))
}

func (l *Logger) pln(level int, v ...any) {
	if level < l.level {
		return
	}
	l.Logger.Output(l.traceLevel, logColorMap[l.isEnableColor][level]+fmt.Sprintln(v...))
}

func (l *Logger) Debugf(f string, v ...any) {
	l.pf(level_debug, f, v...)
}

func (l *Logger) Debug(v ...any) {
	l.pln(level_debug, v...)
}

func (l *Logger) Infof(f string, v ...any) {
	l.pf(level_info, f, v...)
}

func (l *Logger) Info(v ...any) {
	l.pln(level_info, v...)
}

func (l *Logger) Warnf(f string, v ...any) {
	l.pf(level_warn, f, v...)
}

func (l *Logger) Warn(v ...any) {
	l.pln(level_warn, v...)
}

func (l *Logger) Fatalf(f string, v ...any) {
	l.pf(level_fatal, f, v...)
}

func (l *Logger) Fatal(v ...any) {
	l.pln(level_fatal, v...)
}

func (l *Logger) setup_file() {
	var (
		fname    string
		fp       *os.File
		deadline time.Time
		err      error
	)

	if len(l.cfg.Output) <= 0 {
		return
	}

	fname, deadline = parse_log_fname(l.cfg.Output)

	if fp, err = open_log_file(fname); err != nil {
		fp = os.Stderr
	}

	l.Fptr = fp
	l.Fname = fname
	l.Logger.SetOutput(fp)

	go func() {
		select {
		case <-time.After(deadline.Sub(time.Now())):
			l.setup_file()
		}
	}()
}

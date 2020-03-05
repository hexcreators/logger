package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

//Logger it's a root object for logging
type Logger struct {
	sync.Mutex
	logger         *log.Logger
	isDebugEnabled bool
	isTraceEnabled bool
	infoLabel      string
	warnLabel      string
	errorLabel     string
	fatalLabel     string
	debugLabel     string
	traceLabel     string
}

//NewStdLogger create a logger
func NewStdLogger(time, isDebugEnabled, isTraceenabled, colors, pid bool) *Logger {
	flags := 0
	if time {
		flags = log.LstdFlags | log.Lmicroseconds
	}
	pre := ""
	if pid {
		pre = pidPrefix()
	}
	l := &Logger{
		logger:         log.New(os.Stderr, pre, flags),
		isDebugEnabled: isDebugEnabled,
		isTraceEnabled: isTraceenabled,
	}

	if colors {
		setColoredLabelFormat(l)
	} else {
		setPlainLabelFormat(l)
	}

	return l
}

func setColoredLabelFormat(l *Logger) {
	colorFormat := "[\x1b[%sm%s\x1b[0m] "
	l.infoLabel = fmt.Sprintf(colorFormat, "32", "INF")
	l.debugLabel = fmt.Sprintf(colorFormat, "36", "DBG")
	l.warnLabel = fmt.Sprintf(colorFormat, "0;93", "WRN")
	l.errorLabel = fmt.Sprintf(colorFormat, "31", "ERR")
	l.fatalLabel = fmt.Sprintf(colorFormat, "31", "FTL")
	l.traceLabel = fmt.Sprintf(colorFormat, "33", "TRC")
}

func setPlainLabelFormat(l *Logger) {
	l.infoLabel = "[INF] "
	l.debugLabel = "[DBG] "
	l.warnLabel = "[WRN] "
	l.errorLabel = "[ERR] "
	l.fatalLabel = "[FTL] "
	l.traceLabel = "[TRC] "
}
func pidPrefix() string {
	return fmt.Sprintf("[%d] ", os.Getpid())
}

//Noticef ...
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.logger.Printf(l.infoLabel+format, v...)
}

//Warnf ...
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger.Printf(l.warnLabel+format, v...)
}

//Errorf ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(l.errorLabel+format, v...)
}

//Fatalf ...
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Printf(l.fatalLabel+format, v...)
}

//Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.isDebugEnabled {
		l.logger.Printf(l.debugLabel+format, v...)
	}
}

//Tracef ...
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.isTraceEnabled {
		l.logger.Printf(l.traceLabel+format, v...)
	}
}

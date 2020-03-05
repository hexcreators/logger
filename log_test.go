package logger

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

func TestStdLogger(t *testing.T) {
	logger := NewStdLogger(false, false, false, false, false)

	flags := logger.logger.Flags()
	if flags != 0 {
		t.Fatalf("expected %q , recived %q \n", 0, flags)
	}
	if logger.isDebugEnabled {
		t.Fatalf("expected %t , recived %t \n", false, logger.isDebugEnabled)
	}
	if logger.isTraceEnabled {
		t.Fatalf("expected %t , recived %t \n", false, logger.isDebugEnabled)
	}

}

func TestStdLoggerWithDebugTraceAndTime(t *testing.T) {
	logger := NewStdLogger(true, true, true, false, false)
	flags := logger.logger.Flags()
	if flags != log.LstdFlags|log.Lmicroseconds {
		t.Fatalf("expected %q , recived %q \n", log.LstdFlags|log.Lmicroseconds, flags)
	}
	if !logger.isDebugEnabled {
		t.Fatalf("expected %t , recived %t \n", true, logger.isDebugEnabled)
	}
	if !logger.isTraceEnabled {
		t.Fatalf("expected %t , recived %t \n", true, logger.isDebugEnabled)
	}
}

func TestStdLoggerNotice(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Noticef("foo")
	}, "[INF] foo\n")
}

func TestLoggerNoticeWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, false, false, true, false)
		Logger.Noticef("foo")
	}, "[\x1b[32mINF\x1b[0m] foo\n")
}

func TestStdLoggerTrace(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, false, true, false, false)
		logger.Tracef("foo")
	}, "[TRC] foo\n")
}

func TestLoggerTraceWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, false, true, true, false)
		Logger.Tracef("foo")
	}, "[\x1b[33mTRC\x1b[0m] foo\n")
}

func TestStdLoggerDebug(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, true, false, false, false)
		logger.Debugf("foo")
	}, "[DBG] foo\n")
}

func TestLoggerDebugWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, true, false, true, false)
		Logger.Debugf("foo")
	}, "[\x1b[36mDBG\x1b[0m] foo\n")
}

func TestStdLoggerFatal(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Fatalf("foo")
	}, "[FTL] foo\n")
}

func TestLoggerFatalWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, false, false, true, false)
		Logger.Fatalf("foo")
	}, "[\x1b[31mFTL\x1b[0m] foo\n")
}

func TestStdLoggerError(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Errorf("foo")
	}, "[ERR] foo\n")
}

func TestLoggerErrorWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, false, false, true, false)
		Logger.Errorf("foo")
	}, "[\x1b[31mERR\x1b[0m] foo\n")
}

func TestStdLoggerWarnf(t *testing.T) {
	erxpectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Warnf("foo")
	}, "[WRN] foo\n")
}

func TestLoggerWarnfWithColor(t *testing.T) {
	erxpectOutput(t, func() {
		Logger := NewStdLogger(false, false, false, true, false)
		Logger.Warnf("foo")
	}, "[\x1b[0;93mWRN\x1b[0m] foo\n")
}

func erxpectOutput(t *testing.T, f func(), erxpectOutput string) {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	f()
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	os.Stderr.Close()
	os.Stderr = old
	out := <-outC
	if out != erxpectOutput {
		t.Fatalf("Expected '%s' ,recived '%s' ", erxpectOutput, out)
	}

}

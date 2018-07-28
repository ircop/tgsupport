package logger

import (
	"fmt"
	"time"
	"os"
	"log"
)

// logger struct
type logger struct {
	DebugEnabled	bool
	LogPath			string
	LogFile			*os.File
}

// mlog is logger instance
var mlog = logger{DebugEnabled: true, LogPath: ""}

// Write log to file with timestamp
func (l *logger) Write(format string, args ...interface{}) {
	t := time.Now()
	tm := fmt.Sprintf("[%d.%02d.%02d %02d:%02d:%02d]: ", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

	fmt.Printf(tm + format + "\n", args...)
	if mlog.LogFile != nil {
		log.Printf(format+"\n", args...)
	}
}

// Panic - Write panic log to separate file
func Panic(format string, args ...interface{}) {
	t := time.Now()
	tm := fmt.Sprintf("[%d.%02d.%02d %02d:%02d:%02d]:[PANIC]: ", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

	fmt.Printf(tm + format + "\n", args...)
	Err(fmt.Sprintf(tm + format + "\n", args...))
}

// WriteAsIs without timestamp
func (l *logger) WriteAsIs(format string, args ...interface{}){
	if mlog.LogFile != nil {
		//fmt.Printf(format, args...)
		log.Printf(format, args...)
	}
}

// Log - normal log
func Log(format string, args ...interface{}) {
	mlog.Write("[INFO]: " + format, args...)
}

// DebugAsIs - write debug without timestamps
func DebugAsIs( format string, args ...interface{}) {
	if mlog.DebugEnabled == false {
		return
	}
	mlog.WriteAsIs(format, args...)
}

// Debug log
func Debug(format string, args ...interface{}) {
	if mlog.DebugEnabled == false {
		return
	}
	mlog.Write("[DEBUG]: " + format, args...)
}

// Err error log
func Err(format string, args ...interface{}) {
	mlog.Write("[ERROR]: " + format, args...)
}

// SetDebug - is debug enabled
func SetDebug(val bool) {
	mlog.DebugEnabled = val
}

// SetPath - set log files paths
func SetPath(path string) {
	mlog.LogPath = path

	if( mlog.LogPath != "" ) {
		var err error
		mlog.LogFile, err = os.OpenFile(mlog.LogPath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0600)
		if err != nil {
			log.Fatalf("Cannot open logfile %s: %s", mlog.LogPath, err.Error())
		}
		log.SetOutput(mlog.LogFile)
	}
}

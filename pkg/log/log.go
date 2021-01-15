package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

type CLog struct{}

// NewLogger instantiates instance var with the configurations
// debugFlag decides to discard or write debug logs
func NewLogger(debugFlag bool) {
	debugHandle := ioutil.Discard
	// create the info and error handler by default
	infoHandle := os.Stdout
	errorHandle := os.Stderr

	if debugFlag {
		debugHandle = os.Stdout
	}

	// instatiate package level vars
	Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func (c CLog) Debugf(format string, args ...interface{}) {
	Debug.Printf(format, args...)
}

func (logger CLog) Infof(format string, args ...interface{}) {
	Info.Printf(format, args...)
}

func (logger CLog) Warnf(format string, args ...interface{}) {
	Info.Printf(format, args...)
}

func (logger CLog) Errorf(format string, args ...interface{}) {
	Error.Printf(format, args...)
}

package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// Info level
	Info *log.Logger
	// Debug level
	Debug *log.Logger
	// Error level
	Error *log.Logger
)

//CLog control-plane snapshot logger
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
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds)
}

//Debugf Logger for controlplane, funcs used in snapshot logging
func (c CLog) Debugf(format string, args ...interface{}) {
	Debug.Printf(format, args...)
}

//Infof Logger for controlplane, funcs used in snapshot logging
func (c CLog) Infof(format string, args ...interface{}) {
	Info.Printf(format, args...)
}

//Warnf Logger for controlplane, funcs used in snapshot logging
func (c CLog) Warnf(format string, args ...interface{}) {
	Info.Printf(format, args...)
}

//Errorf Logger for controlplane, funcs used in snapshot logging
func (c CLog) Errorf(format string, args ...interface{}) {
	Error.Printf(format, args...)
}

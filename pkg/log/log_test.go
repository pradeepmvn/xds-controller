package log_test

import (
	"testing"

	"github.com/pradeepmvn/xds-controller/pkg/log"
)

func TestDebuglog(t *testing.T) {
	log.NewLogger(true)

	log.Debug.Println("Debug Message")
	log.Info.Println("Info Message")
	log.Error.Println("Error Message")
}

func TestInfolog(t *testing.T) {
	log.NewLogger(false)

	log.Debug.Println("Debug Message")
	log.Info.Println("Info Message")
	log.Error.Println("Error Message")
}

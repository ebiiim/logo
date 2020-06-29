package pkg

import (
	"github.com/ebiiim/logo"
)

// Log is the Logo logger used for this package.
var Log = logo.New(logo.INFO, nil)

// PrintLogs prints some logs.
func PrintLogs() {
	Log.D("this is a DEBUG message: %v", "a")
	Log.I("this is an INFO message: %v", "b")
	Log.W("this is a WARNING message: %v", "c")
	Log.E("this is an ERROR message: %v", "d")
}

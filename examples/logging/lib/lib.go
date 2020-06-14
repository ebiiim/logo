package pkg

import (
	"github.com/ebiiim/logo"
)

// Log is the Logo logger used for this package.
var Log = logo.New(logo.INFO, nil)

// PrintLogs print some logs.
func PrintLogs() {
	Log.D("this is a debug message: %v", "a")
	Log.I("this is a debug message: %v", "b")
	Log.W("this is a warning message: %v", "c")
	Log.E("this is a error message: %v", "d")
}

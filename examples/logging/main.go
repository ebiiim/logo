package main

import (
	"log"
	"os"

	"github.com/ebiiim/logo"
	lib "github.com/ebiiim/logo/examples/logging/lib"
)

func main() {
	lib.PrintLogs()
	// 2000/01/01 00:00:00 [INFO] this is a debug message: b
	// 2000/01/01 00:00:00 [WARNING] this is a warning message: c
	// 2000/01/01 00:00:00 [ERROR] this is a error message: d

	lib.Log.Level = logo.ERROR
	lib.PrintLogs()
	// 2000/01/01 00:00:00 [ERROR] this is a error message: d

	logger := log.New(os.Stdout, "[prefix]", log.Lshortfile)
	lib.Log.Logger = logger
	lib.PrintLogs()
	// [prefix]logo.go:82: [ERROR] this is a error message: d
}

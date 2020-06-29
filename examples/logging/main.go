package main

import (
	"log"
	"os"

	"github.com/ebiiim/logo"
	lib "github.com/ebiiim/logo/examples/logging/lib"
)

func main() {
	lib.PrintLogs()
	// 2000/01/01 00:00:00 [INFO] this is an INFO message: b
	// 2000/01/01 00:00:00 [WARNING] this is a WARNING message: c
	// 2000/01/01 00:00:00 [ERROR] this is an ERROR message: d

	lib.Log.Level = logo.ERROR
	lib.PrintLogs()
	// 2000/01/01 00:00:00 [ERROR] this is an ERROR message: d

	logger := log.New(os.Stdout, "[prefix]", log.Lshortfile)
	lib.Log.Logger = logger
	lib.PrintLogs()
	// [prefix]logo.go:82: [ERROR] this is an ERROR message: d
}

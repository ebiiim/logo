# LOGO


[![GoDoc](https://godoc.org/github.com/ebiiim/logo?status.svg)](https://godoc.org/github.com/ebiiim/logo)
[![Build Status](https://travis-ci.org/ebiiim/logo.svg?branch=master)](https://travis-ci.org/ebiiim/logo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebiiim/logo)](https://goreportcard.com/report/github.com/ebiiim/logo)

## What is this?

A simple logging library.

## Usage

### In your library:

```go
var Log = logo.New(logo.INFO, nil)

func Hello(){
    Log.I("hello")
    fmt.Println("world")
}
```

### In your app:

1. No settings required!

```go
func main(){
    lib.Hello()
}
// [INFO] hello
// world
```

2. You can set the minimum log level to display.

- `DEBUG` < `INFO` < `WARNING` < `ERROR`

```go
func main(){
    lib.Log.Level = Logo.ERROR
    lib.Hello()
}
// world
```

3. You can also set a custom label for each log level.

```go
func main(){
    lib.Log.LogMessageMap[Logo.INFO] = "[情報] "
    lib.Hello()
}
// [情報] hello
// world
```

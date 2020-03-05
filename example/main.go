package main

import "github.com/hexcreators/logger"

func main() {
	//	var logger logger.Logger
	logger := logger.NewStdLogger(true, true, true, false, true)
	logger.Noticef("test")
}

package main

import (
	"github.com/Nutdat/logger"
)

func main() {
	defer logger.RecoverAndFlush()
	logger.Console("SQL", "SELECT * FROM reallife WHERE pain = 0")
	logger.Info("its fine")
	logger.Warn("maybe not")
	logger.Error("pretty sure something wrong")
	logger.Fatal("thats fucked up")
	panic("jesus")
}

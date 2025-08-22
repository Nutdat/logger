package main

import (
	"github.com/Nutdat/logger"
)

type TLSConfig struct {
	Port     int
	CertFile string
	KeyFile  string
	Enabled  bool
}

type HTTPConfig struct {
	Port  int
	Debug bool
	TLS   TLSConfig
}

func main() {
	defer logger.RecoverAndFlush()
	logger.Console("SQL", "SELECT * FROM reallife WHERE pain = 0")
	logger.Info("its fine")
	logger.Warn("maybe not")
	logger.Error("pretty sure something wrong")
	logger.Fatal("thats fucked up")
	defaultConfig := HTTPConfig{
		Port:  8080,
		Debug: true,
		TLS: TLSConfig{
			Port:     8433,
			CertFile: "server.crt",
			KeyFile:  "server.key",
			Enabled:  false,
		},
	}

	logger.PrettyPrintJSON(defaultConfig)
	panic("jesus")
}

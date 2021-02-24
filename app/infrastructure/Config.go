package infrastructure

import (
	"log"
	"os"
)

type Config struct {
	AbsolutePath string
}

func NewConfig() *Config {
	conf := new(Config)
	var err error
	conf.AbsolutePath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

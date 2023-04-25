package main

import (
	"flag"

	log "github.com/inconshreveable/log15"
)

var ConfigFile string

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&ConfigFile, "config-file", "config.yaml", "config file for kafka-faker")
	if err := fs.Parse(flag.Args()); err != nil {
		log.Error("error parsing flags", "err", err)
		return
	}
}

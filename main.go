package main

import (
	"context"
	"flag"
	"os"
	"sync"

	log "github.com/inconshreveable/log15"
)

var configFile string
var wg sync.WaitGroup

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&configFile, "config-file", "config.yaml", "config file for kafka-faker")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Error("error parsing flags", "err", err)
		return
	}

	cfg, err := LoadConfig(configFile)
	if err != nil {
		log.Error("failed to load config", "filename", configFile, "err", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// initialize and start sender
	sender, err := NewSender(cfg.Kafka)
	if err != nil {
		log.Error("failed to create a sender", "err", err)
		return
	}
	go func() {
		sender.Run(ctx)
	}()

	// initialize and start generators
	for i := 0; i < len(cfg.Generators); i++ {
		wg.Add(1)
		go func(cfg GeneratorConfig) {
			defer wg.Done()
			g := NewGenerator(cfg, sender)
			g.Run(ctx)
		}(cfg.Generators[i])
	}

	log.Info("kafka-faker started to generate messages")
	wg.Wait()
	log.Info("kafka-faker terminated, no generators left")
}

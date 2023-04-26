package main

import (
	"context"

	log "github.com/inconshreveable/log15"
)

type Generator struct {
	cfg    GeneratorConfig
	sender *Sender
}

func NewGenerator(cfg GeneratorConfig, sender *Sender) *Generator {
	return &Generator{
		cfg:    cfg,
		sender: sender,
	}
}

func (g *Generator) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			task := func() {
				for i := 0; i < g.cfg.Number; i++ {
					parsed, err := ParseTemplate(g.cfg.Template)
					if err != nil {
						log.Error("failed to generate from template", "template", g.cfg.Template, "err", err)
						continue
					}

					g.sender.Send(&Entry{
						Topic: g.cfg.Topic,
						Key:   g.cfg.PartitionKey,
						Value: parsed,
					})
				}
			}

			if g.cfg.Loop {
				tl := TaskLoop{
					Delay: g.cfg.Delay,
				}
				tl.Do(task)
			} else {
				to := TaskOnce{}
				to.Do(task)
			}
		}
	}
}

package main

import (
	"context"
	"encoding/json"

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
					var key string
					txt := g.cfg.Schema.GenerateJSON()

					b, err := json.Marshal(txt)
					if err != nil {
						log.Error("failed to generate JSON from schema", "err", err)
						panic("failed to generate JSON from schema")
					}

					if g.cfg.PartitionKey != "" && g.cfg.PartitionKey != "nil" {
						var ok bool
						key, ok = txt[g.cfg.PartitionKey].(string)
						if !ok {
							log.Error("partition key is not string", "key", g.cfg.PartitionKey)
							panic("partition key is not string")
						}
					}

					g.sender.Send(&Entry{
						Topic: g.cfg.Topic,
						Key:   key,
						Value: string(b),
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

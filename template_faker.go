package main

import (
	"html/template"
	"math/rand"
	"time"

	"github.com/jaswdr/faker"
)

var TemplateFunctions = template.FuncMap{
	"personName": func() string {
		return fake().Person().Name()
	},
	"address": func() string {
		return fake().Address().Address()
	},
	"loremText": func(size int) string {
		return fake().Lorem().Text(size)
	},
	"email": func() string {
		return fake().Internet().Email()
	},
	"ethereumAddress": func() string {
		return fake().Crypto().EtheriumAddress()
	},
	"intBetween": func(min, max int) int {
		return fake().IntBetween(min, max)
	},
	"timestamp": func() int64 {
		return time.Now().UTC().Unix()
	},
}

func fake() faker.Faker {
	seed := rand.NewSource(time.Now().UnixNano())
	return faker.NewWithSeed(seed)
}

package main

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/jaswdr/faker"
)

type FuncType string

var (
	FuncEthereum  FuncType = "ethereum_address"
	FuncNumber    FuncType = "number"
	FuncTimestamp FuncType = "timestamp"
)

func StringToFuncType(name string) FuncType {
	switch name {
	case "ethereum_address":
		return FuncEthereum
	case "number":
		return FuncNumber
	case "timestamp":
		return FuncTimestamp
	default:
		return ""
	}
}

type Func struct {
	Name string
	Args []string
}

// Unmarshal parses the function from a string.
func (f *Func) Unmarshal(txt string) error {
	if !strings.HasPrefix(txt, "::") {
		return errors.New("invalid function format")
	}

	if !strings.HasSuffix(txt, ")") {
		return errors.New("invalid function format")
	}

	txt = txt[2:]
	txt = txt[:len(txt)-1]
	parts := strings.Split(txt, "(")
	if len(parts) != 2 {
		return errors.New("invalid function format")
	}

	f.Name = parts[0]
	if len(parts) > 1 {
		args := strings.Split(parts[1], ",")
		for _, arg := range args {
			if arg != "" {
				f.Args = append(f.Args, strings.TrimSpace(arg))
			}
		}
	}

	return nil
}

func (f *Func) Generate() interface{} {
	seed := rand.NewSource(time.Now().UnixMicro())
	fake := faker.NewWithSeed(seed)

	switch StringToFuncType(f.Name) {
	case FuncEthereum:
		return fake.Crypto().EtheriumAddress()
	case FuncNumber:
		return fake.IntBetween(10, 10000000)
	case FuncTimestamp:
		return time.Now().UTC().Unix()
	}

	return nil
}

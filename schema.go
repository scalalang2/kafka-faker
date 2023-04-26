package main

import (
	"encoding/json"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

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

// Schema represents JSON schema in the config file.
type Schema map[string]interface{}

func (s *Schema) MarshalYAML() (interface{}, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (s *Schema) UnmarshalYAML(value *yaml.Node) error {
	var txt string
	if err := value.Decode(&txt); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(txt), s); err != nil {
		return err
	}

	if err := parseFunc((*map[string]interface{})(s)); err != nil {
		return err
	}

	return nil
}

// parseFunc traverses the schema and replaces function strings with Func objects.
func parseFunc(m *map[string]interface{}) error {
	for k, v := range *m {
		switch typed := v.(type) {
		case string:
			if strings.HasPrefix(typed, "::") {
				f := &Func{}
				if err := f.Unmarshal(typed); err != nil {
					return err
				}
				(*m)[k] = f
			}
		case map[string]interface{}:
			return parseFunc(&typed)
		}
	}

	return nil
}

package main

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

type JSON map[string]interface{}

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

	if err := walkToParseFunc((*map[string]interface{})(s)); err != nil {
		return err
	}

	return nil
}

func (s *Schema) GenerateJSON() JSON {
	return walkToGenerateSchema((*map[string]interface{})(s))
}

func walkToGenerateSchema(m *map[string]interface{}) JSON {
	kv := make(JSON)
	for k, v := range *m {
		switch typed := v.(type) {
		case string:
			kv[k] = typed
		case int, int8, int16, int32, int64:
			kv[k] = typed
		case *Func:
			kv[k] = typed.Generate()
		case map[string]interface{}:
			kv[k] = walkToGenerateSchema(&typed)
		default:
			kv[k] = typed
		}
	}
	return kv
}

// walkToParseFunc traverses the schema and replaces function strings with Func objects.
func walkToParseFunc(m *map[string]interface{}) error {
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
			return walkToParseFunc(&typed)
		}
	}

	return nil
}

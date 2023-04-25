package main

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Schema represents JSON schema in the config file.
type Schema map[string]string

func (s *Schema) MarshalYAML() (interface{}, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (s *Schema) UnmarshalYAML(value *yaml.Node) error {
	var txt string
	err := value.Decode(&txt)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(txt), s)
}

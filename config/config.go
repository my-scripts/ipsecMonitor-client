package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Rpc struct {
		Port int `json:"port"`
	}
}

func (this *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Config) Save(path string) error {
	content, err := json.Marshal(this)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, content, 0755)
}

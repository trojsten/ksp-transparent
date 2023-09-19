package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Person struct {
	Name   string `yaml:"name"`
	Avatar string `yaml:"avatar"`
	Candle string `yaml:"candle"`
	Color  string `yaml:"color"`
}

type Config struct {
	People []Person `yaml:"people"`
}

var config Config

func Load() error {
	fh, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer fh.Close()

	decoder := yaml.NewDecoder(fh)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}

func Get() *Config {
	return &config
}

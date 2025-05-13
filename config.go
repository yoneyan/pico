package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Patlites []PatliteConfig `yaml:"patlites"`
	Jobs     []Job           `yaml:"jobs"`
}

type PatliteConfig struct {
	Name   string `yaml:"name"`
	Host   string `yaml:"host"`
	IsHttp bool   `yaml:"is_http"`
}

func GetConfig() error {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	return nil
}

func GetJobName(name string) *Job {
	for _, job := range config.Jobs {
		if job.Name == name {
			return &job
		}
	}
	return nil
}

func GetPatliteConfig(name string) *PatliteConfig {
	for _, patlite := range config.Patlites {
		if patlite.Name == name {
			return &patlite
		}
	}
	return nil
}

package config

import (
	"github.com/mraraneda/mrlogger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type StaticConfig struct {
	POSTGRES struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Dbname       string `yaml:"dbname"`
		Sqlstatement string `yaml:"sqlStatement"`
	} `yaml:"postgres"`
}

// Config es la interfaz de lectura de CSV
type Config interface {
	Read() StaticConfig
	Exist() string
}

// YamlInput es el tipo archivo
type YamlInput struct {
	PathInput string
}

func (c YamlInput) Exist() {
	if _, err := os.Stat(c.PathInput); os.IsNotExist(err) {
		mrlogger.Error("No existe el archivo de configuración:", c.PathInput)
		mrlogger.Check(err, mrlogger.InThisPoint())
	}
}

func (c YamlInput) Read() StaticConfig {

	c.Exist() // Evalúa la existencia del archivo

	file, err := ioutil.ReadFile(c.PathInput)
	mrlogger.Check(err, mrlogger.InThisPoint())

	var yamlConfig StaticConfig

	err = yaml.Unmarshal(file, &yamlConfig)
	mrlogger.Check(err, mrlogger.InThisPoint())

	return yamlConfig
}

func NewStaticConfig(s StaticConfig) *StaticConfig {
	var c StaticConfig = s
	return &c
}

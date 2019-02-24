package util

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/onuryartasi/context-manager/types"
	"gopkg.in/yaml.v2"
)

var configFile string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Homedir error %s", err)
	}
	configFile = filepath.Join(home, ".kube", "config")
}

func GetConfig() types.KubeConfig {
	var config types.KubeConfig
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Can't read config %s, Error: %s", configFile, err)
	}
	err = yaml.Unmarshal(configData, &config)

	return config
}

func SetConfig(config types.KubeConfig) {
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Can't serialize config, %s", err)
	}

	err = ioutil.WriteFile(configFile, data, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

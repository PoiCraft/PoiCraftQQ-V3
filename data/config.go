package data

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var AppDir string
var Conf = new(Config)

func Init(dir string) {
	AppDir = dir
	err := initFiles()
	if err != nil {
		panic("无法创建文件," + err.Error())
	}
	err = initConfig()
	if err != nil {
		panic("无法解析配置" + err.Error())
	}
	InitDB()
}

type Config struct {
	MySQL          string   `yaml:"mysql"`
	Administrators []string `yaml:"administrators"`
}

func initConfig() error {
	yamlFile, err := ioutil.ReadFile(filepath.Join(AppDir, "conf.yaml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		return err
	}
	return nil
}

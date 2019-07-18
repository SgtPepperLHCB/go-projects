package conf

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

// MqttTopic topic/qos
type MqttTopic struct {
	Topic string `yaml:"topic"`
	Qos   int    `yaml:"qos"`
}

func (t MqttTopic) String() string {
	return fmt.Sprintf("%s,%d", t.Topic, t.Qos)
}

// MqttConf configuration structure
type MqttConf struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Topics []MqttTopic
}

func (c *MqttConf) getConf(conf string) *MqttConf {
	yfile, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Printf("Read YAML file error #%v", err)
	}
	err = yaml.Unmarshal(yfile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

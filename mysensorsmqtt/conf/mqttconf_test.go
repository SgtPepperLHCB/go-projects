package conf

import (
	"fmt"
	"testing"
)

func TestCanReadConf(t *testing.T) {
	var c MqttConf
	c.getConf("mqtt.yaml")
	for t := range c.Topics {
		fmt.Printf("t=%s\n", c.Topics[t])
	}
}

package mysensors

import (
	"fmt"
	"testing"
)

func TestCanCreateNewMySensor(t *testing.T) {
	var m = "tlmy/gar-out/20/0/1/0/47 12.3"
	var sensor = NewMySensor(m)
	fmt.Println(sensor)
	fmt.Printf("= %s\n", sensor.String())
}

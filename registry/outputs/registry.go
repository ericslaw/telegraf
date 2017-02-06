package outputs

import (
	"log"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/registry"
)

type Creator func() telegraf.Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	if override := registry.GetName(); override != "" {
		name = override
	}
	log.Println("D! Loading plugin: [[outputs." + name + "]]")
	Outputs[name] = creator
}
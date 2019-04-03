package main

import (
	"flag"
	"log"
	"path/filepath"
	"plugin"

	"./reg"
)

var (
	flPluginDir = flag.String("d", "./plugins/", "default directory for plugins")
)

func main() {
	flag.Parse()

	matches, err := filepath.Glob(*flPluginDir + "/*.so")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range matches {
		p, err := plugin.Open(file)
		if err != nil {
			panic(err)
		}
		v, err := p.Lookup("V")
		if err != nil {
			panic(err)
		}
		f, err := p.Lookup("F")
		if err != nil {
			panic(err)
		}
		*v.(*int) = 7
		f.(func())() // prints "Hello, number 7"
	}
	println(len(reg.Known()))
}

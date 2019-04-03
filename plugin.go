// +build ignore

package main

// // No C code needed.
import "C"

import "fmt"
import "./reg"

func init() {
	reg.RegisterPlugin(myplug)
}

var myplug = myPlug{name: "farts"}

type myPlug struct {
	name string
}

func (mp myPlug) Name() string {
	return mp.name
}

var V int

func F() { fmt.Printf("Hello, number %d\n", V) }

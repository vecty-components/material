package component_test

import (
	"log"

	"agamigo.io/gojs/jsdom"
	"agamigo.io/material/mdctest"
)

var (
	dom jsdom.JSDOM
)

func init() {
	err := mdctest.LoadMDCModule()
	if err != nil {
		log.Fatalf("Unable to load MDC JS module: %v", err)
	}

	dom, err = mdctest.EmulateDOM()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

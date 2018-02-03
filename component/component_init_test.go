package component_test

import (
	"log"

	"agamigo.io/material/mdctest"
)

func init() {
	err := mdctest.LoadMDCModule()
	if err != nil {
		log.Fatalf("Unable to load MDC JS module: %v", err)
	}
}

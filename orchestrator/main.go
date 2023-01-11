package main

import (
	"fmt"
)

func main() {
	fmt.Println("Orchestrator has started")

	StartContainerFromImage("dockersamples/101-tutorial")
}

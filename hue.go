package main

import (
	"fmt"
	"io"
)

func connect(logger io.Writer, bridge string) string {
	fmt.Fprintf(logger, "[green](bridge:connect)[white] Connecting to bridge %s...\n", bridge)
	fmt.Fprintf(logger, "[green](bridge:connect)[white] Successfully connected to bridge %s!\n", bridge)
	return bridge
}

func discoverGroups(logger io.Writer, bridge string) []string {
	fmt.Fprintf(logger, "[green](groups:discover)[white] Discovering groups on %s...\n", bridge)
	groups := []string{
		"Bedroom",
		"Closet",
		"Kitchen",
		"Living Room",
		"Office",
	}
	fmt.Fprintf(logger, "[green](groups:discover)[white] Completed group discovery on %s\n", bridge)
	return groups
}

func discoverLights(logger io.Writer, bridge string) []string {
	fmt.Fprintf(logger, "[green](lights:discover)[white] Discovering lights on %s...\n", bridge)
	lights := []string{
		"Light 1",
		"Light 2",
		"Light 3",
		"Light 4",
		"Light 5",
	}
	fmt.Fprintf(logger, "[green](lights:discover)[white] Completed light discovery on %s\n", bridge)
	return lights
}

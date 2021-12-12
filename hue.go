package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func connect(logs *tview.TextView, bridge string) string {
	fmt.Fprintf(logs, "[green](bridge:connect)[white] Connecting to bridge %s...\n", bridge)
	fmt.Fprintf(logs, "[green](bridge:connect)[white] Successfully connected to bridge %s!\n", bridge)

	return bridge
}

func discoverGroups(app *tview.Application, logs *tview.TextView, list *tview.List, bridge string) {
	fmt.Fprintf(logs, "[green](groups:discover)[white] Discovering groups on %s...\n", bridge)
	app.QueueUpdateDraw(func() {
		list.AddItem("Bedroom", "", 0, nil)
		list.AddItem("Living Room", "", 0, nil)
		list.AddItem("Office", "", 0, nil)
	})
	fmt.Fprintf(logs, "[green](groups:discover)[white] Completed group discovery on %s\n", bridge)
}

func discoverLights(app *tview.Application, logs *tview.TextView, list *tview.List, bridge string) {
	fmt.Fprintf(logs, "[green](lights:discover)[white] Discovering lights on %s...\n", bridge)
	app.QueueUpdateDraw(func() {
		list.AddItem("Light 1", "", 0, nil)
		list.AddItem("Light 2", "", 0, nil)
		list.AddItem("Light 3", "", 0, nil)
		list.AddItem("Light 4", "", 0, nil)
		list.AddItem("Light 5", "", 0, nil)
	})
	fmt.Fprintf(logs, "[green](lights:discover)[white] Completed light discovery on %s\n", bridge)
}

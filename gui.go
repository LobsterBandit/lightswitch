package main

import (
	"fmt"
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// Gui represents the ui renderer and state
type Gui struct {
	app    *tview.Application
	logger io.Writer
	panel  *GuiPanels
	state  *GuiState
}

// GuiPanels holds the various ui panels
type GuiPanels struct {
	bridges *tview.List
	logs    *tview.TextView
	query   *tview.List
	results *tview.List
}

// GuiState represents the ui state
type GuiState struct {
	bridge string
}

func initGui() (gui *Gui) {
	gui = &Gui{
		app:   tview.NewApplication(),
		panel: &GuiPanels{},
		state: &GuiState{},
	}
	logsPanel := gui.createLogsPanel()
	gui.logger = logsPanel
	gui.panel.logs = logsPanel

	gui.panel.bridges = gui.createBridgesPanel()
	gui.panel.query = gui.createQueryPanel()
	gui.panel.results = gui.createResultsPanel()
	return
}

func (gui *Gui) run(cmd *cobra.Command, args []string) {
	dashboard := tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(gui.panel.bridges, 0, 1, true).
			AddItem(gui.panel.query, 0, 1, false), 0, 1, true).
		AddItem(gui.panel.results, 0, 2, false).
		AddItem(gui.panel.logs, 0, 2, false)

	gui.app.SetRoot(tview.NewPages().AddPage("Dashboard", dashboard, true, true), true)
	if err := gui.app.Run(); err != nil {
		cmd.PrintErrf("Error running lightswitch: %w", err)
	}
}

func (gui *Gui) createBridgesPanel() *tview.List {
	bridges := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(func(i int, main string, secondary string, shortcut rune) {
			gui.panel.results.Clear()
			gui.panel.results.AddItem("Connecting...", "", 0, nil)

			gui.state.bridge = connect(gui.panel.logs, main)

			gui.panel.results.Clear()
			gui.panel.results.AddItem("Select a query", "", 0, nil)

			gui.app.SetFocus(gui.panel.query)
		})
	bridges.SetBorder(true).SetTitle("Bridges")

	bridges.AddItem("Living room bridge", "192.168.1.1", 0, func() {
		fmt.Fprintln(gui.logger, "[green](bridge:select)[white] Selected 'Living room bridge'")
	})

	return bridges
}

func (gui *Gui) createQueryPanel() *tview.List {
	query := tview.NewList().ShowSecondaryText(false)
	query.SetBorder(true).SetTitle("Query")

	query.SetSelectedFunc(func(i int, main string, secondary string, shortcut rune) {
		gui.panel.results.Clear()

		fmt.Fprintf(gui.logger, "[green](query:select)[white] Query %s on %s...\n", main, gui.state.bridge)

		switch main {
		case "Groups":
			go func() {
				groups := discoverGroups(gui.logger, gui.state.bridge)
				gui.app.QueueUpdateDraw(func() {
					for _, g := range groups {
						gui.panel.results.AddItem(g, "", 0, nil)
					}
				})
			}()
		case "Lights":
			go func() {
				lights := discoverLights(gui.logger, gui.state.bridge)
				gui.app.QueueUpdateDraw(func() {
					for _, l := range lights {
						gui.panel.results.AddItem(l, "", 0, nil)
					}
				})
			}()
		default:
			gui.panel.results.AddItem(fmt.Sprintf("%s query is unsupported", main), "", 0, nil)
			fmt.Fprintf(gui.logger, "[yellow](query:select) %s query is unsupported[white]\n", main)
		}
	})

	query.AddItem("Groups", "", 0, nil)
	query.AddItem("Lights", "", 0, nil)
	query.AddItem("Scenes", "", 0, nil)
	query.AddItem("Rules", "", 0, nil)
	query.AddItem("Accessories", "", 0, nil)

	return query
}

func (gui *Gui) createLogsPanel() *tview.TextView {
	logs := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			gui.app.Draw()
		})
	logs.SetDoneFunc(func(key tcell.Key) {
		logs.ScrollToEnd()
	})
	logs.SetBorder(true).SetTitle("Logs")

	return logs
}

func (gui *Gui) createResultsPanel() *tview.List {
	results := tview.NewList().
		SetSelectedFocusOnly(true).
		ShowSecondaryText(false)
	results.SetBorder(true).SetTitle("Results")
	results.AddItem("Select a bridge", "", 0, nil)

	return results
}

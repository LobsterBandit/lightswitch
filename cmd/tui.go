package cmd

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Device management Terminal UI",
	Long:  `Terminal UI for management of hue devices`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		var bridge string

		logs := tview.NewTextView().
			SetDynamicColors(true).
			SetRegions(true).
			SetChangedFunc(func() {
				app.Draw()
			})
		logs.SetDoneFunc(func(key tcell.Key) {
			logs.ScrollToEnd()
		})
		logs.SetBorder(true).SetTitle("Logs")

		results := tview.NewList().
			SetSelectedFocusOnly(true).
			ShowSecondaryText(false)
		results.SetBorder(true).SetTitle("Results")
		results.AddItem("Select a bridge", "", 0, nil)

		query := queryWidget()
		query.SetSelectedFunc(func(i int, main string, secondary string, shortcut rune) {
			results.Clear()

			fmt.Fprintf(logs, "[green](query:select)[white] Query %s on %s...\n", main, bridge)

			switch main {
			case "Groups":
				go discoverGroups(app, logs, results, bridge)
			case "Lights":
				go discoverLights(app, logs, results, bridge)
			default:
				results.AddItem(fmt.Sprintf("%s query is unsupported", main), "", 0, nil)
				fmt.Fprintf(logs, "[yellow](query:select) %s query is unsupported[white]\n", main)
			}
		})

		bridges := tview.NewList().
			ShowSecondaryText(false).
			SetSelectedFunc(func(i int, main string, secondary string, shortcut rune) {
				results.Clear()
				results.AddItem("Connecting...", "", 0, nil)

				bridge = connect(logs, main)

				results.Clear()
				results.AddItem("Select a query", "", 0, nil)

				app.SetFocus(query)
			})
		bridges.SetBorder(true).SetTitle("Bridges")

		bridges.AddItem("Living room bridge", "192.168.1.1", 0, func() {
			fmt.Fprintln(logs, "[green](bridge:select)[white] Selected 'Living room bridge'")
		})

		dashboard := tview.NewFlex().
			AddItem(tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(bridges, 0, 1, true).
				AddItem(query, 0, 1, false), 0, 1, true).
			AddItem(results, 0, 2, false).
			AddItem(logs, 0, 2, false)

		app.SetRoot(tview.NewPages().AddPage("Dashboard", dashboard, true, true), true)
		if err := app.Run(); err != nil {
			cmd.PrintErrf("Error running lightswitch tui: %w", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}

func queryWidget() *tview.List {
	query := tview.NewList().ShowSecondaryText(false)
	query.SetBorder(true).SetTitle("Query")

	query.AddItem("Groups", "", 0, nil)
	query.AddItem("Lights", "", 0, nil)
	query.AddItem("Scenes", "", 0, nil)
	query.AddItem("Rules", "", 0, nil)
	query.AddItem("Accessories", "", 0, nil)

	return query
}

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

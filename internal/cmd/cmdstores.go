package cmd

import (
	"os"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kyokomi/emoji/v2"
)

func cmdStores(config *api.Config, args ...string) {
	stores, err := config.GetStores()
	if err != nil {
		color.Red("%sAn error occured while getting stores: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	headerFmt := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyanFmt := color.New(color.FgCyan).SprintfFunc()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{headerFmt("ID"), headerFmt("Store")})

	for i := 1; i <= len(stores); i++ {
		t.AppendRow([]interface{}{cyanFmt("[" + stores[i].StoreID + "]"), stores[i].StoreName})
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}

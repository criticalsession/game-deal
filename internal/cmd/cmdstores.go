package cmd

import (
	"fmt"
	"os"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func cmdStores(config *api.Config, args ...string) {
	stores, err := config.GetStores()
	if err != nil {
		utils.PrintError(fmt.Sprint("An error occured while getting stores:", err.Error()))
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

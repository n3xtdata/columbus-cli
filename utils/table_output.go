package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintTable(header []string, rows [][]string) {

	fmt.Println("")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(50)
	table.SetHeaderLine(true)
	table.SetRowLine(true)
	table.SetAutoWrapText(true)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator(" ")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
	fmt.Println("")
}

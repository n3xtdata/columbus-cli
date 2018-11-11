package get

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/n3xtdata/columbus-cli/restclient"
	"github.com/n3xtdata/columbus-cli/utils"
)

type Checks [] struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

func listChecks() {
	var response = restclient.Get("checks")

	var record Checks

	json.Unmarshal([]byte(response), &record)

	var data [][]string

	for _, v := range record {



		row := []string{v.Label, color.GreenString("SUCCESS"), v.Description}
		data = append(data, row)
	}

	header := []string{"Label", "current Status", "Description"}
	utils.PrintTable(header, data)
}

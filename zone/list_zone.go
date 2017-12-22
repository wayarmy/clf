// @package: Zone
// This package is controller for cloudflare's zones on your account
// Power by Wayarmy <quanpc294@gmail.com>
package zone

import (
	"log"
	"github.com/spf13/cobra"
	"clf/authen"
)

// listZoneCmd represents the listZone command

// var api = authen.Api
var ListZoneCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all zones in your account",
	Long: `ListZone: 
	List all zones (sites) in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		listZone()
	},
}

func listZone() {
	authen.Login()
	api := authen.Api
	zones, err := api.ListZones()
	if err != nil {
	    log.Fatal(err)
	}
	output := make([][]string, 0, len(zones))
	for _, z := range zones {
		output = append(output, []string{
			z.ID,
			z.Name,
			z.Plan.Name,
			z.Status,
		})
	}
	authen.WriteTable(output, "ID", "Name", "Plan", "Status")
}
package zone
// listZoneCmd represents the listZone command
import (
	"log"
	"clf/authen"
	"github.com/spf13/cobra"
)

var listZoneCmd = &cobra.Command{
	Use:   "listZone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listZone called")
	},
}

func init() {
	zoneCmd.AddCommand(listZoneCmd)
}

var listZoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Action with zones on cloudflare",
	Long: `Action with zones on cloudflare:

Select a domain that you want to manage`,
	Run: func(cmd *cobra.Command, args []string) {
		ListZones()
	},
}

var api = authen.Api
func ListZones() {
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
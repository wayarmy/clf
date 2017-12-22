// List all Record with specific zone
package dns

import (
	"strconv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"clf/authen"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

// listRecordCmd represents the listRecord command
var ListRecordCmd = &cobra.Command{
	Use:   "ls",
	Short: "List records of an existing zone",
	Long: `List records of an existing zone on Cloudflare. For example:
	clf dns ls --zone example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		listRecord()
	},
}

func init() {
	ListRecordCmd.Flags().StringP("zone", "z", "", "Specific zone that you want to list DNS record")

	// Parse arg to viper
	viper.BindPFlag("zoneName", ListRecordCmd.Flags().Lookup("zone"))
}

func listRecord() {
	var ttl string
	authen.Login()
	api := authen.Api

	zoneName := viper.GetString("zoneName")

	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
    	log.Fatal(err)
	}

	recs, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{})
	if err != nil {
	    log.Fatal(err)
	}

	output := make([][]string, 0, len(recs))
	for _, r := range recs {
		if strconv.FormatInt(int64(r.TTL), 10) == "1" {
			ttl = "Automation"
		} else {
			ttl = strconv.FormatInt(int64(r.TTL), 10)
		}
		output = append(output, []string{
			r.ID,
			r.Type,
			r.Name,
			r.Content,
			ttl,
			strconv.FormatBool(r.Proxied),
		})
	}
	authen.WriteTable(output, "ID", "Type", "Name", "Content", "TTL", "Proxied")
}

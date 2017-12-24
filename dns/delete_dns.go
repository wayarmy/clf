// List all Record with specific zone
package dns

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"clf/authen"
	"log"
	"os"
	"fmt"
)

// listRecordCmd represents the listRecord command
var DeleteRecordCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete an exist DNS record",
	Long: `Delete an exist DNS record with specific zonename and record ID. For example:
	clf dns rm --zone example.com --id {RECORD-ID}`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteRecord()
	},
}

func init() {
	DeleteRecordCmd.Flags().StringP("zone", "z", "", "Specific zone that you want to list DNS record")
	DeleteRecordCmd.Flags().StringP("id", "i", "", "Your record id")
	// Parse arg to viper
	viper.BindPFlag("zoneRm", DeleteRecordCmd.Flags().Lookup("zone"))
	viper.BindPFlag("idRm", DeleteRecordCmd.Flags().Lookup("id"))
}

func sendDeleteDNSRequest(zoneID, recordID string) error {
	authen.Login()
	api := authen.Api

	err := api.DeleteDNSRecord(zoneID, recordID)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	return nil
}

func deleteRecord() {
	zoneID, _ := getZoneID(viper.GetString("zoneRm"))
	recs, _ := getListRecord(zoneID)
	record := findRecordByID(recs, viper.GetString("idRm"))
	checkRecordExist(record)

	sendDeleteDNSRequest(zoneID, viper.GetString("idRm"))

	fmt.Println("DNS record name: " + record.Name + " deleted success!!")
}
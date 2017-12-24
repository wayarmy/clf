// Update an existing record with specific zone
package dns

import (
	"github.com/pkg/errors"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"clf/authen"
	"strconv"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

// listRecordCmd represents the listRecord command
var (
	ttl string
	proxied bool
	content string
)
var UpdateRecordCmd = &cobra.Command{
	Use:   "update",
	Short: "Create new record for a zone",
	Long: `Create new record for a zone. For example:
	clf dns create --id xxxyyyzzz --zone example.com --type A --content 1.2.3.4 --ttl 120 --enable-proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		updateRecord()
	},
}

func init() {
	UpdateRecordCmd.Flags().StringP("zone", "z", "", "Specific zone that you want to list DNS record")
	UpdateRecordCmd.Flags().StringP("id", "i", "", "Your record ID")
	UpdateRecordCmd.Flags().StringP("content", "c", "", "Your new record address")
	UpdateRecordCmd.Flags().IntP("ttl", "", 120,"Your record's TTL, default Cloudflare Automation")
	UpdateRecordCmd.Flags().Bool("enable-proxy", false,"Enable proxy for your record, default enabled")

	// Parse arg to viper
	viper.BindPFlag("zone", UpdateRecordCmd.Flags().Lookup("zone"))
	viper.BindPFlag("id", UpdateRecordCmd.Flags().Lookup("id"))
	viper.BindPFlag("content", UpdateRecordCmd.Flags().Lookup("content"))
	viper.BindPFlag("ttl", UpdateRecordCmd.Flags().Lookup("ttl"))
	viper.BindPFlag("enable-proxy", UpdateRecordCmd.Flags().Lookup("enable-proxy"))
}

func checkUpdateRequire() {
	if viper.GetString("zone") == "" || viper.GetString("id") == "" {
		err := errors.Errorf("error: the required flag was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
	return
}

func diffRecord(oldRec cloudflare.DNSRecord) cloudflare.DNSRecord {
	var (
		ttl2 int
		// proxied bool
		content string
	)
	if oldRec.Content != viper.GetString("content") {
		content = viper.GetString("content")
	} else {
		content = oldRec.Content
	}

	// if oldRec.Proxied != viper.GetBool("enable-proxy") && viper.GetBool("enable-proxy") == true {
	// 	proxied = viper.GetBool("enable-proxy")
	// } else {
	// 	proxied = oldRec.Proxied
	// }

	if oldRec.TTL != viper.GetInt("ttl") {
		ttl2 = viper.GetInt("ttl")
	} else {
		ttl2 = oldRec.TTL
	}

	return cloudflare.DNSRecord{
		Name:    oldRec.Name,
		Type:    oldRec.Type,
		Content: content,
		TTL:     ttl2,
		Proxied: viper.GetBool("enable-proxy"),
	}
}

func findRecordByID(records []cloudflare.DNSRecord, id string) cloudflare.DNSRecord{
	for _, r := range records {
		if r.ID == id {
			return cloudflare.DNSRecord{
				Name: r.Name,
				Type: r.Type,
				Content: r.Content,
				TTL: r.TTL,
				Proxied: r.Proxied,
			}
		}
	}
	return cloudflare.DNSRecord{}
}

func convTTL(ttl int) string{
	if ttl == 1 {
		return "Automation"
	} else {
		return strconv.FormatInt(int64(ttl), 10)
	}
}

func checkRecordExist(record cloudflare.DNSRecord){
	if record.Name == "" {
		err := errors.Errorf("Cannot find the record with provide id")
		log.Fatal(err)
		os.Exit(0)
	}
}

func updateRecord() {
	// var ttl2 string
	checkUpdateRequire()
	authen.Login()
	api := authen.Api

	zoneID, errZone := api.ZoneIDByName(viper.GetString("zone"))
	if errZone != nil {
		fmt.Fprintln(os.Stderr, "Error editing DNS record: ", errZone)
		return
	}

	records, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{})
	if err != nil {
	    log.Fatal(err)
	}

	oldRec := findRecordByID(records, viper.GetString("id"))
	checkRecordExist(oldRec)

	record := diffRecord(oldRec)
	errUpdate := api.UpdateDNSRecord(zoneID, viper.GetString("id"), record)
	if errUpdate != nil {
		fmt.Fprintln(os.Stderr, "Error editing DNS record: ", errUpdate)
		return
	}

	output := [][]string{[]string{
		viper.GetString("id"),
		record.Type,
		record.Name,
		record.Content,
		convTTL(record.TTL),
		strconv.FormatBool(record.Proxied),
	}}

	authen.WriteTable(output, "ID", "Type", "Name", "Content", "TTL", "Proxied")
}

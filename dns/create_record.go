// Create new record with specific zone
package dns

import (
	"github.com/pkg/errors"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wayarmy/clf/authen"
	"strconv"
	"github.com/cloudflare/cloudflare-go"
)

// listRecordCmd represents the listRecord command
var CreateRecordCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new record for a zone",
	Long: `Create new record for a zone. For example:
	clf dns create --name dns --zone example.com --type A --content 1.2.3.4 --ttl 120 --enable-proxy true`,
	Run: func(cmd *cobra.Command, args []string) {
		createRecord()
	},
}

func init() {
	CreateRecordCmd.Flags().StringP("zone", "z", "", "Specific zone that you want to list DNS record")
	CreateRecordCmd.Flags().StringP("type", "t", "", "Type of your record")
	CreateRecordCmd.Flags().StringP("name", "n", "", "Your record name")
	CreateRecordCmd.Flags().StringP("content", "c", "", "Your record address")
	CreateRecordCmd.Flags().IntP("ttl", "", 120,"Your record's TTL, default Cloudflare Automation")
	CreateRecordCmd.Flags().Bool("enable-proxy", false,"Enable proxy for your record, default enabled")

	// Parse arg to viper
	viper.BindPFlag("zoneA", CreateRecordCmd.Flags().Lookup("zone"))
	viper.BindPFlag("typeA", CreateRecordCmd.Flags().Lookup("type"))
	viper.BindPFlag("nameA", CreateRecordCmd.Flags().Lookup("name"))
	viper.BindPFlag("contentA", CreateRecordCmd.Flags().Lookup("content"))
	viper.BindPFlag("ttlA", CreateRecordCmd.Flags().Lookup("ttl"))
	viper.BindPFlag("enable-proxyA", CreateRecordCmd.Flags().Lookup("enable-proxy"))
}

func checkFlagsRequirement() {
	if viper.GetString("zoneA") == "" || viper.GetString("typeA") == "" || viper.GetString("nameA") == "" || viper.GetString("contentA") == "" {
		err := errors.Errorf("error: the required flag was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

func createRecord() {
	checkFlagsRequirement()
	authen.Login()
	api := authen.Api

	zoneID, err := api.ZoneIDByName(viper.GetString("zoneA"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating DNS record: ", err)
		return
	}

	record := cloudflare.DNSRecord{
		Name:    viper.GetString("nameA"),
		Type:    viper.GetString("typeA"),
		Content: viper.GetString("contentA"),
		TTL:     viper.GetInt("ttlA"),
		Proxied: viper.GetBool("enable-proxyA"),
	}
	resp, err := api.CreateDNSRecord(zoneID, record)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating DNS record: ", err)
		return
	}

	if resp.Result.TTL == 1 {
		ttl = "Automation"
	} else {
		ttl = strconv.FormatInt(int64(resp.Result.TTL), 10)
	}
	output := [][]string{
		[]string{
			resp.Result.ID,
			resp.Result.Type,
			resp.Result.Name,
			resp.Result.Content,
			ttl,
			strconv.FormatBool(resp.Result.Proxied),
		},
	}
	authen.WriteTable(output, "ID", "Type", "Name", "Content", "TTL", "Proxied")
}

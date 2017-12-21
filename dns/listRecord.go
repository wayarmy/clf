// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dns

import (
	"strconv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"clf/authen"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

var ttl string
// listRecordCmd represents the listRecord command
var ListRecordCmd = &cobra.Command{
	Use:   "ls",
	Short: "List records of an existing zone",
	Long: `List records of an existing zone on Cloudflare. For example:
	clf dns ls --zone example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		list_record()
	},
}

func init() {
	ListRecordCmd.Flags().StringP("zone", "z", "", "Specific zone that you want to list DNS record")

	// Parse arg to viper
	viper.BindPFlag("zoneName", ListRecordCmd.Flags().Lookup("zone"))
}

func list_record() {
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
			r.Type,
			r.Name,
			r.Content,
			ttl,
			strconv.FormatBool(r.Proxied),
		})
	}
	authen.WriteTable(output, "Type", "Name", "Content", "TTL", "Proxied")
}

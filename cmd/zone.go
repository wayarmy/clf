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

package cmd

import (
	// "github.com/spf13/viper"
	"github.com/spf13/cobra"
	"fmt"
	"clf/zone"
)

// zoneCmd represents the zone command
var listZoneCmd = zone.ListZoneCmd
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Action with zones on cloudflare",
	Long: `Action with zones on cloudflare:

Select a domain that you want to manage`,
	Run: func(cmd *cobra.Command, args []string) {
		getZone()
	},
}

func init() {
	rootCmd.AddCommand(zoneCmd)
	zoneCmd.AddCommand(listZoneCmd)
}

func getZone() {
	fmt.Println("Zone")
}
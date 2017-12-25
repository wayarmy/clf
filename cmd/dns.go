// Copyright © 2017 Wayarmy <quanpc294@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wayarmy/clf/dns"
)

// dnsCmd represents the dns command
var ListRecordCmd = dns.ListRecordCmd
var CreateRecordCmd = dns.CreateRecordCmd
var UpdateRecordCmd = dns.UpdateRecordCmd
var DeleteRecordCmd = dns.DeleteRecordCmd
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "A Cloudflare DNS manager CLI",
	Long: `Cloudflare DNS Manager CLI. For example:
	clf dns create --type A --address --zone example.com --enable-cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns called")
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(ListRecordCmd)
	dnsCmd.AddCommand(CreateRecordCmd)
	dnsCmd.AddCommand(UpdateRecordCmd)
	dnsCmd.AddCommand(DeleteRecordCmd)
}

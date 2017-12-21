// Copyright © 2017 Wayarmy <quanpc294@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "A Cloudflare DNS manager CLI",
	Long: `Cloudflare DNS Manager CLI. For example:
	clf dns create record dns --type A --address --zone example.com --enable-cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns called")
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

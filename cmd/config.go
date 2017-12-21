// Copyright © 2017 Wayarmy <quanpc294@gmail.com>
// Configuration of clf-cli
// Functions: create, edit, delete credentials


package cmd

import (
	// "github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"fmt"
	"bufio"
	"os"
	"strings"
	"io/ioutil"
)
var (
	_cfEmail string
	_cfKey string
)

var cfgFile string = "/.clf.yaml"

// zoneCmd represents the zone command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage clf config on your local machine",
	Long: `Manage clf config on your local machine`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigure()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}
func initConfigure() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please specific email address: ")
	_cfEmail, _ = reader.ReadString('\n')
	fmt.Println("Please specific key: ")
	_cfKey, _ = reader.ReadString('\n')
	writeConfig(strings.TrimSpace(_cfKey), strings.TrimSpace(_cfEmail))	
}

func writeConfig(cfEmail string, cfKey string) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := []byte("accounts:\n- EMAIL: "+ cfEmail +"\n  KEY: "+ cfKey)
	err2 := ioutil.WriteFile(home + cfgFile, content, 0644)
	check(err2)
	fmt.Println("\nYour Configuration is created!")
}
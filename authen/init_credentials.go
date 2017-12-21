package authen

import (
	"github.com/pkg/errors"
	"fmt"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/cloudflare/cloudflare-go"
	"log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfApiEmail string
var cfApiKey string
var cfgFile string
var Api *cloudflare.API
func Login () {
	initConfig()
	initCredential()
}

func initCredential() {
	if Api == nil {
		var err error
		Api, err = cloudflare.New(cfApiKey, cfApiEmail)
		if err != nil {
			log.Fatal(err)
		}
	}

	if Api.APIKey == "" {
		fmt.Println("API key not defined")
	}
	if Api.APIEmail == "" {
		fmt.Println("API email not defined")
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".clf" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".clf")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cfApiEmail = viper.Get("EMAIL").(string)
		cfApiKey = viper.Get("KEY").(string)
	} else {
		err := errors.Errorf("Cannot find the credential config, please run `clf config create` to create new credential config")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

func WriteTable(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}
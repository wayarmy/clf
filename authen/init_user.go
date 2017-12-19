package authen

import (
	"fmt"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

var Api *cloudflare.API
func init() {
	if Api == nil {
		var err error
		Api, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
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

func WriteTable(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}
package zone

import (
	"log"
	"github.com/cloudflare/cloudflare-go"
)

func CreateNewZone() {
	listOrg, _, err := api.ListOrganizations()
	var org cloudflare.Organization
	org.ID = listOrg[0].ID
	// fmt.Println(org)
	_, err = api.CreateZone("quanphuong.net", true, org)
	if err != nil {
	    log.Fatal(err)
	}
}
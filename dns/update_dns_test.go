package dns

import (
	"fmt"
	"testing"
	// "github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/cloudflare/cloudflare-go"
	"reflect"
)

func setupViper() {
	viper.Set("zone", "quanphuong.net")
	viper.Set("id", "aad29b7a4334c2574b432a73f574ac22")
	viper.Set("content", "test2.quanphuong.net")
	viper.Set("ttl", 10)
	viper.Set("enable_proxy", true)
}

func setupRecordForTest() cloudflare.DNSRecord {
	return cloudflare.DNSRecord{
		ID: "aad29b7a4334c2574b432a73f574ac22",
		Name:    "test",
		Type:    "CNAME",
		Content: "test.quanphuong.net",
		TTL:     120,
		Proxied: false,
	}
}

func setupListRecords() []cloudflare.DNSRecord {
	r1 := setupRecordForTest()
	output := []cloudflare.DNSRecord{
		cloudflare.DNSRecord{
		ID: "e6ccd4a94611ac26840b3b0c40a9c01f",
		Name:    "test2",
		Type:    "CNAME",
		Content: "test2.quanphuong.net",
		TTL:     60,
		Proxied: true,
	}, r1}
	return output
}

func TestFindRecordByID(t *testing.T) {
	rs := setupListRecords()
	id := "aad29b7a4334c2574b432a73f574ac22"
	r := findRecordByID(rs, id)
	if r.Name != "test" {
		t.Fatalf("Expected Name: %v - Got %v", "test", r.Name)
	}
}

func TestDiffRecord(t *testing.T) {
	setupViper()
	r := setupRecordForTest()
	newR := diffRecord(r)

	expectR := cloudflare.DNSRecord{
		Name: "test",
		Type: "CNAME",
		Content: "test2.quanphuong.net",
		TTL: 10,
		Proxied: false,
	}

	fmt.Println(reflect.TypeOf(newR))
	if newR.Name != expectR.Name {
		t.Fatalf("Expected Name: %v - Got %v", expectR.Name, newR.Name)
	}

	if newR.Type != expectR.Type {
		t.Fatalf("Expected Type: %v - Got %v", expectR.Type, newR.Type)
	}

	if newR.Content != expectR.Content {
		t.Fatalf("Expected Content: %v - Got %v", expectR.Content, newR.Content)
	}

	if newR.TTL != expectR.TTL {
		t.Fatalf("Expected TTL: %v - Got %v", expectR.TTL, newR.TTL)
	}

	if newR.Proxied != expectR.Proxied {
		t.Fatalf("Expected Proxied: %v - Got %v", expectR.Proxied, newR.Proxied)
	}
}
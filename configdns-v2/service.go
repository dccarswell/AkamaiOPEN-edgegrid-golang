package dnsv2

import (
	
//	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/dccarswell/AkamaiOPEN-edgegrid-golang/edgegrid"
	//"akamai/edgegrid"
)

var (
	// Config contains the Akamai OPEN Edgegrid API credentials
	// for automatic signing of requests
	Config edgegrid.Config
)

// Init sets the DNSv2 edgegrid Config
func Init(config edgegrid.Config) {
	Config = config
}

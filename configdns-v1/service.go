package dns

import (
	
//	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/dccarswell/AkamaiOPEN-edgegrid-golang/edgegrid"
//	"akamai/edgegrid"
)

var (
	// Config contains the Akamai OPEN Edgegrid API credentials
	// for automatic signing of requests
	Config edgegrid.Config
)

// Init sets the FastDNS edgegrid Config
func Init(config edgegrid.Config) {
	Config = config
}

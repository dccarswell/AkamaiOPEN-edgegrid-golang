// Package papi provides a simple wrapper for the Akamai Property Manager API
package papi


//import "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
import "github.com/dccarswell/AkamaiOPEN-edgegrid-golang/edgegrid"
//import "akamai/edgegrid"

// Init sets the PAPI edgegrid Config
func Init(config edgegrid.Config) {
	Config = config
}

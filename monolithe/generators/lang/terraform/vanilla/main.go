package main

import (
	"github.com/hashicorp/terraform/plugin"
	nuagenetworks "github.com/terraform-providers/terraform-provider-nuagenetworks/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}

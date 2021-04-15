package main

import (
	. "github.com/yuvarajhere/terraform-provider-appdynamics/appdynamics"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

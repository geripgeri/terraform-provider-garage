// Modified by geripgeri on 2026-11-21.
// Source: AGPL-3.0 (see LICENSE-code and NOTICE.md).

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/geripgeri/terraform-provider-garage/garage"
)

func main() {
	pluginServe(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return garage.Provider()
		},
	})
}

var pluginServe = plugin.Serve

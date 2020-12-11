package main

import (
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/webcerebrium/terraform-provider-openssl/openssl"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openssl.Provider,
	})
}

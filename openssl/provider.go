package openssl

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider creates Terraform resouce provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
		},
		ResourcesMap: map[string]*schema.Resource{
			"openssl_passwd":  resourceOpenSSLPasswd(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

// ProviderConfig contains provider configuration
type ProviderConfig struct {
	Version string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cmd := exec.Command("openssl", "version")
	var out bytes.Buffer
	cmd.Stdout = &out
 	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("OpenSSL must be locally available: %s", err)
	}
	return &ProviderConfig{
		Version: string(out.Bytes()),
	}, nil
}
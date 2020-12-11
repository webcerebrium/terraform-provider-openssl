package openssl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"openssl": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
func TestAccProvider_Simple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testProviderSimple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("openssl_passwd.password", "hash", "$apr1$xxxxxxx$BBUvl1pIHmSzdHA.SvV7n1"),
				),
			},
		},
	})
}

const testProviderSimple = `
provider "openssl" {
}

resource "openssl_passwd" password {
    value = "mysecret"
    algorithm = "apr1"
    salt = "xxxxxxx"
}
`

func testAccPreCheck(t *testing.T) {
	fmt.Println("testAccPreCheck")
}

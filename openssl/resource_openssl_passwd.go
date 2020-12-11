package openssl

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var letters = []rune("01234567890abcdefghijkmnoprstuvwxyz")

func newRandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func resourceOpenSSLPasswd() *schema.Resource {
	return &schema.Resource{
		Create: onPasswdCreate,
		Read:   onPasswdRead,
		Update: onPasswdUpdate,
		Delete: onPasswdDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "unencrypted password - input value",
			},
			"algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(optional) algorithm, apr1 will be default if not specified",
			},
			"salt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(optional) salt for the hash generation",
			},
			"hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "generated hash of the password",
			},
		},
	}
}

func createHash(d *schema.ResourceData) (string, error) {
	args := make([]string, 0)
	args = append(args, "passwd")
	algorithm := "-apr1"
	if d.Get("algorithm") != nil {
		allowed := map[string]int{"1": 1, "5": 1, "6": 1, "apr1": 1, "aixmd5": 1, "crypt": 1}
		if _, ok := allowed[strings.ToLower(d.Get("algorithm").(string))]; !ok {
			return "", fmt.Errorf("Bad Algorithm (allowed: 1,5,6,apr1,aixmd5,crypt)")
		}
		algorithm = "-" + strings.ToLower(d.Get("algorithm").(string))
	}
	args = append(args, algorithm)
	if d.Get("salt") != nil {
		args = append(args, "-salt")
		args = append(args, d.Get("salt").(string))
	}
	args = append(args, d.Get("value").(string))

	cmd := exec.Command("openssl", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("OpenSSL command fauilure %v", err)
	}
	return strings.TrimSpace(string(out.Bytes())), nil
}

func onPasswdCreate(d *schema.ResourceData, m interface{}) error {
	ID := newRandSeq(8)
	d.SetId(ID)
	hash, err := createHash(d)
	if err != nil {
		return fmt.Errorf("Hash creation failure %v", err)
	}
	d.Set("hash", hash)
	return nil
}

func onPasswdRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func onPasswdUpdate(d *schema.ResourceData, m interface{}) error {
	hash, err := createHash(d)
	if err != nil {
		return fmt.Errorf("Hash creation failure %v", err)
	}
	d.Set("hash", hash)
	return nil
}

func onPasswdDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

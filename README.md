:warning: This repo is **DEPRECATED**. We recommend to use [htpasswd from Terraform Registry](https://registry.terraform.io/providers/loafoe/htpasswd) in new projects.

# terraform-provider-openssl

Terraform Provider to wrap `openssl passwd` command

This provider can become handy to including password hashes into you terraform plan.

Its main resource, `openssl_passwd` has password value on input and and hash on output.

## Example Usage
```
provider "openssl" {
}

resource "openssl_passwd" password {
    value = "mysecret"
    algorithm = "apr1"
    salt = "xxxxxxx"
}
```

will produce `openssl_passwd.password.hash`

### License 

MIT

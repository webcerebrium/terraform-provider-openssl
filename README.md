# terraform-provider-openssl-passwd

Terraform Provider to wrap `openssl passwd` command

This provider can become handy to including password hashes into you terraform plan.

Its main resource, `openssl_passwd` has password value on input and and hash on output.

## Example Usage
```
provider "passwd" {
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

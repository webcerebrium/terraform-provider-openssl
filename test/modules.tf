# terraform {
#  required_version = "<= 0.12"
# }

provider "openssl" {
}

resource "openssl_passwd" password {
    value = "mysecret"
    algorithm = "apr1"
    salt = "xxxxxxx"
}

FROM golang:1.15
ARG TF_VERSION=0.12.29

RUN set -x; apt-get update && apt-get -yf install unzip && curl -L -o terraform.zip \
    "https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip" && \
    unzip terraform.zip && mv terraform /usr/bin/terraform
WORKDIR /app
COPY test/modules.tf /app/modules.tf
COPY bin/terraform-provider-openssl /root/.terraform.d/plugins/linux_amd64/terraform-provider-openssl_v0.0.1

CMD ["bash", "-c", "TF_LOG=DEBUG terraform init"]
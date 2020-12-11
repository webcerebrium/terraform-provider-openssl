all: test build install

.PHONY: build test install

test:
	TF_LOG=WARN TF_ACC=true go test -v ./...

build:
	mkdir -p ./bin
	go build -o ./bin/terraform-provider-openssl .

install:
	mkdir -p ~/.terraform.d/plugins/linux_amd64 ./test/.terraform/plugins/linux_amd64
	cp -f ./bin/terraform-provider-openssl ./test/.terraform/plugins/linux_amd64/terraform-provider-openssl_v0.0.1
	cp -f ./bin/terraform-provider-openssl ~/.terraform.d/plugins/terraform-provider-openssl_v0.0.1
	cp -f ./bin/terraform-provider-openssl ~/.terraform.d/plugins/linux_amd64/terraform-provider-openssl_v0.0.1
	# cp -f ./bin/terraform-provider-openssl ${GOPATH}/bin/terraform-provider-openssl_v0.0.1

clean:
	rm -f ~/.terraform.d/plugins/linux_amd64/terraform-provider-openssl_v0.0.1 || true
	rm -f ./bin/terraform-provider-openssl || true
	rm -rf ./test/.terraform || true

test-install:
	docker build . -t tpopenssl && docker run -it --rm tpopenssl
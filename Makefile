NAMESPACE="gastrodon.io/gastrodon"
VERSION=0.0.1
ARCH=amd64
OS=linux

PACKAGE_NAME=database
WEBSITE_REPO=github.com/hashicorp/terraform-website

default: install

# docs:
# 	[ -d "./render/website" ] || git clone https://$(WEBSITE_REPO) ./render/website
# 	@$(MAKE) -C ./render/website \
# 		website-provider \
# 		PROVIDER_PATH=$(shell pwd) \
# 		PROVIDER_NAME=$(PACKAGE_NAME)

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build \
		-ldflags="-w -s" \
		-o ./terraform-provider-database

build-debug:
	go build -o ./terraform-provider-database

install: build
	mkdir -p ~/.terraform.d/plugins/$(NAMESPACE)/database/$(VERSION)/$(OS)_$(ARCH)
	cp ./terraform-provider-database \
		 ~/.terraform.d/plugins/$(NAMESPACE)/database/$(VERSION)/$(OS)_$(ARCH)/terraform-provider-database

install-debug: build-debug install

PKG_NAME=intersight
VERSION=1.0.74
TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
SWAGGER_SPEC=spec/swagger.json
GENERATED_FOLDERS = $(PKG_NAME) models website

export GOOS=$(shell go env GOOS)
export GO_BUILD=env go build
export GO_RUN=env go run
export GO_VET=env go vet
export GO_TEST=env go test
export GO111MODULE=on


default: build

build: fmt fmtcheck
	go get -v golang.org/x/tools/cmd/goimports
	goimports -w ./intersight
	go mod tidy
	go mod vendor
	go install
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o .build/linux_amd64/terraform-provider-intersight_v$(VERSION)
	#GOOS=windows GOARCH=amd64 $(GO_BUILD) -o .build/windows/terraform-provider-intersight_v$(VERSION).exe
	#GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o .build/darwin_amd64/terraform-provider-intersight_v$(VERSION)

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	tfproviderlint ./intersight
	golangci-lint run ./...

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

website:
	ifeq (,$(wildcard ../$(WEBSITE_REPO)))
		echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
		git clone https://$(WEBSITE_REPO) ../$(WEBSITE_REPO)
	endif
		@$(MAKE) -C $../$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/

website-test:
ifeq (,$(wildcard ../$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) ../$(WEBSITE_REPO)
endif
	@$(MAKE) -C ../$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

clean:
	go clean --cache
	rm -rf vendor .build

clobber:
	go clean --cache
	rm -rf $(GENERATED_FOLDERS) vendor $(SWAGGER_SPEC) .build

.PHONY: build test testacc vet fmt fmtcheck lint tools test-compile website website-lint website-test

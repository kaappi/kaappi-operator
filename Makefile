.DEFAULT_GOAL := help

## -- Utility targets --

## Print help message for all Makefile targets
## Run `make` or `make help` to see the help
.PHONY: help
help: ## Credit: https://gist.github.com/prwhite/8168133#gistcomment-2749866

	@printf "Makefile targets usage:\n";

	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                     "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                     "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)

## -- Quality assurance targets --

## Run linters
.PHONY: lint
lint:
	gometalinter $(shell go list ./...) --deadline=2m  --disable-all --enable=errcheck --enable=ineffassign \
	--enable=gofmt --enable=vet --enable=deadcode --enable=varcheck \
	--enable=structcheck --enable=maligned --enable=unconvert \
	--enable=goconst --enable=gosec --enable=unparam --enable=staticcheck \
	--enable=interfacer --enable=vetshadow --enable=golint

## Run unit tests
.PHONY: test
test:
	go test $(shell go list ./... | grep -v /test/e2e) -v -race -failfast -coverprofile=cover.out -covermode=atomic -outputdir=.

## Run e2e tests
test-e2e:
	operator-sdk test local ./test/e2e --namespace operator-test --up-local --go-test-flags "-v"

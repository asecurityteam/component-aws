TAG := $(shell git rev-parse --short HEAD)
DIR := $(shell pwd -L)
SDCLI_VERSION :=v1.4.0
SDCLI=docker run --rm -v "$(DIR):$(DIR)" -w "$(DIR)" asecurityteam/sdcli:$(SDCLI_VERSION)


dep:
	$(SDCLI) go dep

lint:
	$(SDCLI) go lint

test:
	$(SDCLI) go test

integration:
	DIR=$(DIR) \
	PROJECT_PATH=/go/src/$(PROJECT_PATH) \
	docker-compose \
		-f docker-compose.it.yml \
		up \
			--abort-on-container-exit \
			--build \
			--exit-code-from test

coverage:
	$(SDCLI) go coverage

doc: ;

build-dev: ;

build: ;

run: ;

deploy-dev: ;

deploy: ;

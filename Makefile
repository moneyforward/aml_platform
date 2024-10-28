PROJ_DIR:=$(shell pwd)
export GOBIN := $(PROJ_DIR)/bin

OAPI_DIR:=$(PROJ_DIR)/openapi

## Generate golang source code from openapi file.
gen-from-oapi:
	cd $(OAPI_DIR) && make gen-from-oapi

## Install necessary tools for local development
install-tools:
	@$(GOMOD) /bin/bash ./tools/install_tools.sh

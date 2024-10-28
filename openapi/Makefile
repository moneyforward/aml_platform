PROJ_DIR := $(PWD)/..

GOBIN?=$(PROJ_DIR)/bin
export GOBIN:=$(GOBIN)
export PATH:=$(GOBIN):$(PATH)

OAPI_VERSION         :=v1
NORMAL_OAPI_DIR      := $(PWD)
NORMAL_OAPI_FILE     := $(NORMAL_OAPI_DIR)/openapi.json
TMP_WORKING_OAPI_DIR := ${NORMAL_OAPI_DIR}/tmp
TMP_WORKING_OAPI_FILE := ${TMP_WORKING_OAPI_DIR}/openapi.json
OAPI_GEN_DST_DIR     := $(PROJ_DIR)/app/openapi/$(OAPI_VERSION)
OAPI_GO_PACKAGE      := openapi

# Generate go files from openapi
gen-from-oapi: prepare-working-dir gen-oapi-from-working-dir clean-working-dir

# Prepare a temporary working dir with a clone file of openapi.yml
prepare-working-dir:
	make clean-working-dir
	mkdir ${TMP_WORKING_OAPI_DIR}
	cp ${NORMAL_OAPI_FILE} ${TMP_WORKING_OAPI_FILE}

# Generate .go code from the edited openapi.yml
gen-oapi-from-working-dir:
	$(call gen_oapi_server_helper, $(OAPI_GEN_DST_DIR),aml_api, $(OAPI_GO_PACKAGE), $(TMP_WORKING_OAPI_FILE))

define gen_oapi_server_helper
	rm -f $(1)/$(2)_*.gen.go
	$(GOBIN)/oapi-codegen -package $(3) -generate spec   $(4) > $(1)/$(2)_spec.gen.go
	$(GOBIN)/oapi-codegen -package $(3) -generate types  $(4) > $(1)/$(2)_types.gen.go
	$(GOBIN)/oapi-codegen -package $(3) -generate server $(4) > $(1)/$(2)_server.gen.go
endef

clean-working-dir:
	rm -rf ${TMP_WORKING_OAPI_DIR}

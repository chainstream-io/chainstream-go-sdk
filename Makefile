#
# Project Metadata
# These act as inputs to other makefile targets.
#
ORG := chainstream-io
PROJECT := chainstream-go-sdk
VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)

#
# Tool Prerequisites Check
# This ensures that you have the necessary executables installed to run this makefile.
#
BUILD_PREREQUISITES = git go
VALIDATION_PREREQUISITES = golangci-lint

#
# Build Options
# Typical inputs to the build targets found below.
#
TARGET=target
BIN=$(TARGET)/bin

.PHONY: usage
usage:
	@ echo "Usage: make [`cat Makefile | grep "^[A-z\%\-]*:" | awk '{print $$1}' | sed "s/://g" | sed "s/%/[1-3]/g" | xargs`]"

.PHONY: info
info:
	@ echo ORG: $(ORG)
	@ echo PROJECT: $(PROJECT)
	@ echo VERSION: $(VERSION)
	@ echo BUILD: $(BUILD)

#
# Build Targets
# These are the primary developer workflow targets for building the software.
#

.PHONY: clean
clean: info
	@ rm -rf target
	@ go mod tidy

.PHONY: build_deps
build_deps: info clean
	@ printf $(foreach exec,$(BUILD_PREREQUISITES), \
        $(if $(shell which $(exec)),"", \
        $(error "No $(exec) in PATH. Prerequisites are: $(BUILD_PREREQUISITES)")))

.PHONY: validation_deps
validation_deps: info clean
	@ printf $(foreach exec,$(VALIDATION_PREREQUISITES), \
        $(if $(shell which $(exec)),"", \
        $(error "No $(exec) in PATH. Prerequisites are: $(VALIDATION_PREREQUISITES)")))

.PHONY: lint
lint: clean validation_deps
	@ printf "\nLint App\n"
	@golangci-lint run --timeout=5m --config=.golangci.yaml ./...

.PHONY: lint-fix
lint-fix: clean validation_deps
	@ printf "\nFixing lint issues\n"
	@golangci-lint run --timeout=5m --config=.golangci.yaml --fix ./...

.PHONY: docs
docs:
	@printf "\nGenerating docs\n"
	@golds -gen -dir=./docs -emphasize-wdpkgs

.PHONY: test
test: build_deps
	@ printf "\nRunning Go tests\n"
	@ go test -v ./...

# Default client generation (modular)
.PHONY: client
client: client-execution

# Module directories (under openapi/)
MODULES := bridgeaggregator swapaggregator blockchain dex dexpool ipfs job kyt prediction ranking redpacket token trade transaction wallet watchlist webhook
OPENAPI_DIR := ./openapi

# OpenAPI preprocessed file
OPENAPI_PREPROCESSED := ../openapi-preprocessed.yaml

OAPI_CODEGEN_VERSION := v2.7.0

# Generate all clients by module
.PHONY: client-modules
client-modules:
	@ oapi-codegen -version | grep -q "$(OAPI_CODEGEN_VERSION)" || { echo "oapi-codegen $(OAPI_CODEGEN_VERSION) required"; exit 1; }
	@ printf "\nGenerating modular clients...\n"
	@ cd ../python && uv run python ../scripts/preprocess_openapi.py ../openapi.yaml ../openapi-preprocessed.yaml
	@ $(foreach mod,$(MODULES),mkdir -p $(OPENAPI_DIR)/$(mod);)
	@ oapi-codegen -package bridgeaggregator -generate types,client -include-tags "Bridge Aggregator" -o $(OPENAPI_DIR)/bridgeaggregator/client.gen.go $(OPENAPI_PREPROCESSED)
	@ oapi-codegen -package swapaggregator -generate types,client -include-tags "Swap Aggregator" -o $(OPENAPI_DIR)/swapaggregator/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating blockchain client...\n"
	@ oapi-codegen -package blockchain -generate types,client -include-tags Blockchain -o $(OPENAPI_DIR)/blockchain/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating dex client...\n"
	@ oapi-codegen -package dex -generate types,client -include-tags Dex -o $(OPENAPI_DIR)/dex/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating dexpool client...\n"
	@ oapi-codegen -package dexpool -generate types,client -include-tags DexPool -o $(OPENAPI_DIR)/dexpool/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating ipfs client...\n"
	@ oapi-codegen -package ipfs -generate types,client -include-tags IPFS -o $(OPENAPI_DIR)/ipfs/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating job client...\n"
	@ oapi-codegen -package job -generate types,client -include-tags Job -o $(OPENAPI_DIR)/job/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating kyt client...\n"
	@ oapi-codegen -package kyt -generate types,client -include-tags KYT -o $(OPENAPI_DIR)/kyt/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating prediction client...\n"
	@ oapi-codegen -package prediction -generate types,client -include-tags Prediction -o $(OPENAPI_DIR)/prediction/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating ranking client...\n"
	@ oapi-codegen -package ranking -generate types,client -include-tags Ranking -o $(OPENAPI_DIR)/ranking/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating redpacket client...\n"
	@ oapi-codegen -package redpacket -generate types,client -include-tags RedPacket -o $(OPENAPI_DIR)/redpacket/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating token client...\n"
	@ oapi-codegen -package token -generate types,client -include-tags Token -o $(OPENAPI_DIR)/token/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating trade client...\n"
	@ oapi-codegen -package trade -generate types,client -include-tags Trade -o $(OPENAPI_DIR)/trade/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating transaction client...\n"
	@ oapi-codegen -package transaction -generate types,client -include-tags Transaction -o $(OPENAPI_DIR)/transaction/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating wallet client...\n"
	@ oapi-codegen -package wallet -generate types,client -include-tags Wallet -o $(OPENAPI_DIR)/wallet/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating watchlist client...\n"
	@ oapi-codegen -package watchlist -generate types,client -include-tags Watchlist -o $(OPENAPI_DIR)/watchlist/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating webhook client...\n"
	@ oapi-codegen -package webhook -generate types,client -include-tags Webhook -o $(OPENAPI_DIR)/webhook/client.gen.go $(OPENAPI_PREPROCESSED)
	@ cd ../python && uv run python ../scripts/fix_execution_go_unions.py ../go/openapi
	@ gofmt -w openapi/dex/client.gen.go openapi/job/client.gen.go openapi/transaction/client.gen.go
	@ printf "\nAll modular clients generated successfully!\n"

# Clean generator output only; preserve custom contract decoders and tests.
.PHONY: clean-generated
clean-generated:
	@ $(foreach mod,$(MODULES),rm -f $(OPENAPI_DIR)/$(mod)/client.gen.go;)

# Scope-aware execution update: other modules retain their compatibility baseline.
.PHONY: client-execution
client-execution:
	@ oapi-codegen -version | grep -q "$(OAPI_CODEGEN_VERSION)" || { echo "oapi-codegen $(OAPI_CODEGEN_VERSION) required"; exit 1; }
	@ cd ../python && uv run python ../scripts/preprocess_openapi.py ../openapi.yaml ../openapi-preprocessed.yaml
	@ mkdir -p openapi/bridgeaggregator openapi/swapaggregator
	@ oapi-codegen -package dex -generate types,client -include-tags Dex -o openapi/dex/client.gen.go $(OPENAPI_PREPROCESSED)
	@ oapi-codegen -package job -generate types,client -include-tags Job -o openapi/job/client.gen.go $(OPENAPI_PREPROCESSED)
	@ oapi-codegen -package transaction -generate types,client -include-tags Transaction -o openapi/transaction/client.gen.go $(OPENAPI_PREPROCESSED)
	@ oapi-codegen -package bridgeaggregator -generate types,client -include-tags "Bridge Aggregator" -o openapi/bridgeaggregator/client.gen.go $(OPENAPI_PREPROCESSED)
	@ oapi-codegen -package swapaggregator -generate types,client -include-tags "Swap Aggregator" -o openapi/swapaggregator/client.gen.go $(OPENAPI_PREPROCESSED)
	@ cd ../python && uv run python ../scripts/fix_execution_go_unions.py ../go/openapi
	@ gofmt -w openapi/dex/client.gen.go openapi/job/client.gen.go openapi/transaction/client.gen.go

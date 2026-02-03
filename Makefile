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
client: client-modules

# Module directories (under openapi/)
MODULES := jobs ipfs blockchain transaction dex defi_moonshot defi_pumpfun trade token dexpool ranking wallet redpacket watchlist kyt endpoint
OPENAPI_DIR := ./openapi

# OpenAPI preprocessed file
OPENAPI_PREPROCESSED := ../openapi-preprocessed.yaml

# Generate all clients by module
.PHONY: client-modules
client-modules: clean-generated
	@ printf "\nGenerating modular clients...\n"
	@ python3 ../scripts/preprocess_openapi.py ../openapi.yaml $(OPENAPI_PREPROCESSED)
	@ $(foreach mod,$(MODULES),mkdir -p $(OPENAPI_DIR)/$(mod);)
	@ printf "Generating jobs client...\n"
	@ oapi-codegen -package jobs -generate types,client -include-tags Jobs -o $(OPENAPI_DIR)/jobs/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating ipfs client...\n"
	@ oapi-codegen -package ipfs -generate types,client -include-tags Ipfs -o $(OPENAPI_DIR)/ipfs/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating blockchain client...\n"
	@ oapi-codegen -package blockchain -generate types,client -include-tags Blockchain -o $(OPENAPI_DIR)/blockchain/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating transaction client...\n"
	@ oapi-codegen -package transaction -generate types,client -include-tags Transaction -o $(OPENAPI_DIR)/transaction/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating dex client...\n"
	@ oapi-codegen -package dex -generate types,client -include-tags Dex -o $(OPENAPI_DIR)/dex/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating defi_moonshot client...\n"
	@ oapi-codegen -package defi_moonshot -generate types,client -include-tags "Defi/sol/Moonshot" -o $(OPENAPI_DIR)/defi_moonshot/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating defi_pumpfun client...\n"
	@ oapi-codegen -package defi_pumpfun -generate types,client -include-tags "Defi/sol/Pumpfun" -o $(OPENAPI_DIR)/defi_pumpfun/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating trade client...\n"
	@ oapi-codegen -package trade -generate types,client -include-tags Trade -o $(OPENAPI_DIR)/trade/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating token client...\n"
	@ oapi-codegen -package token -generate types,client -include-tags Token -o $(OPENAPI_DIR)/token/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating dexpool client...\n"
	@ oapi-codegen -package dexpool -generate types,client -include-tags DexPool -o $(OPENAPI_DIR)/dexpool/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating ranking client...\n"
	@ oapi-codegen -package ranking -generate types,client -include-tags Ranking -o $(OPENAPI_DIR)/ranking/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating wallet client...\n"
	@ oapi-codegen -package wallet -generate types,client -include-tags Wallet -o $(OPENAPI_DIR)/wallet/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating redpacket client...\n"
	@ oapi-codegen -package redpacket -generate types,client -include-tags RedPacket -o $(OPENAPI_DIR)/redpacket/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating watchlist client...\n"
	@ oapi-codegen -package watchlist -generate types,client -include-tags Watchlist -o $(OPENAPI_DIR)/watchlist/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating kyt client...\n"
	@ oapi-codegen -package kyt -generate types,client -include-tags KYT -o $(OPENAPI_DIR)/kyt/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "Generating endpoint client...\n"
	@ oapi-codegen -package endpoint -generate types,client -include-tags Endpoint -o $(OPENAPI_DIR)/endpoint/client.gen.go $(OPENAPI_PREPROCESSED)
	@ printf "\nAll modular clients generated successfully!\n"

# Clean all generated files
.PHONY: clean-generated
clean-generated:
	@ rm -rf $(OPENAPI_DIR)/*

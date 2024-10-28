#!/bin/bash
# this script is used to install all dependent tools
declare -a tools=(
  "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1"
)

for package in "${tools[@]}"; do
  echo "installing" "$package"
  go install $package
done

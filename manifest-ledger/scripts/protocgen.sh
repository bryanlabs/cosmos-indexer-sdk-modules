#!/usr/bin/env bash

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    # this regex checks if a proto file has its go_package set to github.com/liftedinit/manifest-ledger/...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*github.com/liftedinit/manifest-ledger/api' "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo.yaml "$file"
    fi
  done
done

cd ..

cp -r github.com/liftedinit/manifest-ledger/* ./
rm -rf github.com

# Copy files over for dep injection
rm -rf api && mkdir api
custom_modules=$(find . -name 'module' -type d -not -path "./proto/*")
for module in $custom_modules; do
  dirPath=$(dirname $module)
  mkdir -p api/$dirPath

  mv $dirPath/* ./api/$dirPath/

  # incorrect ref
  find api/$dirPath -type f -name '*.go' -exec sed -i '' -e 's|types "github.com/cosmos/cosmos-sdk/types"|types "cosmossdk.io/api/cosmos/base/v1beta1"|g' {} \;

  rm -rf $dirPath
done
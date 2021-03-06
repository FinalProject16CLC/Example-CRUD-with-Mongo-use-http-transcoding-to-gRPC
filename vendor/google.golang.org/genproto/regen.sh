#!/bin/bash
#
# Copyright 2016 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script rebuilds the generated code for the protocol buffers.
# To run this you will need protoc and goprotobuf installed;
# see https://github.com/golang/protobuf for instructions.
# You also need Go and Git installed.

set -e

PKG=google.golang.org/genproto
PROTO_REPO=https://github.com/google/protobuf
GOOGLEAPIS_REPO=https://github.com/googleapis/googleapis

function die() {
  echo 1>&2 $*
  exit 1
}

# Sanity check that the right tools are accessible.
for tool in go git protoc protoc-gen-go; do
  q=$(which $tool) || die "didn't find $tool"
  echo 1>&2 "$tool: $q"
done

if [ -z "$PROTOBUF" ]; then
  proto_repo_dir=$(mktemp -d -t regen-cds-proto.XXXXXX)
  git clone $PROTO_REPO $proto_repo_dir
  remove_dirs="$proto_repo_dir"
  # The protoc include directory is actually the "src" directory of the repo.
  protodir="$proto_repo_dir/src"
else
  protodir="$PROTOBUF/src"
fi

if [ -z "$GOOGLEAPIS" ]; then
  apidir=$(mktemp -d -t regen-cds-api.XXXXXX)
  git clone $GOOGLEAPIS_REPO $apidir
  remove_dirs="$remove_dirs $apidir"
else
  apidir="$GOOGLEAPIS"
fi

wait

# Nuke everything, we'll generate them back
rm -r googleapis protobuf generated || true

mkdir generated
go run regen.go -go_out generated -pkg_prefix "$PKG" "$apidir" "$protodir"
mv generated/google.golang.org/genproto/googleapis googleapis
mv generated/google.golang.org/genproto/protobuf protobuf
rm -rf generated

# throw away changes to some special libs
for d in "googleapis/grafeas/v1" "googleapis/devtools/containeranalysis/v1"; do
  git checkout $d
  git clean -df $d
done

# Sanity check the build.
echo 1>&2 "Checking that the libraries build..."
go build -v ./...

gofmt -s -l -w . && goimports -w .

echo 1>&2 "All done!"

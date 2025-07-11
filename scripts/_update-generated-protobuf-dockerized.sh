#!/usr/bin/env bash

# Copyright 2015 The Kubernetes Authors.
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

# This script genertates `*/api.pb.go` from the protobuf file `*/api.proto`.
# Usage:
#     scripts/update-generated-protobuf-dockerized.sh "${APIROOTS[@]}"
#     An example APIROOT is: "k8s.io/api/admissionregistration/v1"

set -o errexit
set -o nounset
set -o pipefail

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/protoc.sh"

onex::protoc::check_protoc
onex::golang::setup_env

if ! command -v go-to-protobuf &> /dev/null ; then
  GOPROXY=off go install k8s.io/code-generator/cmd/go-to-protobuf
fi

if ! command -v protoc-gen-gogo &> /dev/null ; then
  GOPROXY=off go install k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo
fi

# requires the 'proto' tag to build (will remove when ready)
# searches for the protoc-gen-gogo extension in the output directory
# satisfies import of github.com/gogo/protobuf/gogoproto/gogo.proto and the
# core Google protobuf types
PATH="${PROJ_ROOT_DIR}/_output/bin:${PATH}" \
  go-to-protobuf \
  -v "${KUBE_VERBOSE}" \
  --go-header-file "${PROJ_ROOT_DIR}/scripts/boilerplate/boilerplate.generatego.txt" \
  --proto-import="${PROJ_ROOT_DIR}/third_party/protobuf" \
  --proto-import="${GOPATH}/src" \
  --output-dir="${GOPATH}/src" \
  --apimachinery-packages '-k8s.io/apimachinery/pkg/util/intstr,-k8s.io/apimachinery/pkg/api/resource,-k8s.io/apimachinery/pkg/runtime/schema,-k8s.io/apimachinery/pkg/runtime,-k8s.io/apimachinery/pkg/apis/meta/v1,-k8s.io/apimachinery/pkg/apis/meta/v1beta1,-k8s.io/apimachinery/pkg/apis/testapigroup/v1' \
  --packages="$(IFS=, ; echo "$*")"

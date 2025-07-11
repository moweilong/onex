# ==============================================================================
# Versions used by all Makefiles
#

KUSTOMIZE_VERSION ?= v3.8.7
CONTROLLER_TOOLS_VERSION ?= v0.9.2
CONTROLLER_TOOLS_VERSION ?= v0.8.0
HADOLINT_VER ?= v2.10.0
CHART_VERSION ?= 1.0.0
#KIND_VERSION ?= v0.22.0
KIND_VERSION ?= v0.19.0
HELM_VERSION ?= v3.15.1
GO_TESTS_SUM_VERSION ?=v1.6.4
GOLANGCI_LINT_VERSION := v1.55.2
YQ_VERSION ?= v4.25.2
GRPCURL_VERSION ?= v1.8.9
LOGCHECK_VERSION ?= v0.8.1
GO_FUMPT_VERSION ?= v0.5.0
GO_MODIFY_TAGS_VERSION ?= v1.16.0
GO_THANKS_VERSION ?= V0.5.0
LICENSE_VERSION ?= v5.0.4
AIR_VERSION ?= v1.49.0
PROTOC_GO_INJECT_TAG_VERSION ?= v1.4.0
DB_TO_STRUCT_VERSION ?= v1.0.2
GEN_TOOL_VERSION ?= v0.3.25
HELM_DOCS_VERSION ?= v1.12.0
KUBE_CONFORM_VERSION ?= v0.6.4
KUBE_LINTER_VERSION ?= v0.6.6
GO_SWAGGER_VERSION ?= v0.31.0
GO_JUNIT_REPORT_VERSION ?= v2.1.0
GO_TESTS_VERSION ?= v1.6.0
GO_IMPORTS_VERSION ?= v0.14.2
GO_GIT_LINT_VERSION ?= v1.1.1
GSEMVER_VERSION ?= v0.9.1
GIT_CHGLOG_VERSION ?= v0.15.4
ADDLICENSE_VERSION ?= v0.0.2
KUSTOMIZE_VERSION ?= v5.3.0
GO_APIDIFF_VERSION ?= v0.8.2
PROTOC_GEN_OPENAPI_VERSION ?= v0.7.0
BUF_VERSION ?= v1.28.1
KAFKACTL_VERSION ?= v4.0.0
GRPC_GATEWAY_VERSION ?= $(call get_go_version,github.com/grpc-ecosystem/grpc-gateway/v2)
KRATOS_VERSION ?= $(call get_go_version,github.com/go-kratos/kratos/v2)
PROTOC_GEN_VALIDATE_VERSION ?= $(call get_go_version,github.com/envoyproxy/protoc-gen-validate)
PROTOC_GEN_GO_VERSION ?= $(call get_go_version,google.golang.org/protobuf)
PROTOC_GEN_GO_GRPC_VERSION ?= $(call get_go_version,google.golang.org/grpc)
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
WIRE_VERSION ?= $(call get_go_version,github.com/google/wire)
MOCKGEN_VERSION ?= $(call get_go_version,github.com/golang/mock)

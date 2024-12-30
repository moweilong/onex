// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeConfigSourceApplyConfiguration represents a declarative configuration of the NodeConfigSource type for use
// with apply.
type NodeConfigSourceApplyConfiguration struct {
	ConfigMap *ConfigMapNodeConfigSourceApplyConfiguration `json:"configMap,omitempty"`
}

// NodeConfigSourceApplyConfiguration constructs a declarative configuration of the NodeConfigSource type for use with
// apply.
func NodeConfigSource() *NodeConfigSourceApplyConfiguration {
	return &NodeConfigSourceApplyConfiguration{}
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *NodeConfigSourceApplyConfiguration) WithConfigMap(value *ConfigMapNodeConfigSourceApplyConfiguration) *NodeConfigSourceApplyConfiguration {
	b.ConfigMap = value
	return b
}

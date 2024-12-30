// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeDaemonEndpointsApplyConfiguration represents a declarative configuration of the NodeDaemonEndpoints type for use
// with apply.
type NodeDaemonEndpointsApplyConfiguration struct {
	KubeletEndpoint *DaemonEndpointApplyConfiguration `json:"kubeletEndpoint,omitempty"`
}

// NodeDaemonEndpointsApplyConfiguration constructs a declarative configuration of the NodeDaemonEndpoints type for use with
// apply.
func NodeDaemonEndpoints() *NodeDaemonEndpointsApplyConfiguration {
	return &NodeDaemonEndpointsApplyConfiguration{}
}

// WithKubeletEndpoint sets the KubeletEndpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubeletEndpoint field is set to the value of the last call.
func (b *NodeDaemonEndpointsApplyConfiguration) WithKubeletEndpoint(value *DaemonEndpointApplyConfiguration) *NodeDaemonEndpointsApplyConfiguration {
	b.KubeletEndpoint = value
	return b
}

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// LoadBalancerIngressApplyConfiguration represents a declarative configuration of the LoadBalancerIngress type for use
// with apply.
type LoadBalancerIngressApplyConfiguration struct {
	IP       *string                        `json:"ip,omitempty"`
	Hostname *string                        `json:"hostname,omitempty"`
	IPMode   *corev1.LoadBalancerIPMode     `json:"ipMode,omitempty"`
	Ports    []PortStatusApplyConfiguration `json:"ports,omitempty"`
}

// LoadBalancerIngressApplyConfiguration constructs a declarative configuration of the LoadBalancerIngress type for use with
// apply.
func LoadBalancerIngress() *LoadBalancerIngressApplyConfiguration {
	return &LoadBalancerIngressApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithIP(value string) *LoadBalancerIngressApplyConfiguration {
	b.IP = &value
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithHostname(value string) *LoadBalancerIngressApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithIPMode sets the IPMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPMode field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithIPMode(value corev1.LoadBalancerIPMode) *LoadBalancerIngressApplyConfiguration {
	b.IPMode = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *LoadBalancerIngressApplyConfiguration) WithPorts(values ...*PortStatusApplyConfiguration) *LoadBalancerIngressApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

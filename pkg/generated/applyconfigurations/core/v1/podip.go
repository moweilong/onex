// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PodIPApplyConfiguration represents a declarative configuration of the PodIP type for use
// with apply.
type PodIPApplyConfiguration struct {
	IP *string `json:"ip,omitempty"`
}

// PodIPApplyConfiguration constructs a declarative configuration of the PodIP type for use with
// apply.
func PodIP() *PodIPApplyConfiguration {
	return &PodIPApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *PodIPApplyConfiguration) WithIP(value string) *PodIPApplyConfiguration {
	b.IP = &value
	return b
}

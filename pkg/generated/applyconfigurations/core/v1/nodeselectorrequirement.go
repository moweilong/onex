// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// NodeSelectorRequirementApplyConfiguration represents a declarative configuration of the NodeSelectorRequirement type for use
// with apply.
type NodeSelectorRequirementApplyConfiguration struct {
	Key      *string                  `json:"key,omitempty"`
	Operator *v1.NodeSelectorOperator `json:"operator,omitempty"`
	Values   []string                 `json:"values,omitempty"`
}

// NodeSelectorRequirementApplyConfiguration constructs a declarative configuration of the NodeSelectorRequirement type for use with
// apply.
func NodeSelectorRequirement() *NodeSelectorRequirementApplyConfiguration {
	return &NodeSelectorRequirementApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *NodeSelectorRequirementApplyConfiguration) WithKey(value string) *NodeSelectorRequirementApplyConfiguration {
	b.Key = &value
	return b
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *NodeSelectorRequirementApplyConfiguration) WithOperator(value v1.NodeSelectorOperator) *NodeSelectorRequirementApplyConfiguration {
	b.Operator = &value
	return b
}

// WithValues adds the given value to the Values field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Values field.
func (b *NodeSelectorRequirementApplyConfiguration) WithValues(values ...string) *NodeSelectorRequirementApplyConfiguration {
	for i := range values {
		b.Values = append(b.Values, values[i])
	}
	return b
}

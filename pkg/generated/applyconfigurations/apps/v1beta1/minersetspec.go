// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/meta/v1"
)

// MinerSetSpecApplyConfiguration represents a declarative configuration of the MinerSetSpec type for use
// with apply.
type MinerSetSpecApplyConfiguration struct {
	Replicas                *int32                               `json:"replicas,omitempty"`
	Selector                *v1.LabelSelectorApplyConfiguration  `json:"selector,omitempty"`
	Template                *MinerTemplateSpecApplyConfiguration `json:"template,omitempty"`
	DisplayName             *string                              `json:"displayName,omitempty"`
	DeletePolicy            *string                              `json:"deletePolicy,omitempty"`
	MinReadySeconds         *int32                               `json:"minReadySeconds,omitempty"`
	ProgressDeadlineSeconds *int32                               `json:"progressDeadlineSeconds,omitempty"`
}

// MinerSetSpecApplyConfiguration constructs a declarative configuration of the MinerSetSpec type for use with
// apply.
func MinerSetSpec() *MinerSetSpecApplyConfiguration {
	return &MinerSetSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithReplicas(value int32) *MinerSetSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *MinerSetSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithTemplate(value *MinerTemplateSpecApplyConfiguration) *MinerSetSpecApplyConfiguration {
	b.Template = value
	return b
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithDisplayName(value string) *MinerSetSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithDeletePolicy sets the DeletePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletePolicy field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithDeletePolicy(value string) *MinerSetSpecApplyConfiguration {
	b.DeletePolicy = &value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithMinReadySeconds(value int32) *MinerSetSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithProgressDeadlineSeconds sets the ProgressDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProgressDeadlineSeconds field is set to the value of the last call.
func (b *MinerSetSpecApplyConfiguration) WithProgressDeadlineSeconds(value int32) *MinerSetSpecApplyConfiguration {
	b.ProgressDeadlineSeconds = &value
	return b
}

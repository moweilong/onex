// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// EvaluateSpecApplyConfiguration represents an declarative configuration of the EvaluateSpec type for use
// with apply.
type EvaluateSpecApplyConfiguration struct {
	DisplayName *string `json:"displayName,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	ModelID     *int64  `json:"modelID,omitempty"`
	SampleID    *int64  `json:"sampleID,omitempty"`
}

// EvaluateSpecApplyConfiguration constructs an declarative configuration of the EvaluateSpec type for use with
// apply.
func EvaluateSpec() *EvaluateSpecApplyConfiguration {
	return &EvaluateSpecApplyConfiguration{}
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *EvaluateSpecApplyConfiguration) WithDisplayName(value string) *EvaluateSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *EvaluateSpecApplyConfiguration) WithProvider(value string) *EvaluateSpecApplyConfiguration {
	b.Provider = &value
	return b
}

// WithModelID sets the ModelID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ModelID field is set to the value of the last call.
func (b *EvaluateSpecApplyConfiguration) WithModelID(value int64) *EvaluateSpecApplyConfiguration {
	b.ModelID = &value
	return b
}

// WithSampleID sets the SampleID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleID field is set to the value of the last call.
func (b *EvaluateSpecApplyConfiguration) WithSampleID(value int64) *EvaluateSpecApplyConfiguration {
	b.SampleID = &value
	return b
}

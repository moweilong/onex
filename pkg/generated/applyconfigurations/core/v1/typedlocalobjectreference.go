// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// TypedLocalObjectReferenceApplyConfiguration represents an declarative configuration of the TypedLocalObjectReference type for use
// with apply.
type TypedLocalObjectReferenceApplyConfiguration struct {
	APIGroup *string `json:"apiGroup,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// TypedLocalObjectReferenceApplyConfiguration constructs an declarative configuration of the TypedLocalObjectReference type for use with
// apply.
func TypedLocalObjectReference() *TypedLocalObjectReferenceApplyConfiguration {
	return &TypedLocalObjectReferenceApplyConfiguration{}
}

// WithAPIGroup sets the APIGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIGroup field is set to the value of the last call.
func (b *TypedLocalObjectReferenceApplyConfiguration) WithAPIGroup(value string) *TypedLocalObjectReferenceApplyConfiguration {
	b.APIGroup = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *TypedLocalObjectReferenceApplyConfiguration) WithKind(value string) *TypedLocalObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TypedLocalObjectReferenceApplyConfiguration) WithName(value string) *TypedLocalObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}

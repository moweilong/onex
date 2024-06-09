// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SecretVolumeSourceApplyConfiguration represents an declarative configuration of the SecretVolumeSource type for use
// with apply.
type SecretVolumeSourceApplyConfiguration struct {
	SecretName  *string                       `json:"secretName,omitempty"`
	Items       []KeyToPathApplyConfiguration `json:"items,omitempty"`
	DefaultMode *int32                        `json:"defaultMode,omitempty"`
	Optional    *bool                         `json:"optional,omitempty"`
}

// SecretVolumeSourceApplyConfiguration constructs an declarative configuration of the SecretVolumeSource type for use with
// apply.
func SecretVolumeSource() *SecretVolumeSourceApplyConfiguration {
	return &SecretVolumeSourceApplyConfiguration{}
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *SecretVolumeSourceApplyConfiguration) WithSecretName(value string) *SecretVolumeSourceApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithItems adds the given value to the Items field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Items field.
func (b *SecretVolumeSourceApplyConfiguration) WithItems(values ...*KeyToPathApplyConfiguration) *SecretVolumeSourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithItems")
		}
		b.Items = append(b.Items, *values[i])
	}
	return b
}

// WithDefaultMode sets the DefaultMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultMode field is set to the value of the last call.
func (b *SecretVolumeSourceApplyConfiguration) WithDefaultMode(value int32) *SecretVolumeSourceApplyConfiguration {
	b.DefaultMode = &value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *SecretVolumeSourceApplyConfiguration) WithOptional(value bool) *SecretVolumeSourceApplyConfiguration {
	b.Optional = &value
	return b
}

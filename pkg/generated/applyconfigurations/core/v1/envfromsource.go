// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// EnvFromSourceApplyConfiguration represents a declarative configuration of the EnvFromSource type for use
// with apply.
type EnvFromSourceApplyConfiguration struct {
	Prefix       *string                               `json:"prefix,omitempty"`
	ConfigMapRef *ConfigMapEnvSourceApplyConfiguration `json:"configMapRef,omitempty"`
	SecretRef    *SecretEnvSourceApplyConfiguration    `json:"secretRef,omitempty"`
}

// EnvFromSourceApplyConfiguration constructs a declarative configuration of the EnvFromSource type for use with
// apply.
func EnvFromSource() *EnvFromSourceApplyConfiguration {
	return &EnvFromSourceApplyConfiguration{}
}

// WithPrefix sets the Prefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prefix field is set to the value of the last call.
func (b *EnvFromSourceApplyConfiguration) WithPrefix(value string) *EnvFromSourceApplyConfiguration {
	b.Prefix = &value
	return b
}

// WithConfigMapRef sets the ConfigMapRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMapRef field is set to the value of the last call.
func (b *EnvFromSourceApplyConfiguration) WithConfigMapRef(value *ConfigMapEnvSourceApplyConfiguration) *EnvFromSourceApplyConfiguration {
	b.ConfigMapRef = value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *EnvFromSourceApplyConfiguration) WithSecretRef(value *SecretEnvSourceApplyConfiguration) *EnvFromSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

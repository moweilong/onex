// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// EphemeralVolumeSourceApplyConfiguration represents a declarative configuration of the EphemeralVolumeSource type for use
// with apply.
type EphemeralVolumeSourceApplyConfiguration struct {
	VolumeClaimTemplate *PersistentVolumeClaimTemplateApplyConfiguration `json:"volumeClaimTemplate,omitempty"`
}

// EphemeralVolumeSourceApplyConfiguration constructs a declarative configuration of the EphemeralVolumeSource type for use with
// apply.
func EphemeralVolumeSource() *EphemeralVolumeSourceApplyConfiguration {
	return &EphemeralVolumeSourceApplyConfiguration{}
}

// WithVolumeClaimTemplate sets the VolumeClaimTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeClaimTemplate field is set to the value of the last call.
func (b *EphemeralVolumeSourceApplyConfiguration) WithVolumeClaimTemplate(value *PersistentVolumeClaimTemplateApplyConfiguration) *EphemeralVolumeSourceApplyConfiguration {
	b.VolumeClaimTemplate = value
	return b
}

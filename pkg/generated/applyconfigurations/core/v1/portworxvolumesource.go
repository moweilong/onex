// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PortworxVolumeSourceApplyConfiguration represents an declarative configuration of the PortworxVolumeSource type for use
// with apply.
type PortworxVolumeSourceApplyConfiguration struct {
	VolumeID *string `json:"volumeID,omitempty"`
	FSType   *string `json:"fsType,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// PortworxVolumeSourceApplyConfiguration constructs an declarative configuration of the PortworxVolumeSource type for use with
// apply.
func PortworxVolumeSource() *PortworxVolumeSourceApplyConfiguration {
	return &PortworxVolumeSourceApplyConfiguration{}
}

// WithVolumeID sets the VolumeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeID field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithVolumeID(value string) *PortworxVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithFSType(value string) *PortworxVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithReadOnly(value bool) *PortworxVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

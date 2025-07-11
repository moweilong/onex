// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// VolumeMountStatusApplyConfiguration represents a declarative configuration of the VolumeMountStatus type for use
// with apply.
type VolumeMountStatusApplyConfiguration struct {
	Name              *string                       `json:"name,omitempty"`
	MountPath         *string                       `json:"mountPath,omitempty"`
	ReadOnly          *bool                         `json:"readOnly,omitempty"`
	RecursiveReadOnly *corev1.RecursiveReadOnlyMode `json:"recursiveReadOnly,omitempty"`
}

// VolumeMountStatusApplyConfiguration constructs a declarative configuration of the VolumeMountStatus type for use with
// apply.
func VolumeMountStatus() *VolumeMountStatusApplyConfiguration {
	return &VolumeMountStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VolumeMountStatusApplyConfiguration) WithName(value string) *VolumeMountStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithMountPath sets the MountPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountPath field is set to the value of the last call.
func (b *VolumeMountStatusApplyConfiguration) WithMountPath(value string) *VolumeMountStatusApplyConfiguration {
	b.MountPath = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *VolumeMountStatusApplyConfiguration) WithReadOnly(value bool) *VolumeMountStatusApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithRecursiveReadOnly sets the RecursiveReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RecursiveReadOnly field is set to the value of the last call.
func (b *VolumeMountStatusApplyConfiguration) WithRecursiveReadOnly(value corev1.RecursiveReadOnlyMode) *VolumeMountStatusApplyConfiguration {
	b.RecursiveReadOnly = &value
	return b
}

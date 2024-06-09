// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ScaleSpecApplyConfiguration represents an declarative configuration of the ScaleSpec type for use
// with apply.
type ScaleSpecApplyConfiguration struct {
	Replicas *int32 `json:"replicas,omitempty"`
}

// ScaleSpecApplyConfiguration constructs an declarative configuration of the ScaleSpec type for use with
// apply.
func ScaleSpec() *ScaleSpecApplyConfiguration {
	return &ScaleSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ScaleSpecApplyConfiguration) WithReplicas(value int32) *ScaleSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

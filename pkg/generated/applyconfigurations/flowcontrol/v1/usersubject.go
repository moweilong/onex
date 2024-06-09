// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// UserSubjectApplyConfiguration represents an declarative configuration of the UserSubject type for use
// with apply.
type UserSubjectApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// UserSubjectApplyConfiguration constructs an declarative configuration of the UserSubject type for use with
// apply.
func UserSubject() *UserSubjectApplyConfiguration {
	return &UserSubjectApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *UserSubjectApplyConfiguration) WithName(value string) *UserSubjectApplyConfiguration {
	b.Name = &value
	return b
}

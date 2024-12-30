//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinerSetControllerConfiguration) DeepCopyInto(out *MinerSetControllerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.SyncPeriod = in.SyncPeriod
	out.LeaderElection = in.LeaderElection
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinerSetControllerConfiguration.
func (in *MinerSetControllerConfiguration) DeepCopy() *MinerSetControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(MinerSetControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinerSetControllerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

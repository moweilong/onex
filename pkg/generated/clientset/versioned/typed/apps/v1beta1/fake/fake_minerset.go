// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1beta1 "github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	appsv1beta1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/apps/v1beta1"
	applyconfigurationsautoscalingv1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/autoscaling/v1"
	typedappsv1beta1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"
)

// fakeMinerSets implements MinerSetInterface
type fakeMinerSets struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.MinerSet, *v1beta1.MinerSetList, *appsv1beta1.MinerSetApplyConfiguration]
	Fake *FakeAppsV1beta1
}

func newFakeMinerSets(fake *FakeAppsV1beta1, namespace string) typedappsv1beta1.MinerSetInterface {
	return &fakeMinerSets{
		gentype.NewFakeClientWithListAndApply[*v1beta1.MinerSet, *v1beta1.MinerSetList, *appsv1beta1.MinerSetApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("minersets"),
			v1beta1.SchemeGroupVersion.WithKind("MinerSet"),
			func() *v1beta1.MinerSet { return &v1beta1.MinerSet{} },
			func() *v1beta1.MinerSetList { return &v1beta1.MinerSetList{} },
			func(dst, src *v1beta1.MinerSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.MinerSetList) []*v1beta1.MinerSet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.MinerSetList, items []*v1beta1.MinerSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

// GetScale takes name of the minerSet, and returns the corresponding scale object, and an error if there is any.
func (c *fakeMinerSets) GetScale(ctx context.Context, minerSetName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(c.Resource(), c.Namespace(), "scale", minerSetName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *fakeMinerSets) UpdateScale(ctx context.Context, minerSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(c.Resource(), "scale", c.Namespace(), scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *fakeMinerSets) ApplyScale(ctx context.Context, minerSetName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(c.Resource(), c.Namespace(), minerSetName, types.ApplyPatchType, data, opts.ToPatchOptions(), "scale"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	applyconfigurationscorev1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/core/v1"
	scheme "github.com/onexstack/onex/pkg/generated/clientset/versioned/scheme"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ComponentStatusesGetter has a method to return a ComponentStatusInterface.
// A group's client should implement this interface.
type ComponentStatusesGetter interface {
	ComponentStatuses() ComponentStatusInterface
}

// ComponentStatusInterface has methods to work with ComponentStatus resources.
type ComponentStatusInterface interface {
	Create(ctx context.Context, componentStatus *corev1.ComponentStatus, opts metav1.CreateOptions) (*corev1.ComponentStatus, error)
	Update(ctx context.Context, componentStatus *corev1.ComponentStatus, opts metav1.UpdateOptions) (*corev1.ComponentStatus, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ComponentStatus, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ComponentStatusList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.ComponentStatus, err error)
	Apply(ctx context.Context, componentStatus *applyconfigurationscorev1.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.ComponentStatus, err error)
	ComponentStatusExpansion
}

// componentStatuses implements ComponentStatusInterface
type componentStatuses struct {
	*gentype.ClientWithListAndApply[*corev1.ComponentStatus, *corev1.ComponentStatusList, *applyconfigurationscorev1.ComponentStatusApplyConfiguration]
}

// newComponentStatuses returns a ComponentStatuses
func newComponentStatuses(c *CoreV1Client) *componentStatuses {
	return &componentStatuses{
		gentype.NewClientWithListAndApply[*corev1.ComponentStatus, *corev1.ComponentStatusList, *applyconfigurationscorev1.ComponentStatusApplyConfiguration](
			"componentstatuses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *corev1.ComponentStatus { return &corev1.ComponentStatus{} },
			func() *corev1.ComponentStatusList { return &corev1.ComponentStatusList{} },
			gentype.PrefersProtobuf[*corev1.ComponentStatus](),
		),
	}
}

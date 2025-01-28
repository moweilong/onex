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

// ResourceQuotasGetter has a method to return a ResourceQuotaInterface.
// A group's client should implement this interface.
type ResourceQuotasGetter interface {
	ResourceQuotas(namespace string) ResourceQuotaInterface
}

// ResourceQuotaInterface has methods to work with ResourceQuota resources.
type ResourceQuotaInterface interface {
	Create(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts metav1.CreateOptions) (*corev1.ResourceQuota, error)
	Update(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts metav1.UpdateOptions) (*corev1.ResourceQuota, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts metav1.UpdateOptions) (*corev1.ResourceQuota, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ResourceQuota, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ResourceQuotaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.ResourceQuota, err error)
	Apply(ctx context.Context, resourceQuota *applyconfigurationscorev1.ResourceQuotaApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.ResourceQuota, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, resourceQuota *applyconfigurationscorev1.ResourceQuotaApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.ResourceQuota, err error)
	ResourceQuotaExpansion
}

// resourceQuotas implements ResourceQuotaInterface
type resourceQuotas struct {
	*gentype.ClientWithListAndApply[*corev1.ResourceQuota, *corev1.ResourceQuotaList, *applyconfigurationscorev1.ResourceQuotaApplyConfiguration]
}

// newResourceQuotas returns a ResourceQuotas
func newResourceQuotas(c *CoreV1Client, namespace string) *resourceQuotas {
	return &resourceQuotas{
		gentype.NewClientWithListAndApply[*corev1.ResourceQuota, *corev1.ResourceQuotaList, *applyconfigurationscorev1.ResourceQuotaApplyConfiguration](
			"resourcequotas",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.ResourceQuota { return &corev1.ResourceQuota{} },
			func() *corev1.ResourceQuotaList { return &corev1.ResourceQuotaList{} },
			gentype.PrefersProtobuf[*corev1.ResourceQuota](),
		),
	}
}

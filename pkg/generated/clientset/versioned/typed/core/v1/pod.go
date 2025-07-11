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

// PodsGetter has a method to return a PodInterface.
// A group's client should implement this interface.
type PodsGetter interface {
	Pods(namespace string) PodInterface
}

// PodInterface has methods to work with Pod resources.
type PodInterface interface {
	Create(ctx context.Context, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error)
	Update(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Pod, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.Pod, err error)
	Apply(ctx context.Context, pod *applyconfigurationscorev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Pod, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, pod *applyconfigurationscorev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Pod, err error)
	UpdateEphemeralContainers(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error)
	UpdateResize(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error)

	PodExpansion
}

// pods implements PodInterface
type pods struct {
	*gentype.ClientWithListAndApply[*corev1.Pod, *corev1.PodList, *applyconfigurationscorev1.PodApplyConfiguration]
}

// newPods returns a Pods
func newPods(c *CoreV1Client, namespace string) *pods {
	return &pods{
		gentype.NewClientWithListAndApply[*corev1.Pod, *corev1.PodList, *applyconfigurationscorev1.PodApplyConfiguration](
			"pods",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.Pod { return &corev1.Pod{} },
			func() *corev1.PodList { return &corev1.PodList{} },
			gentype.PrefersProtobuf[*corev1.Pod](),
		),
	}
}

// UpdateEphemeralContainers takes the top resource name and the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *pods) UpdateEphemeralContainers(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions) (result *corev1.Pod, err error) {
	result = &corev1.Pod{}
	err = c.GetClient().Put().
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("pods").
		Name(podName).
		SubResource("ephemeralcontainers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pod).
		Do(ctx).
		Into(result)
	return
}

// UpdateResize takes the top resource name and the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *pods) UpdateResize(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions) (result *corev1.Pod, err error) {
	result = &corev1.Pod{}
	err = c.GetClient().Put().
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("pods").
		Name(podName).
		SubResource("resize").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pod).
		Do(ctx).
		Into(result)
	return
}

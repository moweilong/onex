// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	versioned "github.com/onexstack/onex/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/onexstack/onex/pkg/generated/informers/internalinterfaces"
	corev1 "github.com/onexstack/onex/pkg/generated/listers/core/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodTemplateInformer provides access to a shared informer and lister for
// PodTemplates.
type PodTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1.PodTemplateLister
}

type podTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodTemplateInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodTemplateInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&apicorev1.PodTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *podTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apicorev1.PodTemplate{}, f.defaultInformer)
}

func (f *podTemplateInformer) Lister() corev1.PodTemplateLister {
	return corev1.NewPodTemplateLister(f.Informer().GetIndexer())
}

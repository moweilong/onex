// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by informer-gen. DO NOT EDIT.

package informers

import (
	"fmt"

	v1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
	v1 "github.com/superproj/onex/pkg/apis/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.onex.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("chains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Chains().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("chargerequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ChargeRequests().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("evaluates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Evaluates().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("miners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Miners().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("minersets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().MinerSets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("modelcompares"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ModelCompares().Informer()}, nil

		// Group=coordination.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("leases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Coordination().V1().Leases().Informer()}, nil

		// Group=core, Version=v1
	case corev1.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ConfigMaps().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Events().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Namespaces().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Secrets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

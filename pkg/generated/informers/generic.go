// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by informer-gen. DO NOT EDIT.

package informers

import (
	fmt "fmt"

	v1beta1 "github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	coordinationv1 "k8s.io/api/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	flowcontrolv1 "k8s.io/api/flowcontrol/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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
	// Group=apiextensions.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("customresourcedefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiextensions().V1().CustomResourceDefinitions().Informer()}, nil

		// Group=apps.onex.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("chains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Chains().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("chargerequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ChargeRequests().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("miners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().Miners().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("minersets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().MinerSets().Informer()}, nil

		// Group=coordination.k8s.io, Version=v1
	case coordinationv1.SchemeGroupVersion.WithResource("leases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Coordination().V1().Leases().Informer()}, nil

		// Group=core, Version=v1
	case corev1.SchemeGroupVersion.WithResource("componentstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ComponentStatuses().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ConfigMaps().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Endpoints().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Events().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("limitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().LimitRanges().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Namespaces().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Nodes().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("persistentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumes().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PersistentVolumeClaims().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Pods().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("podtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().PodTemplates().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("replicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ReplicationControllers().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ResourceQuotas().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Secrets().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Services().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("serviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ServiceAccounts().Informer()}, nil

		// Group=flowcontrol.apiserver.k8s.io, Version=v1
	case flowcontrolv1.SchemeGroupVersion.WithResource("flowschemas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flowcontrol().V1().FlowSchemas().Informer()}, nil
	case flowcontrolv1.SchemeGroupVersion.WithResource("prioritylevelconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flowcontrol().V1().PriorityLevelConfigurations().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

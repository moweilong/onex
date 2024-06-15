// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ModelCompareLister helps list ModelCompares.
// All objects returned here must be treated as read-only.
type ModelCompareLister interface {
	// List lists all ModelCompares in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ModelCompare, err error)
	// ModelCompares returns an object that can list and get ModelCompares.
	ModelCompares(namespace string) ModelCompareNamespaceLister
	ModelCompareListerExpansion
}

// modelCompareLister implements the ModelCompareLister interface.
type modelCompareLister struct {
	indexer cache.Indexer
}

// NewModelCompareLister returns a new ModelCompareLister.
func NewModelCompareLister(indexer cache.Indexer) ModelCompareLister {
	return &modelCompareLister{indexer: indexer}
}

// List lists all ModelCompares in the indexer.
func (s *modelCompareLister) List(selector labels.Selector) (ret []*v1beta1.ModelCompare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ModelCompare))
	})
	return ret, err
}

// ModelCompares returns an object that can list and get ModelCompares.
func (s *modelCompareLister) ModelCompares(namespace string) ModelCompareNamespaceLister {
	return modelCompareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ModelCompareNamespaceLister helps list and get ModelCompares.
// All objects returned here must be treated as read-only.
type ModelCompareNamespaceLister interface {
	// List lists all ModelCompares in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ModelCompare, err error)
	// Get retrieves the ModelCompare from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ModelCompare, error)
	ModelCompareNamespaceListerExpansion
}

// modelCompareNamespaceLister implements the ModelCompareNamespaceLister
// interface.
type modelCompareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ModelCompares in the indexer for a given namespace.
func (s modelCompareNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ModelCompare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ModelCompare))
	})
	return ret, err
}

// Get retrieves the ModelCompare from the indexer for a given namespace and name.
func (s modelCompareNamespaceLister) Get(name string) (*v1beta1.ModelCompare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("modelcompare"), name)
	}
	return obj.(*v1beta1.ModelCompare), nil
}

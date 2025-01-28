// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	apiextensionsv1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/apiextensions/v1"
	typedapiextensionsv1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/apiextensions/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCustomResourceDefinitions implements CustomResourceDefinitionInterface
type fakeCustomResourceDefinitions struct {
	*gentype.FakeClientWithListAndApply[*v1.CustomResourceDefinition, *v1.CustomResourceDefinitionList, *apiextensionsv1.CustomResourceDefinitionApplyConfiguration]
	Fake *FakeApiextensionsV1
}

func newFakeCustomResourceDefinitions(fake *FakeApiextensionsV1) typedapiextensionsv1.CustomResourceDefinitionInterface {
	return &fakeCustomResourceDefinitions{
		gentype.NewFakeClientWithListAndApply[*v1.CustomResourceDefinition, *v1.CustomResourceDefinitionList, *apiextensionsv1.CustomResourceDefinitionApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("customresourcedefinitions"),
			v1.SchemeGroupVersion.WithKind("CustomResourceDefinition"),
			func() *v1.CustomResourceDefinition { return &v1.CustomResourceDefinition{} },
			func() *v1.CustomResourceDefinitionList { return &v1.CustomResourceDefinitionList{} },
			func(dst, src *v1.CustomResourceDefinitionList) { dst.ListMeta = src.ListMeta },
			func(list *v1.CustomResourceDefinitionList) []*v1.CustomResourceDefinition {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.CustomResourceDefinitionList, items []*v1.CustomResourceDefinition) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

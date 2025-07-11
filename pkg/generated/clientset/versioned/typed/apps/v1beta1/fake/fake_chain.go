// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	appsv1beta1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/apps/v1beta1"
	typedappsv1beta1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/apps/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeChains implements ChainInterface
type fakeChains struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.Chain, *v1beta1.ChainList, *appsv1beta1.ChainApplyConfiguration]
	Fake *FakeAppsV1beta1
}

func newFakeChains(fake *FakeAppsV1beta1, namespace string) typedappsv1beta1.ChainInterface {
	return &fakeChains{
		gentype.NewFakeClientWithListAndApply[*v1beta1.Chain, *v1beta1.ChainList, *appsv1beta1.ChainApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("chains"),
			v1beta1.SchemeGroupVersion.WithKind("Chain"),
			func() *v1beta1.Chain { return &v1beta1.Chain{} },
			func() *v1beta1.ChainList { return &v1beta1.ChainList{} },
			func(dst, src *v1beta1.ChainList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.ChainList) []*v1beta1.Chain { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.ChainList, items []*v1beta1.Chain) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

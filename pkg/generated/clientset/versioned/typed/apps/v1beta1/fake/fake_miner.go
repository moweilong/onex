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

// fakeMiners implements MinerInterface
type fakeMiners struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.Miner, *v1beta1.MinerList, *appsv1beta1.MinerApplyConfiguration]
	Fake *FakeAppsV1beta1
}

func newFakeMiners(fake *FakeAppsV1beta1, namespace string) typedappsv1beta1.MinerInterface {
	return &fakeMiners{
		gentype.NewFakeClientWithListAndApply[*v1beta1.Miner, *v1beta1.MinerList, *appsv1beta1.MinerApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("miners"),
			v1beta1.SchemeGroupVersion.WithKind("Miner"),
			func() *v1beta1.Miner { return &v1beta1.Miner{} },
			func() *v1beta1.MinerList { return &v1beta1.MinerList{} },
			func(dst, src *v1beta1.MinerList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.MinerList) []*v1beta1.Miner { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.MinerList, items []*v1beta1.Miner) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

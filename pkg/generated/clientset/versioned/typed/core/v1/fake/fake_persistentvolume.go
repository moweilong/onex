// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	corev1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/core/v1"
	typedcorev1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/core/v1"
	v1 "k8s.io/api/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakePersistentVolumes implements PersistentVolumeInterface
type fakePersistentVolumes struct {
	*gentype.FakeClientWithListAndApply[*v1.PersistentVolume, *v1.PersistentVolumeList, *corev1.PersistentVolumeApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakePersistentVolumes(fake *FakeCoreV1) typedcorev1.PersistentVolumeInterface {
	return &fakePersistentVolumes{
		gentype.NewFakeClientWithListAndApply[*v1.PersistentVolume, *v1.PersistentVolumeList, *corev1.PersistentVolumeApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("persistentvolumes"),
			v1.SchemeGroupVersion.WithKind("PersistentVolume"),
			func() *v1.PersistentVolume { return &v1.PersistentVolume{} },
			func() *v1.PersistentVolumeList { return &v1.PersistentVolumeList{} },
			func(dst, src *v1.PersistentVolumeList) { dst.ListMeta = src.ListMeta },
			func(list *v1.PersistentVolumeList) []*v1.PersistentVolume { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.PersistentVolumeList, items []*v1.PersistentVolume) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//nolint:gomodguard
package contract

import (
	"strings"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Ref provide a helper struct for working with references in Unstructured objects.
type Ref struct {
	path Path
}

// Path returns the path of the reference.
func (r *Ref) Path() Path {
	return r.path
}

// Get gets the reference value from the Unstructured object.
func (r *Ref) Get(obj *unstructured.Unstructured) (*corev1.ObjectReference, error) {
	return GetNestedRef(obj, r.path...)
}

// Set sets the reference value in the Unstructured object.
func (r *Ref) Set(obj, refObj *unstructured.Unstructured) error {
	return SetNestedRef(obj, refObj, r.path...)
}

// GetNestedRef returns the ref value from a nested field in an Unstructured object.
func GetNestedRef(obj *unstructured.Unstructured, fields ...string) (*corev1.ObjectReference, error) {
	ref := &corev1.ObjectReference{}
	if v, ok, err := unstructured.NestedString(obj.UnstructuredContent(), append(fields, "apiVersion")...); ok && err == nil {
		ref.APIVersion = v
	} else {
		return nil, errors.Errorf("failed to get %s.apiVersion from %s", strings.Join(fields, "."), obj.GetKind())
	}
	if v, ok, err := unstructured.NestedString(obj.UnstructuredContent(), append(fields, "kind")...); ok && err == nil {
		ref.Kind = v
	} else {
		return nil, errors.Errorf("failed to get %s.kind from %s", strings.Join(fields, "."), obj.GetKind())
	}
	if v, ok, err := unstructured.NestedString(obj.UnstructuredContent(), append(fields, "name")...); ok && err == nil {
		ref.Name = v
	} else {
		return nil, errors.Errorf("failed to get %s.name from %s", strings.Join(fields, "."), obj.GetKind())
	}
	if v, ok, err := unstructured.NestedString(obj.UnstructuredContent(), append(fields, "namespace")...); ok && err == nil {
		ref.Namespace = v
	} else {
		return nil, errors.Errorf("failed to get %s.namespace from %s", strings.Join(fields, "."), obj.GetKind())
	}
	return ref, nil
}

// SetNestedRef sets the value of a nested field in an Unstructured to a reference to the refObj provided.
func SetNestedRef(obj, refObj *unstructured.Unstructured, fields ...string) error {
	ref := map[string]any{
		"kind":       refObj.GetKind(),
		"namespace":  refObj.GetNamespace(),
		"name":       refObj.GetName(),
		"apiVersion": refObj.GetAPIVersion(),
	}
	if err := unstructured.SetNestedField(obj.UnstructuredContent(), ref, fields...); err != nil {
		return errors.Wrapf(err, "failed to set object reference on object %v %s",
			obj.GroupVersionKind(), klog.KObj(obj))
	}
	return nil
}

// ObjToRef returns a reference to the given object.
func ObjToRef(obj client.Object) *corev1.ObjectReference {
	gvk := obj.GetObjectKind().GroupVersionKind()
	return &corev1.ObjectReference{
		Kind:       gvk.Kind,
		APIVersion: gvk.GroupVersion().String(),
		Namespace:  obj.GetNamespace(),
		Name:       obj.GetName(),
	}
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	config "github.com/superproj/onex/pkg/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*GroupResource)(nil), (*config.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_GroupResource_To_config_GroupResource(a.(*GroupResource), b.(*config.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GroupResource)(nil), (*GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GroupResource_To_v1beta1_GroupResource(a.(*config.GroupResource), b.(*GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.GarbageCollectorControllerConfiguration)(nil), (*GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(a.(*config.GarbageCollectorControllerConfiguration), b.(*GarbageCollectorControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.GenericControllerManagerConfiguration)(nil), (*GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(a.(*config.GenericControllerManagerConfiguration), b.(*GenericControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.MySQLConfiguration)(nil), (*MySQLConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(a.(*config.MySQLConfiguration), b.(*MySQLConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.RedisConfiguration)(nil), (*RedisConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(a.(*config.RedisConfiguration), b.(*RedisConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*GarbageCollectorControllerConfiguration)(nil), (*config.GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(a.(*GarbageCollectorControllerConfiguration), b.(*config.GarbageCollectorControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*GenericControllerManagerConfiguration)(nil), (*config.GenericControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(a.(*GenericControllerManagerConfiguration), b.(*config.GenericControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MySQLConfiguration)(nil), (*config.MySQLConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(a.(*MySQLConfiguration), b.(*config.MySQLConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*RedisConfiguration)(nil), (*config.RedisConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(a.(*RedisConfiguration), b.(*config.RedisConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]config.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	return nil
}

func autoConvert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	return nil
}

func autoConvert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.BindAddress = in.BindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.PprofBindAddress = in.PprofBindAddress
	out.Parallelism = in.Parallelism
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	return nil
}

func autoConvert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *GenericControllerManagerConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.BindAddress = in.BindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.PprofBindAddress = in.PprofBindAddress
	out.Parallelism = in.Parallelism
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	return nil
}

func autoConvert_v1beta1_GroupResource_To_config_GroupResource(in *GroupResource, out *config.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1beta1_GroupResource_To_config_GroupResource is an autogenerated conversion function.
func Convert_v1beta1_GroupResource_To_config_GroupResource(in *GroupResource, out *config.GroupResource, s conversion.Scope) error {
	return autoConvert_v1beta1_GroupResource_To_config_GroupResource(in, out, s)
}

func autoConvert_config_GroupResource_To_v1beta1_GroupResource(in *config.GroupResource, out *GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_config_GroupResource_To_v1beta1_GroupResource is an autogenerated conversion function.
func Convert_config_GroupResource_To_v1beta1_GroupResource(in *config.GroupResource, out *GroupResource, s conversion.Scope) error {
	return autoConvert_config_GroupResource_To_v1beta1_GroupResource(in, out, s)
}

func autoConvert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(in *MySQLConfiguration, out *config.MySQLConfiguration, s conversion.Scope) error {
	out.Host = in.Host
	out.Username = in.Username
	out.Password = in.Password
	out.Database = in.Database
	out.MaxIdleConnections = in.MaxIdleConnections
	out.MaxOpenConnections = in.MaxOpenConnections
	out.MaxConnectionLifeTime = in.MaxConnectionLifeTime
	return nil
}

func autoConvert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(in *config.MySQLConfiguration, out *MySQLConfiguration, s conversion.Scope) error {
	out.Host = in.Host
	out.Username = in.Username
	out.Password = in.Password
	out.Database = in.Database
	out.MaxIdleConnections = in.MaxIdleConnections
	out.MaxOpenConnections = in.MaxOpenConnections
	out.MaxConnectionLifeTime = in.MaxConnectionLifeTime
	return nil
}

func autoConvert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(in *RedisConfiguration, out *config.RedisConfiguration, s conversion.Scope) error {
	out.Addr = in.Addr
	out.Username = in.Username
	out.Password = in.Password
	out.Database = in.Database
	out.MaxRetries = in.MaxRetries
	out.Timeout = in.Timeout
	return nil
}

func autoConvert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(in *config.RedisConfiguration, out *RedisConfiguration, s conversion.Scope) error {
	out.Addr = in.Addr
	out.Username = in.Username
	out.Password = in.Password
	out.Database = in.Database
	out.MaxRetries = in.MaxRetries
	out.Timeout = in.Timeout
	return nil
}

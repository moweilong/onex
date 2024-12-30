// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/flowcontrol/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFlowcontrolV1 struct {
	*testing.Fake
}

func (c *FakeFlowcontrolV1) FlowSchemas() v1.FlowSchemaInterface {
	return &FakeFlowSchemas{c}
}

func (c *FakeFlowcontrolV1) PriorityLevelConfigurations() v1.PriorityLevelConfigurationInterface {
	return &FakePriorityLevelConfigurations{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFlowcontrolV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

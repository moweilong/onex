// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ExemptPriorityLevelConfigurationApplyConfiguration represents a declarative configuration of the ExemptPriorityLevelConfiguration type for use
// with apply.
type ExemptPriorityLevelConfigurationApplyConfiguration struct {
	NominalConcurrencyShares *int32 `json:"nominalConcurrencyShares,omitempty"`
	LendablePercent          *int32 `json:"lendablePercent,omitempty"`
}

// ExemptPriorityLevelConfigurationApplyConfiguration constructs a declarative configuration of the ExemptPriorityLevelConfiguration type for use with
// apply.
func ExemptPriorityLevelConfiguration() *ExemptPriorityLevelConfigurationApplyConfiguration {
	return &ExemptPriorityLevelConfigurationApplyConfiguration{}
}

// WithNominalConcurrencyShares sets the NominalConcurrencyShares field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NominalConcurrencyShares field is set to the value of the last call.
func (b *ExemptPriorityLevelConfigurationApplyConfiguration) WithNominalConcurrencyShares(value int32) *ExemptPriorityLevelConfigurationApplyConfiguration {
	b.NominalConcurrencyShares = &value
	return b
}

// WithLendablePercent sets the LendablePercent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LendablePercent field is set to the value of the last call.
func (b *ExemptPriorityLevelConfigurationApplyConfiguration) WithLendablePercent(value int32) *ExemptPriorityLevelConfigurationApplyConfiguration {
	b.LendablePercent = &value
	return b
}

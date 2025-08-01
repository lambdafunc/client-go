/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// AllocationResultApplyConfiguration represents a declarative configuration of the AllocationResult type for use
// with apply.
type AllocationResultApplyConfiguration struct {
	Devices             *DeviceAllocationResultApplyConfiguration `json:"devices,omitempty"`
	NodeSelector        *v1.NodeSelectorApplyConfiguration        `json:"nodeSelector,omitempty"`
	AllocationTimestamp *metav1.Time                              `json:"allocationTimestamp,omitempty"`
}

// AllocationResultApplyConfiguration constructs a declarative configuration of the AllocationResult type for use with
// apply.
func AllocationResult() *AllocationResultApplyConfiguration {
	return &AllocationResultApplyConfiguration{}
}

// WithDevices sets the Devices field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Devices field is set to the value of the last call.
func (b *AllocationResultApplyConfiguration) WithDevices(value *DeviceAllocationResultApplyConfiguration) *AllocationResultApplyConfiguration {
	b.Devices = value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *AllocationResultApplyConfiguration) WithNodeSelector(value *v1.NodeSelectorApplyConfiguration) *AllocationResultApplyConfiguration {
	b.NodeSelector = value
	return b
}

// WithAllocationTimestamp sets the AllocationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocationTimestamp field is set to the value of the last call.
func (b *AllocationResultApplyConfiguration) WithAllocationTimestamp(value metav1.Time) *AllocationResultApplyConfiguration {
	b.AllocationTimestamp = &value
	return b
}

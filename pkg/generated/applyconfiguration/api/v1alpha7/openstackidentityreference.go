/*
Copyright 2024 The Kubernetes Authors.

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

package v1alpha7

// OpenStackIdentityReferenceApplyConfiguration represents a declarative configuration of the OpenStackIdentityReference type for use
// with apply.
type OpenStackIdentityReferenceApplyConfiguration struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
}

// OpenStackIdentityReferenceApplyConfiguration constructs a declarative configuration of the OpenStackIdentityReference type for use with
// apply.
func OpenStackIdentityReference() *OpenStackIdentityReferenceApplyConfiguration {
	return &OpenStackIdentityReferenceApplyConfiguration{}
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *OpenStackIdentityReferenceApplyConfiguration) WithKind(value string) *OpenStackIdentityReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OpenStackIdentityReferenceApplyConfiguration) WithName(value string) *OpenStackIdentityReferenceApplyConfiguration {
	b.Name = &value
	return b
}

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

package v1alpha6

// ExternalRouterIPParamApplyConfiguration represents a declarative configuration of the ExternalRouterIPParam type for use
// with apply.
type ExternalRouterIPParamApplyConfiguration struct {
	FixedIP *string                        `json:"fixedIP,omitempty"`
	Subnet  *SubnetParamApplyConfiguration `json:"subnet,omitempty"`
}

// ExternalRouterIPParamApplyConfiguration constructs a declarative configuration of the ExternalRouterIPParam type for use with
// apply.
func ExternalRouterIPParam() *ExternalRouterIPParamApplyConfiguration {
	return &ExternalRouterIPParamApplyConfiguration{}
}

// WithFixedIP sets the FixedIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FixedIP field is set to the value of the last call.
func (b *ExternalRouterIPParamApplyConfiguration) WithFixedIP(value string) *ExternalRouterIPParamApplyConfiguration {
	b.FixedIP = &value
	return b
}

// WithSubnet sets the Subnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subnet field is set to the value of the last call.
func (b *ExternalRouterIPParamApplyConfiguration) WithSubnet(value *SubnetParamApplyConfiguration) *ExternalRouterIPParamApplyConfiguration {
	b.Subnet = value
	return b
}

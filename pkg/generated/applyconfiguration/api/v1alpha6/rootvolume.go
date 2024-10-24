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

// RootVolumeApplyConfiguration represents a declarative configuration of the RootVolume type for use
// with apply.
type RootVolumeApplyConfiguration struct {
	Size             *int    `json:"diskSize,omitempty"`
	VolumeType       *string `json:"volumeType,omitempty"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
}

// RootVolumeApplyConfiguration constructs a declarative configuration of the RootVolume type for use with
// apply.
func RootVolume() *RootVolumeApplyConfiguration {
	return &RootVolumeApplyConfiguration{}
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithSize(value int) *RootVolumeApplyConfiguration {
	b.Size = &value
	return b
}

// WithVolumeType sets the VolumeType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeType field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithVolumeType(value string) *RootVolumeApplyConfiguration {
	b.VolumeType = &value
	return b
}

// WithAvailabilityZone sets the AvailabilityZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailabilityZone field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithAvailabilityZone(value string) *RootVolumeApplyConfiguration {
	b.AvailabilityZone = &value
	return b
}

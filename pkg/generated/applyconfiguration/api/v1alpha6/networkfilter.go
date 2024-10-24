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

// NetworkFilterApplyConfiguration represents a declarative configuration of the NetworkFilter type for use
// with apply.
type NetworkFilterApplyConfiguration struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ProjectID   *string `json:"projectId,omitempty"`
	ID          *string `json:"id,omitempty"`
	Tags        *string `json:"tags,omitempty"`
	TagsAny     *string `json:"tagsAny,omitempty"`
	NotTags     *string `json:"notTags,omitempty"`
	NotTagsAny  *string `json:"notTagsAny,omitempty"`
}

// NetworkFilterApplyConfiguration constructs a declarative configuration of the NetworkFilter type for use with
// apply.
func NetworkFilter() *NetworkFilterApplyConfiguration {
	return &NetworkFilterApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithName(value string) *NetworkFilterApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithDescription(value string) *NetworkFilterApplyConfiguration {
	b.Description = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithProjectID(value string) *NetworkFilterApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithID(value string) *NetworkFilterApplyConfiguration {
	b.ID = &value
	return b
}

// WithTags sets the Tags field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tags field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithTags(value string) *NetworkFilterApplyConfiguration {
	b.Tags = &value
	return b
}

// WithTagsAny sets the TagsAny field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TagsAny field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithTagsAny(value string) *NetworkFilterApplyConfiguration {
	b.TagsAny = &value
	return b
}

// WithNotTags sets the NotTags field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotTags field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithNotTags(value string) *NetworkFilterApplyConfiguration {
	b.NotTags = &value
	return b
}

// WithNotTagsAny sets the NotTagsAny field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotTagsAny field is set to the value of the last call.
func (b *NetworkFilterApplyConfiguration) WithNotTagsAny(value string) *NetworkFilterApplyConfiguration {
	b.NotTagsAny = &value
	return b
}

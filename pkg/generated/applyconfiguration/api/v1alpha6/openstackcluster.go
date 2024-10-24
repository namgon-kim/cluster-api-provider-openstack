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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	apiv1alpha6 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha6"
	internal "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/internal"
)

// OpenStackClusterApplyConfiguration represents a declarative configuration of the OpenStackCluster type for use
// with apply.
type OpenStackClusterApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *OpenStackClusterSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *OpenStackClusterStatusApplyConfiguration `json:"status,omitempty"`
}

// OpenStackCluster constructs a declarative configuration of the OpenStackCluster type for use with
// apply.
func OpenStackCluster(name, namespace string) *OpenStackClusterApplyConfiguration {
	b := &OpenStackClusterApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("OpenStackCluster")
	b.WithAPIVersion("infrastructure.cluster.x-k8s.io/v1alpha6")
	return b
}

// ExtractOpenStackCluster extracts the applied configuration owned by fieldManager from
// openStackCluster. If no managedFields are found in openStackCluster for fieldManager, a
// OpenStackClusterApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// openStackCluster must be a unmodified OpenStackCluster API object that was retrieved from the Kubernetes API.
// ExtractOpenStackCluster provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractOpenStackCluster(openStackCluster *apiv1alpha6.OpenStackCluster, fieldManager string) (*OpenStackClusterApplyConfiguration, error) {
	return extractOpenStackCluster(openStackCluster, fieldManager, "")
}

// ExtractOpenStackClusterStatus is the same as ExtractOpenStackCluster except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractOpenStackClusterStatus(openStackCluster *apiv1alpha6.OpenStackCluster, fieldManager string) (*OpenStackClusterApplyConfiguration, error) {
	return extractOpenStackCluster(openStackCluster, fieldManager, "status")
}

func extractOpenStackCluster(openStackCluster *apiv1alpha6.OpenStackCluster, fieldManager string, subresource string) (*OpenStackClusterApplyConfiguration, error) {
	b := &OpenStackClusterApplyConfiguration{}
	err := managedfields.ExtractInto(openStackCluster, internal.Parser().Type("io.k8s.sigs.cluster-api-provider-openstack.api.v1alpha6.OpenStackCluster"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(openStackCluster.Name)
	b.WithNamespace(openStackCluster.Namespace)

	b.WithKind("OpenStackCluster")
	b.WithAPIVersion("infrastructure.cluster.x-k8s.io/v1alpha6")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithKind(value string) *OpenStackClusterApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithAPIVersion(value string) *OpenStackClusterApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithName(value string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithGenerateName(value string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithNamespace(value string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithUID(value types.UID) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithResourceVersion(value string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithGeneration(value int64) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithCreationTimestamp(value metav1.Time) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *OpenStackClusterApplyConfiguration) WithLabels(entries map[string]string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *OpenStackClusterApplyConfiguration) WithAnnotations(entries map[string]string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *OpenStackClusterApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *OpenStackClusterApplyConfiguration) WithFinalizers(values ...string) *OpenStackClusterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *OpenStackClusterApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithSpec(value *OpenStackClusterSpecApplyConfiguration) *OpenStackClusterApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *OpenStackClusterApplyConfiguration) WithStatus(value *OpenStackClusterStatusApplyConfiguration) *OpenStackClusterApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *OpenStackClusterApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}

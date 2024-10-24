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

package v1beta1

import (
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// OpenStackClusterSpecApplyConfiguration represents a declarative configuration of the OpenStackClusterSpec type for use
// with apply.
type OpenStackClusterSpecApplyConfiguration struct {
	ManagedSubnets                   []SubnetSpecApplyConfiguration                `json:"managedSubnets,omitempty"`
	Router                           *RouterParamApplyConfiguration                `json:"router,omitempty"`
	Network                          *NetworkParamApplyConfiguration               `json:"network,omitempty"`
	Subnets                          []SubnetParamApplyConfiguration               `json:"subnets,omitempty"`
	NetworkMTU                       *int                                          `json:"networkMTU,omitempty"`
	ExternalRouterIPs                []ExternalRouterIPParamApplyConfiguration     `json:"externalRouterIPs,omitempty"`
	ExternalNetwork                  *NetworkParamApplyConfiguration               `json:"externalNetwork,omitempty"`
	DisableExternalNetwork           *bool                                         `json:"disableExternalNetwork,omitempty"`
	APIServerLoadBalancer            *APIServerLoadBalancerApplyConfiguration      `json:"apiServerLoadBalancer,omitempty"`
	DisableAPIServerFloatingIP       *bool                                         `json:"disableAPIServerFloatingIP,omitempty"`
	APIServerFloatingIP              *string                                       `json:"apiServerFloatingIP,omitempty"`
	APIServerFixedIP                 *string                                       `json:"apiServerFixedIP,omitempty"`
	APIServerPort                    *uint16                                       `json:"apiServerPort,omitempty"`
	ManagedSecurityGroups            *ManagedSecurityGroupsApplyConfiguration      `json:"managedSecurityGroups,omitempty"`
	DisablePortSecurity              *bool                                         `json:"disablePortSecurity,omitempty"`
	Tags                             []string                                      `json:"tags,omitempty"`
	ControlPlaneEndpoint             *apiv1beta1.APIEndpoint                       `json:"controlPlaneEndpoint,omitempty"`
	ControlPlaneAvailabilityZones    []string                                      `json:"controlPlaneAvailabilityZones,omitempty"`
	ControlPlaneOmitAvailabilityZone *bool                                         `json:"controlPlaneOmitAvailabilityZone,omitempty"`
	Bastion                          *BastionApplyConfiguration                    `json:"bastion,omitempty"`
	IdentityRef                      *OpenStackIdentityReferenceApplyConfiguration `json:"identityRef,omitempty"`
}

// OpenStackClusterSpecApplyConfiguration constructs a declarative configuration of the OpenStackClusterSpec type for use with
// apply.
func OpenStackClusterSpec() *OpenStackClusterSpecApplyConfiguration {
	return &OpenStackClusterSpecApplyConfiguration{}
}

// WithManagedSubnets adds the given value to the ManagedSubnets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ManagedSubnets field.
func (b *OpenStackClusterSpecApplyConfiguration) WithManagedSubnets(values ...*SubnetSpecApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithManagedSubnets")
		}
		b.ManagedSubnets = append(b.ManagedSubnets, *values[i])
	}
	return b
}

// WithRouter sets the Router field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Router field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithRouter(value *RouterParamApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.Router = value
	return b
}

// WithNetwork sets the Network field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Network field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithNetwork(value *NetworkParamApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.Network = value
	return b
}

// WithSubnets adds the given value to the Subnets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Subnets field.
func (b *OpenStackClusterSpecApplyConfiguration) WithSubnets(values ...*SubnetParamApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSubnets")
		}
		b.Subnets = append(b.Subnets, *values[i])
	}
	return b
}

// WithNetworkMTU sets the NetworkMTU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkMTU field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithNetworkMTU(value int) *OpenStackClusterSpecApplyConfiguration {
	b.NetworkMTU = &value
	return b
}

// WithExternalRouterIPs adds the given value to the ExternalRouterIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalRouterIPs field.
func (b *OpenStackClusterSpecApplyConfiguration) WithExternalRouterIPs(values ...*ExternalRouterIPParamApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalRouterIPs")
		}
		b.ExternalRouterIPs = append(b.ExternalRouterIPs, *values[i])
	}
	return b
}

// WithExternalNetwork sets the ExternalNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalNetwork field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithExternalNetwork(value *NetworkParamApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.ExternalNetwork = value
	return b
}

// WithDisableExternalNetwork sets the DisableExternalNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableExternalNetwork field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithDisableExternalNetwork(value bool) *OpenStackClusterSpecApplyConfiguration {
	b.DisableExternalNetwork = &value
	return b
}

// WithAPIServerLoadBalancer sets the APIServerLoadBalancer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerLoadBalancer field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithAPIServerLoadBalancer(value *APIServerLoadBalancerApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.APIServerLoadBalancer = value
	return b
}

// WithDisableAPIServerFloatingIP sets the DisableAPIServerFloatingIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableAPIServerFloatingIP field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithDisableAPIServerFloatingIP(value bool) *OpenStackClusterSpecApplyConfiguration {
	b.DisableAPIServerFloatingIP = &value
	return b
}

// WithAPIServerFloatingIP sets the APIServerFloatingIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerFloatingIP field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithAPIServerFloatingIP(value string) *OpenStackClusterSpecApplyConfiguration {
	b.APIServerFloatingIP = &value
	return b
}

// WithAPIServerFixedIP sets the APIServerFixedIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerFixedIP field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithAPIServerFixedIP(value string) *OpenStackClusterSpecApplyConfiguration {
	b.APIServerFixedIP = &value
	return b
}

// WithAPIServerPort sets the APIServerPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerPort field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithAPIServerPort(value uint16) *OpenStackClusterSpecApplyConfiguration {
	b.APIServerPort = &value
	return b
}

// WithManagedSecurityGroups sets the ManagedSecurityGroups field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagedSecurityGroups field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithManagedSecurityGroups(value *ManagedSecurityGroupsApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.ManagedSecurityGroups = value
	return b
}

// WithDisablePortSecurity sets the DisablePortSecurity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisablePortSecurity field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithDisablePortSecurity(value bool) *OpenStackClusterSpecApplyConfiguration {
	b.DisablePortSecurity = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *OpenStackClusterSpecApplyConfiguration) WithTags(values ...string) *OpenStackClusterSpecApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithControlPlaneEndpoint sets the ControlPlaneEndpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneEndpoint field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithControlPlaneEndpoint(value apiv1beta1.APIEndpoint) *OpenStackClusterSpecApplyConfiguration {
	b.ControlPlaneEndpoint = &value
	return b
}

// WithControlPlaneAvailabilityZones adds the given value to the ControlPlaneAvailabilityZones field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ControlPlaneAvailabilityZones field.
func (b *OpenStackClusterSpecApplyConfiguration) WithControlPlaneAvailabilityZones(values ...string) *OpenStackClusterSpecApplyConfiguration {
	for i := range values {
		b.ControlPlaneAvailabilityZones = append(b.ControlPlaneAvailabilityZones, values[i])
	}
	return b
}

// WithControlPlaneOmitAvailabilityZone sets the ControlPlaneOmitAvailabilityZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneOmitAvailabilityZone field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithControlPlaneOmitAvailabilityZone(value bool) *OpenStackClusterSpecApplyConfiguration {
	b.ControlPlaneOmitAvailabilityZone = &value
	return b
}

// WithBastion sets the Bastion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bastion field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithBastion(value *BastionApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.Bastion = value
	return b
}

// WithIdentityRef sets the IdentityRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IdentityRef field is set to the value of the last call.
func (b *OpenStackClusterSpecApplyConfiguration) WithIdentityRef(value *OpenStackIdentityReferenceApplyConfiguration) *OpenStackClusterSpecApplyConfiguration {
	b.IdentityRef = value
	return b
}

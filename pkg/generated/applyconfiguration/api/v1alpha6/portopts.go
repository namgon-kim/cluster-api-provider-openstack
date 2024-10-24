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

// PortOptsApplyConfiguration represents a declarative configuration of the PortOpts type for use
// with apply.
type PortOptsApplyConfiguration struct {
	Network              *NetworkFilterApplyConfiguration       `json:"network,omitempty"`
	NameSuffix           *string                                `json:"nameSuffix,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	AdminStateUp         *bool                                  `json:"adminStateUp,omitempty"`
	MACAddress           *string                                `json:"macAddress,omitempty"`
	FixedIPs             []FixedIPApplyConfiguration            `json:"fixedIPs,omitempty"`
	TenantID             *string                                `json:"tenantId,omitempty"`
	ProjectID            *string                                `json:"projectId,omitempty"`
	SecurityGroups       []string                               `json:"securityGroups,omitempty"`
	SecurityGroupFilters []SecurityGroupParamApplyConfiguration `json:"securityGroupFilters,omitempty"`
	AllowedAddressPairs  []AddressPairApplyConfiguration        `json:"allowedAddressPairs,omitempty"`
	Trunk                *bool                                  `json:"trunk,omitempty"`
	HostID               *string                                `json:"hostId,omitempty"`
	VNICType             *string                                `json:"vnicType,omitempty"`
	Profile              map[string]string                      `json:"profile,omitempty"`
	DisablePortSecurity  *bool                                  `json:"disablePortSecurity,omitempty"`
	Tags                 []string                               `json:"tags,omitempty"`
	ValueSpecs           []ValueSpecApplyConfiguration          `json:"valueSpecs,omitempty"`
}

// PortOptsApplyConfiguration constructs a declarative configuration of the PortOpts type for use with
// apply.
func PortOpts() *PortOptsApplyConfiguration {
	return &PortOptsApplyConfiguration{}
}

// WithNetwork sets the Network field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Network field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithNetwork(value *NetworkFilterApplyConfiguration) *PortOptsApplyConfiguration {
	b.Network = value
	return b
}

// WithNameSuffix sets the NameSuffix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NameSuffix field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithNameSuffix(value string) *PortOptsApplyConfiguration {
	b.NameSuffix = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithDescription(value string) *PortOptsApplyConfiguration {
	b.Description = &value
	return b
}

// WithAdminStateUp sets the AdminStateUp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminStateUp field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithAdminStateUp(value bool) *PortOptsApplyConfiguration {
	b.AdminStateUp = &value
	return b
}

// WithMACAddress sets the MACAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MACAddress field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithMACAddress(value string) *PortOptsApplyConfiguration {
	b.MACAddress = &value
	return b
}

// WithFixedIPs adds the given value to the FixedIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FixedIPs field.
func (b *PortOptsApplyConfiguration) WithFixedIPs(values ...*FixedIPApplyConfiguration) *PortOptsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFixedIPs")
		}
		b.FixedIPs = append(b.FixedIPs, *values[i])
	}
	return b
}

// WithTenantID sets the TenantID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TenantID field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithTenantID(value string) *PortOptsApplyConfiguration {
	b.TenantID = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithProjectID(value string) *PortOptsApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithSecurityGroups adds the given value to the SecurityGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SecurityGroups field.
func (b *PortOptsApplyConfiguration) WithSecurityGroups(values ...string) *PortOptsApplyConfiguration {
	for i := range values {
		b.SecurityGroups = append(b.SecurityGroups, values[i])
	}
	return b
}

// WithSecurityGroupFilters adds the given value to the SecurityGroupFilters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SecurityGroupFilters field.
func (b *PortOptsApplyConfiguration) WithSecurityGroupFilters(values ...*SecurityGroupParamApplyConfiguration) *PortOptsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSecurityGroupFilters")
		}
		b.SecurityGroupFilters = append(b.SecurityGroupFilters, *values[i])
	}
	return b
}

// WithAllowedAddressPairs adds the given value to the AllowedAddressPairs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedAddressPairs field.
func (b *PortOptsApplyConfiguration) WithAllowedAddressPairs(values ...*AddressPairApplyConfiguration) *PortOptsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllowedAddressPairs")
		}
		b.AllowedAddressPairs = append(b.AllowedAddressPairs, *values[i])
	}
	return b
}

// WithTrunk sets the Trunk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Trunk field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithTrunk(value bool) *PortOptsApplyConfiguration {
	b.Trunk = &value
	return b
}

// WithHostID sets the HostID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostID field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithHostID(value string) *PortOptsApplyConfiguration {
	b.HostID = &value
	return b
}

// WithVNICType sets the VNICType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VNICType field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithVNICType(value string) *PortOptsApplyConfiguration {
	b.VNICType = &value
	return b
}

// WithProfile puts the entries into the Profile field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Profile field,
// overwriting an existing map entries in Profile field with the same key.
func (b *PortOptsApplyConfiguration) WithProfile(entries map[string]string) *PortOptsApplyConfiguration {
	if b.Profile == nil && len(entries) > 0 {
		b.Profile = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Profile[k] = v
	}
	return b
}

// WithDisablePortSecurity sets the DisablePortSecurity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisablePortSecurity field is set to the value of the last call.
func (b *PortOptsApplyConfiguration) WithDisablePortSecurity(value bool) *PortOptsApplyConfiguration {
	b.DisablePortSecurity = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *PortOptsApplyConfiguration) WithTags(values ...string) *PortOptsApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithValueSpecs adds the given value to the ValueSpecs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ValueSpecs field.
func (b *PortOptsApplyConfiguration) WithValueSpecs(values ...*ValueSpecApplyConfiguration) *PortOptsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithValueSpecs")
		}
		b.ValueSpecs = append(b.ValueSpecs, *values[i])
	}
	return b
}

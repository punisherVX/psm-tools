/*
 * Security API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// SecurityNetworkSecurityPolicySpec struct for SecurityNetworkSecurityPolicySpec
type SecurityNetworkSecurityPolicySpec struct {
	// list of security groups this policy is attached to.
	AttachGroups *[]string `json:"attach-groups,omitempty"`
	// specifies if the set of rules need to be attached globally to a tenant.
	AttachTenant *bool `json:"attach-tenant,omitempty"`
	// Policy priority. If not specified, it will be assigned automatically in increments of 1000.
	Priority *int64 `json:"priority,omitempty"`
	// list of rules.
	Rules *[]SecuritySGRule `json:"rules,omitempty"`
}

// NewSecurityNetworkSecurityPolicySpec instantiates a new SecurityNetworkSecurityPolicySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityNetworkSecurityPolicySpec() *SecurityNetworkSecurityPolicySpec {
	this := SecurityNetworkSecurityPolicySpec{}
	return &this
}

// NewSecurityNetworkSecurityPolicySpecWithDefaults instantiates a new SecurityNetworkSecurityPolicySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityNetworkSecurityPolicySpecWithDefaults() *SecurityNetworkSecurityPolicySpec {
	this := SecurityNetworkSecurityPolicySpec{}
	return &this
}

// GetAttachGroups returns the AttachGroups field value if set, zero value otherwise.
func (o *SecurityNetworkSecurityPolicySpec) GetAttachGroups() []string {
	if o == nil || o.AttachGroups == nil {
		var ret []string
		return ret
	}
	return *o.AttachGroups
}

// GetAttachGroupsOk returns a tuple with the AttachGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityNetworkSecurityPolicySpec) GetAttachGroupsOk() (*[]string, bool) {
	if o == nil || o.AttachGroups == nil {
		return nil, false
	}
	return o.AttachGroups, true
}

// HasAttachGroups returns a boolean if a field has been set.
func (o *SecurityNetworkSecurityPolicySpec) HasAttachGroups() bool {
	if o != nil && o.AttachGroups != nil {
		return true
	}

	return false
}

// SetAttachGroups gets a reference to the given []string and assigns it to the AttachGroups field.
func (o *SecurityNetworkSecurityPolicySpec) SetAttachGroups(v []string) {
	o.AttachGroups = &v
}

// GetAttachTenant returns the AttachTenant field value if set, zero value otherwise.
func (o *SecurityNetworkSecurityPolicySpec) GetAttachTenant() bool {
	if o == nil || o.AttachTenant == nil {
		var ret bool
		return ret
	}
	return *o.AttachTenant
}

// GetAttachTenantOk returns a tuple with the AttachTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityNetworkSecurityPolicySpec) GetAttachTenantOk() (*bool, bool) {
	if o == nil || o.AttachTenant == nil {
		return nil, false
	}
	return o.AttachTenant, true
}

// HasAttachTenant returns a boolean if a field has been set.
func (o *SecurityNetworkSecurityPolicySpec) HasAttachTenant() bool {
	if o != nil && o.AttachTenant != nil {
		return true
	}

	return false
}

// SetAttachTenant gets a reference to the given bool and assigns it to the AttachTenant field.
func (o *SecurityNetworkSecurityPolicySpec) SetAttachTenant(v bool) {
	o.AttachTenant = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SecurityNetworkSecurityPolicySpec) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityNetworkSecurityPolicySpec) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SecurityNetworkSecurityPolicySpec) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *SecurityNetworkSecurityPolicySpec) SetPriority(v int64) {
	o.Priority = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SecurityNetworkSecurityPolicySpec) GetRules() []SecuritySGRule {
	if o == nil || o.Rules == nil {
		var ret []SecuritySGRule
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityNetworkSecurityPolicySpec) GetRulesOk() (*[]SecuritySGRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SecurityNetworkSecurityPolicySpec) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SecuritySGRule and assigns it to the Rules field.
func (o *SecurityNetworkSecurityPolicySpec) SetRules(v []SecuritySGRule) {
	o.Rules = &v
}

func (o SecurityNetworkSecurityPolicySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachGroups != nil {
		toSerialize["attach-groups"] = o.AttachGroups
	}
	if o.AttachTenant != nil {
		toSerialize["attach-tenant"] = o.AttachTenant
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityNetworkSecurityPolicySpec struct {
	value *SecurityNetworkSecurityPolicySpec
	isSet bool
}

func (v NullableSecurityNetworkSecurityPolicySpec) Get() *SecurityNetworkSecurityPolicySpec {
	return v.value
}

func (v *NullableSecurityNetworkSecurityPolicySpec) Set(val *SecurityNetworkSecurityPolicySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityNetworkSecurityPolicySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityNetworkSecurityPolicySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityNetworkSecurityPolicySpec(val *SecurityNetworkSecurityPolicySpec) *NullableSecurityNetworkSecurityPolicySpec {
	return &NullableSecurityNetworkSecurityPolicySpec{value: val, isSet: true}
}

func (v NullableSecurityNetworkSecurityPolicySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityNetworkSecurityPolicySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


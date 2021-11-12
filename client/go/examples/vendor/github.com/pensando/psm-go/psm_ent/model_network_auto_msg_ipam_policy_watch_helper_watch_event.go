/*
 * Network API reference
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

// NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent struct for NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent
type NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent struct {
	Object *NetworkIPAMPolicy `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent instantiates a new NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent() *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent {
	this := NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgIPAMPolicyWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgIPAMPolicyWatchHelperWatchEventWithDefaults() *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent {
	this := NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) GetObject() NetworkIPAMPolicy {
	if o == nil || o.Object == nil {
		var ret NetworkIPAMPolicy
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) GetObjectOk() (*NetworkIPAMPolicy, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkIPAMPolicy and assigns it to the Object field.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) SetObject(v NetworkIPAMPolicy) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent struct {
	value *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) Get() *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) Set(val *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent(val *NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) *NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// SecurityAutoMsgSecurityGroupWatchHelperWatchEvent struct for SecurityAutoMsgSecurityGroupWatchHelperWatchEvent
type SecurityAutoMsgSecurityGroupWatchHelperWatchEvent struct {
	Object *SecuritySecurityGroup `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSecurityAutoMsgSecurityGroupWatchHelperWatchEvent instantiates a new SecurityAutoMsgSecurityGroupWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgSecurityGroupWatchHelperWatchEvent() *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent {
	this := SecurityAutoMsgSecurityGroupWatchHelperWatchEvent{}
	return &this
}

// NewSecurityAutoMsgSecurityGroupWatchHelperWatchEventWithDefaults instantiates a new SecurityAutoMsgSecurityGroupWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgSecurityGroupWatchHelperWatchEventWithDefaults() *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent {
	this := SecurityAutoMsgSecurityGroupWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) GetObject() SecuritySecurityGroup {
	if o == nil || o.Object == nil {
		var ret SecuritySecurityGroup
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) GetObjectOk() (*SecuritySecurityGroup, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given SecuritySecurityGroup and assigns it to the Object field.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) SetObject(v SecuritySecurityGroup) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent struct {
	value *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent
	isSet bool
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) Get() *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent {
	return v.value
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) Set(val *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent(val *SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) *NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent {
	return &NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


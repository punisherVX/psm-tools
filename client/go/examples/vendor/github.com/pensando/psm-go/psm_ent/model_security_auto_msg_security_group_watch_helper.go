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

// SecurityAutoMsgSecurityGroupWatchHelper AutoMsgSecurityGroupWatchHelper is a wrapper object for watch events for SecurityGroup objects.
type SecurityAutoMsgSecurityGroupWatchHelper struct {
	Events *[]SecurityAutoMsgSecurityGroupWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewSecurityAutoMsgSecurityGroupWatchHelper instantiates a new SecurityAutoMsgSecurityGroupWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgSecurityGroupWatchHelper() *SecurityAutoMsgSecurityGroupWatchHelper {
	this := SecurityAutoMsgSecurityGroupWatchHelper{}
	return &this
}

// NewSecurityAutoMsgSecurityGroupWatchHelperWithDefaults instantiates a new SecurityAutoMsgSecurityGroupWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgSecurityGroupWatchHelperWithDefaults() *SecurityAutoMsgSecurityGroupWatchHelper {
	this := SecurityAutoMsgSecurityGroupWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SecurityAutoMsgSecurityGroupWatchHelper) GetEvents() []SecurityAutoMsgSecurityGroupWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []SecurityAutoMsgSecurityGroupWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelper) GetEventsOk() (*[]SecurityAutoMsgSecurityGroupWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SecurityAutoMsgSecurityGroupWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SecurityAutoMsgSecurityGroupWatchHelperWatchEvent and assigns it to the Events field.
func (o *SecurityAutoMsgSecurityGroupWatchHelper) SetEvents(v []SecurityAutoMsgSecurityGroupWatchHelperWatchEvent) {
	o.Events = &v
}

func (o SecurityAutoMsgSecurityGroupWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgSecurityGroupWatchHelper struct {
	value *SecurityAutoMsgSecurityGroupWatchHelper
	isSet bool
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelper) Get() *SecurityAutoMsgSecurityGroupWatchHelper {
	return v.value
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelper) Set(val *SecurityAutoMsgSecurityGroupWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgSecurityGroupWatchHelper(val *SecurityAutoMsgSecurityGroupWatchHelper) *NullableSecurityAutoMsgSecurityGroupWatchHelper {
	return &NullableSecurityAutoMsgSecurityGroupWatchHelper{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgSecurityGroupWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgSecurityGroupWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



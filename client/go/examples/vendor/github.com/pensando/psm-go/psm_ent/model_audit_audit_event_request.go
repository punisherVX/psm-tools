/*
 * Audit API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// AuditAuditEventRequest Request for an audit event.
type AuditAuditEventRequest struct {
	Uuid *string `json:"uuid,omitempty"`
}

// NewAuditAuditEventRequest instantiates a new AuditAuditEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditAuditEventRequest() *AuditAuditEventRequest {
	this := AuditAuditEventRequest{}
	return &this
}

// NewAuditAuditEventRequestWithDefaults instantiates a new AuditAuditEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditAuditEventRequestWithDefaults() *AuditAuditEventRequest {
	this := AuditAuditEventRequest{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AuditAuditEventRequest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEventRequest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AuditAuditEventRequest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AuditAuditEventRequest) SetUuid(v string) {
	o.Uuid = &v
}

func (o AuditAuditEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableAuditAuditEventRequest struct {
	value *AuditAuditEventRequest
	isSet bool
}

func (v NullableAuditAuditEventRequest) Get() *AuditAuditEventRequest {
	return v.value
}

func (v *NullableAuditAuditEventRequest) Set(val *AuditAuditEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditAuditEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditAuditEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditAuditEventRequest(val *AuditAuditEventRequest) *NullableAuditAuditEventRequest {
	return &NullableAuditAuditEventRequest{value: val, isSet: true}
}

func (v NullableAuditAuditEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditAuditEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


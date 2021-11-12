/*
 * Routing API reference
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

// RoutingNeighborList struct for RoutingNeighborList
type RoutingNeighborList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Items *[]RoutingNeighbor `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
}

// NewRoutingNeighborList instantiates a new RoutingNeighborList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingNeighborList() *RoutingNeighborList {
	this := RoutingNeighborList{}
	return &this
}

// NewRoutingNeighborListWithDefaults instantiates a new RoutingNeighborList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingNeighborListWithDefaults() *RoutingNeighborList {
	this := RoutingNeighborList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *RoutingNeighborList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighborList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *RoutingNeighborList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *RoutingNeighborList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RoutingNeighborList) GetItems() []RoutingNeighbor {
	if o == nil || o.Items == nil {
		var ret []RoutingNeighbor
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighborList) GetItemsOk() (*[]RoutingNeighbor, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RoutingNeighborList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RoutingNeighbor and assigns it to the Items field.
func (o *RoutingNeighborList) SetItems(v []RoutingNeighbor) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RoutingNeighborList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighborList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RoutingNeighborList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RoutingNeighborList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *RoutingNeighborList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighborList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *RoutingNeighborList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *RoutingNeighborList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

func (o RoutingNeighborList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ListMeta != nil {
		toSerialize["list-meta"] = o.ListMeta
	}
	return json.Marshal(toSerialize)
}

type NullableRoutingNeighborList struct {
	value *RoutingNeighborList
	isSet bool
}

func (v NullableRoutingNeighborList) Get() *RoutingNeighborList {
	return v.value
}

func (v *NullableRoutingNeighborList) Set(val *RoutingNeighborList) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingNeighborList) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingNeighborList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingNeighborList(val *RoutingNeighborList) *NullableRoutingNeighborList {
	return &NullableRoutingNeighborList{value: val, isSet: true}
}

func (v NullableRoutingNeighborList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingNeighborList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


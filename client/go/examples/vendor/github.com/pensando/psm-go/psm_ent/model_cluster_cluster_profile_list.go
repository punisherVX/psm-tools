/*
 * Cluster API reference
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

// ClusterClusterProfileList ClusterProfileList is a container object for list of ClusterProfile objects.
type ClusterClusterProfileList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// List of ClusterProfile objects.
	Items *[]ClusterClusterProfile `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
}

// NewClusterClusterProfileList instantiates a new ClusterClusterProfileList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterClusterProfileList() *ClusterClusterProfileList {
	this := ClusterClusterProfileList{}
	return &this
}

// NewClusterClusterProfileListWithDefaults instantiates a new ClusterClusterProfileList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterClusterProfileListWithDefaults() *ClusterClusterProfileList {
	this := ClusterClusterProfileList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterClusterProfileList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterProfileList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterClusterProfileList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterClusterProfileList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterClusterProfileList) GetItems() []ClusterClusterProfile {
	if o == nil || o.Items == nil {
		var ret []ClusterClusterProfile
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterProfileList) GetItemsOk() (*[]ClusterClusterProfile, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterClusterProfileList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClusterClusterProfile and assigns it to the Items field.
func (o *ClusterClusterProfileList) SetItems(v []ClusterClusterProfile) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterClusterProfileList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterProfileList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterClusterProfileList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterClusterProfileList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *ClusterClusterProfileList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterProfileList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *ClusterClusterProfileList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *ClusterClusterProfileList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

func (o ClusterClusterProfileList) MarshalJSON() ([]byte, error) {
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

type NullableClusterClusterProfileList struct {
	value *ClusterClusterProfileList
	isSet bool
}

func (v NullableClusterClusterProfileList) Get() *ClusterClusterProfileList {
	return v.value
}

func (v *NullableClusterClusterProfileList) Set(val *ClusterClusterProfileList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterClusterProfileList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterClusterProfileList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterClusterProfileList(val *ClusterClusterProfileList) *NullableClusterClusterProfileList {
	return &NullableClusterClusterProfileList{value: val, isSet: true}
}

func (v NullableClusterClusterProfileList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterClusterProfileList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Data views

OpenAPI schema for data view endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_views

import (
	"encoding/json"
)

// checks if the GetAllDataViews200ResponseDataViewInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllDataViews200ResponseDataViewInner{}

// GetAllDataViews200ResponseDataViewInner struct for GetAllDataViews200ResponseDataViewInner
type GetAllDataViews200ResponseDataViewInner struct {
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Namespaces []string               `json:"namespaces,omitempty"`
	Title      *string                `json:"title,omitempty"`
	TypeMeta   map[string]interface{} `json:"typeMeta,omitempty"`
}

// NewGetAllDataViews200ResponseDataViewInner instantiates a new GetAllDataViews200ResponseDataViewInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllDataViews200ResponseDataViewInner() *GetAllDataViews200ResponseDataViewInner {
	this := GetAllDataViews200ResponseDataViewInner{}
	return &this
}

// NewGetAllDataViews200ResponseDataViewInnerWithDefaults instantiates a new GetAllDataViews200ResponseDataViewInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllDataViews200ResponseDataViewInnerWithDefaults() *GetAllDataViews200ResponseDataViewInner {
	this := GetAllDataViews200ResponseDataViewInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAllDataViews200ResponseDataViewInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDataViews200ResponseDataViewInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAllDataViews200ResponseDataViewInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAllDataViews200ResponseDataViewInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAllDataViews200ResponseDataViewInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDataViews200ResponseDataViewInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAllDataViews200ResponseDataViewInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAllDataViews200ResponseDataViewInner) SetName(v string) {
	o.Name = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *GetAllDataViews200ResponseDataViewInner) GetNamespaces() []string {
	if o == nil || IsNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDataViews200ResponseDataViewInner) GetNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *GetAllDataViews200ResponseDataViewInner) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *GetAllDataViews200ResponseDataViewInner) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetAllDataViews200ResponseDataViewInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDataViews200ResponseDataViewInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetAllDataViews200ResponseDataViewInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetAllDataViews200ResponseDataViewInner) SetTitle(v string) {
	o.Title = &v
}

// GetTypeMeta returns the TypeMeta field value if set, zero value otherwise.
func (o *GetAllDataViews200ResponseDataViewInner) GetTypeMeta() map[string]interface{} {
	if o == nil || IsNil(o.TypeMeta) {
		var ret map[string]interface{}
		return ret
	}
	return o.TypeMeta
}

// GetTypeMetaOk returns a tuple with the TypeMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDataViews200ResponseDataViewInner) GetTypeMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TypeMeta) {
		return map[string]interface{}{}, false
	}
	return o.TypeMeta, true
}

// HasTypeMeta returns a boolean if a field has been set.
func (o *GetAllDataViews200ResponseDataViewInner) HasTypeMeta() bool {
	if o != nil && !IsNil(o.TypeMeta) {
		return true
	}

	return false
}

// SetTypeMeta gets a reference to the given map[string]interface{} and assigns it to the TypeMeta field.
func (o *GetAllDataViews200ResponseDataViewInner) SetTypeMeta(v map[string]interface{}) {
	o.TypeMeta = v
}

func (o GetAllDataViews200ResponseDataViewInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllDataViews200ResponseDataViewInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TypeMeta) {
		toSerialize["typeMeta"] = o.TypeMeta
	}
	return toSerialize, nil
}

type NullableGetAllDataViews200ResponseDataViewInner struct {
	value *GetAllDataViews200ResponseDataViewInner
	isSet bool
}

func (v NullableGetAllDataViews200ResponseDataViewInner) Get() *GetAllDataViews200ResponseDataViewInner {
	return v.value
}

func (v *NullableGetAllDataViews200ResponseDataViewInner) Set(val *GetAllDataViews200ResponseDataViewInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllDataViews200ResponseDataViewInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllDataViews200ResponseDataViewInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllDataViews200ResponseDataViewInner(val *GetAllDataViews200ResponseDataViewInner) *NullableGetAllDataViews200ResponseDataViewInner {
	return &NullableGetAllDataViews200ResponseDataViewInner{value: val, isSet: true}
}

func (v NullableGetAllDataViews200ResponseDataViewInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllDataViews200ResponseDataViewInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

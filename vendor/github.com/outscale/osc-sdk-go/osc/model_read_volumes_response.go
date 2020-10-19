/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadVolumesResponse struct for ReadVolumesResponse
type ReadVolumesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more volumes.
	Volumes *[]Volume `json:"Volumes,omitempty"`
}

// NewReadVolumesResponse instantiates a new ReadVolumesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVolumesResponse() *ReadVolumesResponse {
	this := ReadVolumesResponse{}
	return &this
}

// NewReadVolumesResponseWithDefaults instantiates a new ReadVolumesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVolumesResponseWithDefaults() *ReadVolumesResponse {
	this := ReadVolumesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVolumesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVolumesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVolumesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVolumesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ReadVolumesResponse) GetVolumes() []Volume {
	if o == nil || o.Volumes == nil {
		var ret []Volume
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVolumesResponse) GetVolumesOk() (*[]Volume, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ReadVolumesResponse) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volume and assigns it to the Volumes field.
func (o *ReadVolumesResponse) SetVolumes(v []Volume) {
	o.Volumes = &v
}

func (o ReadVolumesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}
	return json.Marshal(toSerialize)
}

type NullableReadVolumesResponse struct {
	value *ReadVolumesResponse
	isSet bool
}

func (v NullableReadVolumesResponse) Get() *ReadVolumesResponse {
	return v.value
}

func (v *NullableReadVolumesResponse) Set(val *ReadVolumesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVolumesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVolumesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVolumesResponse(val *ReadVolumesResponse) *NullableReadVolumesResponse {
	return &NullableReadVolumesResponse{value: val, isSet: true}
}

func (v NullableReadVolumesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVolumesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



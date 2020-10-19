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

// CreateVpnConnectionResponse struct for CreateVpnConnectionResponse
type CreateVpnConnectionResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	VpnConnection *VpnConnection `json:"VpnConnection,omitempty"`
}

// NewCreateVpnConnectionResponse instantiates a new CreateVpnConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVpnConnectionResponse() *CreateVpnConnectionResponse {
	this := CreateVpnConnectionResponse{}
	return &this
}

// NewCreateVpnConnectionResponseWithDefaults instantiates a new CreateVpnConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVpnConnectionResponseWithDefaults() *CreateVpnConnectionResponse {
	this := CreateVpnConnectionResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateVpnConnectionResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateVpnConnectionResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateVpnConnectionResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVpnConnection returns the VpnConnection field value if set, zero value otherwise.
func (o *CreateVpnConnectionResponse) GetVpnConnection() VpnConnection {
	if o == nil || o.VpnConnection == nil {
		var ret VpnConnection
		return ret
	}
	return *o.VpnConnection
}

// GetVpnConnectionOk returns a tuple with the VpnConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVpnConnectionResponse) GetVpnConnectionOk() (*VpnConnection, bool) {
	if o == nil || o.VpnConnection == nil {
		return nil, false
	}
	return o.VpnConnection, true
}

// HasVpnConnection returns a boolean if a field has been set.
func (o *CreateVpnConnectionResponse) HasVpnConnection() bool {
	if o != nil && o.VpnConnection != nil {
		return true
	}

	return false
}

// SetVpnConnection gets a reference to the given VpnConnection and assigns it to the VpnConnection field.
func (o *CreateVpnConnectionResponse) SetVpnConnection(v VpnConnection) {
	o.VpnConnection = &v
}

func (o CreateVpnConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VpnConnection != nil {
		toSerialize["VpnConnection"] = o.VpnConnection
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVpnConnectionResponse struct {
	value *CreateVpnConnectionResponse
	isSet bool
}

func (v NullableCreateVpnConnectionResponse) Get() *CreateVpnConnectionResponse {
	return v.value
}

func (v *NullableCreateVpnConnectionResponse) Set(val *CreateVpnConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVpnConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVpnConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVpnConnectionResponse(val *CreateVpnConnectionResponse) *NullableCreateVpnConnectionResponse {
	return &NullableCreateVpnConnectionResponse{value: val, isSet: true}
}

func (v NullableCreateVpnConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVpnConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



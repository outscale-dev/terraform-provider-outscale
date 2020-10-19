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

// Account Information about the account.
type Account struct {
	// The ID of the account.
	AccountId *string `json:"AccountId,omitempty"`
	// The city of the account owner.
	City *string `json:"City,omitempty"`
	// The name of the company for the account.
	CompanyName *string `json:"CompanyName,omitempty"`
	// The country of the account owner.
	Country *string `json:"Country,omitempty"`
	// The ID of the customer.
	CustomerId *string `json:"CustomerId,omitempty"`
	// The email address for the account.
	Email *string `json:"Email,omitempty"`
	// The first name of the account owner.
	FirstName *string `json:"FirstName,omitempty"`
	// The job title of the account owner.
	JobTitle *string `json:"JobTitle,omitempty"`
	// The last name of the account owner.
	LastName *string `json:"LastName,omitempty"`
	// The mobile phone number of the account owner.
	MobileNumber *string `json:"MobileNumber,omitempty"`
	// The landline phone number of the account owner.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The state/province of the account.
	StateProvince *string `json:"StateProvince,omitempty"`
	// The value added tax (VAT) number for the account.
	VatNumber *string `json:"VatNumber,omitempty"`
	// The ZIP code of the city.
	ZipCode *string `json:"ZipCode,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Account) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Account) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Account) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Account) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Account) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Account) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Account) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Account) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Account) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Account) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Account) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Account) SetCountry(v string) {
	o.Country = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *Account) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Account) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *Account) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Account) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Account) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Account) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Account) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Account) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Account) SetFirstName(v string) {
	o.FirstName = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Account) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Account) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Account) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Account) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Account) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Account) SetLastName(v string) {
	o.LastName = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *Account) GetMobileNumber() string {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetMobileNumberOk() (*string, bool) {
	if o == nil || o.MobileNumber == nil {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *Account) HasMobileNumber() bool {
	if o != nil && o.MobileNumber != nil {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *Account) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Account) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Account) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Account) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *Account) GetStateProvince() string {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetStateProvinceOk() (*string, bool) {
	if o == nil || o.StateProvince == nil {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *Account) HasStateProvince() bool {
	if o != nil && o.StateProvince != nil {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *Account) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Account) GetVatNumber() string {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetVatNumberOk() (*string, bool) {
	if o == nil || o.VatNumber == nil {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Account) HasVatNumber() bool {
	if o != nil && o.VatNumber != nil {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Account) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *Account) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *Account) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *Account) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.City != nil {
		toSerialize["City"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["CompanyName"] = o.CompanyName
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.CustomerId != nil {
		toSerialize["CustomerId"] = o.CustomerId
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["FirstName"] = o.FirstName
	}
	if o.JobTitle != nil {
		toSerialize["JobTitle"] = o.JobTitle
	}
	if o.LastName != nil {
		toSerialize["LastName"] = o.LastName
	}
	if o.MobileNumber != nil {
		toSerialize["MobileNumber"] = o.MobileNumber
	}
	if o.PhoneNumber != nil {
		toSerialize["PhoneNumber"] = o.PhoneNumber
	}
	if o.StateProvince != nil {
		toSerialize["StateProvince"] = o.StateProvince
	}
	if o.VatNumber != nil {
		toSerialize["VatNumber"] = o.VatNumber
	}
	if o.ZipCode != nil {
		toSerialize["ZipCode"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



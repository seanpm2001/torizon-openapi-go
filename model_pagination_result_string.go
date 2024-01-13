/*
Torizon OTA

 This API is rate limited and will return the following headers for each API call.    - X-RateLimit-Limit - The total number of requests allowed within a time period   - X-RateLimit-Remaining - The total number of requests still allowed until the end of the rate limiting period   - X-RateLimit-Reset - The number of seconds until the limit is fully reset  In addition, if an API client is rate limited, it will receive a HTTP 420 response with the following header:     - Retry-After - The number of seconds to wait until this request is allowed  

API version: 2.0-Beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PaginationResultString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResultString{}

// PaginationResultString struct for PaginationResultString
type PaginationResultString struct {
	Values []string `json:"values,omitempty"`
	Total int64 `json:"total"`
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
}

type _PaginationResultString PaginationResultString

// NewPaginationResultString instantiates a new PaginationResultString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResultString(total int64, offset int64, limit int64) *PaginationResultString {
	this := PaginationResultString{}
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	return &this
}

// NewPaginationResultStringWithDefaults instantiates a new PaginationResultString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResultStringWithDefaults() *PaginationResultString {
	this := PaginationResultString{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PaginationResultString) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResultString) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PaginationResultString) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *PaginationResultString) SetValues(v []string) {
	o.Values = v
}

// GetTotal returns the Total field value
func (o *PaginationResultString) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationResultString) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginationResultString) SetTotal(v int64) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *PaginationResultString) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PaginationResultString) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PaginationResultString) SetOffset(v int64) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *PaginationResultString) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginationResultString) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PaginationResultString) SetLimit(v int64) {
	o.Limit = v
}

func (o PaginationResultString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationResultString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	toSerialize["total"] = o.Total
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *PaginationResultString) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"offset",
		"limit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginationResultString := _PaginationResultString{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationResultString)

	if err != nil {
		return err
	}

	*o = PaginationResultString(varPaginationResultString)

	return err
}

type NullablePaginationResultString struct {
	value *PaginationResultString
	isSet bool
}

func (v NullablePaginationResultString) Get() *PaginationResultString {
	return v.value
}

func (v *NullablePaginationResultString) Set(val *PaginationResultString) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResultString) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResultString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResultString(val *PaginationResultString) *NullablePaginationResultString {
	return &NullablePaginationResultString{value: val, isSet: true}
}

func (v NullablePaginationResultString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResultString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


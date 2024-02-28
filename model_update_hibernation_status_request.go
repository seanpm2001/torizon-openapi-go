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

// checks if the UpdateHibernationStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHibernationStatusRequest{}

// UpdateHibernationStatusRequest struct for UpdateHibernationStatusRequest
type UpdateHibernationStatusRequest struct {
	Status bool `json:"status"`
}

type _UpdateHibernationStatusRequest UpdateHibernationStatusRequest

// NewUpdateHibernationStatusRequest instantiates a new UpdateHibernationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHibernationStatusRequest(status bool) *UpdateHibernationStatusRequest {
	this := UpdateHibernationStatusRequest{}
	this.Status = status
	return &this
}

// NewUpdateHibernationStatusRequestWithDefaults instantiates a new UpdateHibernationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHibernationStatusRequestWithDefaults() *UpdateHibernationStatusRequest {
	this := UpdateHibernationStatusRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateHibernationStatusRequest) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateHibernationStatusRequest) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateHibernationStatusRequest) SetStatus(v bool) {
	o.Status = v
}

func (o UpdateHibernationStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHibernationStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *UpdateHibernationStatusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varUpdateHibernationStatusRequest := _UpdateHibernationStatusRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateHibernationStatusRequest)

	if err != nil {
		return err
	}

	*o = UpdateHibernationStatusRequest(varUpdateHibernationStatusRequest)

	return err
}

type NullableUpdateHibernationStatusRequest struct {
	value *UpdateHibernationStatusRequest
	isSet bool
}

func (v NullableUpdateHibernationStatusRequest) Get() *UpdateHibernationStatusRequest {
	return v.value
}

func (v *NullableUpdateHibernationStatusRequest) Set(val *UpdateHibernationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHibernationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHibernationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHibernationStatusRequest(val *UpdateHibernationStatusRequest) *NullableUpdateHibernationStatusRequest {
	return &NullableUpdateHibernationStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateHibernationStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHibernationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


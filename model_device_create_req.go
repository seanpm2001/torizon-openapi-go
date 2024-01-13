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

// checks if the DeviceCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCreateReq{}

// DeviceCreateReq struct for DeviceCreateReq
type DeviceCreateReq struct {
	DeviceName *string `json:"deviceName,omitempty"`
	DeviceId string `json:"deviceId"`
}

type _DeviceCreateReq DeviceCreateReq

// NewDeviceCreateReq instantiates a new DeviceCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateReq(deviceId string) *DeviceCreateReq {
	this := DeviceCreateReq{}
	this.DeviceId = deviceId
	return &this
}

// NewDeviceCreateReqWithDefaults instantiates a new DeviceCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateReqWithDefaults() *DeviceCreateReq {
	this := DeviceCreateReq{}
	return &this
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *DeviceCreateReq) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateReq) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *DeviceCreateReq) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *DeviceCreateReq) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceId returns the DeviceId field value
func (o *DeviceCreateReq) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateReq) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *DeviceCreateReq) SetDeviceId(v string) {
	o.DeviceId = v
}

func (o DeviceCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	toSerialize["deviceId"] = o.DeviceId
	return toSerialize, nil
}

func (o *DeviceCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceId",
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

	varDeviceCreateReq := _DeviceCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceCreateReq)

	if err != nil {
		return err
	}

	*o = DeviceCreateReq(varDeviceCreateReq)

	return err
}

type NullableDeviceCreateReq struct {
	value *DeviceCreateReq
	isSet bool
}

func (v NullableDeviceCreateReq) Get() *DeviceCreateReq {
	return v.value
}

func (v *NullableDeviceCreateReq) Set(val *DeviceCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateReq(val *DeviceCreateReq) *NullableDeviceCreateReq {
	return &NullableDeviceCreateReq{value: val, isSet: true}
}

func (v NullableDeviceCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


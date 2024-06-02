/*
Torizon OTA

 This API is rate limited and will return the following headers for each API call.    - X-RateLimit-Limit - The total number of requests allowed within a time period   - X-RateLimit-Remaining - The total number of requests still allowed until the end of the rate limiting period   - X-RateLimit-Reset - The number of seconds until the limit is fully reset  In addition, if an API client is rate limited, it will receive a HTTP 420 response with the following header:     - Retry-After - The number of seconds to wait until this request is allowed  

API version: 2.0-Beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the DeviceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfo{}

// DeviceInfo struct for DeviceInfo
type DeviceInfo struct {
	DeviceId string `json:"deviceId"`
	LastApiConnect time.Time `json:"lastApiConnect"`
	LastSshSessionConnect LastSshSession `json:"lastSshSessionConnect"`
	LastVersion string `json:"lastVersion"`
	LastUserConnect UserInfo `json:"lastUserConnect"`
}

type _DeviceInfo DeviceInfo

// NewDeviceInfo instantiates a new DeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfo(deviceId string, lastApiConnect time.Time, lastSshSessionConnect LastSshSession, lastVersion string, lastUserConnect UserInfo) *DeviceInfo {
	this := DeviceInfo{}
	this.DeviceId = deviceId
	this.LastApiConnect = lastApiConnect
	this.LastSshSessionConnect = lastSshSessionConnect
	this.LastVersion = lastVersion
	this.LastUserConnect = lastUserConnect
	return &this
}

// NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoWithDefaults() *DeviceInfo {
	this := DeviceInfo{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *DeviceInfo) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *DeviceInfo) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetLastApiConnect returns the LastApiConnect field value
func (o *DeviceInfo) GetLastApiConnect() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastApiConnect
}

// GetLastApiConnectOk returns a tuple with the LastApiConnect field value
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetLastApiConnectOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastApiConnect, true
}

// SetLastApiConnect sets field value
func (o *DeviceInfo) SetLastApiConnect(v time.Time) {
	o.LastApiConnect = v
}

// GetLastSshSessionConnect returns the LastSshSessionConnect field value
func (o *DeviceInfo) GetLastSshSessionConnect() LastSshSession {
	if o == nil {
		var ret LastSshSession
		return ret
	}

	return o.LastSshSessionConnect
}

// GetLastSshSessionConnectOk returns a tuple with the LastSshSessionConnect field value
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetLastSshSessionConnectOk() (*LastSshSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSshSessionConnect, true
}

// SetLastSshSessionConnect sets field value
func (o *DeviceInfo) SetLastSshSessionConnect(v LastSshSession) {
	o.LastSshSessionConnect = v
}

// GetLastVersion returns the LastVersion field value
func (o *DeviceInfo) GetLastVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetLastVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastVersion, true
}

// SetLastVersion sets field value
func (o *DeviceInfo) SetLastVersion(v string) {
	o.LastVersion = v
}

// GetLastUserConnect returns the LastUserConnect field value
func (o *DeviceInfo) GetLastUserConnect() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.LastUserConnect
}

// GetLastUserConnectOk returns a tuple with the LastUserConnect field value
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetLastUserConnectOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUserConnect, true
}

// SetLastUserConnect sets field value
func (o *DeviceInfo) SetLastUserConnect(v UserInfo) {
	o.LastUserConnect = v
}

func (o DeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["lastApiConnect"] = o.LastApiConnect
	toSerialize["lastSshSessionConnect"] = o.LastSshSessionConnect
	toSerialize["lastVersion"] = o.LastVersion
	toSerialize["lastUserConnect"] = o.LastUserConnect
	return toSerialize, nil
}

func (o *DeviceInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceId",
		"lastApiConnect",
		"lastSshSessionConnect",
		"lastVersion",
		"lastUserConnect",
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

	varDeviceInfo := _DeviceInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceInfo)

	if err != nil {
		return err
	}

	*o = DeviceInfo(varDeviceInfo)

	return err
}

type NullableDeviceInfo struct {
	value *DeviceInfo
	isSet bool
}

func (v NullableDeviceInfo) Get() *DeviceInfo {
	return v.value
}

func (v *NullableDeviceInfo) Set(val *DeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfo(val *DeviceInfo) *NullableDeviceInfo {
	return &NullableDeviceInfo{value: val, isSet: true}
}

func (v NullableDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



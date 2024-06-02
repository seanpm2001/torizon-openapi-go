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

// checks if the UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfo{}

// UserInfo struct for UserInfo
type UserInfo struct {
	RemoteIp string `json:"RemoteIp"`
	ConnectedAt time.Time `json:"ConnectedAt"`
}

type _UserInfo UserInfo

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo(remoteIp string, connectedAt time.Time) *UserInfo {
	this := UserInfo{}
	this.RemoteIp = remoteIp
	this.ConnectedAt = connectedAt
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetRemoteIp returns the RemoteIp field value
func (o *UserInfo) GetRemoteIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetRemoteIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIp, true
}

// SetRemoteIp sets field value
func (o *UserInfo) SetRemoteIp(v string) {
	o.RemoteIp = v
}

// GetConnectedAt returns the ConnectedAt field value
func (o *UserInfo) GetConnectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetConnectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedAt, true
}

// SetConnectedAt sets field value
func (o *UserInfo) SetConnectedAt(v time.Time) {
	o.ConnectedAt = v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RemoteIp"] = o.RemoteIp
	toSerialize["ConnectedAt"] = o.ConnectedAt
	return toSerialize, nil
}

func (o *UserInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"RemoteIp",
		"ConnectedAt",
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

	varUserInfo := _UserInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInfo)

	if err != nil {
		return err
	}

	*o = UserInfo(varUserInfo)

	return err
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



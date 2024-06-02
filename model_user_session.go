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

// checks if the UserSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSession{}

// UserSession struct for UserSession
type UserSession struct {
	DeviceId string `json:"deviceId"`
	Session DeviceSession `json:"session"`
}

type _UserSession UserSession

// NewUserSession instantiates a new UserSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSession(deviceId string, session DeviceSession) *UserSession {
	this := UserSession{}
	this.DeviceId = deviceId
	this.Session = session
	return &this
}

// NewUserSessionWithDefaults instantiates a new UserSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSessionWithDefaults() *UserSession {
	this := UserSession{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *UserSession) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *UserSession) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *UserSession) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetSession returns the Session field value
func (o *UserSession) GetSession() DeviceSession {
	if o == nil {
		var ret DeviceSession
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *UserSession) GetSessionOk() (*DeviceSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *UserSession) SetSession(v DeviceSession) {
	o.Session = v
}

func (o UserSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["session"] = o.Session
	return toSerialize, nil
}

func (o *UserSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceId",
		"session",
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

	varUserSession := _UserSession{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSession)

	if err != nil {
		return err
	}

	*o = UserSession(varUserSession)

	return err
}

type NullableUserSession struct {
	value *UserSession
	isSet bool
}

func (v NullableUserSession) Get() *UserSession {
	return v.value
}

func (v *NullableUserSession) Set(val *UserSession) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSession) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSession(val *UserSession) *NullableUserSession {
	return &NullableUserSession{value: val, isSet: true}
}

func (v NullableUserSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



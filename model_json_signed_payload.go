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

// checks if the JsonSignedPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonSignedPayload{}

// JsonSignedPayload struct for JsonSignedPayload
type JsonSignedPayload struct {
	Signatures []ClientSignature `json:"signatures,omitempty"`
	Signed interface{} `json:"signed"`
}

type _JsonSignedPayload JsonSignedPayload

// NewJsonSignedPayload instantiates a new JsonSignedPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonSignedPayload(signed interface{}) *JsonSignedPayload {
	this := JsonSignedPayload{}
	this.Signed = signed
	return &this
}

// NewJsonSignedPayloadWithDefaults instantiates a new JsonSignedPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonSignedPayloadWithDefaults() *JsonSignedPayload {
	this := JsonSignedPayload{}
	return &this
}

// GetSignatures returns the Signatures field value if set, zero value otherwise.
func (o *JsonSignedPayload) GetSignatures() []ClientSignature {
	if o == nil || IsNil(o.Signatures) {
		var ret []ClientSignature
		return ret
	}
	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonSignedPayload) GetSignaturesOk() ([]ClientSignature, bool) {
	if o == nil || IsNil(o.Signatures) {
		return nil, false
	}
	return o.Signatures, true
}

// HasSignatures returns a boolean if a field has been set.
func (o *JsonSignedPayload) HasSignatures() bool {
	if o != nil && !IsNil(o.Signatures) {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given []ClientSignature and assigns it to the Signatures field.
func (o *JsonSignedPayload) SetSignatures(v []ClientSignature) {
	o.Signatures = v
}

// GetSigned returns the Signed field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JsonSignedPayload) GetSigned() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Signed
}

// GetSignedOk returns a tuple with the Signed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JsonSignedPayload) GetSignedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Signed) {
		return nil, false
	}
	return &o.Signed, true
}

// SetSigned sets field value
func (o *JsonSignedPayload) SetSigned(v interface{}) {
	o.Signed = v
}

func (o JsonSignedPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonSignedPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signatures) {
		toSerialize["signatures"] = o.Signatures
	}
	if o.Signed != nil {
		toSerialize["signed"] = o.Signed
	}
	return toSerialize, nil
}

func (o *JsonSignedPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signed",
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

	varJsonSignedPayload := _JsonSignedPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJsonSignedPayload)

	if err != nil {
		return err
	}

	*o = JsonSignedPayload(varJsonSignedPayload)

	return err
}

type NullableJsonSignedPayload struct {
	value *JsonSignedPayload
	isSet bool
}

func (v NullableJsonSignedPayload) Get() *JsonSignedPayload {
	return v.value
}

func (v *NullableJsonSignedPayload) Set(val *JsonSignedPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonSignedPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonSignedPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonSignedPayload(val *JsonSignedPayload) *NullableJsonSignedPayload {
	return &NullableJsonSignedPayload{value: val, isSet: true}
}

func (v NullableJsonSignedPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonSignedPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


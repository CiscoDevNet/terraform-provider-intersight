/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2025051220
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// checks if the TelemetryDruidError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidError{}

// TelemetryDruidError The error response when a Druid query fails.
type TelemetryDruidError struct {
	// A well-defined error code.
	Error *string `json:"error,omitempty"`
	// A free-form message with more information about the error. May be null.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The class of the exception that caused this error. May be null.
	ErrorClass           interface{} `json:"errorClass,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidError TelemetryDruidError

// NewTelemetryDruidError instantiates a new TelemetryDruidError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidError() *TelemetryDruidError {
	this := TelemetryDruidError{}
	return &this
}

// NewTelemetryDruidErrorWithDefaults instantiates a new TelemetryDruidError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidErrorWithDefaults() *TelemetryDruidError {
	this := TelemetryDruidError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TelemetryDruidError) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidError) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TelemetryDruidError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TelemetryDruidError) SetError(v string) {
	o.Error = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TelemetryDruidError) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidError) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TelemetryDruidError) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TelemetryDruidError) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorClass returns the ErrorClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelemetryDruidError) GetErrorClass() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorClass
}

// GetErrorClassOk returns a tuple with the ErrorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryDruidError) GetErrorClassOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorClass) {
		return nil, false
	}
	return &o.ErrorClass, true
}

// HasErrorClass returns a boolean if a field has been set.
func (o *TelemetryDruidError) HasErrorClass() bool {
	if o != nil && !IsNil(o.ErrorClass) {
		return true
	}

	return false
}

// SetErrorClass gets a reference to the given interface{} and assigns it to the ErrorClass field.
func (o *TelemetryDruidError) SetErrorClass(v interface{}) {
	o.ErrorClass = v
}

func (o TelemetryDruidError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ErrorClass != nil {
		toSerialize["errorClass"] = o.ErrorClass
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidError) UnmarshalJSON(data []byte) (err error) {
	varTelemetryDruidError := _TelemetryDruidError{}

	err = json.Unmarshal(data, &varTelemetryDruidError)

	if err != nil {
		return err
	}

	*o = TelemetryDruidError(varTelemetryDruidError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "errorClass")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidError struct {
	value *TelemetryDruidError
	isSet bool
}

func (v NullableTelemetryDruidError) Get() *TelemetryDruidError {
	return v.value
}

func (v *NullableTelemetryDruidError) Set(val *TelemetryDruidError) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidError) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidError(val *TelemetryDruidError) *NullableTelemetryDruidError {
	return &NullableTelemetryDruidError{value: val, isSet: true}
}

func (v NullableTelemetryDruidError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

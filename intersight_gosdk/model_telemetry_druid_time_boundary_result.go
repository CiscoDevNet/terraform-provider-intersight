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
	"time"
)

// checks if the TelemetryDruidTimeBoundaryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidTimeBoundaryResult{}

// TelemetryDruidTimeBoundaryResult The earliest and latest timestamps for a dataset
type TelemetryDruidTimeBoundaryResult struct {
	// The ISO 8601 timestamp.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The corresponding timestamps for query. May contain maxTime, minTime, or both (default).
	Result               map[string]interface{} `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTimeBoundaryResult TelemetryDruidTimeBoundaryResult

// NewTelemetryDruidTimeBoundaryResult instantiates a new TelemetryDruidTimeBoundaryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTimeBoundaryResult() *TelemetryDruidTimeBoundaryResult {
	this := TelemetryDruidTimeBoundaryResult{}
	return &this
}

// NewTelemetryDruidTimeBoundaryResultWithDefaults instantiates a new TelemetryDruidTimeBoundaryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTimeBoundaryResultWithDefaults() *TelemetryDruidTimeBoundaryResult {
	this := TelemetryDruidTimeBoundaryResult{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryResult) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryResult) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryResult) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TelemetryDruidTimeBoundaryResult) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryResult) GetResult() map[string]interface{} {
	if o == nil || IsNil(o.Result) {
		var ret map[string]interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryResult) GetResultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return map[string]interface{}{}, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *TelemetryDruidTimeBoundaryResult) SetResult(v map[string]interface{}) {
	o.Result = v
}

func (o TelemetryDruidTimeBoundaryResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidTimeBoundaryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidTimeBoundaryResult) UnmarshalJSON(data []byte) (err error) {
	varTelemetryDruidTimeBoundaryResult := _TelemetryDruidTimeBoundaryResult{}

	err = json.Unmarshal(data, &varTelemetryDruidTimeBoundaryResult)

	if err != nil {
		return err
	}

	*o = TelemetryDruidTimeBoundaryResult(varTelemetryDruidTimeBoundaryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTimeBoundaryResult struct {
	value *TelemetryDruidTimeBoundaryResult
	isSet bool
}

func (v NullableTelemetryDruidTimeBoundaryResult) Get() *TelemetryDruidTimeBoundaryResult {
	return v.value
}

func (v *NullableTelemetryDruidTimeBoundaryResult) Set(val *TelemetryDruidTimeBoundaryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTimeBoundaryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTimeBoundaryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTimeBoundaryResult(val *TelemetryDruidTimeBoundaryResult) *NullableTelemetryDruidTimeBoundaryResult {
	return &NullableTelemetryDruidTimeBoundaryResult{value: val, isSet: true}
}

func (v NullableTelemetryDruidTimeBoundaryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTimeBoundaryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

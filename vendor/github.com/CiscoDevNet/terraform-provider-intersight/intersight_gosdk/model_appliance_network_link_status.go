/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ApplianceNetworkLinkStatus Link between two nodes of an Intersight Appliance cluster.
type ApplianceNetworkLinkStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Hostname of the destination endpoint.
	DestinationHostname *string `json:"DestinationHostname,omitempty"`
	// Time to reach the destination endpoint in milliseconds from the source endpoint.
	PingTime *float32 `json:"PingTime,omitempty"`
	// Hostname of the source endpoint.
	SourceHostname       *string                              `json:"SourceHostname,omitempty"`
	NodeOpStatus         *ApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNetworkLinkStatus ApplianceNetworkLinkStatus

// NewApplianceNetworkLinkStatus instantiates a new ApplianceNetworkLinkStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNetworkLinkStatus(classId string, objectType string) *ApplianceNetworkLinkStatus {
	this := ApplianceNetworkLinkStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNetworkLinkStatusWithDefaults instantiates a new ApplianceNetworkLinkStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNetworkLinkStatusWithDefaults() *ApplianceNetworkLinkStatus {
	this := ApplianceNetworkLinkStatus{}
	var classId string = "appliance.NetworkLinkStatus"
	this.ClassId = classId
	var objectType string = "appliance.NetworkLinkStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNetworkLinkStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNetworkLinkStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNetworkLinkStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNetworkLinkStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDestinationHostname returns the DestinationHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetDestinationHostname() string {
	if o == nil || o.DestinationHostname == nil {
		var ret string
		return ret
	}
	return *o.DestinationHostname
}

// GetDestinationHostnameOk returns a tuple with the DestinationHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetDestinationHostnameOk() (*string, bool) {
	if o == nil || o.DestinationHostname == nil {
		return nil, false
	}
	return o.DestinationHostname, true
}

// HasDestinationHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasDestinationHostname() bool {
	if o != nil && o.DestinationHostname != nil {
		return true
	}

	return false
}

// SetDestinationHostname gets a reference to the given string and assigns it to the DestinationHostname field.
func (o *ApplianceNetworkLinkStatus) SetDestinationHostname(v string) {
	o.DestinationHostname = &v
}

// GetPingTime returns the PingTime field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetPingTime() float32 {
	if o == nil || o.PingTime == nil {
		var ret float32
		return ret
	}
	return *o.PingTime
}

// GetPingTimeOk returns a tuple with the PingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetPingTimeOk() (*float32, bool) {
	if o == nil || o.PingTime == nil {
		return nil, false
	}
	return o.PingTime, true
}

// HasPingTime returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasPingTime() bool {
	if o != nil && o.PingTime != nil {
		return true
	}

	return false
}

// SetPingTime gets a reference to the given float32 and assigns it to the PingTime field.
func (o *ApplianceNetworkLinkStatus) SetPingTime(v float32) {
	o.PingTime = &v
}

// GetSourceHostname returns the SourceHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetSourceHostname() string {
	if o == nil || o.SourceHostname == nil {
		var ret string
		return ret
	}
	return *o.SourceHostname
}

// GetSourceHostnameOk returns a tuple with the SourceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetSourceHostnameOk() (*string, bool) {
	if o == nil || o.SourceHostname == nil {
		return nil, false
	}
	return o.SourceHostname, true
}

// HasSourceHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasSourceHostname() bool {
	if o != nil && o.SourceHostname != nil {
		return true
	}

	return false
}

// SetSourceHostname gets a reference to the given string and assigns it to the SourceHostname field.
func (o *ApplianceNetworkLinkStatus) SetSourceHostname(v string) {
	o.SourceHostname = &v
}

// GetNodeOpStatus returns the NodeOpStatus field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetNodeOpStatus() ApplianceNodeOpStatusRelationship {
	if o == nil || o.NodeOpStatus == nil {
		var ret ApplianceNodeOpStatusRelationship
		return ret
	}
	return *o.NodeOpStatus
}

// GetNodeOpStatusOk returns a tuple with the NodeOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool) {
	if o == nil || o.NodeOpStatus == nil {
		return nil, false
	}
	return o.NodeOpStatus, true
}

// HasNodeOpStatus returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasNodeOpStatus() bool {
	if o != nil && o.NodeOpStatus != nil {
		return true
	}

	return false
}

// SetNodeOpStatus gets a reference to the given ApplianceNodeOpStatusRelationship and assigns it to the NodeOpStatus field.
func (o *ApplianceNetworkLinkStatus) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship) {
	o.NodeOpStatus = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceNetworkLinkStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ApplianceNetworkLinkStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DestinationHostname != nil {
		toSerialize["DestinationHostname"] = o.DestinationHostname
	}
	if o.PingTime != nil {
		toSerialize["PingTime"] = o.PingTime
	}
	if o.SourceHostname != nil {
		toSerialize["SourceHostname"] = o.SourceHostname
	}
	if o.NodeOpStatus != nil {
		toSerialize["NodeOpStatus"] = o.NodeOpStatus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNetworkLinkStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceNetworkLinkStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Hostname of the destination endpoint.
		DestinationHostname *string `json:"DestinationHostname,omitempty"`
		// Time to reach the destination endpoint in milliseconds from the source endpoint.
		PingTime *float32 `json:"PingTime,omitempty"`
		// Hostname of the source endpoint.
		SourceHostname   *string                              `json:"SourceHostname,omitempty"`
		NodeOpStatus     *ApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varApplianceNetworkLinkStatusWithoutEmbeddedStruct := ApplianceNetworkLinkStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceNetworkLinkStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceNetworkLinkStatus := _ApplianceNetworkLinkStatus{}
		varApplianceNetworkLinkStatus.ClassId = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.ClassId
		varApplianceNetworkLinkStatus.ObjectType = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.ObjectType
		varApplianceNetworkLinkStatus.DestinationHostname = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.DestinationHostname
		varApplianceNetworkLinkStatus.PingTime = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.PingTime
		varApplianceNetworkLinkStatus.SourceHostname = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.SourceHostname
		varApplianceNetworkLinkStatus.NodeOpStatus = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.NodeOpStatus
		varApplianceNetworkLinkStatus.RegisteredDevice = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.RegisteredDevice
		*o = ApplianceNetworkLinkStatus(varApplianceNetworkLinkStatus)
	} else {
		return err
	}

	varApplianceNetworkLinkStatus := _ApplianceNetworkLinkStatus{}

	err = json.Unmarshal(bytes, &varApplianceNetworkLinkStatus)
	if err == nil {
		o.MoBaseMo = varApplianceNetworkLinkStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestinationHostname")
		delete(additionalProperties, "PingTime")
		delete(additionalProperties, "SourceHostname")
		delete(additionalProperties, "NodeOpStatus")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceNetworkLinkStatus struct {
	value *ApplianceNetworkLinkStatus
	isSet bool
}

func (v NullableApplianceNetworkLinkStatus) Get() *ApplianceNetworkLinkStatus {
	return v.value
}

func (v *NullableApplianceNetworkLinkStatus) Set(val *ApplianceNetworkLinkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNetworkLinkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNetworkLinkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNetworkLinkStatus(val *ApplianceNetworkLinkStatus) *NullableApplianceNetworkLinkStatus {
	return &NullableApplianceNetworkLinkStatus{value: val, isSet: true}
}

func (v NullableApplianceNetworkLinkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNetworkLinkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

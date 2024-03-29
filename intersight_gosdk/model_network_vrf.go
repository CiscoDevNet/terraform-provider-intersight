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

// NetworkVrf Concrete class for virtual routing and forwarding.
type NetworkVrf struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the adminState of the given component.
	AdminState *string `json:"AdminState,omitempty"`
	// This field identifies the name of the given component.
	Name *string `json:"Name,omitempty"`
	// This field identifies the id of the given component.
	VrfId                *string                              `json:"VrfId,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkVrf NetworkVrf

// NewNetworkVrf instantiates a new NetworkVrf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVrf(classId string, objectType string) *NetworkVrf {
	this := NetworkVrf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVrfWithDefaults instantiates a new NetworkVrf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVrfWithDefaults() *NetworkVrf {
	this := NetworkVrf{}
	var classId string = "network.Vrf"
	this.ClassId = classId
	var objectType string = "network.Vrf"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVrf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVrf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVrf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVrf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NetworkVrf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NetworkVrf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NetworkVrf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkVrf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkVrf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkVrf) SetName(v string) {
	o.Name = &v
}

// GetVrfId returns the VrfId field value if set, zero value otherwise.
func (o *NetworkVrf) GetVrfId() string {
	if o == nil || o.VrfId == nil {
		var ret string
		return ret
	}
	return *o.VrfId
}

// GetVrfIdOk returns a tuple with the VrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetVrfIdOk() (*string, bool) {
	if o == nil || o.VrfId == nil {
		return nil, false
	}
	return o.VrfId, true
}

// HasVrfId returns a boolean if a field has been set.
func (o *NetworkVrf) HasVrfId() bool {
	if o != nil && o.VrfId != nil {
		return true
	}

	return false
}

// SetVrfId gets a reference to the given string and assigns it to the VrfId field.
func (o *NetworkVrf) SetVrfId(v string) {
	o.VrfId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkVrf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVrf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVrf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkVrf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVrf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVrf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVrf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkVrf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VrfId != nil {
		toSerialize["VrfId"] = o.VrfId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkVrf) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkVrfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field identifies the adminState of the given component.
		AdminState *string `json:"AdminState,omitempty"`
		// This field identifies the name of the given component.
		Name *string `json:"Name,omitempty"`
		// This field identifies the id of the given component.
		VrfId            *string                              `json:"VrfId,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkVrfWithoutEmbeddedStruct := NetworkVrfWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkVrfWithoutEmbeddedStruct)
	if err == nil {
		varNetworkVrf := _NetworkVrf{}
		varNetworkVrf.ClassId = varNetworkVrfWithoutEmbeddedStruct.ClassId
		varNetworkVrf.ObjectType = varNetworkVrfWithoutEmbeddedStruct.ObjectType
		varNetworkVrf.AdminState = varNetworkVrfWithoutEmbeddedStruct.AdminState
		varNetworkVrf.Name = varNetworkVrfWithoutEmbeddedStruct.Name
		varNetworkVrf.VrfId = varNetworkVrfWithoutEmbeddedStruct.VrfId
		varNetworkVrf.NetworkElement = varNetworkVrfWithoutEmbeddedStruct.NetworkElement
		varNetworkVrf.RegisteredDevice = varNetworkVrfWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkVrf(varNetworkVrf)
	} else {
		return err
	}

	varNetworkVrf := _NetworkVrf{}

	err = json.Unmarshal(bytes, &varNetworkVrf)
	if err == nil {
		o.InventoryBase = varNetworkVrf.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VrfId")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkVrf struct {
	value *NetworkVrf
	isSet bool
}

func (v NullableNetworkVrf) Get() *NetworkVrf {
	return v.value
}

func (v *NullableNetworkVrf) Set(val *NetworkVrf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVrf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVrf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVrf(val *NetworkVrf) *NullableNetworkVrf {
	return &NullableNetworkVrf{value: val, isSet: true}
}

func (v NullableNetworkVrf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVrf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
)

// NiatelemetryLeafPolGrpDetailsAllOf Definition of the list of properties defined in 'niatelemetry.LeafPolGrpDetails', excluding properties defined in parent classes.
type NiatelemetryLeafPolGrpDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Pol group and leaf profile relational object.
	Dn *string `json:"Dn,omitempty"`
	// Fabric node control dn associated with the pol group.
	FabricNodeControlDn *string `json:"FabricNodeControlDn,omitempty"`
	// Fabric node control policy name associated with the pol group.
	FabricNodeControlPolName *string `json:"FabricNodeControlPolName,omitempty"`
	// Leaf policy group name in APIC.
	LeafPolGroupName *string `json:"LeafPolGroupName,omitempty"`
	// Leaf profile group name in APIC.
	LeafProfileName *string `json:"LeafProfileName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// State of fabric node control pol.
	State                *string                              `json:"State,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryLeafPolGrpDetailsAllOf NiatelemetryLeafPolGrpDetailsAllOf

// NewNiatelemetryLeafPolGrpDetailsAllOf instantiates a new NiatelemetryLeafPolGrpDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryLeafPolGrpDetailsAllOf(classId string, objectType string) *NiatelemetryLeafPolGrpDetailsAllOf {
	this := NiatelemetryLeafPolGrpDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryLeafPolGrpDetailsAllOfWithDefaults instantiates a new NiatelemetryLeafPolGrpDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryLeafPolGrpDetailsAllOfWithDefaults() *NiatelemetryLeafPolGrpDetailsAllOf {
	this := NiatelemetryLeafPolGrpDetailsAllOf{}
	var classId string = "niatelemetry.LeafPolGrpDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.LeafPolGrpDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetFabricNodeControlDn returns the FabricNodeControlDn field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlDn() string {
	if o == nil || o.FabricNodeControlDn == nil {
		var ret string
		return ret
	}
	return *o.FabricNodeControlDn
}

// GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlDnOk() (*string, bool) {
	if o == nil || o.FabricNodeControlDn == nil {
		return nil, false
	}
	return o.FabricNodeControlDn, true
}

// HasFabricNodeControlDn returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasFabricNodeControlDn() bool {
	if o != nil && o.FabricNodeControlDn != nil {
		return true
	}

	return false
}

// SetFabricNodeControlDn gets a reference to the given string and assigns it to the FabricNodeControlDn field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetFabricNodeControlDn(v string) {
	o.FabricNodeControlDn = &v
}

// GetFabricNodeControlPolName returns the FabricNodeControlPolName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlPolName() string {
	if o == nil || o.FabricNodeControlPolName == nil {
		var ret string
		return ret
	}
	return *o.FabricNodeControlPolName
}

// GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlPolNameOk() (*string, bool) {
	if o == nil || o.FabricNodeControlPolName == nil {
		return nil, false
	}
	return o.FabricNodeControlPolName, true
}

// HasFabricNodeControlPolName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasFabricNodeControlPolName() bool {
	if o != nil && o.FabricNodeControlPolName != nil {
		return true
	}

	return false
}

// SetFabricNodeControlPolName gets a reference to the given string and assigns it to the FabricNodeControlPolName field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetFabricNodeControlPolName(v string) {
	o.FabricNodeControlPolName = &v
}

// GetLeafPolGroupName returns the LeafPolGroupName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafPolGroupName() string {
	if o == nil || o.LeafPolGroupName == nil {
		var ret string
		return ret
	}
	return *o.LeafPolGroupName
}

// GetLeafPolGroupNameOk returns a tuple with the LeafPolGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafPolGroupNameOk() (*string, bool) {
	if o == nil || o.LeafPolGroupName == nil {
		return nil, false
	}
	return o.LeafPolGroupName, true
}

// HasLeafPolGroupName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasLeafPolGroupName() bool {
	if o != nil && o.LeafPolGroupName != nil {
		return true
	}

	return false
}

// SetLeafPolGroupName gets a reference to the given string and assigns it to the LeafPolGroupName field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetLeafPolGroupName(v string) {
	o.LeafPolGroupName = &v
}

// GetLeafProfileName returns the LeafProfileName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafProfileName() string {
	if o == nil || o.LeafProfileName == nil {
		var ret string
		return ret
	}
	return *o.LeafProfileName
}

// GetLeafProfileNameOk returns a tuple with the LeafProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafProfileNameOk() (*string, bool) {
	if o == nil || o.LeafProfileName == nil {
		return nil, false
	}
	return o.LeafProfileName, true
}

// HasLeafProfileName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasLeafProfileName() bool {
	if o != nil && o.LeafProfileName != nil {
		return true
	}

	return false
}

// SetLeafProfileName gets a reference to the given string and assigns it to the LeafProfileName field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetLeafProfileName(v string) {
	o.LeafProfileName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetState(v string) {
	o.State = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryLeafPolGrpDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.FabricNodeControlDn != nil {
		toSerialize["FabricNodeControlDn"] = o.FabricNodeControlDn
	}
	if o.FabricNodeControlPolName != nil {
		toSerialize["FabricNodeControlPolName"] = o.FabricNodeControlPolName
	}
	if o.LeafPolGroupName != nil {
		toSerialize["LeafPolGroupName"] = o.LeafPolGroupName
	}
	if o.LeafProfileName != nil {
		toSerialize["LeafProfileName"] = o.LeafProfileName
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryLeafPolGrpDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryLeafPolGrpDetailsAllOf := _NiatelemetryLeafPolGrpDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryLeafPolGrpDetailsAllOf); err == nil {
		*o = NiatelemetryLeafPolGrpDetailsAllOf(varNiatelemetryLeafPolGrpDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricNodeControlDn")
		delete(additionalProperties, "FabricNodeControlPolName")
		delete(additionalProperties, "LeafPolGroupName")
		delete(additionalProperties, "LeafProfileName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "State")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryLeafPolGrpDetailsAllOf struct {
	value *NiatelemetryLeafPolGrpDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryLeafPolGrpDetailsAllOf) Get() *NiatelemetryLeafPolGrpDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryLeafPolGrpDetailsAllOf) Set(val *NiatelemetryLeafPolGrpDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryLeafPolGrpDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryLeafPolGrpDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryLeafPolGrpDetailsAllOf(val *NiatelemetryLeafPolGrpDetailsAllOf) *NullableNiatelemetryLeafPolGrpDetailsAllOf {
	return &NullableNiatelemetryLeafPolGrpDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryLeafPolGrpDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryLeafPolGrpDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

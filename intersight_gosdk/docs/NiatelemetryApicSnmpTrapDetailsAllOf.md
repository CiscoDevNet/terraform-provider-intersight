# NiatelemetryApicSnmpTrapDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicSnmpTrapDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicSnmpTrapDetails"]
**Dn** | Pointer to **string** | Dn of SNMP Trap attribute in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**V3SecLevel** | Pointer to **string** | V3SecLevel of SNMP Trap in APIC. | [optional] 
**Ver** | Pointer to **string** | Version of SNMP Trap in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicSnmpTrapDetailsAllOf

`func NewNiatelemetryApicSnmpTrapDetailsAllOf(classId string, objectType string, ) *NiatelemetryApicSnmpTrapDetailsAllOf`

NewNiatelemetryApicSnmpTrapDetailsAllOf instantiates a new NiatelemetryApicSnmpTrapDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicSnmpTrapDetailsAllOfWithDefaults

`func NewNiatelemetryApicSnmpTrapDetailsAllOfWithDefaults() *NiatelemetryApicSnmpTrapDetailsAllOf`

NewNiatelemetryApicSnmpTrapDetailsAllOfWithDefaults instantiates a new NiatelemetryApicSnmpTrapDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetV3SecLevel() string`

GetV3SecLevel returns the V3SecLevel field if non-nil, zero value otherwise.

### GetV3SecLevelOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetV3SecLevelOk() (*string, bool)`

GetV3SecLevelOk returns a tuple with the V3SecLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetV3SecLevel(v string)`

SetV3SecLevel sets V3SecLevel field to given value.

### HasV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasV3SecLevel() bool`

HasV3SecLevel returns a boolean if a field has been set.

### GetVer

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetVer() string`

GetVer returns the Ver field if non-nil, zero value otherwise.

### GetVerOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetVerOk() (*string, bool)`

GetVerOk returns a tuple with the Ver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVer

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetVer(v string)`

SetVer sets Ver field to given value.

### HasVer

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasVer() bool`

HasVer returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



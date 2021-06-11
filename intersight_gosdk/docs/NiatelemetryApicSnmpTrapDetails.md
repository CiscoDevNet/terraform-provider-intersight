# NiatelemetryApicSnmpTrapDetails

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

### NewNiatelemetryApicSnmpTrapDetails

`func NewNiatelemetryApicSnmpTrapDetails(classId string, objectType string, ) *NiatelemetryApicSnmpTrapDetails`

NewNiatelemetryApicSnmpTrapDetails instantiates a new NiatelemetryApicSnmpTrapDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicSnmpTrapDetailsWithDefaults

`func NewNiatelemetryApicSnmpTrapDetailsWithDefaults() *NiatelemetryApicSnmpTrapDetails`

NewNiatelemetryApicSnmpTrapDetailsWithDefaults instantiates a new NiatelemetryApicSnmpTrapDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicSnmpTrapDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicSnmpTrapDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicSnmpTrapDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicSnmpTrapDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryApicSnmpTrapDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicSnmpTrapDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicSnmpTrapDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicSnmpTrapDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicSnmpTrapDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicSnmpTrapDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicSnmpTrapDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicSnmpTrapDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicSnmpTrapDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicSnmpTrapDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetails) GetV3SecLevel() string`

GetV3SecLevel returns the V3SecLevel field if non-nil, zero value otherwise.

### GetV3SecLevelOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetV3SecLevelOk() (*string, bool)`

GetV3SecLevelOk returns a tuple with the V3SecLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetails) SetV3SecLevel(v string)`

SetV3SecLevel sets V3SecLevel field to given value.

### HasV3SecLevel

`func (o *NiatelemetryApicSnmpTrapDetails) HasV3SecLevel() bool`

HasV3SecLevel returns a boolean if a field has been set.

### GetVer

`func (o *NiatelemetryApicSnmpTrapDetails) GetVer() string`

GetVer returns the Ver field if non-nil, zero value otherwise.

### GetVerOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetVerOk() (*string, bool)`

GetVerOk returns a tuple with the Ver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVer

`func (o *NiatelemetryApicSnmpTrapDetails) SetVer(v string)`

SetVer sets Ver field to given value.

### HasVer

`func (o *NiatelemetryApicSnmpTrapDetails) HasVer() bool`

HasVer returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicSnmpTrapDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicSnmpTrapDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NiatelemetryApicFlashDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicFlashDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicFlashDetails"]
**Dn** | Pointer to **string** | Dn for the flash in the inventory. | [optional] 
**ModelNumber** | Pointer to **string** | Model number of the flash in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Revision** | Pointer to **string** | Revision of the flash Mo in APIC. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the flash in APIC. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicFlashDetails

`func NewNiatelemetryApicFlashDetails(classId string, objectType string, ) *NiatelemetryApicFlashDetails`

NewNiatelemetryApicFlashDetails instantiates a new NiatelemetryApicFlashDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicFlashDetailsWithDefaults

`func NewNiatelemetryApicFlashDetailsWithDefaults() *NiatelemetryApicFlashDetails`

NewNiatelemetryApicFlashDetailsWithDefaults instantiates a new NiatelemetryApicFlashDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicFlashDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicFlashDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicFlashDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicFlashDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicFlashDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicFlashDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryApicFlashDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicFlashDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicFlashDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicFlashDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetModelNumber

`func (o *NiatelemetryApicFlashDetails) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *NiatelemetryApicFlashDetails) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *NiatelemetryApicFlashDetails) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *NiatelemetryApicFlashDetails) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicFlashDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicFlashDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicFlashDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicFlashDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicFlashDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicFlashDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicFlashDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicFlashDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRevision

`func (o *NiatelemetryApicFlashDetails) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NiatelemetryApicFlashDetails) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NiatelemetryApicFlashDetails) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NiatelemetryApicFlashDetails) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryApicFlashDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryApicFlashDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryApicFlashDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryApicFlashDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicFlashDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicFlashDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicFlashDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicFlashDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicFlashDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicFlashDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicFlashDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicFlashDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NiatelemetryFabricModuleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.FabricModuleDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.FabricModuleDetails"]
**Dn** | Pointer to **string** | Dn of the fabric module in APIC. | [optional] 
**HwVer** | Pointer to **string** | Hardware version of fabric module. | [optional] 
**Model** | Pointer to **string** | Model of the fabric module. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Revision** | Pointer to **string** | Revision of the fabric module. | [optional] 
**Serial** | Pointer to **string** | Serial number of the fabric module. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**Type** | Pointer to **string** | Type of the fabric module. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryFabricModuleDetails

`func NewNiatelemetryFabricModuleDetails(classId string, objectType string, ) *NiatelemetryFabricModuleDetails`

NewNiatelemetryFabricModuleDetails instantiates a new NiatelemetryFabricModuleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFabricModuleDetailsWithDefaults

`func NewNiatelemetryFabricModuleDetailsWithDefaults() *NiatelemetryFabricModuleDetails`

NewNiatelemetryFabricModuleDetailsWithDefaults instantiates a new NiatelemetryFabricModuleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFabricModuleDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFabricModuleDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFabricModuleDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFabricModuleDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFabricModuleDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFabricModuleDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryFabricModuleDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryFabricModuleDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryFabricModuleDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryFabricModuleDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHwVer

`func (o *NiatelemetryFabricModuleDetails) GetHwVer() string`

GetHwVer returns the HwVer field if non-nil, zero value otherwise.

### GetHwVerOk

`func (o *NiatelemetryFabricModuleDetails) GetHwVerOk() (*string, bool)`

GetHwVerOk returns a tuple with the HwVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVer

`func (o *NiatelemetryFabricModuleDetails) SetHwVer(v string)`

SetHwVer sets HwVer field to given value.

### HasHwVer

`func (o *NiatelemetryFabricModuleDetails) HasHwVer() bool`

HasHwVer returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryFabricModuleDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryFabricModuleDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryFabricModuleDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryFabricModuleDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryFabricModuleDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryFabricModuleDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryFabricModuleDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryFabricModuleDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryFabricModuleDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryFabricModuleDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryFabricModuleDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryFabricModuleDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRevision

`func (o *NiatelemetryFabricModuleDetails) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NiatelemetryFabricModuleDetails) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NiatelemetryFabricModuleDetails) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NiatelemetryFabricModuleDetails) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryFabricModuleDetails) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryFabricModuleDetails) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryFabricModuleDetails) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryFabricModuleDetails) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryFabricModuleDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryFabricModuleDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryFabricModuleDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryFabricModuleDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryFabricModuleDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryFabricModuleDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryFabricModuleDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryFabricModuleDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryFabricModuleDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryFabricModuleDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryFabricModuleDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryFabricModuleDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



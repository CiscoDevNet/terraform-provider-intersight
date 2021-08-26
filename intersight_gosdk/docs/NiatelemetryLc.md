# NiatelemetryLc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Lc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Lc"]
**Description** | Pointer to **string** | Description of the line cards present. | [optional] 
**Dn** | Pointer to **string** | Dn value for the line cards present. | [optional] 
**HardwareVersion** | Pointer to **string** | Hardware version of the line cards present. | [optional] 
**Model** | Pointer to **string** | Model of the line cards present. | [optional] 
**NodeId** | Pointer to **int64** | Node Id of the line card present. | [optional] 
**OperationalState** | Pointer to **string** | Opretaional state of the line cards present. | [optional] 
**PowerState** | Pointer to **string** | Power state of the line cards present. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RedundancyState** | Pointer to **string** | Redundancy state of the line cards present. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the line card present. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. | [optional] 
**Vid** | Pointer to **string** | VID for the line card in the inventory. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryLc

`func NewNiatelemetryLc(classId string, objectType string, ) *NiatelemetryLc`

NewNiatelemetryLc instantiates a new NiatelemetryLc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLcWithDefaults

`func NewNiatelemetryLcWithDefaults() *NiatelemetryLc`

NewNiatelemetryLcWithDefaults instantiates a new NiatelemetryLc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NiatelemetryLc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiatelemetryLc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiatelemetryLc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiatelemetryLc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryLc) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryLc) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryLc) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryLc) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *NiatelemetryLc) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *NiatelemetryLc) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *NiatelemetryLc) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *NiatelemetryLc) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryLc) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryLc) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryLc) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryLc) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNodeId

`func (o *NiatelemetryLc) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryLc) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryLc) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryLc) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOperationalState

`func (o *NiatelemetryLc) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NiatelemetryLc) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NiatelemetryLc) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NiatelemetryLc) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPowerState

`func (o *NiatelemetryLc) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *NiatelemetryLc) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *NiatelemetryLc) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *NiatelemetryLc) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryLc) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryLc) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryLc) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryLc) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryLc) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryLc) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryLc) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryLc) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRedundancyState

`func (o *NiatelemetryLc) GetRedundancyState() string`

GetRedundancyState returns the RedundancyState field if non-nil, zero value otherwise.

### GetRedundancyStateOk

`func (o *NiatelemetryLc) GetRedundancyStateOk() (*string, bool)`

GetRedundancyStateOk returns a tuple with the RedundancyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyState

`func (o *NiatelemetryLc) SetRedundancyState(v string)`

SetRedundancyState sets RedundancyState field to given value.

### HasRedundancyState

`func (o *NiatelemetryLc) HasRedundancyState() bool`

HasRedundancyState returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryLc) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryLc) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryLc) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryLc) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryLc) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryLc) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryLc) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryLc) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetVid

`func (o *NiatelemetryLc) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *NiatelemetryLc) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *NiatelemetryLc) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *NiatelemetryLc) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryLc) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryLc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryLc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryLc) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



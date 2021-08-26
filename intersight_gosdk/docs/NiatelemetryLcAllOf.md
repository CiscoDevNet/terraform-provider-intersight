# NiatelemetryLcAllOf

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

### NewNiatelemetryLcAllOf

`func NewNiatelemetryLcAllOf(classId string, objectType string, ) *NiatelemetryLcAllOf`

NewNiatelemetryLcAllOf instantiates a new NiatelemetryLcAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLcAllOfWithDefaults

`func NewNiatelemetryLcAllOfWithDefaults() *NiatelemetryLcAllOf`

NewNiatelemetryLcAllOfWithDefaults instantiates a new NiatelemetryLcAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLcAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLcAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLcAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLcAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLcAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLcAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NiatelemetryLcAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiatelemetryLcAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiatelemetryLcAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiatelemetryLcAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryLcAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryLcAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryLcAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryLcAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *NiatelemetryLcAllOf) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *NiatelemetryLcAllOf) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *NiatelemetryLcAllOf) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *NiatelemetryLcAllOf) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryLcAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryLcAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryLcAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryLcAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNodeId

`func (o *NiatelemetryLcAllOf) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryLcAllOf) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryLcAllOf) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryLcAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOperationalState

`func (o *NiatelemetryLcAllOf) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *NiatelemetryLcAllOf) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *NiatelemetryLcAllOf) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *NiatelemetryLcAllOf) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetPowerState

`func (o *NiatelemetryLcAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *NiatelemetryLcAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *NiatelemetryLcAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *NiatelemetryLcAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryLcAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryLcAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryLcAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryLcAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryLcAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryLcAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryLcAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryLcAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRedundancyState

`func (o *NiatelemetryLcAllOf) GetRedundancyState() string`

GetRedundancyState returns the RedundancyState field if non-nil, zero value otherwise.

### GetRedundancyStateOk

`func (o *NiatelemetryLcAllOf) GetRedundancyStateOk() (*string, bool)`

GetRedundancyStateOk returns a tuple with the RedundancyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyState

`func (o *NiatelemetryLcAllOf) SetRedundancyState(v string)`

SetRedundancyState sets RedundancyState field to given value.

### HasRedundancyState

`func (o *NiatelemetryLcAllOf) HasRedundancyState() bool`

HasRedundancyState returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryLcAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryLcAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryLcAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryLcAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryLcAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryLcAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryLcAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryLcAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetVid

`func (o *NiatelemetryLcAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *NiatelemetryLcAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *NiatelemetryLcAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *NiatelemetryLcAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryLcAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryLcAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryLcAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryLcAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



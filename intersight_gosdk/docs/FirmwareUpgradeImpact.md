# FirmwareUpgradeImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.UpgradeImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.UpgradeImpact"]
**Chassis** | Pointer to [**[]EquipmentChassisRelationship**](EquipmentChassisRelationship.md) | An array of relationships to equipmentChassis resources. | [optional] 
**Device** | Pointer to [**[]AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) | An array of relationships to assetDeviceRegistration resources. | [optional] [readonly] 
**Distributable** | Pointer to [**FirmwareDistributableRelationship**](FirmwareDistributableRelationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](NetworkElementRelationship.md) | An array of relationships to networkElement resources. | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](SoftwarerepositoryReleaseRelationship.md) |  | [optional] 
**Server** | Pointer to [**[]ComputePhysicalRelationship**](ComputePhysicalRelationship.md) | An array of relationships to computePhysical resources. | [optional] 

## Methods

### NewFirmwareUpgradeImpact

`func NewFirmwareUpgradeImpact(classId string, objectType string, ) *FirmwareUpgradeImpact`

NewFirmwareUpgradeImpact instantiates a new FirmwareUpgradeImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactWithDefaults

`func NewFirmwareUpgradeImpactWithDefaults() *FirmwareUpgradeImpact`

NewFirmwareUpgradeImpactWithDefaults instantiates a new FirmwareUpgradeImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeImpact) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeImpact) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeImpact) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeImpact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeImpact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeImpact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassis

`func (o *FirmwareUpgradeImpact) GetChassis() []EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *FirmwareUpgradeImpact) GetChassisOk() (*[]EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *FirmwareUpgradeImpact) SetChassis(v []EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *FirmwareUpgradeImpact) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *FirmwareUpgradeImpact) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *FirmwareUpgradeImpact) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetDevice

`func (o *FirmwareUpgradeImpact) GetDevice() []AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgradeImpact) GetDeviceOk() (*[]AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgradeImpact) SetDevice(v []AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgradeImpact) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwareUpgradeImpact) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwareUpgradeImpact) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDistributable

`func (o *FirmwareUpgradeImpact) GetDistributable() FirmwareDistributableRelationship`

GetDistributable returns the Distributable field if non-nil, zero value otherwise.

### GetDistributableOk

`func (o *FirmwareUpgradeImpact) GetDistributableOk() (*FirmwareDistributableRelationship, bool)`

GetDistributableOk returns a tuple with the Distributable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributable

`func (o *FirmwareUpgradeImpact) SetDistributable(v FirmwareDistributableRelationship)`

SetDistributable sets Distributable field to given value.

### HasDistributable

`func (o *FirmwareUpgradeImpact) HasDistributable() bool`

HasDistributable returns a boolean if a field has been set.

### GetNetworkElements

`func (o *FirmwareUpgradeImpact) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareUpgradeImpact) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareUpgradeImpact) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareUpgradeImpact) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareUpgradeImpact) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareUpgradeImpact) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetRelease

`func (o *FirmwareUpgradeImpact) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareUpgradeImpact) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareUpgradeImpact) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareUpgradeImpact) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgradeImpact) GetServer() []ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgradeImpact) GetServerOk() (*[]ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgradeImpact) SetServer(v []ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgradeImpact) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *FirmwareUpgradeImpact) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *FirmwareUpgradeImpact) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



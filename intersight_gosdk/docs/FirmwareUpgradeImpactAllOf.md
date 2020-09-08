# FirmwareUpgradeImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**[]EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) | An array of relationships to equipmentChassis resources. | [optional] 
**Device** | Pointer to [**[]AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) | An array of relationships to assetDeviceRegistration resources. | [optional] [readonly] 
**Distributable** | Pointer to [**FirmwareDistributableRelationship**](firmware.Distributable.Relationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](network.Element.Relationship.md) | An array of relationships to networkElement resources. | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 
**Server** | Pointer to [**[]ComputePhysicalRelationship**](compute.Physical.Relationship.md) | An array of relationships to computePhysical resources. | [optional] 

## Methods

### NewFirmwareUpgradeImpactAllOf

`func NewFirmwareUpgradeImpactAllOf() *FirmwareUpgradeImpactAllOf`

NewFirmwareUpgradeImpactAllOf instantiates a new FirmwareUpgradeImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactAllOfWithDefaults

`func NewFirmwareUpgradeImpactAllOfWithDefaults() *FirmwareUpgradeImpactAllOf`

NewFirmwareUpgradeImpactAllOfWithDefaults instantiates a new FirmwareUpgradeImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *FirmwareUpgradeImpactAllOf) GetChassis() []EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *FirmwareUpgradeImpactAllOf) GetChassisOk() (*[]EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *FirmwareUpgradeImpactAllOf) SetChassis(v []EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *FirmwareUpgradeImpactAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *FirmwareUpgradeImpactAllOf) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *FirmwareUpgradeImpactAllOf) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetDevice

`func (o *FirmwareUpgradeImpactAllOf) GetDevice() []AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgradeImpactAllOf) GetDeviceOk() (*[]AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgradeImpactAllOf) SetDevice(v []AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgradeImpactAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwareUpgradeImpactAllOf) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwareUpgradeImpactAllOf) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDistributable

`func (o *FirmwareUpgradeImpactAllOf) GetDistributable() FirmwareDistributableRelationship`

GetDistributable returns the Distributable field if non-nil, zero value otherwise.

### GetDistributableOk

`func (o *FirmwareUpgradeImpactAllOf) GetDistributableOk() (*FirmwareDistributableRelationship, bool)`

GetDistributableOk returns a tuple with the Distributable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributable

`func (o *FirmwareUpgradeImpactAllOf) SetDistributable(v FirmwareDistributableRelationship)`

SetDistributable sets Distributable field to given value.

### HasDistributable

`func (o *FirmwareUpgradeImpactAllOf) HasDistributable() bool`

HasDistributable returns a boolean if a field has been set.

### GetNetworkElements

`func (o *FirmwareUpgradeImpactAllOf) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareUpgradeImpactAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareUpgradeImpactAllOf) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareUpgradeImpactAllOf) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareUpgradeImpactAllOf) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareUpgradeImpactAllOf) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetRelease

`func (o *FirmwareUpgradeImpactAllOf) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareUpgradeImpactAllOf) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareUpgradeImpactAllOf) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareUpgradeImpactAllOf) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgradeImpactAllOf) GetServer() []ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgradeImpactAllOf) GetServerOk() (*[]ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgradeImpactAllOf) SetServer(v []ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgradeImpactAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *FirmwareUpgradeImpactAllOf) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *FirmwareUpgradeImpactAllOf) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



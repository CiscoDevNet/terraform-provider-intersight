# EquipmentFex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryState** | Pointer to **string** | Discovery state of IO card or fabric extender. | [optional] 
**Fans** | Pointer to [**[]EquipmentFanRelationship**](equipment.Fan.Relationship.md) | An array of relationships to equipmentFan resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentFex

`func NewEquipmentFex() *EquipmentFex`

NewEquipmentFex instantiates a new EquipmentFex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFexWithDefaults

`func NewEquipmentFexWithDefaults() *EquipmentFex`

NewEquipmentFexWithDefaults instantiates a new EquipmentFex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryState

`func (o *EquipmentFex) GetDiscoveryState() string`

GetDiscoveryState returns the DiscoveryState field if non-nil, zero value otherwise.

### GetDiscoveryStateOk

`func (o *EquipmentFex) GetDiscoveryStateOk() (*string, bool)`

GetDiscoveryStateOk returns a tuple with the DiscoveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryState

`func (o *EquipmentFex) SetDiscoveryState(v string)`

SetDiscoveryState sets DiscoveryState field to given value.

### HasDiscoveryState

`func (o *EquipmentFex) HasDiscoveryState() bool`

HasDiscoveryState returns a boolean if a field has been set.

### GetFans

`func (o *EquipmentFex) GetFans() []EquipmentFanRelationship`

GetFans returns the Fans field if non-nil, zero value otherwise.

### GetFansOk

`func (o *EquipmentFex) GetFansOk() (*[]EquipmentFanRelationship, bool)`

GetFansOk returns a tuple with the Fans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFans

`func (o *EquipmentFex) SetFans(v []EquipmentFanRelationship)`

SetFans sets Fans field to given value.

### HasFans

`func (o *EquipmentFex) HasFans() bool`

HasFans returns a boolean if a field has been set.

### SetFansNil

`func (o *EquipmentFex) SetFansNil(b bool)`

 SetFansNil sets the value for Fans to be an explicit nil

### UnsetFans
`func (o *EquipmentFex) UnsetFans()`

UnsetFans ensures that no value is present for Fans, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentFex) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFex) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFex) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFex) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *EquipmentFex) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *EquipmentFex) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *EquipmentFex) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *EquipmentFex) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentFex) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentFex) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentFex) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentFex) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentFex) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentFex) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentFex) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentFex) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentFex) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentFex) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentFex) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFex) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFex) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFex) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



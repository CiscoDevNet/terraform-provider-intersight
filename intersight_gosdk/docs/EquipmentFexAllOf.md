# EquipmentFexAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Fex"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Fex"]
**ConnectionPath** | Pointer to **string** | Switch Id to which the FEX is connected to. The value can be A or B or AB in case of active-active topology. | [optional] [readonly] 
**DiscoveryState** | Pointer to **string** | Discovery state of IO card or fabric extender. | [optional] 
**Fans** | Pointer to [**[]EquipmentFanRelationship**](EquipmentFanRelationship.md) | An array of relationships to equipmentFan resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**Ioms** | Pointer to [**[]EquipmentIoCardRelationship**](EquipmentIoCardRelationship.md) | An array of relationships to equipmentIoCard resources. | [optional] [readonly] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](EquipmentPsuRelationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentFexAllOf

`func NewEquipmentFexAllOf(classId string, objectType string, ) *EquipmentFexAllOf`

NewEquipmentFexAllOf instantiates a new EquipmentFexAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFexAllOfWithDefaults

`func NewEquipmentFexAllOfWithDefaults() *EquipmentFexAllOf`

NewEquipmentFexAllOfWithDefaults instantiates a new EquipmentFexAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentFexAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentFexAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentFexAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentFexAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentFexAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentFexAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionPath

`func (o *EquipmentFexAllOf) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentFexAllOf) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentFexAllOf) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentFexAllOf) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetDiscoveryState

`func (o *EquipmentFexAllOf) GetDiscoveryState() string`

GetDiscoveryState returns the DiscoveryState field if non-nil, zero value otherwise.

### GetDiscoveryStateOk

`func (o *EquipmentFexAllOf) GetDiscoveryStateOk() (*string, bool)`

GetDiscoveryStateOk returns a tuple with the DiscoveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryState

`func (o *EquipmentFexAllOf) SetDiscoveryState(v string)`

SetDiscoveryState sets DiscoveryState field to given value.

### HasDiscoveryState

`func (o *EquipmentFexAllOf) HasDiscoveryState() bool`

HasDiscoveryState returns a boolean if a field has been set.

### GetFans

`func (o *EquipmentFexAllOf) GetFans() []EquipmentFanRelationship`

GetFans returns the Fans field if non-nil, zero value otherwise.

### GetFansOk

`func (o *EquipmentFexAllOf) GetFansOk() (*[]EquipmentFanRelationship, bool)`

GetFansOk returns a tuple with the Fans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFans

`func (o *EquipmentFexAllOf) SetFans(v []EquipmentFanRelationship)`

SetFans sets Fans field to given value.

### HasFans

`func (o *EquipmentFexAllOf) HasFans() bool`

HasFans returns a boolean if a field has been set.

### SetFansNil

`func (o *EquipmentFexAllOf) SetFansNil(b bool)`

 SetFansNil sets the value for Fans to be an explicit nil

### UnsetFans
`func (o *EquipmentFexAllOf) UnsetFans()`

UnsetFans ensures that no value is present for Fans, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentFexAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFexAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFexAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFexAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetIoms

`func (o *EquipmentFexAllOf) GetIoms() []EquipmentIoCardRelationship`

GetIoms returns the Ioms field if non-nil, zero value otherwise.

### GetIomsOk

`func (o *EquipmentFexAllOf) GetIomsOk() (*[]EquipmentIoCardRelationship, bool)`

GetIomsOk returns a tuple with the Ioms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoms

`func (o *EquipmentFexAllOf) SetIoms(v []EquipmentIoCardRelationship)`

SetIoms sets Ioms field to given value.

### HasIoms

`func (o *EquipmentFexAllOf) HasIoms() bool`

HasIoms returns a boolean if a field has been set.

### SetIomsNil

`func (o *EquipmentFexAllOf) SetIomsNil(b bool)`

 SetIomsNil sets the value for Ioms to be an explicit nil

### UnsetIoms
`func (o *EquipmentFexAllOf) UnsetIoms()`

UnsetIoms ensures that no value is present for Ioms, not even an explicit nil
### GetLocatorLed

`func (o *EquipmentFexAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *EquipmentFexAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *EquipmentFexAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *EquipmentFexAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentFexAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentFexAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentFexAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentFexAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentFexAllOf) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentFexAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentFexAllOf) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentFexAllOf) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentFexAllOf) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentFexAllOf) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentFexAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFexAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFexAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFexAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



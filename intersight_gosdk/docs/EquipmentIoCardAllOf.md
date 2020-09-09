# EquipmentIoCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionPath** | Pointer to **string** | Switch Id to which the IOM is connected to. The value can be A or B. | [optional] [readonly] 
**DcSupported** | Pointer to **bool** | IOM device connector support. | [optional] [readonly] 
**Side** | Pointer to **string** | Location of IOM within a chassis. The value can be left or right. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIoCardAllOf

`func NewEquipmentIoCardAllOf() *EquipmentIoCardAllOf`

NewEquipmentIoCardAllOf instantiates a new EquipmentIoCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardAllOfWithDefaults

`func NewEquipmentIoCardAllOfWithDefaults() *EquipmentIoCardAllOf`

NewEquipmentIoCardAllOfWithDefaults instantiates a new EquipmentIoCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionPath

`func (o *EquipmentIoCardAllOf) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentIoCardAllOf) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentIoCardAllOf) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentIoCardAllOf) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetDcSupported

`func (o *EquipmentIoCardAllOf) GetDcSupported() bool`

GetDcSupported returns the DcSupported field if non-nil, zero value otherwise.

### GetDcSupportedOk

`func (o *EquipmentIoCardAllOf) GetDcSupportedOk() (*bool, bool)`

GetDcSupportedOk returns a tuple with the DcSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSupported

`func (o *EquipmentIoCardAllOf) SetDcSupported(v bool)`

SetDcSupported sets DcSupported field to given value.

### HasDcSupported

`func (o *EquipmentIoCardAllOf) HasDcSupported() bool`

HasDcSupported returns a boolean if a field has been set.

### GetSide

`func (o *EquipmentIoCardAllOf) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquipmentIoCardAllOf) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquipmentIoCardAllOf) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquipmentIoCardAllOf) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentIoCardAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentIoCardAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentIoCardAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentIoCardAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentIoCardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentIoCardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentIoCardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentIoCardAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDeviceRegistration

`func (o *EquipmentIoCardAllOf) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship`

GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field if non-nil, zero value otherwise.

### GetPhysicalDeviceRegistrationOk

`func (o *EquipmentIoCardAllOf) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDeviceRegistration

`func (o *EquipmentIoCardAllOf) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetPhysicalDeviceRegistration sets PhysicalDeviceRegistration field to given value.

### HasPhysicalDeviceRegistration

`func (o *EquipmentIoCardAllOf) HasPhysicalDeviceRegistration() bool`

HasPhysicalDeviceRegistration returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentIoCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentIoCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentIoCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentIoCardAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



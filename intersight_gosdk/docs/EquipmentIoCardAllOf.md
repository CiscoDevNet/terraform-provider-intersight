# EquipmentIoCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.IoCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.IoCard"]
**ConnectionPath** | Pointer to **string** | Switch Id to which the IOM is connected to. The value can be A or B. | [optional] [readonly] 
**DcSupported** | Pointer to **bool** | IOM device connector support. | [optional] [readonly] 
**InbandIpAddresses** | Pointer to [**[]ComputeIpAddress**](ComputeIpAddress.md) |  | [optional] 
**Side** | Pointer to **string** | Location of IOM within a chassis. The value can be left or right. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIoCardAllOf

`func NewEquipmentIoCardAllOf(classId string, objectType string, ) *EquipmentIoCardAllOf`

NewEquipmentIoCardAllOf instantiates a new EquipmentIoCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardAllOfWithDefaults

`func NewEquipmentIoCardAllOfWithDefaults() *EquipmentIoCardAllOf`

NewEquipmentIoCardAllOfWithDefaults instantiates a new EquipmentIoCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIoCardAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIoCardAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIoCardAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIoCardAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIoCardAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIoCardAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetInbandIpAddresses

`func (o *EquipmentIoCardAllOf) GetInbandIpAddresses() []ComputeIpAddress`

GetInbandIpAddresses returns the InbandIpAddresses field if non-nil, zero value otherwise.

### GetInbandIpAddressesOk

`func (o *EquipmentIoCardAllOf) GetInbandIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetInbandIpAddressesOk returns a tuple with the InbandIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpAddresses

`func (o *EquipmentIoCardAllOf) SetInbandIpAddresses(v []ComputeIpAddress)`

SetInbandIpAddresses sets InbandIpAddresses field to given value.

### HasInbandIpAddresses

`func (o *EquipmentIoCardAllOf) HasInbandIpAddresses() bool`

HasInbandIpAddresses returns a boolean if a field has been set.

### SetInbandIpAddressesNil

`func (o *EquipmentIoCardAllOf) SetInbandIpAddressesNil(b bool)`

 SetInbandIpAddressesNil sets the value for InbandIpAddresses to be an explicit nil

### UnsetInbandIpAddresses
`func (o *EquipmentIoCardAllOf) UnsetInbandIpAddresses()`

UnsetInbandIpAddresses ensures that no value is present for InbandIpAddresses, not even an explicit nil
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

### GetEquipmentFex

`func (o *EquipmentIoCardAllOf) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentIoCardAllOf) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentIoCardAllOf) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentIoCardAllOf) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

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



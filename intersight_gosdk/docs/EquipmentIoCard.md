# EquipmentIoCard

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

### NewEquipmentIoCard

`func NewEquipmentIoCard(classId string, objectType string, ) *EquipmentIoCard`

NewEquipmentIoCard instantiates a new EquipmentIoCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardWithDefaults

`func NewEquipmentIoCardWithDefaults() *EquipmentIoCard`

NewEquipmentIoCardWithDefaults instantiates a new EquipmentIoCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIoCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIoCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIoCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIoCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIoCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIoCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionPath

`func (o *EquipmentIoCard) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentIoCard) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentIoCard) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentIoCard) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetDcSupported

`func (o *EquipmentIoCard) GetDcSupported() bool`

GetDcSupported returns the DcSupported field if non-nil, zero value otherwise.

### GetDcSupportedOk

`func (o *EquipmentIoCard) GetDcSupportedOk() (*bool, bool)`

GetDcSupportedOk returns a tuple with the DcSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSupported

`func (o *EquipmentIoCard) SetDcSupported(v bool)`

SetDcSupported sets DcSupported field to given value.

### HasDcSupported

`func (o *EquipmentIoCard) HasDcSupported() bool`

HasDcSupported returns a boolean if a field has been set.

### GetInbandIpAddresses

`func (o *EquipmentIoCard) GetInbandIpAddresses() []ComputeIpAddress`

GetInbandIpAddresses returns the InbandIpAddresses field if non-nil, zero value otherwise.

### GetInbandIpAddressesOk

`func (o *EquipmentIoCard) GetInbandIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetInbandIpAddressesOk returns a tuple with the InbandIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpAddresses

`func (o *EquipmentIoCard) SetInbandIpAddresses(v []ComputeIpAddress)`

SetInbandIpAddresses sets InbandIpAddresses field to given value.

### HasInbandIpAddresses

`func (o *EquipmentIoCard) HasInbandIpAddresses() bool`

HasInbandIpAddresses returns a boolean if a field has been set.

### SetInbandIpAddressesNil

`func (o *EquipmentIoCard) SetInbandIpAddressesNil(b bool)`

 SetInbandIpAddressesNil sets the value for InbandIpAddresses to be an explicit nil

### UnsetInbandIpAddresses
`func (o *EquipmentIoCard) UnsetInbandIpAddresses()`

UnsetInbandIpAddresses ensures that no value is present for InbandIpAddresses, not even an explicit nil
### GetSide

`func (o *EquipmentIoCard) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquipmentIoCard) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquipmentIoCard) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquipmentIoCard) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentIoCard) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentIoCard) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentIoCard) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentIoCard) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentIoCard) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentIoCard) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentIoCard) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentIoCard) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentIoCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentIoCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentIoCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentIoCard) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDeviceRegistration

`func (o *EquipmentIoCard) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship`

GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field if non-nil, zero value otherwise.

### GetPhysicalDeviceRegistrationOk

`func (o *EquipmentIoCard) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDeviceRegistration

`func (o *EquipmentIoCard) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetPhysicalDeviceRegistration sets PhysicalDeviceRegistration field to given value.

### HasPhysicalDeviceRegistration

`func (o *EquipmentIoCard) HasPhysicalDeviceRegistration() bool`

HasPhysicalDeviceRegistration returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentIoCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentIoCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentIoCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentIoCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



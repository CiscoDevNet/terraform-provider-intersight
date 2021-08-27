# ComputeVmediaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.Vmedia"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.Vmedia"]
**Enabled** | Pointer to **bool** | State of the Virtual Media service on the server. | [optional] [readonly] [default to true]
**Encryption** | Pointer to **bool** | If enabled, allows encryption of all Virtual Media communications. | [optional] [readonly] 
**LowPowerUsb** | Pointer to **bool** | If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. | [optional] [readonly] [default to true]
**ComputePhysicalUnit** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**Mappings** | Pointer to [**[]ComputeMappingRelationship**](ComputeMappingRelationship.md) | An array of relationships to computeMapping resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewComputeVmediaAllOf

`func NewComputeVmediaAllOf(classId string, objectType string, ) *ComputeVmediaAllOf`

NewComputeVmediaAllOf instantiates a new ComputeVmediaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeVmediaAllOfWithDefaults

`func NewComputeVmediaAllOfWithDefaults() *ComputeVmediaAllOf`

NewComputeVmediaAllOfWithDefaults instantiates a new ComputeVmediaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeVmediaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeVmediaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeVmediaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeVmediaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeVmediaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeVmediaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *ComputeVmediaAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComputeVmediaAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComputeVmediaAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ComputeVmediaAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryption

`func (o *ComputeVmediaAllOf) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ComputeVmediaAllOf) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ComputeVmediaAllOf) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ComputeVmediaAllOf) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLowPowerUsb

`func (o *ComputeVmediaAllOf) GetLowPowerUsb() bool`

GetLowPowerUsb returns the LowPowerUsb field if non-nil, zero value otherwise.

### GetLowPowerUsbOk

`func (o *ComputeVmediaAllOf) GetLowPowerUsbOk() (*bool, bool)`

GetLowPowerUsbOk returns a tuple with the LowPowerUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPowerUsb

`func (o *ComputeVmediaAllOf) SetLowPowerUsb(v bool)`

SetLowPowerUsb sets LowPowerUsb field to given value.

### HasLowPowerUsb

`func (o *ComputeVmediaAllOf) HasLowPowerUsb() bool`

HasLowPowerUsb returns a boolean if a field has been set.

### GetComputePhysicalUnit

`func (o *ComputeVmediaAllOf) GetComputePhysicalUnit() ComputePhysicalRelationship`

GetComputePhysicalUnit returns the ComputePhysicalUnit field if non-nil, zero value otherwise.

### GetComputePhysicalUnitOk

`func (o *ComputeVmediaAllOf) GetComputePhysicalUnitOk() (*ComputePhysicalRelationship, bool)`

GetComputePhysicalUnitOk returns a tuple with the ComputePhysicalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePhysicalUnit

`func (o *ComputeVmediaAllOf) SetComputePhysicalUnit(v ComputePhysicalRelationship)`

SetComputePhysicalUnit sets ComputePhysicalUnit field to given value.

### HasComputePhysicalUnit

`func (o *ComputeVmediaAllOf) HasComputePhysicalUnit() bool`

HasComputePhysicalUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ComputeVmediaAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeVmediaAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeVmediaAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeVmediaAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMappings

`func (o *ComputeVmediaAllOf) GetMappings() []ComputeMappingRelationship`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ComputeVmediaAllOf) GetMappingsOk() (*[]ComputeMappingRelationship, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ComputeVmediaAllOf) SetMappings(v []ComputeMappingRelationship)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ComputeVmediaAllOf) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *ComputeVmediaAllOf) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *ComputeVmediaAllOf) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeVmediaAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeVmediaAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeVmediaAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeVmediaAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



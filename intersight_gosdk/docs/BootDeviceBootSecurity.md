# BootDeviceBootSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.DeviceBootSecurity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.DeviceBootSecurity"]
**SecureBoot** | Pointer to **string** | The user desired BIOS secure boot as configured in the boot policy. | [optional] [readonly] 
**ComputePhysical** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBootDeviceBootSecurity

`func NewBootDeviceBootSecurity(classId string, objectType string, ) *BootDeviceBootSecurity`

NewBootDeviceBootSecurity instantiates a new BootDeviceBootSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootDeviceBootSecurityWithDefaults

`func NewBootDeviceBootSecurityWithDefaults() *BootDeviceBootSecurity`

NewBootDeviceBootSecurityWithDefaults instantiates a new BootDeviceBootSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootDeviceBootSecurity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootDeviceBootSecurity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootDeviceBootSecurity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootDeviceBootSecurity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootDeviceBootSecurity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootDeviceBootSecurity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSecureBoot

`func (o *BootDeviceBootSecurity) GetSecureBoot() string`

GetSecureBoot returns the SecureBoot field if non-nil, zero value otherwise.

### GetSecureBootOk

`func (o *BootDeviceBootSecurity) GetSecureBootOk() (*string, bool)`

GetSecureBootOk returns a tuple with the SecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBoot

`func (o *BootDeviceBootSecurity) SetSecureBoot(v string)`

SetSecureBoot sets SecureBoot field to given value.

### HasSecureBoot

`func (o *BootDeviceBootSecurity) HasSecureBoot() bool`

HasSecureBoot returns a boolean if a field has been set.

### GetComputePhysical

`func (o *BootDeviceBootSecurity) GetComputePhysical() ComputePhysicalRelationship`

GetComputePhysical returns the ComputePhysical field if non-nil, zero value otherwise.

### GetComputePhysicalOk

`func (o *BootDeviceBootSecurity) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool)`

GetComputePhysicalOk returns a tuple with the ComputePhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePhysical

`func (o *BootDeviceBootSecurity) SetComputePhysical(v ComputePhysicalRelationship)`

SetComputePhysical sets ComputePhysical field to given value.

### HasComputePhysical

`func (o *BootDeviceBootSecurity) HasComputePhysical() bool`

HasComputePhysical returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BootDeviceBootSecurity) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BootDeviceBootSecurity) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BootDeviceBootSecurity) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BootDeviceBootSecurity) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BootDeviceBootSecurity) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BootDeviceBootSecurity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BootDeviceBootSecurity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BootDeviceBootSecurity) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BootPchStorageDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.PchStorageDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.PchStorageDevice"]
**ComputePhysical** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBootPchStorageDeviceAllOf

`func NewBootPchStorageDeviceAllOf(classId string, objectType string, ) *BootPchStorageDeviceAllOf`

NewBootPchStorageDeviceAllOf instantiates a new BootPchStorageDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPchStorageDeviceAllOfWithDefaults

`func NewBootPchStorageDeviceAllOfWithDefaults() *BootPchStorageDeviceAllOf`

NewBootPchStorageDeviceAllOfWithDefaults instantiates a new BootPchStorageDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootPchStorageDeviceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootPchStorageDeviceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootPchStorageDeviceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootPchStorageDeviceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootPchStorageDeviceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootPchStorageDeviceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComputePhysical

`func (o *BootPchStorageDeviceAllOf) GetComputePhysical() ComputePhysicalRelationship`

GetComputePhysical returns the ComputePhysical field if non-nil, zero value otherwise.

### GetComputePhysicalOk

`func (o *BootPchStorageDeviceAllOf) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool)`

GetComputePhysicalOk returns a tuple with the ComputePhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePhysical

`func (o *BootPchStorageDeviceAllOf) SetComputePhysical(v ComputePhysicalRelationship)`

SetComputePhysical sets ComputePhysical field to given value.

### HasComputePhysical

`func (o *BootPchStorageDeviceAllOf) HasComputePhysical() bool`

HasComputePhysical returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BootPchStorageDeviceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BootPchStorageDeviceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BootPchStorageDeviceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BootPchStorageDeviceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BootPchStorageDeviceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BootPchStorageDeviceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BootPchStorageDeviceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BootPchStorageDeviceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MemoryUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.Unit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.Unit"]
**MemoryId** | Pointer to **int64** | This represents the ID of a regular DIMM on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryArray** | Pointer to [**MemoryArrayRelationship**](MemoryArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMemoryUnitAllOf

`func NewMemoryUnitAllOf(classId string, objectType string, ) *MemoryUnitAllOf`

NewMemoryUnitAllOf instantiates a new MemoryUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryUnitAllOfWithDefaults

`func NewMemoryUnitAllOfWithDefaults() *MemoryUnitAllOf`

NewMemoryUnitAllOfWithDefaults instantiates a new MemoryUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMemoryId

`func (o *MemoryUnitAllOf) GetMemoryId() int64`

GetMemoryId returns the MemoryId field if non-nil, zero value otherwise.

### GetMemoryIdOk

`func (o *MemoryUnitAllOf) GetMemoryIdOk() (*int64, bool)`

GetMemoryIdOk returns a tuple with the MemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryId

`func (o *MemoryUnitAllOf) SetMemoryId(v int64)`

SetMemoryId sets MemoryId field to given value.

### HasMemoryId

`func (o *MemoryUnitAllOf) HasMemoryId() bool`

HasMemoryId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryArray

`func (o *MemoryUnitAllOf) GetMemoryArray() MemoryArrayRelationship`

GetMemoryArray returns the MemoryArray field if non-nil, zero value otherwise.

### GetMemoryArrayOk

`func (o *MemoryUnitAllOf) GetMemoryArrayOk() (*MemoryArrayRelationship, bool)`

GetMemoryArrayOk returns a tuple with the MemoryArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArray

`func (o *MemoryUnitAllOf) SetMemoryArray(v MemoryArrayRelationship)`

SetMemoryArray sets MemoryArray field to given value.

### HasMemoryArray

`func (o *MemoryUnitAllOf) HasMemoryArray() bool`

HasMemoryArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



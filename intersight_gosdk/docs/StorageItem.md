# StorageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.Item"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.Item"]
**AlarmType** | Pointer to **string** | The alarmType of the Local storage in FI. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Local storage in FI. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operState of the Local storage in FI. | [optional] [readonly] 
**Size** | Pointer to **string** | The size (MB) of the Local storage in FI. | [optional] [readonly] 
**Used** | Pointer to **string** | The used percent of the Local storage in FI. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageItem

`func NewStorageItem(classId string, objectType string, ) *StorageItem`

NewStorageItem instantiates a new StorageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageItemWithDefaults

`func NewStorageItemWithDefaults() *StorageItem`

NewStorageItemWithDefaults instantiates a new StorageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmType

`func (o *StorageItem) GetAlarmType() string`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *StorageItem) GetAlarmTypeOk() (*string, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *StorageItem) SetAlarmType(v string)`

SetAlarmType sets AlarmType field to given value.

### HasAlarmType

`func (o *StorageItem) HasAlarmType() bool`

HasAlarmType returns a boolean if a field has been set.

### GetName

`func (o *StorageItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *StorageItem) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageItem) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageItem) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageItem) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetSize

`func (o *StorageItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUsed

`func (o *StorageItem) GetUsed() string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageItem) GetUsedOk() (*string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageItem) SetUsed(v string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageItem) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageItem) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageItem) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageItem) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageItem) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *StorageItem) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *StorageItem) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *StorageItem) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *StorageItem) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageItem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageItem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageItem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageItem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



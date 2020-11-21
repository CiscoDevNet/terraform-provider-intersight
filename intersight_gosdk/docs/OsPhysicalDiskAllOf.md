# OsPhysicalDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.PhysicalDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.PhysicalDisk"]
**Name** | Pointer to **string** | The Physical Disk Name to be used as Install Target. | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number of the Physical Disk Target. | [optional] 
**StorageControllerSlotId** | Pointer to **string** | The SlotID of the Storage Controller associated to the physical disk. | [optional] 

## Methods

### NewOsPhysicalDiskAllOf

`func NewOsPhysicalDiskAllOf(classId string, objectType string, ) *OsPhysicalDiskAllOf`

NewOsPhysicalDiskAllOf instantiates a new OsPhysicalDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsPhysicalDiskAllOfWithDefaults

`func NewOsPhysicalDiskAllOfWithDefaults() *OsPhysicalDiskAllOf`

NewOsPhysicalDiskAllOfWithDefaults instantiates a new OsPhysicalDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsPhysicalDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsPhysicalDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsPhysicalDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsPhysicalDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsPhysicalDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsPhysicalDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *OsPhysicalDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsPhysicalDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsPhysicalDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsPhysicalDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *OsPhysicalDiskAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OsPhysicalDiskAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OsPhysicalDiskAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OsPhysicalDiskAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStorageControllerSlotId

`func (o *OsPhysicalDiskAllOf) GetStorageControllerSlotId() string`

GetStorageControllerSlotId returns the StorageControllerSlotId field if non-nil, zero value otherwise.

### GetStorageControllerSlotIdOk

`func (o *OsPhysicalDiskAllOf) GetStorageControllerSlotIdOk() (*string, bool)`

GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerSlotId

`func (o *OsPhysicalDiskAllOf) SetStorageControllerSlotId(v string)`

SetStorageControllerSlotId sets StorageControllerSlotId field to given value.

### HasStorageControllerSlotId

`func (o *OsPhysicalDiskAllOf) HasStorageControllerSlotId() bool`

HasStorageControllerSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



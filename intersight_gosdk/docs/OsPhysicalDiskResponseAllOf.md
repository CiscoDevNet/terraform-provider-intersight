# OsPhysicalDiskResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.PhysicalDiskResponse"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.PhysicalDiskResponse"]
**Bootable** | Pointer to **string** | Bootable field of the Physical Drive target. | [optional] 
**Name** | Pointer to **string** | The Physical Disk Name to be used as Install Target. | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number of the Physical Disk Target. | [optional] 
**StorageControllerSlotId** | Pointer to **string** | The Storage Controller associated to the physical disk. | [optional] 

## Methods

### NewOsPhysicalDiskResponseAllOf

`func NewOsPhysicalDiskResponseAllOf(classId string, objectType string, ) *OsPhysicalDiskResponseAllOf`

NewOsPhysicalDiskResponseAllOf instantiates a new OsPhysicalDiskResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsPhysicalDiskResponseAllOfWithDefaults

`func NewOsPhysicalDiskResponseAllOfWithDefaults() *OsPhysicalDiskResponseAllOf`

NewOsPhysicalDiskResponseAllOfWithDefaults instantiates a new OsPhysicalDiskResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsPhysicalDiskResponseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsPhysicalDiskResponseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsPhysicalDiskResponseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsPhysicalDiskResponseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsPhysicalDiskResponseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsPhysicalDiskResponseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootable

`func (o *OsPhysicalDiskResponseAllOf) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *OsPhysicalDiskResponseAllOf) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *OsPhysicalDiskResponseAllOf) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *OsPhysicalDiskResponseAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetName

`func (o *OsPhysicalDiskResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsPhysicalDiskResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsPhysicalDiskResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsPhysicalDiskResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *OsPhysicalDiskResponseAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OsPhysicalDiskResponseAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OsPhysicalDiskResponseAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OsPhysicalDiskResponseAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStorageControllerSlotId

`func (o *OsPhysicalDiskResponseAllOf) GetStorageControllerSlotId() string`

GetStorageControllerSlotId returns the StorageControllerSlotId field if non-nil, zero value otherwise.

### GetStorageControllerSlotIdOk

`func (o *OsPhysicalDiskResponseAllOf) GetStorageControllerSlotIdOk() (*string, bool)`

GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerSlotId

`func (o *OsPhysicalDiskResponseAllOf) SetStorageControllerSlotId(v string)`

SetStorageControllerSlotId sets StorageControllerSlotId field to given value.

### HasStorageControllerSlotId

`func (o *OsPhysicalDiskResponseAllOf) HasStorageControllerSlotId() bool`

HasStorageControllerSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



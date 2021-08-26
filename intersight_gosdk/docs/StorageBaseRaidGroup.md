# StorageBaseRaidGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.HitachiParityGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.HitachiParityGroup"]
**Level** | Pointer to **string** | The RAID level associated with the group of disks. * &#x60;Unknown&#x60; - Default unknown RAID type. * &#x60;RAID0&#x60; - RAID0 splits (\&quot;stripes\&quot;) data evenly across two or more disks, without parity information. * &#x60;RAID1&#x60; - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * &#x60;RAID4&#x60; - RAID4 stripes block level data and dedicates a disk to parity. * &#x60;RAID5&#x60; - RAID5  distributes striping and parity at a block level. * &#x60;RAID6&#x60; - RAID6 level operates like RAID5 with distributed parity and striping. * &#x60;RAID10&#x60; - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * &#x60;RAIDDP&#x60; - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * &#x60;RAIDTEC&#x60; - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group. | [optional] [readonly] [default to "Unknown"]
**Name** | Pointer to **string** | Human readable name of the RAID group. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**NullableStorageBaseCapacity**](StorageBaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseRaidGroup

`func NewStorageBaseRaidGroup(classId string, objectType string, ) *StorageBaseRaidGroup`

NewStorageBaseRaidGroup instantiates a new StorageBaseRaidGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseRaidGroupWithDefaults

`func NewStorageBaseRaidGroupWithDefaults() *StorageBaseRaidGroup`

NewStorageBaseRaidGroupWithDefaults instantiates a new StorageBaseRaidGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseRaidGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseRaidGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseRaidGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseRaidGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseRaidGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseRaidGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLevel

`func (o *StorageBaseRaidGroup) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *StorageBaseRaidGroup) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *StorageBaseRaidGroup) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *StorageBaseRaidGroup) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseRaidGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseRaidGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseRaidGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseRaidGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseRaidGroup) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseRaidGroup) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseRaidGroup) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseRaidGroup) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageBaseRaidGroup) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageBaseRaidGroup) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



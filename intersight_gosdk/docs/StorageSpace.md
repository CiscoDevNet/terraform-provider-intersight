# StorageSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.Space"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.Space"]
**LbaLocation** | Pointer to **string** | Starting location of the LBA of the partition in the external parity group (in a multiple of 512 bytes). | [optional] [readonly] 
**LbaSize** | Pointer to **string** | Size of the partition in the external parity group (in a multiple of 512 bytes). | [optional] [readonly] 
**LdevId** | Pointer to **string** | LDEV number, property is output only if LDEV has been implemented. | [optional] [readonly] 
**PartitionNumber** | Pointer to **int64** | Number of a partition created as a result of partitioning of an external parity group. | [optional] [readonly] 
**Status** | Pointer to **string** | Status about LDEV for partition. | [optional] [readonly] 

## Methods

### NewStorageSpace

`func NewStorageSpace(classId string, objectType string, ) *StorageSpace`

NewStorageSpace instantiates a new StorageSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpaceWithDefaults

`func NewStorageSpaceWithDefaults() *StorageSpace`

NewStorageSpaceWithDefaults instantiates a new StorageSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageSpace) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageSpace) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageSpace) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageSpace) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageSpace) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageSpace) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLbaLocation

`func (o *StorageSpace) GetLbaLocation() string`

GetLbaLocation returns the LbaLocation field if non-nil, zero value otherwise.

### GetLbaLocationOk

`func (o *StorageSpace) GetLbaLocationOk() (*string, bool)`

GetLbaLocationOk returns a tuple with the LbaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbaLocation

`func (o *StorageSpace) SetLbaLocation(v string)`

SetLbaLocation sets LbaLocation field to given value.

### HasLbaLocation

`func (o *StorageSpace) HasLbaLocation() bool`

HasLbaLocation returns a boolean if a field has been set.

### GetLbaSize

`func (o *StorageSpace) GetLbaSize() string`

GetLbaSize returns the LbaSize field if non-nil, zero value otherwise.

### GetLbaSizeOk

`func (o *StorageSpace) GetLbaSizeOk() (*string, bool)`

GetLbaSizeOk returns a tuple with the LbaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbaSize

`func (o *StorageSpace) SetLbaSize(v string)`

SetLbaSize sets LbaSize field to given value.

### HasLbaSize

`func (o *StorageSpace) HasLbaSize() bool`

HasLbaSize returns a boolean if a field has been set.

### GetLdevId

`func (o *StorageSpace) GetLdevId() string`

GetLdevId returns the LdevId field if non-nil, zero value otherwise.

### GetLdevIdOk

`func (o *StorageSpace) GetLdevIdOk() (*string, bool)`

GetLdevIdOk returns a tuple with the LdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdevId

`func (o *StorageSpace) SetLdevId(v string)`

SetLdevId sets LdevId field to given value.

### HasLdevId

`func (o *StorageSpace) HasLdevId() bool`

HasLdevId returns a boolean if a field has been set.

### GetPartitionNumber

`func (o *StorageSpace) GetPartitionNumber() int64`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *StorageSpace) GetPartitionNumberOk() (*int64, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *StorageSpace) SetPartitionNumber(v int64)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *StorageSpace) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### GetStatus

`func (o *StorageSpace) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageSpace) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageSpace) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageSpace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



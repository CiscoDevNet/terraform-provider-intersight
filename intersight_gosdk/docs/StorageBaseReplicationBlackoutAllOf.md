# StorageBaseReplicationBlackoutAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureReplicationBlackout"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureReplicationBlackout"]
**End** | Pointer to **string** | The end time of day for replication blackout window. Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds. | [optional] [readonly] 
**Start** | Pointer to **string** | The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication. Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds. The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds. | [optional] [readonly] 

## Methods

### NewStorageBaseReplicationBlackoutAllOf

`func NewStorageBaseReplicationBlackoutAllOf(classId string, objectType string, ) *StorageBaseReplicationBlackoutAllOf`

NewStorageBaseReplicationBlackoutAllOf instantiates a new StorageBaseReplicationBlackoutAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseReplicationBlackoutAllOfWithDefaults

`func NewStorageBaseReplicationBlackoutAllOfWithDefaults() *StorageBaseReplicationBlackoutAllOf`

NewStorageBaseReplicationBlackoutAllOfWithDefaults instantiates a new StorageBaseReplicationBlackoutAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseReplicationBlackoutAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseReplicationBlackoutAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseReplicationBlackoutAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseReplicationBlackoutAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseReplicationBlackoutAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseReplicationBlackoutAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnd

`func (o *StorageBaseReplicationBlackoutAllOf) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *StorageBaseReplicationBlackoutAllOf) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *StorageBaseReplicationBlackoutAllOf) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *StorageBaseReplicationBlackoutAllOf) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *StorageBaseReplicationBlackoutAllOf) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *StorageBaseReplicationBlackoutAllOf) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *StorageBaseReplicationBlackoutAllOf) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *StorageBaseReplicationBlackoutAllOf) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



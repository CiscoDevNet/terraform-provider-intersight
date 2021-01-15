# HyperflexReplicationScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationSchedule"]
**BackupInterval** | Pointer to **int64** | Time interval between two copies in minutes. | [optional] [default to 240]

## Methods

### NewHyperflexReplicationScheduleAllOf

`func NewHyperflexReplicationScheduleAllOf(classId string, objectType string, ) *HyperflexReplicationScheduleAllOf`

NewHyperflexReplicationScheduleAllOf instantiates a new HyperflexReplicationScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationScheduleAllOfWithDefaults

`func NewHyperflexReplicationScheduleAllOfWithDefaults() *HyperflexReplicationScheduleAllOf`

NewHyperflexReplicationScheduleAllOfWithDefaults instantiates a new HyperflexReplicationScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupInterval

`func (o *HyperflexReplicationScheduleAllOf) GetBackupInterval() int64`

GetBackupInterval returns the BackupInterval field if non-nil, zero value otherwise.

### GetBackupIntervalOk

`func (o *HyperflexReplicationScheduleAllOf) GetBackupIntervalOk() (*int64, bool)`

GetBackupIntervalOk returns a tuple with the BackupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterval

`func (o *HyperflexReplicationScheduleAllOf) SetBackupInterval(v int64)`

SetBackupInterval sets BackupInterval field to given value.

### HasBackupInterval

`func (o *HyperflexReplicationScheduleAllOf) HasBackupInterval() bool`

HasBackupInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



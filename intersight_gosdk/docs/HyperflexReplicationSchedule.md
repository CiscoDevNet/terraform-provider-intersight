# HyperflexReplicationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationSchedule"]
**BackupInterval** | Pointer to **int64** | Time interval between two copies in minutes. | [optional] [default to 240]

## Methods

### NewHyperflexReplicationSchedule

`func NewHyperflexReplicationSchedule(classId string, objectType string, ) *HyperflexReplicationSchedule`

NewHyperflexReplicationSchedule instantiates a new HyperflexReplicationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationScheduleWithDefaults

`func NewHyperflexReplicationScheduleWithDefaults() *HyperflexReplicationSchedule`

NewHyperflexReplicationScheduleWithDefaults instantiates a new HyperflexReplicationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupInterval

`func (o *HyperflexReplicationSchedule) GetBackupInterval() int64`

GetBackupInterval returns the BackupInterval field if non-nil, zero value otherwise.

### GetBackupIntervalOk

`func (o *HyperflexReplicationSchedule) GetBackupIntervalOk() (*int64, bool)`

GetBackupIntervalOk returns a tuple with the BackupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterval

`func (o *HyperflexReplicationSchedule) SetBackupInterval(v int64)`

SetBackupInterval sets BackupInterval field to given value.

### HasBackupInterval

`func (o *HyperflexReplicationSchedule) HasBackupInterval() bool`

HasBackupInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



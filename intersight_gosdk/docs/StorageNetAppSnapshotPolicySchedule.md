# StorageNetAppSnapshotPolicySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSnapshotPolicySchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSnapshotPolicySchedule"]
**Count** | Pointer to **int64** | The number of Snapshot copies to maintain for this schedule. | [optional] [readonly] 
**Label** | Pointer to **string** | Label for SnapMirror operations. | [optional] [readonly] 
**ScheduleName** | Pointer to **string** | Schedule at which Snapshot copies are captured on the volume. | [optional] [readonly] 

## Methods

### NewStorageNetAppSnapshotPolicySchedule

`func NewStorageNetAppSnapshotPolicySchedule(classId string, objectType string, ) *StorageNetAppSnapshotPolicySchedule`

NewStorageNetAppSnapshotPolicySchedule instantiates a new StorageNetAppSnapshotPolicySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppSnapshotPolicyScheduleWithDefaults

`func NewStorageNetAppSnapshotPolicyScheduleWithDefaults() *StorageNetAppSnapshotPolicySchedule`

NewStorageNetAppSnapshotPolicyScheduleWithDefaults instantiates a new StorageNetAppSnapshotPolicySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSnapshotPolicySchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSnapshotPolicySchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSnapshotPolicySchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSnapshotPolicySchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSnapshotPolicySchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSnapshotPolicySchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *StorageNetAppSnapshotPolicySchedule) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StorageNetAppSnapshotPolicySchedule) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StorageNetAppSnapshotPolicySchedule) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *StorageNetAppSnapshotPolicySchedule) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLabel

`func (o *StorageNetAppSnapshotPolicySchedule) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StorageNetAppSnapshotPolicySchedule) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StorageNetAppSnapshotPolicySchedule) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StorageNetAppSnapshotPolicySchedule) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScheduleName

`func (o *StorageNetAppSnapshotPolicySchedule) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *StorageNetAppSnapshotPolicySchedule) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *StorageNetAppSnapshotPolicySchedule) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *StorageNetAppSnapshotPolicySchedule) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



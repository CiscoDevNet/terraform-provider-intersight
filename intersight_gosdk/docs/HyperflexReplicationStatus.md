# HyperflexReplicationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationStatus"]
**BytesReplicated** | Pointer to **int64** | Number of bytes currently replicated. | [optional] [readonly] 
**EndTime** | Pointer to **int64** | Replication end time for this snapshot. | [optional] [readonly] 
**Error** | Pointer to [**NullableHyperflexErrorStack**](HyperflexErrorStack.md) |  | [optional] 
**PackEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**PctComplete** | Pointer to **int64** | Completion percentage for the replication job. | [optional] [readonly] 
**RpoStatus** | Pointer to [**NullableHyperflexRpoStatus**](HyperflexRpoStatus.md) |  | [optional] 
**StartTime** | Pointer to **int64** | Replication start time for this snapshot. | [optional] [readonly] 
**Status** | Pointer to **string** | Current replication state for a particular snapshot. * &#x60;NONE&#x60; - Snapshot replication state is none. * &#x60;SUCCESS&#x60; - Snapshot completed successfully. * &#x60;FAILED&#x60; - Snapshot failed replication status code. * &#x60;PAUSED&#x60; - Snapshot replication paused status code. * &#x60;IN_USE&#x60; - Snapshot replica in use status code. * &#x60;STARTING&#x60; - Snapshot replication starting. * &#x60;REPLICATING&#x60; - Snapshot replication in progress. | [optional] [readonly] [default to "NONE"]

## Methods

### NewHyperflexReplicationStatus

`func NewHyperflexReplicationStatus(classId string, objectType string, ) *HyperflexReplicationStatus`

NewHyperflexReplicationStatus instantiates a new HyperflexReplicationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationStatusWithDefaults

`func NewHyperflexReplicationStatusWithDefaults() *HyperflexReplicationStatus`

NewHyperflexReplicationStatusWithDefaults instantiates a new HyperflexReplicationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBytesReplicated

`func (o *HyperflexReplicationStatus) GetBytesReplicated() int64`

GetBytesReplicated returns the BytesReplicated field if non-nil, zero value otherwise.

### GetBytesReplicatedOk

`func (o *HyperflexReplicationStatus) GetBytesReplicatedOk() (*int64, bool)`

GetBytesReplicatedOk returns a tuple with the BytesReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReplicated

`func (o *HyperflexReplicationStatus) SetBytesReplicated(v int64)`

SetBytesReplicated sets BytesReplicated field to given value.

### HasBytesReplicated

`func (o *HyperflexReplicationStatus) HasBytesReplicated() bool`

HasBytesReplicated returns a boolean if a field has been set.

### GetEndTime

`func (o *HyperflexReplicationStatus) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HyperflexReplicationStatus) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HyperflexReplicationStatus) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HyperflexReplicationStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetError

`func (o *HyperflexReplicationStatus) GetError() HyperflexErrorStack`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperflexReplicationStatus) GetErrorOk() (*HyperflexErrorStack, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperflexReplicationStatus) SetError(v HyperflexErrorStack)`

SetError sets Error field to given value.

### HasError

`func (o *HyperflexReplicationStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HyperflexReplicationStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HyperflexReplicationStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPackEntityReference

`func (o *HyperflexReplicationStatus) GetPackEntityReference() HyperflexEntityReference`

GetPackEntityReference returns the PackEntityReference field if non-nil, zero value otherwise.

### GetPackEntityReferenceOk

`func (o *HyperflexReplicationStatus) GetPackEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetPackEntityReferenceOk returns a tuple with the PackEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackEntityReference

`func (o *HyperflexReplicationStatus) SetPackEntityReference(v HyperflexEntityReference)`

SetPackEntityReference sets PackEntityReference field to given value.

### HasPackEntityReference

`func (o *HyperflexReplicationStatus) HasPackEntityReference() bool`

HasPackEntityReference returns a boolean if a field has been set.

### SetPackEntityReferenceNil

`func (o *HyperflexReplicationStatus) SetPackEntityReferenceNil(b bool)`

 SetPackEntityReferenceNil sets the value for PackEntityReference to be an explicit nil

### UnsetPackEntityReference
`func (o *HyperflexReplicationStatus) UnsetPackEntityReference()`

UnsetPackEntityReference ensures that no value is present for PackEntityReference, not even an explicit nil
### GetPctComplete

`func (o *HyperflexReplicationStatus) GetPctComplete() int64`

GetPctComplete returns the PctComplete field if non-nil, zero value otherwise.

### GetPctCompleteOk

`func (o *HyperflexReplicationStatus) GetPctCompleteOk() (*int64, bool)`

GetPctCompleteOk returns a tuple with the PctComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctComplete

`func (o *HyperflexReplicationStatus) SetPctComplete(v int64)`

SetPctComplete sets PctComplete field to given value.

### HasPctComplete

`func (o *HyperflexReplicationStatus) HasPctComplete() bool`

HasPctComplete returns a boolean if a field has been set.

### GetRpoStatus

`func (o *HyperflexReplicationStatus) GetRpoStatus() HyperflexRpoStatus`

GetRpoStatus returns the RpoStatus field if non-nil, zero value otherwise.

### GetRpoStatusOk

`func (o *HyperflexReplicationStatus) GetRpoStatusOk() (*HyperflexRpoStatus, bool)`

GetRpoStatusOk returns a tuple with the RpoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoStatus

`func (o *HyperflexReplicationStatus) SetRpoStatus(v HyperflexRpoStatus)`

SetRpoStatus sets RpoStatus field to given value.

### HasRpoStatus

`func (o *HyperflexReplicationStatus) HasRpoStatus() bool`

HasRpoStatus returns a boolean if a field has been set.

### SetRpoStatusNil

`func (o *HyperflexReplicationStatus) SetRpoStatusNil(b bool)`

 SetRpoStatusNil sets the value for RpoStatus to be an explicit nil

### UnsetRpoStatus
`func (o *HyperflexReplicationStatus) UnsetRpoStatus()`

UnsetRpoStatus ensures that no value is present for RpoStatus, not even an explicit nil
### GetStartTime

`func (o *HyperflexReplicationStatus) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexReplicationStatus) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexReplicationStatus) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexReplicationStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexReplicationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexReplicationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexReplicationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexReplicationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



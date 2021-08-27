# HyperflexSnapshotStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SnapshotStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SnapshotStatus"]
**Description** | Pointer to **string** | Description of this Snapshot Point. | [optional] [readonly] 
**Error** | Pointer to [**NullableHyperflexErrorStack**](HyperflexErrorStack.md) |  | [optional] 
**PctComplete** | Pointer to **int64** | Completion percentage for this snapshot. | [optional] [readonly] 
**Status** | Pointer to **string** | Current snapshot state for this snapshot. * &#x60;SUCCESS&#x60; - This snapshot status code is success. * &#x60;FAILED&#x60; - This snapshot status code is failed. * &#x60;IN_PROGRESS&#x60; - This snapshot status code is in progress. * &#x60;DELETING&#x60; - This snapshot status code is deleting. * &#x60;DELETED&#x60; - This snapshot status code is deleted. * &#x60;NONE&#x60; - This snapshot status code is none. | [optional] [readonly] [default to "SUCCESS"]
**Timestamp** | Pointer to **int64** | Timestamp at which the Snapshot is taken. | [optional] [readonly] 
**UsedSpace** | Pointer to **int64** | Space Used by this Snapshot Point. | [optional] [readonly] 

## Methods

### NewHyperflexSnapshotStatus

`func NewHyperflexSnapshotStatus(classId string, objectType string, ) *HyperflexSnapshotStatus`

NewHyperflexSnapshotStatus instantiates a new HyperflexSnapshotStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSnapshotStatusWithDefaults

`func NewHyperflexSnapshotStatusWithDefaults() *HyperflexSnapshotStatus`

NewHyperflexSnapshotStatusWithDefaults instantiates a new HyperflexSnapshotStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSnapshotStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSnapshotStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSnapshotStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSnapshotStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSnapshotStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSnapshotStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *HyperflexSnapshotStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexSnapshotStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexSnapshotStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexSnapshotStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetError

`func (o *HyperflexSnapshotStatus) GetError() HyperflexErrorStack`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperflexSnapshotStatus) GetErrorOk() (*HyperflexErrorStack, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperflexSnapshotStatus) SetError(v HyperflexErrorStack)`

SetError sets Error field to given value.

### HasError

`func (o *HyperflexSnapshotStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HyperflexSnapshotStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HyperflexSnapshotStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPctComplete

`func (o *HyperflexSnapshotStatus) GetPctComplete() int64`

GetPctComplete returns the PctComplete field if non-nil, zero value otherwise.

### GetPctCompleteOk

`func (o *HyperflexSnapshotStatus) GetPctCompleteOk() (*int64, bool)`

GetPctCompleteOk returns a tuple with the PctComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctComplete

`func (o *HyperflexSnapshotStatus) SetPctComplete(v int64)`

SetPctComplete sets PctComplete field to given value.

### HasPctComplete

`func (o *HyperflexSnapshotStatus) HasPctComplete() bool`

HasPctComplete returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexSnapshotStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexSnapshotStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexSnapshotStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexSnapshotStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *HyperflexSnapshotStatus) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HyperflexSnapshotStatus) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HyperflexSnapshotStatus) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HyperflexSnapshotStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsedSpace

`func (o *HyperflexSnapshotStatus) GetUsedSpace() int64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *HyperflexSnapshotStatus) GetUsedSpaceOk() (*int64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *HyperflexSnapshotStatus) SetUsedSpace(v int64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *HyperflexSnapshotStatus) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



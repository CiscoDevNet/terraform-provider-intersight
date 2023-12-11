# HyperflexReduceReSyncAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReduceReSync"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReduceReSync"]
**CompletionStatus** | Pointer to **bool** | The HyperFlex reduce re-sync script execution completion status. | [optional] 
**ExecutionOutput** | Pointer to **string** | The execution output of the script. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexReduceReSyncAllOf

`func NewHyperflexReduceReSyncAllOf(classId string, objectType string, ) *HyperflexReduceReSyncAllOf`

NewHyperflexReduceReSyncAllOf instantiates a new HyperflexReduceReSyncAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReduceReSyncAllOfWithDefaults

`func NewHyperflexReduceReSyncAllOfWithDefaults() *HyperflexReduceReSyncAllOf`

NewHyperflexReduceReSyncAllOfWithDefaults instantiates a new HyperflexReduceReSyncAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReduceReSyncAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReduceReSyncAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReduceReSyncAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReduceReSyncAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReduceReSyncAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReduceReSyncAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletionStatus

`func (o *HyperflexReduceReSyncAllOf) GetCompletionStatus() bool`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *HyperflexReduceReSyncAllOf) GetCompletionStatusOk() (*bool, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *HyperflexReduceReSyncAllOf) SetCompletionStatus(v bool)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *HyperflexReduceReSyncAllOf) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetExecutionOutput

`func (o *HyperflexReduceReSyncAllOf) GetExecutionOutput() string`

GetExecutionOutput returns the ExecutionOutput field if non-nil, zero value otherwise.

### GetExecutionOutputOk

`func (o *HyperflexReduceReSyncAllOf) GetExecutionOutputOk() (*string, bool)`

GetExecutionOutputOk returns a tuple with the ExecutionOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOutput

`func (o *HyperflexReduceReSyncAllOf) SetExecutionOutput(v string)`

SetExecutionOutput sets ExecutionOutput field to given value.

### HasExecutionOutput

`func (o *HyperflexReduceReSyncAllOf) HasExecutionOutput() bool`

HasExecutionOutput returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexReduceReSyncAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexReduceReSyncAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexReduceReSyncAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexReduceReSyncAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



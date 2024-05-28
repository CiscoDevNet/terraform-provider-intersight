# HyperflexReduceReSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReduceReSync"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReduceReSync"]
**CompletionStatus** | Pointer to **bool** | The HyperFlex reduce re-sync script execution completion status. | [optional] 
**ExecutionOutput** | Pointer to **string** | The execution output of the script. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexReduceReSync

`func NewHyperflexReduceReSync(classId string, objectType string, ) *HyperflexReduceReSync`

NewHyperflexReduceReSync instantiates a new HyperflexReduceReSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReduceReSyncWithDefaults

`func NewHyperflexReduceReSyncWithDefaults() *HyperflexReduceReSync`

NewHyperflexReduceReSyncWithDefaults instantiates a new HyperflexReduceReSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReduceReSync) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReduceReSync) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReduceReSync) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReduceReSync) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReduceReSync) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReduceReSync) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletionStatus

`func (o *HyperflexReduceReSync) GetCompletionStatus() bool`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *HyperflexReduceReSync) GetCompletionStatusOk() (*bool, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *HyperflexReduceReSync) SetCompletionStatus(v bool)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *HyperflexReduceReSync) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetExecutionOutput

`func (o *HyperflexReduceReSync) GetExecutionOutput() string`

GetExecutionOutput returns the ExecutionOutput field if non-nil, zero value otherwise.

### GetExecutionOutputOk

`func (o *HyperflexReduceReSync) GetExecutionOutputOk() (*string, bool)`

GetExecutionOutputOk returns a tuple with the ExecutionOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOutput

`func (o *HyperflexReduceReSync) SetExecutionOutput(v string)`

SetExecutionOutput sets ExecutionOutput field to given value.

### HasExecutionOutput

`func (o *HyperflexReduceReSync) HasExecutionOutput() bool`

HasExecutionOutput returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexReduceReSync) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexReduceReSync) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexReduceReSync) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexReduceReSync) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HyperflexReduceReSync) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HyperflexReduceReSync) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



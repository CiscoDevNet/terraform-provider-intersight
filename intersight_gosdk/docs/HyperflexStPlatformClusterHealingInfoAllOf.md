# HyperflexStPlatformClusterHealingInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.StPlatformClusterHealingInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.StPlatformClusterHealingInfo"]
**EstimatedCompletionTimeInSeconds** | Pointer to **int64** | The estimated time in seconds it will take to complete the auto-healing process. | [optional] [readonly] 
**InProgress** | Pointer to **bool** | The status of the cluster&#39;s auto-healing process. If &#39;true&#39;, auto-healing is in progress for the cluster. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**MessagesIterator** | Pointer to **interface{}** | The current message describing the auto-healing process of the cluster. | [optional] [readonly] 
**MessagesSize** | Pointer to **int64** | The number of elements in the messages collection. | [optional] [readonly] 
**PercentComplete** | Pointer to **int64** | The progress of the auto-healing process as a percentage. | [optional] [readonly] 

## Methods

### NewHyperflexStPlatformClusterHealingInfoAllOf

`func NewHyperflexStPlatformClusterHealingInfoAllOf(classId string, objectType string, ) *HyperflexStPlatformClusterHealingInfoAllOf`

NewHyperflexStPlatformClusterHealingInfoAllOf instantiates a new HyperflexStPlatformClusterHealingInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStPlatformClusterHealingInfoAllOfWithDefaults

`func NewHyperflexStPlatformClusterHealingInfoAllOfWithDefaults() *HyperflexStPlatformClusterHealingInfoAllOf`

NewHyperflexStPlatformClusterHealingInfoAllOfWithDefaults instantiates a new HyperflexStPlatformClusterHealingInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetEstimatedCompletionTimeInSeconds() int64`

GetEstimatedCompletionTimeInSeconds returns the EstimatedCompletionTimeInSeconds field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeInSecondsOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetEstimatedCompletionTimeInSecondsOk() (*int64, bool)`

GetEstimatedCompletionTimeInSecondsOk returns a tuple with the EstimatedCompletionTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetEstimatedCompletionTimeInSeconds(v int64)`

SetEstimatedCompletionTimeInSeconds sets EstimatedCompletionTimeInSeconds field to given value.

### HasEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasEstimatedCompletionTimeInSeconds() bool`

HasEstimatedCompletionTimeInSeconds returns a boolean if a field has been set.

### GetInProgress

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *HyperflexStPlatformClusterHealingInfoAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesIterator() interface{}`

GetMessagesIterator returns the MessagesIterator field if non-nil, zero value otherwise.

### GetMessagesIteratorOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesIteratorOk() (*interface{}, bool)`

GetMessagesIteratorOk returns a tuple with the MessagesIterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesIterator(v interface{})`

SetMessagesIterator sets MessagesIterator field to given value.

### HasMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessagesIterator() bool`

HasMessagesIterator returns a boolean if a field has been set.

### SetMessagesIteratorNil

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesIteratorNil(b bool)`

 SetMessagesIteratorNil sets the value for MessagesIterator to be an explicit nil

### UnsetMessagesIterator
`func (o *HyperflexStPlatformClusterHealingInfoAllOf) UnsetMessagesIterator()`

UnsetMessagesIterator ensures that no value is present for MessagesIterator, not even an explicit nil
### GetMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesSize() int64`

GetMessagesSize returns the MessagesSize field if non-nil, zero value otherwise.

### GetMessagesSizeOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesSizeOk() (*int64, bool)`

GetMessagesSizeOk returns a tuple with the MessagesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesSize(v int64)`

SetMessagesSize sets MessagesSize field to given value.

### HasMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessagesSize() bool`

HasMessagesSize returns a boolean if a field has been set.

### GetPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetPercentComplete() int64`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetPercentCompleteOk() (*int64, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetPercentComplete(v int64)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



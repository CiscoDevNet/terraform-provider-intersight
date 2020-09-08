# HyperflexStPlatformClusterHealingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedCompletionTimeInSeconds** | Pointer to **int64** |  | [optional] [readonly] 
**InProgress** | Pointer to **bool** |  | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**MessagesIterator** | Pointer to **interface{}** |  | [optional] [readonly] 
**MessagesSize** | Pointer to **int64** |  | [optional] [readonly] 
**PercentComplete** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewHyperflexStPlatformClusterHealingInfo

`func NewHyperflexStPlatformClusterHealingInfo() *HyperflexStPlatformClusterHealingInfo`

NewHyperflexStPlatformClusterHealingInfo instantiates a new HyperflexStPlatformClusterHealingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStPlatformClusterHealingInfoWithDefaults

`func NewHyperflexStPlatformClusterHealingInfoWithDefaults() *HyperflexStPlatformClusterHealingInfo`

NewHyperflexStPlatformClusterHealingInfoWithDefaults instantiates a new HyperflexStPlatformClusterHealingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfo) GetEstimatedCompletionTimeInSeconds() int64`

GetEstimatedCompletionTimeInSeconds returns the EstimatedCompletionTimeInSeconds field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeInSecondsOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetEstimatedCompletionTimeInSecondsOk() (*int64, bool)`

GetEstimatedCompletionTimeInSecondsOk returns a tuple with the EstimatedCompletionTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfo) SetEstimatedCompletionTimeInSeconds(v int64)`

SetEstimatedCompletionTimeInSeconds sets EstimatedCompletionTimeInSeconds field to given value.

### HasEstimatedCompletionTimeInSeconds

`func (o *HyperflexStPlatformClusterHealingInfo) HasEstimatedCompletionTimeInSeconds() bool`

HasEstimatedCompletionTimeInSeconds returns a boolean if a field has been set.

### GetInProgress

`func (o *HyperflexStPlatformClusterHealingInfo) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *HyperflexStPlatformClusterHealingInfo) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *HyperflexStPlatformClusterHealingInfo) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexStPlatformClusterHealingInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexStPlatformClusterHealingInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesIterator() interface{}`

GetMessagesIterator returns the MessagesIterator field if non-nil, zero value otherwise.

### GetMessagesIteratorOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesIteratorOk() (*interface{}, bool)`

GetMessagesIteratorOk returns a tuple with the MessagesIterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfo) SetMessagesIterator(v interface{})`

SetMessagesIterator sets MessagesIterator field to given value.

### HasMessagesIterator

`func (o *HyperflexStPlatformClusterHealingInfo) HasMessagesIterator() bool`

HasMessagesIterator returns a boolean if a field has been set.

### SetMessagesIteratorNil

`func (o *HyperflexStPlatformClusterHealingInfo) SetMessagesIteratorNil(b bool)`

 SetMessagesIteratorNil sets the value for MessagesIterator to be an explicit nil

### UnsetMessagesIterator
`func (o *HyperflexStPlatformClusterHealingInfo) UnsetMessagesIterator()`

UnsetMessagesIterator ensures that no value is present for MessagesIterator, not even an explicit nil
### GetMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesSize() int64`

GetMessagesSize returns the MessagesSize field if non-nil, zero value otherwise.

### GetMessagesSizeOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesSizeOk() (*int64, bool)`

GetMessagesSizeOk returns a tuple with the MessagesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfo) SetMessagesSize(v int64)`

SetMessagesSize sets MessagesSize field to given value.

### HasMessagesSize

`func (o *HyperflexStPlatformClusterHealingInfo) HasMessagesSize() bool`

HasMessagesSize returns a boolean if a field has been set.

### GetPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfo) GetPercentComplete() int64`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *HyperflexStPlatformClusterHealingInfo) GetPercentCompleteOk() (*int64, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfo) SetPercentComplete(v int64)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *HyperflexStPlatformClusterHealingInfo) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



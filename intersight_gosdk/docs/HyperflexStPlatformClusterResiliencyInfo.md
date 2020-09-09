# HyperflexStPlatformClusterResiliencyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HddFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**MessagesIterator** | Pointer to **interface{}** |  | [optional] [readonly] 
**MessagesSize** | Pointer to **int64** |  | [optional] [readonly] 
**NodeFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**SsdFailuresTolerable** | Pointer to **int64** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewHyperflexStPlatformClusterResiliencyInfo

`func NewHyperflexStPlatformClusterResiliencyInfo() *HyperflexStPlatformClusterResiliencyInfo`

NewHyperflexStPlatformClusterResiliencyInfo instantiates a new HyperflexStPlatformClusterResiliencyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStPlatformClusterResiliencyInfoWithDefaults

`func NewHyperflexStPlatformClusterResiliencyInfoWithDefaults() *HyperflexStPlatformClusterResiliencyInfo`

NewHyperflexStPlatformClusterResiliencyInfoWithDefaults instantiates a new HyperflexStPlatformClusterResiliencyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetHddFailuresTolerable() int64`

GetHddFailuresTolerable returns the HddFailuresTolerable field if non-nil, zero value otherwise.

### GetHddFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetHddFailuresTolerableOk() (*int64, bool)`

GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetHddFailuresTolerable(v int64)`

SetHddFailuresTolerable sets HddFailuresTolerable field to given value.

### HasHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasHddFailuresTolerable() bool`

HasHddFailuresTolerable returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesIterator() interface{}`

GetMessagesIterator returns the MessagesIterator field if non-nil, zero value otherwise.

### GetMessagesIteratorOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesIteratorOk() (*interface{}, bool)`

GetMessagesIteratorOk returns a tuple with the MessagesIterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessagesIterator(v interface{})`

SetMessagesIterator sets MessagesIterator field to given value.

### HasMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessagesIterator() bool`

HasMessagesIterator returns a boolean if a field has been set.

### SetMessagesIteratorNil

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessagesIteratorNil(b bool)`

 SetMessagesIteratorNil sets the value for MessagesIterator to be an explicit nil

### UnsetMessagesIterator
`func (o *HyperflexStPlatformClusterResiliencyInfo) UnsetMessagesIterator()`

UnsetMessagesIterator ensures that no value is present for MessagesIterator, not even an explicit nil
### GetMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesSize() int64`

GetMessagesSize returns the MessagesSize field if non-nil, zero value otherwise.

### GetMessagesSizeOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetMessagesSizeOk() (*int64, bool)`

GetMessagesSizeOk returns a tuple with the MessagesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetMessagesSize(v int64)`

SetMessagesSize sets MessagesSize field to given value.

### HasMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasMessagesSize() bool`

HasMessagesSize returns a boolean if a field has been set.

### GetNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetNodeFailuresTolerable() int64`

GetNodeFailuresTolerable returns the NodeFailuresTolerable field if non-nil, zero value otherwise.

### GetNodeFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetNodeFailuresTolerableOk() (*int64, bool)`

GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetNodeFailuresTolerable(v int64)`

SetNodeFailuresTolerable sets NodeFailuresTolerable field to given value.

### HasNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasNodeFailuresTolerable() bool`

HasNodeFailuresTolerable returns a boolean if a field has been set.

### GetSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetSsdFailuresTolerable() int64`

GetSsdFailuresTolerable returns the SsdFailuresTolerable field if non-nil, zero value otherwise.

### GetSsdFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetSsdFailuresTolerableOk() (*int64, bool)`

GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetSsdFailuresTolerable(v int64)`

SetSsdFailuresTolerable sets SsdFailuresTolerable field to given value.

### HasSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasSsdFailuresTolerable() bool`

HasSsdFailuresTolerable returns a boolean if a field has been set.

### GetState

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexStPlatformClusterResiliencyInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexStPlatformClusterResiliencyInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexStPlatformClusterResiliencyInfo) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



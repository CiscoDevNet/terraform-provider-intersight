# HyperflexStPlatformClusterResiliencyInfoAllOf

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

### NewHyperflexStPlatformClusterResiliencyInfoAllOf

`func NewHyperflexStPlatformClusterResiliencyInfoAllOf() *HyperflexStPlatformClusterResiliencyInfoAllOf`

NewHyperflexStPlatformClusterResiliencyInfoAllOf instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults

`func NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults() *HyperflexStPlatformClusterResiliencyInfoAllOf`

NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetHddFailuresTolerable() int64`

GetHddFailuresTolerable returns the HddFailuresTolerable field if non-nil, zero value otherwise.

### GetHddFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetHddFailuresTolerableOk() (*int64, bool)`

GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetHddFailuresTolerable(v int64)`

SetHddFailuresTolerable sets HddFailuresTolerable field to given value.

### HasHddFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasHddFailuresTolerable() bool`

HasHddFailuresTolerable returns a boolean if a field has been set.

### GetMessages

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesIterator() interface{}`

GetMessagesIterator returns the MessagesIterator field if non-nil, zero value otherwise.

### GetMessagesIteratorOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesIteratorOk() (*interface{}, bool)`

GetMessagesIteratorOk returns a tuple with the MessagesIterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesIterator(v interface{})`

SetMessagesIterator sets MessagesIterator field to given value.

### HasMessagesIterator

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessagesIterator() bool`

HasMessagesIterator returns a boolean if a field has been set.

### SetMessagesIteratorNil

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesIteratorNil(b bool)`

 SetMessagesIteratorNil sets the value for MessagesIterator to be an explicit nil

### UnsetMessagesIterator
`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) UnsetMessagesIterator()`

UnsetMessagesIterator ensures that no value is present for MessagesIterator, not even an explicit nil
### GetMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesSize() int64`

GetMessagesSize returns the MessagesSize field if non-nil, zero value otherwise.

### GetMessagesSizeOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetMessagesSizeOk() (*int64, bool)`

GetMessagesSizeOk returns a tuple with the MessagesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesSize(v int64)`

SetMessagesSize sets MessagesSize field to given value.

### HasMessagesSize

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasMessagesSize() bool`

HasMessagesSize returns a boolean if a field has been set.

### GetNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetNodeFailuresTolerable() int64`

GetNodeFailuresTolerable returns the NodeFailuresTolerable field if non-nil, zero value otherwise.

### GetNodeFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetNodeFailuresTolerableOk() (*int64, bool)`

GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetNodeFailuresTolerable(v int64)`

SetNodeFailuresTolerable sets NodeFailuresTolerable field to given value.

### HasNodeFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasNodeFailuresTolerable() bool`

HasNodeFailuresTolerable returns a boolean if a field has been set.

### GetSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetSsdFailuresTolerable() int64`

GetSsdFailuresTolerable returns the SsdFailuresTolerable field if non-nil, zero value otherwise.

### GetSsdFailuresTolerableOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetSsdFailuresTolerableOk() (*int64, bool)`

GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetSsdFailuresTolerable(v int64)`

SetSsdFailuresTolerable sets SsdFailuresTolerable field to given value.

### HasSsdFailuresTolerable

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasSsdFailuresTolerable() bool`

HasSsdFailuresTolerable returns a boolean if a field has been set.

### GetState

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



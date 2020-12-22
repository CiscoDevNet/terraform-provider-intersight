# HyperflexStPlatformClusterResiliencyInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.StPlatformClusterResiliencyInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.StPlatformClusterResiliencyInfo"]
**HddFailuresTolerable** | Pointer to **int64** | The number of persistent storage device failures tolerable before the storage cluster becomes offline. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**MessagesIterator** | Pointer to **interface{}** | The current message describing the auto-healing process of the cluster. | [optional] [readonly] 
**MessagesSize** | Pointer to **int64** | The number of elements in the messages collection. | [optional] [readonly] 
**NodeFailuresTolerable** | Pointer to **int64** | The number of node failures tolerable before the storage cluster becomes offline. | [optional] [readonly] 
**SsdFailuresTolerable** | Pointer to **int64** | The number of caching device failures tolerable before the storage cluster becomes offline. | [optional] [readonly] 
**State** | Pointer to **string** | The resiliency state of the cluster. The resiliency status is &#39;HEALTHY&#39; if there are no failures and the storage cluster is fully operational. The resiliency status is &#39;WARNING&#39; when the cluster has experienced failures that may adversely affect the cluster. It is &#39;UNKNOWN&#39; if the cluster is offline or if the state cannot be determined. | [optional] [readonly] 

## Methods

### NewHyperflexStPlatformClusterResiliencyInfoAllOf

`func NewHyperflexStPlatformClusterResiliencyInfoAllOf(classId string, objectType string, ) *HyperflexStPlatformClusterResiliencyInfoAllOf`

NewHyperflexStPlatformClusterResiliencyInfoAllOf instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults

`func NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults() *HyperflexStPlatformClusterResiliencyInfoAllOf`

NewHyperflexStPlatformClusterResiliencyInfoAllOfWithDefaults instantiates a new HyperflexStPlatformClusterResiliencyInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetMessagesNil

`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *HyperflexStPlatformClusterResiliencyInfoAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
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



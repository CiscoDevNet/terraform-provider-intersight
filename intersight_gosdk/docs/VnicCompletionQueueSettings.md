# VnicCompletionQueueSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.CompletionQueueSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.CompletionQueueSettings"]
**Count** | Pointer to **int64** | The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources. | [optional] [default to 5]
**RingSize** | Pointer to **int64** | The number of descriptors in each completion queue. | [optional] [readonly] [default to 1]

## Methods

### NewVnicCompletionQueueSettings

`func NewVnicCompletionQueueSettings(classId string, objectType string, ) *VnicCompletionQueueSettings`

NewVnicCompletionQueueSettings instantiates a new VnicCompletionQueueSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicCompletionQueueSettingsWithDefaults

`func NewVnicCompletionQueueSettingsWithDefaults() *VnicCompletionQueueSettings`

NewVnicCompletionQueueSettingsWithDefaults instantiates a new VnicCompletionQueueSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicCompletionQueueSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicCompletionQueueSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicCompletionQueueSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicCompletionQueueSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicCompletionQueueSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicCompletionQueueSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *VnicCompletionQueueSettings) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicCompletionQueueSettings) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicCompletionQueueSettings) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicCompletionQueueSettings) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRingSize

`func (o *VnicCompletionQueueSettings) GetRingSize() int64`

GetRingSize returns the RingSize field if non-nil, zero value otherwise.

### GetRingSizeOk

`func (o *VnicCompletionQueueSettings) GetRingSizeOk() (*int64, bool)`

GetRingSizeOk returns a tuple with the RingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSize

`func (o *VnicCompletionQueueSettings) SetRingSize(v int64)`

SetRingSize sets RingSize field to given value.

### HasRingSize

`func (o *VnicCompletionQueueSettings) HasRingSize() bool`

HasRingSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BulkSubRequestsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.SubRequestsGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.SubRequestsGroup"]
**Requests** | Pointer to [**[]BulkSubRequest**](BulkSubRequest.md) |  | [optional] 

## Methods

### NewBulkSubRequestsGroup

`func NewBulkSubRequestsGroup(classId string, objectType string, ) *BulkSubRequestsGroup`

NewBulkSubRequestsGroup instantiates a new BulkSubRequestsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSubRequestsGroupWithDefaults

`func NewBulkSubRequestsGroupWithDefaults() *BulkSubRequestsGroup`

NewBulkSubRequestsGroupWithDefaults instantiates a new BulkSubRequestsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkSubRequestsGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkSubRequestsGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkSubRequestsGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkSubRequestsGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkSubRequestsGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkSubRequestsGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRequests

`func (o *BulkSubRequestsGroup) GetRequests() []BulkSubRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BulkSubRequestsGroup) GetRequestsOk() (*[]BulkSubRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BulkSubRequestsGroup) SetRequests(v []BulkSubRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BulkSubRequestsGroup) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *BulkSubRequestsGroup) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *BulkSubRequestsGroup) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



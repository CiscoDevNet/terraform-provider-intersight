# FeedbackFeedbackPostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "feedback.FeedbackPost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "feedback.FeedbackPost"]
**FeedbackData** | Pointer to [**NullableFeedbackFeedbackData**](FeedbackFeedbackData.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewFeedbackFeedbackPostAllOf

`func NewFeedbackFeedbackPostAllOf(classId string, objectType string, ) *FeedbackFeedbackPostAllOf`

NewFeedbackFeedbackPostAllOf instantiates a new FeedbackFeedbackPostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFeedbackPostAllOfWithDefaults

`func NewFeedbackFeedbackPostAllOfWithDefaults() *FeedbackFeedbackPostAllOf`

NewFeedbackFeedbackPostAllOfWithDefaults instantiates a new FeedbackFeedbackPostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FeedbackFeedbackPostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FeedbackFeedbackPostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FeedbackFeedbackPostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FeedbackFeedbackPostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FeedbackFeedbackPostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FeedbackFeedbackPostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeedbackData

`func (o *FeedbackFeedbackPostAllOf) GetFeedbackData() FeedbackFeedbackData`

GetFeedbackData returns the FeedbackData field if non-nil, zero value otherwise.

### GetFeedbackDataOk

`func (o *FeedbackFeedbackPostAllOf) GetFeedbackDataOk() (*FeedbackFeedbackData, bool)`

GetFeedbackDataOk returns a tuple with the FeedbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackData

`func (o *FeedbackFeedbackPostAllOf) SetFeedbackData(v FeedbackFeedbackData)`

SetFeedbackData sets FeedbackData field to given value.

### HasFeedbackData

`func (o *FeedbackFeedbackPostAllOf) HasFeedbackData() bool`

HasFeedbackData returns a boolean if a field has been set.

### SetFeedbackDataNil

`func (o *FeedbackFeedbackPostAllOf) SetFeedbackDataNil(b bool)`

 SetFeedbackDataNil sets the value for FeedbackData to be an explicit nil

### UnsetFeedbackData
`func (o *FeedbackFeedbackPostAllOf) UnsetFeedbackData()`

UnsetFeedbackData ensures that no value is present for FeedbackData, not even an explicit nil
### GetAccount

`func (o *FeedbackFeedbackPostAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *FeedbackFeedbackPostAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *FeedbackFeedbackPostAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *FeedbackFeedbackPostAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FeedbackFeedbackPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackData** | Pointer to [**FeedbackFeedbackData**](feedback.FeedbackData.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewFeedbackFeedbackPost

`func NewFeedbackFeedbackPost() *FeedbackFeedbackPost`

NewFeedbackFeedbackPost instantiates a new FeedbackFeedbackPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFeedbackPostWithDefaults

`func NewFeedbackFeedbackPostWithDefaults() *FeedbackFeedbackPost`

NewFeedbackFeedbackPostWithDefaults instantiates a new FeedbackFeedbackPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackData

`func (o *FeedbackFeedbackPost) GetFeedbackData() FeedbackFeedbackData`

GetFeedbackData returns the FeedbackData field if non-nil, zero value otherwise.

### GetFeedbackDataOk

`func (o *FeedbackFeedbackPost) GetFeedbackDataOk() (*FeedbackFeedbackData, bool)`

GetFeedbackDataOk returns a tuple with the FeedbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackData

`func (o *FeedbackFeedbackPost) SetFeedbackData(v FeedbackFeedbackData)`

SetFeedbackData sets FeedbackData field to given value.

### HasFeedbackData

`func (o *FeedbackFeedbackPost) HasFeedbackData() bool`

HasFeedbackData returns a boolean if a field has been set.

### GetAccount

`func (o *FeedbackFeedbackPost) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *FeedbackFeedbackPost) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *FeedbackFeedbackPost) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *FeedbackFeedbackPost) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



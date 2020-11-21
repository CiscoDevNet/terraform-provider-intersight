# FeedbackFeedbackData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "feedback.FeedbackData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "feedback.FeedbackData"]
**AccountName** | Pointer to **string** | Account name of the feedback sender. Copied in order to be persisted in case of account removal. | [optional] 
**AlternativeFollowUpEmails** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** | Text of the feedback as provided by the user, if it is a bug or a comment. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address details. | [optional] 
**Evaluation** | Pointer to **string** | Evalation rating as provided by the user to capture user sentiment regarding the issue. * &#x60;Excellent&#x60; - Option that specifies user&#39;s excelent evaluation. * &#x60;Poor&#x60; - Option that specifies user&#39;s poor evaluation. * &#x60;Fair&#x60; - Option that specifies user&#39;s fair evaluation. * &#x60;Good&#x60; - Option that specifies user&#39;s good evaluation. | [optional] [default to "Excellent"]
**FollowUp** | Pointer to **bool** | If a user is open for follow-up or not. | [optional] 
**TraceIds** | Pointer to **interface{}** | Bunch of last traceId for reproducing user last activity. | [optional] 
**Type** | Pointer to **string** | Type of the feedback from user. * &#x60;Evaluation&#x60; - User&#39;s feedback classified as an evaluation. * &#x60;Bug&#x60; - User&#39;s feedback classified as a bug. | [optional] [default to "Evaluation"]

## Methods

### NewFeedbackFeedbackData

`func NewFeedbackFeedbackData(classId string, objectType string, ) *FeedbackFeedbackData`

NewFeedbackFeedbackData instantiates a new FeedbackFeedbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFeedbackDataWithDefaults

`func NewFeedbackFeedbackDataWithDefaults() *FeedbackFeedbackData`

NewFeedbackFeedbackDataWithDefaults instantiates a new FeedbackFeedbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FeedbackFeedbackData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FeedbackFeedbackData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FeedbackFeedbackData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FeedbackFeedbackData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FeedbackFeedbackData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FeedbackFeedbackData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountName

`func (o *FeedbackFeedbackData) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *FeedbackFeedbackData) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *FeedbackFeedbackData) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *FeedbackFeedbackData) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAlternativeFollowUpEmails

`func (o *FeedbackFeedbackData) GetAlternativeFollowUpEmails() []string`

GetAlternativeFollowUpEmails returns the AlternativeFollowUpEmails field if non-nil, zero value otherwise.

### GetAlternativeFollowUpEmailsOk

`func (o *FeedbackFeedbackData) GetAlternativeFollowUpEmailsOk() (*[]string, bool)`

GetAlternativeFollowUpEmailsOk returns a tuple with the AlternativeFollowUpEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeFollowUpEmails

`func (o *FeedbackFeedbackData) SetAlternativeFollowUpEmails(v []string)`

SetAlternativeFollowUpEmails sets AlternativeFollowUpEmails field to given value.

### HasAlternativeFollowUpEmails

`func (o *FeedbackFeedbackData) HasAlternativeFollowUpEmails() bool`

HasAlternativeFollowUpEmails returns a boolean if a field has been set.

### SetAlternativeFollowUpEmailsNil

`func (o *FeedbackFeedbackData) SetAlternativeFollowUpEmailsNil(b bool)`

 SetAlternativeFollowUpEmailsNil sets the value for AlternativeFollowUpEmails to be an explicit nil

### UnsetAlternativeFollowUpEmails
`func (o *FeedbackFeedbackData) UnsetAlternativeFollowUpEmails()`

UnsetAlternativeFollowUpEmails ensures that no value is present for AlternativeFollowUpEmails, not even an explicit nil
### GetComment

`func (o *FeedbackFeedbackData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FeedbackFeedbackData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FeedbackFeedbackData) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FeedbackFeedbackData) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEmail

`func (o *FeedbackFeedbackData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FeedbackFeedbackData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FeedbackFeedbackData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FeedbackFeedbackData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEvaluation

`func (o *FeedbackFeedbackData) GetEvaluation() string`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *FeedbackFeedbackData) GetEvaluationOk() (*string, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *FeedbackFeedbackData) SetEvaluation(v string)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *FeedbackFeedbackData) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetFollowUp

`func (o *FeedbackFeedbackData) GetFollowUp() bool`

GetFollowUp returns the FollowUp field if non-nil, zero value otherwise.

### GetFollowUpOk

`func (o *FeedbackFeedbackData) GetFollowUpOk() (*bool, bool)`

GetFollowUpOk returns a tuple with the FollowUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUp

`func (o *FeedbackFeedbackData) SetFollowUp(v bool)`

SetFollowUp sets FollowUp field to given value.

### HasFollowUp

`func (o *FeedbackFeedbackData) HasFollowUp() bool`

HasFollowUp returns a boolean if a field has been set.

### GetTraceIds

`func (o *FeedbackFeedbackData) GetTraceIds() interface{}`

GetTraceIds returns the TraceIds field if non-nil, zero value otherwise.

### GetTraceIdsOk

`func (o *FeedbackFeedbackData) GetTraceIdsOk() (*interface{}, bool)`

GetTraceIdsOk returns a tuple with the TraceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceIds

`func (o *FeedbackFeedbackData) SetTraceIds(v interface{})`

SetTraceIds sets TraceIds field to given value.

### HasTraceIds

`func (o *FeedbackFeedbackData) HasTraceIds() bool`

HasTraceIds returns a boolean if a field has been set.

### SetTraceIdsNil

`func (o *FeedbackFeedbackData) SetTraceIdsNil(b bool)`

 SetTraceIdsNil sets the value for TraceIds to be an explicit nil

### UnsetTraceIds
`func (o *FeedbackFeedbackData) UnsetTraceIds()`

UnsetTraceIds ensures that no value is present for TraceIds, not even an explicit nil
### GetType

`func (o *FeedbackFeedbackData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedbackFeedbackData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedbackFeedbackData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeedbackFeedbackData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



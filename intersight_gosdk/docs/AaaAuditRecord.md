# AaaAuditRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information. | [optional] 
**InstId** | Pointer to **string** | The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before. | [optional] 
**SourceIp** | Pointer to **string** | The source IP of the client. | [optional] 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted. | [optional] [readonly] 
**UserIdOrEmail** | Pointer to **string** | The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Sessions** | Pointer to [**IamSessionRelationship**](iam.Session.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewAaaAuditRecord

`func NewAaaAuditRecord() *AaaAuditRecord`

NewAaaAuditRecord instantiates a new AaaAuditRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaAuditRecordWithDefaults

`func NewAaaAuditRecordWithDefaults() *AaaAuditRecord`

NewAaaAuditRecordWithDefaults instantiates a new AaaAuditRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AaaAuditRecord) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AaaAuditRecord) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AaaAuditRecord) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AaaAuditRecord) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInstId

`func (o *AaaAuditRecord) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *AaaAuditRecord) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *AaaAuditRecord) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *AaaAuditRecord) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSourceIp

`func (o *AaaAuditRecord) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *AaaAuditRecord) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *AaaAuditRecord) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *AaaAuditRecord) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetTimestamp

`func (o *AaaAuditRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AaaAuditRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AaaAuditRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AaaAuditRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *AaaAuditRecord) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *AaaAuditRecord) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *AaaAuditRecord) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *AaaAuditRecord) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetAccount

`func (o *AaaAuditRecord) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AaaAuditRecord) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AaaAuditRecord) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AaaAuditRecord) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSessions

`func (o *AaaAuditRecord) GetSessions() IamSessionRelationship`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AaaAuditRecord) GetSessionsOk() (*IamSessionRelationship, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AaaAuditRecord) SetSessions(v IamSessionRelationship)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AaaAuditRecord) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetUser

`func (o *AaaAuditRecord) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AaaAuditRecord) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AaaAuditRecord) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *AaaAuditRecord) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



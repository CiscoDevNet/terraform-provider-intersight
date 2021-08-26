# AaaAuditRecordAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "aaa.AuditRecord"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "aaa.AuditRecord"]
**Email** | Pointer to **string** | The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information. | [optional] 
**InstId** | Pointer to **string** | The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before. | [optional] 
**SessionId** | Pointer to **string** | The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information. | [optional] 
**SourceIp** | Pointer to **string** | The source IP of the client. | [optional] 
**Timestamp** | Pointer to **time.Time** | The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted. | [optional] [readonly] 
**UserIdOrEmail** | Pointer to **string** | The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Sessions** | Pointer to [**IamSessionRelationship**](IamSessionRelationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewAaaAuditRecordAllOf

`func NewAaaAuditRecordAllOf(classId string, objectType string, ) *AaaAuditRecordAllOf`

NewAaaAuditRecordAllOf instantiates a new AaaAuditRecordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaAuditRecordAllOfWithDefaults

`func NewAaaAuditRecordAllOfWithDefaults() *AaaAuditRecordAllOf`

NewAaaAuditRecordAllOfWithDefaults instantiates a new AaaAuditRecordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AaaAuditRecordAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AaaAuditRecordAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AaaAuditRecordAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AaaAuditRecordAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AaaAuditRecordAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AaaAuditRecordAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEmail

`func (o *AaaAuditRecordAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AaaAuditRecordAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AaaAuditRecordAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AaaAuditRecordAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInstId

`func (o *AaaAuditRecordAllOf) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *AaaAuditRecordAllOf) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *AaaAuditRecordAllOf) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *AaaAuditRecordAllOf) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSessionId

`func (o *AaaAuditRecordAllOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AaaAuditRecordAllOf) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AaaAuditRecordAllOf) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AaaAuditRecordAllOf) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceIp

`func (o *AaaAuditRecordAllOf) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *AaaAuditRecordAllOf) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *AaaAuditRecordAllOf) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *AaaAuditRecordAllOf) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetTimestamp

`func (o *AaaAuditRecordAllOf) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AaaAuditRecordAllOf) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AaaAuditRecordAllOf) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AaaAuditRecordAllOf) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *AaaAuditRecordAllOf) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *AaaAuditRecordAllOf) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *AaaAuditRecordAllOf) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *AaaAuditRecordAllOf) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetAccount

`func (o *AaaAuditRecordAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AaaAuditRecordAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AaaAuditRecordAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AaaAuditRecordAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSessions

`func (o *AaaAuditRecordAllOf) GetSessions() IamSessionRelationship`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AaaAuditRecordAllOf) GetSessionsOk() (*IamSessionRelationship, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AaaAuditRecordAllOf) SetSessions(v IamSessionRelationship)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AaaAuditRecordAllOf) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetUser

`func (o *AaaAuditRecordAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AaaAuditRecordAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AaaAuditRecordAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *AaaAuditRecordAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



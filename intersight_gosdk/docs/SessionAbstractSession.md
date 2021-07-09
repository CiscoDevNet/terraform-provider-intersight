# SessionAbstractSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ClientIpAddress** | Pointer to **string** | The user agent IP address from which the session is launched. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | The time at which the session ended. | [optional] [readonly] 
**Role** | Pointer to **string** | Role of the user who launched the session. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the session. * &#x60;Active&#x60; - The session is currently active. * &#x60;Ended&#x60; - The session has ended normally. * &#x60;Terminated&#x60; - The session was terminated by an admin. | [optional] [default to "Active"]
**UserIdOrEmail** | Pointer to **string** | User ID or E-mail Address of the user who launched the session. | [optional] [readonly] 

## Methods

### NewSessionAbstractSession

`func NewSessionAbstractSession(classId string, objectType string, ) *SessionAbstractSession`

NewSessionAbstractSession instantiates a new SessionAbstractSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAbstractSessionWithDefaults

`func NewSessionAbstractSessionWithDefaults() *SessionAbstractSession`

NewSessionAbstractSessionWithDefaults instantiates a new SessionAbstractSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SessionAbstractSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SessionAbstractSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SessionAbstractSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SessionAbstractSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SessionAbstractSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SessionAbstractSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientIpAddress

`func (o *SessionAbstractSession) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *SessionAbstractSession) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *SessionAbstractSession) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *SessionAbstractSession) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetEndTime

`func (o *SessionAbstractSession) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SessionAbstractSession) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SessionAbstractSession) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SessionAbstractSession) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRole

`func (o *SessionAbstractSession) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SessionAbstractSession) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SessionAbstractSession) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SessionAbstractSession) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *SessionAbstractSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionAbstractSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionAbstractSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionAbstractSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *SessionAbstractSession) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *SessionAbstractSession) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *SessionAbstractSession) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *SessionAbstractSession) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamUserQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserQualifier"]
**LastNotificationTime** | Pointer to **time.Time** | Holds the information on when last email notification was sent to the guest users. | [optional] [readonly] 
**NotificationScope** | Pointer to **string** | Defines the scope of email notifications for guest users. It determines which guest users  should receive email notifications about account access details. Options include notifying all users or only  newly added users. * &#x60;All&#x60; - Email Notification is sent to all users. * &#x60;New&#x60; - Email Notification is sent to newly added users. | [optional] [default to "All"]
**NotifyGuestUsers** | Pointer to **bool** | NotifyGuestUsers holds information on whether guest users have been notified about the guest access information. If set to true, all guest users will receive a email notification about the details of guest access link and instructions. | [optional] 
**UserDetails** | Pointer to [**[]IamUserDetails**](IamUserDetails.md) |  | [optional] 

## Methods

### NewIamUserQualifier

`func NewIamUserQualifier(classId string, objectType string, ) *IamUserQualifier`

NewIamUserQualifier instantiates a new IamUserQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserQualifierWithDefaults

`func NewIamUserQualifierWithDefaults() *IamUserQualifier`

NewIamUserQualifierWithDefaults instantiates a new IamUserQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastNotificationTime

`func (o *IamUserQualifier) GetLastNotificationTime() time.Time`

GetLastNotificationTime returns the LastNotificationTime field if non-nil, zero value otherwise.

### GetLastNotificationTimeOk

`func (o *IamUserQualifier) GetLastNotificationTimeOk() (*time.Time, bool)`

GetLastNotificationTimeOk returns a tuple with the LastNotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotificationTime

`func (o *IamUserQualifier) SetLastNotificationTime(v time.Time)`

SetLastNotificationTime sets LastNotificationTime field to given value.

### HasLastNotificationTime

`func (o *IamUserQualifier) HasLastNotificationTime() bool`

HasLastNotificationTime returns a boolean if a field has been set.

### GetNotificationScope

`func (o *IamUserQualifier) GetNotificationScope() string`

GetNotificationScope returns the NotificationScope field if non-nil, zero value otherwise.

### GetNotificationScopeOk

`func (o *IamUserQualifier) GetNotificationScopeOk() (*string, bool)`

GetNotificationScopeOk returns a tuple with the NotificationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationScope

`func (o *IamUserQualifier) SetNotificationScope(v string)`

SetNotificationScope sets NotificationScope field to given value.

### HasNotificationScope

`func (o *IamUserQualifier) HasNotificationScope() bool`

HasNotificationScope returns a boolean if a field has been set.

### GetNotifyGuestUsers

`func (o *IamUserQualifier) GetNotifyGuestUsers() bool`

GetNotifyGuestUsers returns the NotifyGuestUsers field if non-nil, zero value otherwise.

### GetNotifyGuestUsersOk

`func (o *IamUserQualifier) GetNotifyGuestUsersOk() (*bool, bool)`

GetNotifyGuestUsersOk returns a tuple with the NotifyGuestUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyGuestUsers

`func (o *IamUserQualifier) SetNotifyGuestUsers(v bool)`

SetNotifyGuestUsers sets NotifyGuestUsers field to given value.

### HasNotifyGuestUsers

`func (o *IamUserQualifier) HasNotifyGuestUsers() bool`

HasNotifyGuestUsers returns a boolean if a field has been set.

### GetUserDetails

`func (o *IamUserQualifier) GetUserDetails() []IamUserDetails`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *IamUserQualifier) GetUserDetailsOk() (*[]IamUserDetails, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *IamUserQualifier) SetUserDetails(v []IamUserDetails)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *IamUserQualifier) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.

### SetUserDetailsNil

`func (o *IamUserQualifier) SetUserDetailsNil(b bool)`

 SetUserDetailsNil sets the value for UserDetails to be an explicit nil

### UnsetUserDetails
`func (o *IamUserQualifier) UnsetUserDetails()`

UnsetUserDetails ensures that no value is present for UserDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



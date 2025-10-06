# IamUserDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserDetails"]
**EmailNotificationStatus** | Pointer to **string** | Stores the information on whether user has been sent an email about access details or not. * &#x60;NotSent&#x60; - The Account Admin might not choose send email notification to any of their guest users. * &#x60;Sent&#x60; - An email notification sent to notified users who are configured per guest access link. | [optional] [readonly] [default to "NotSent"]
**UserEmail** | Pointer to **string** | Stores the user email ID of the guest user. | [optional] 
**UserName** | Pointer to **string** | Stores the user name of the guest user. | [optional] 

## Methods

### NewIamUserDetails

`func NewIamUserDetails(classId string, objectType string, ) *IamUserDetails`

NewIamUserDetails instantiates a new IamUserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserDetailsWithDefaults

`func NewIamUserDetailsWithDefaults() *IamUserDetails`

NewIamUserDetailsWithDefaults instantiates a new IamUserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEmailNotificationStatus

`func (o *IamUserDetails) GetEmailNotificationStatus() string`

GetEmailNotificationStatus returns the EmailNotificationStatus field if non-nil, zero value otherwise.

### GetEmailNotificationStatusOk

`func (o *IamUserDetails) GetEmailNotificationStatusOk() (*string, bool)`

GetEmailNotificationStatusOk returns a tuple with the EmailNotificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationStatus

`func (o *IamUserDetails) SetEmailNotificationStatus(v string)`

SetEmailNotificationStatus sets EmailNotificationStatus field to given value.

### HasEmailNotificationStatus

`func (o *IamUserDetails) HasEmailNotificationStatus() bool`

HasEmailNotificationStatus returns a boolean if a field has been set.

### GetUserEmail

`func (o *IamUserDetails) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *IamUserDetails) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *IamUserDetails) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *IamUserDetails) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserName

`func (o *IamUserDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IamUserDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IamUserDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *IamUserDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



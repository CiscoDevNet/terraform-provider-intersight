# IamEndPointPasswordProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointPasswordProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointPasswordProperties"]
**EnablePasswordExpiry** | Pointer to **bool** | Enables password expiry on the endpoint. | [optional] [default to false]
**EnforceStrongPassword** | Pointer to **bool** | Enables a strong password policy. Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User&#39;s Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &amp;, *, -, _, +, &#x3D;). | [optional] [default to true]
**ForceSendPassword** | Pointer to **bool** | User password will always be sent to endpoint device. If the option is not selected, then user password will be sent to endpoint device for new users and if user password is changed for existing users. | [optional] [default to false]
**GracePeriod** | Pointer to **int64** | Time period until when you can use the existing password, after it expires. | [optional] [default to 0]
**NotificationPeriod** | Pointer to **int64** | The duration after which the password will expire. | [optional] [default to 15]
**PasswordExpiryDuration** | Pointer to **int64** | Set time period for password expiration. Value should be greater than notification period and grace period. | [optional] [default to 90]
**PasswordHistory** | Pointer to **int64** | Tracks password change history. Specifies in number of instances, that the new password was already used. | [optional] [default to 5]

## Methods

### NewIamEndPointPasswordProperties

`func NewIamEndPointPasswordProperties(classId string, objectType string, ) *IamEndPointPasswordProperties`

NewIamEndPointPasswordProperties instantiates a new IamEndPointPasswordProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointPasswordPropertiesWithDefaults

`func NewIamEndPointPasswordPropertiesWithDefaults() *IamEndPointPasswordProperties`

NewIamEndPointPasswordPropertiesWithDefaults instantiates a new IamEndPointPasswordProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointPasswordProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointPasswordProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointPasswordProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointPasswordProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointPasswordProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointPasswordProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnablePasswordExpiry

`func (o *IamEndPointPasswordProperties) GetEnablePasswordExpiry() bool`

GetEnablePasswordExpiry returns the EnablePasswordExpiry field if non-nil, zero value otherwise.

### GetEnablePasswordExpiryOk

`func (o *IamEndPointPasswordProperties) GetEnablePasswordExpiryOk() (*bool, bool)`

GetEnablePasswordExpiryOk returns a tuple with the EnablePasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordExpiry

`func (o *IamEndPointPasswordProperties) SetEnablePasswordExpiry(v bool)`

SetEnablePasswordExpiry sets EnablePasswordExpiry field to given value.

### HasEnablePasswordExpiry

`func (o *IamEndPointPasswordProperties) HasEnablePasswordExpiry() bool`

HasEnablePasswordExpiry returns a boolean if a field has been set.

### GetEnforceStrongPassword

`func (o *IamEndPointPasswordProperties) GetEnforceStrongPassword() bool`

GetEnforceStrongPassword returns the EnforceStrongPassword field if non-nil, zero value otherwise.

### GetEnforceStrongPasswordOk

`func (o *IamEndPointPasswordProperties) GetEnforceStrongPasswordOk() (*bool, bool)`

GetEnforceStrongPasswordOk returns a tuple with the EnforceStrongPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceStrongPassword

`func (o *IamEndPointPasswordProperties) SetEnforceStrongPassword(v bool)`

SetEnforceStrongPassword sets EnforceStrongPassword field to given value.

### HasEnforceStrongPassword

`func (o *IamEndPointPasswordProperties) HasEnforceStrongPassword() bool`

HasEnforceStrongPassword returns a boolean if a field has been set.

### GetForceSendPassword

`func (o *IamEndPointPasswordProperties) GetForceSendPassword() bool`

GetForceSendPassword returns the ForceSendPassword field if non-nil, zero value otherwise.

### GetForceSendPasswordOk

`func (o *IamEndPointPasswordProperties) GetForceSendPasswordOk() (*bool, bool)`

GetForceSendPasswordOk returns a tuple with the ForceSendPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSendPassword

`func (o *IamEndPointPasswordProperties) SetForceSendPassword(v bool)`

SetForceSendPassword sets ForceSendPassword field to given value.

### HasForceSendPassword

`func (o *IamEndPointPasswordProperties) HasForceSendPassword() bool`

HasForceSendPassword returns a boolean if a field has been set.

### GetGracePeriod

`func (o *IamEndPointPasswordProperties) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *IamEndPointPasswordProperties) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *IamEndPointPasswordProperties) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *IamEndPointPasswordProperties) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetNotificationPeriod

`func (o *IamEndPointPasswordProperties) GetNotificationPeriod() int64`

GetNotificationPeriod returns the NotificationPeriod field if non-nil, zero value otherwise.

### GetNotificationPeriodOk

`func (o *IamEndPointPasswordProperties) GetNotificationPeriodOk() (*int64, bool)`

GetNotificationPeriodOk returns a tuple with the NotificationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPeriod

`func (o *IamEndPointPasswordProperties) SetNotificationPeriod(v int64)`

SetNotificationPeriod sets NotificationPeriod field to given value.

### HasNotificationPeriod

`func (o *IamEndPointPasswordProperties) HasNotificationPeriod() bool`

HasNotificationPeriod returns a boolean if a field has been set.

### GetPasswordExpiryDuration

`func (o *IamEndPointPasswordProperties) GetPasswordExpiryDuration() int64`

GetPasswordExpiryDuration returns the PasswordExpiryDuration field if non-nil, zero value otherwise.

### GetPasswordExpiryDurationOk

`func (o *IamEndPointPasswordProperties) GetPasswordExpiryDurationOk() (*int64, bool)`

GetPasswordExpiryDurationOk returns a tuple with the PasswordExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiryDuration

`func (o *IamEndPointPasswordProperties) SetPasswordExpiryDuration(v int64)`

SetPasswordExpiryDuration sets PasswordExpiryDuration field to given value.

### HasPasswordExpiryDuration

`func (o *IamEndPointPasswordProperties) HasPasswordExpiryDuration() bool`

HasPasswordExpiryDuration returns a boolean if a field has been set.

### GetPasswordHistory

`func (o *IamEndPointPasswordProperties) GetPasswordHistory() int64`

GetPasswordHistory returns the PasswordHistory field if non-nil, zero value otherwise.

### GetPasswordHistoryOk

`func (o *IamEndPointPasswordProperties) GetPasswordHistoryOk() (*int64, bool)`

GetPasswordHistoryOk returns a tuple with the PasswordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistory

`func (o *IamEndPointPasswordProperties) SetPasswordHistory(v int64)`

SetPasswordHistory sets PasswordHistory field to given value.

### HasPasswordHistory

`func (o *IamEndPointPasswordProperties) HasPasswordHistory() bool`

HasPasswordHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



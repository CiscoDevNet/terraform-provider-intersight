# IamEndPointPasswordPropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePasswordExpiry** | Pointer to **bool** | Enables password expiry on the endpoint. | [optional] 
**EnforceStrongPassword** | Pointer to **bool** | Enables a strong password policy Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User&#39;s Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &amp;, *, -, _, +, &#x3D;). | [optional] 
**ForceSendPassword** | Pointer to **bool** | User password will always be sent to endpoint device. If the option is not selected, then users password will be sent to endpoint device if password is changed for existing users and for new users. | [optional] 
**GracePeriod** | Pointer to **int64** | Time period until when you can use the existing password, after it expires. | [optional] 
**NotificationPeriod** | Pointer to **int64** | The duration after which the password will expire. | [optional] 
**PasswordExpiryDuration** | Pointer to **int64** | Set time period for password expiration. Value should be greater than notification period and grace period. | [optional] 
**PasswordHistory** | Pointer to **int64** | Tracks password change history. Specifies in number of instances, that the new password was already used. | [optional] 

## Methods

### NewIamEndPointPasswordPropertiesAllOf

`func NewIamEndPointPasswordPropertiesAllOf() *IamEndPointPasswordPropertiesAllOf`

NewIamEndPointPasswordPropertiesAllOf instantiates a new IamEndPointPasswordPropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointPasswordPropertiesAllOfWithDefaults

`func NewIamEndPointPasswordPropertiesAllOfWithDefaults() *IamEndPointPasswordPropertiesAllOf`

NewIamEndPointPasswordPropertiesAllOfWithDefaults instantiates a new IamEndPointPasswordPropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePasswordExpiry

`func (o *IamEndPointPasswordPropertiesAllOf) GetEnablePasswordExpiry() bool`

GetEnablePasswordExpiry returns the EnablePasswordExpiry field if non-nil, zero value otherwise.

### GetEnablePasswordExpiryOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetEnablePasswordExpiryOk() (*bool, bool)`

GetEnablePasswordExpiryOk returns a tuple with the EnablePasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordExpiry

`func (o *IamEndPointPasswordPropertiesAllOf) SetEnablePasswordExpiry(v bool)`

SetEnablePasswordExpiry sets EnablePasswordExpiry field to given value.

### HasEnablePasswordExpiry

`func (o *IamEndPointPasswordPropertiesAllOf) HasEnablePasswordExpiry() bool`

HasEnablePasswordExpiry returns a boolean if a field has been set.

### GetEnforceStrongPassword

`func (o *IamEndPointPasswordPropertiesAllOf) GetEnforceStrongPassword() bool`

GetEnforceStrongPassword returns the EnforceStrongPassword field if non-nil, zero value otherwise.

### GetEnforceStrongPasswordOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetEnforceStrongPasswordOk() (*bool, bool)`

GetEnforceStrongPasswordOk returns a tuple with the EnforceStrongPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceStrongPassword

`func (o *IamEndPointPasswordPropertiesAllOf) SetEnforceStrongPassword(v bool)`

SetEnforceStrongPassword sets EnforceStrongPassword field to given value.

### HasEnforceStrongPassword

`func (o *IamEndPointPasswordPropertiesAllOf) HasEnforceStrongPassword() bool`

HasEnforceStrongPassword returns a boolean if a field has been set.

### GetForceSendPassword

`func (o *IamEndPointPasswordPropertiesAllOf) GetForceSendPassword() bool`

GetForceSendPassword returns the ForceSendPassword field if non-nil, zero value otherwise.

### GetForceSendPasswordOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetForceSendPasswordOk() (*bool, bool)`

GetForceSendPasswordOk returns a tuple with the ForceSendPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSendPassword

`func (o *IamEndPointPasswordPropertiesAllOf) SetForceSendPassword(v bool)`

SetForceSendPassword sets ForceSendPassword field to given value.

### HasForceSendPassword

`func (o *IamEndPointPasswordPropertiesAllOf) HasForceSendPassword() bool`

HasForceSendPassword returns a boolean if a field has been set.

### GetGracePeriod

`func (o *IamEndPointPasswordPropertiesAllOf) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *IamEndPointPasswordPropertiesAllOf) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *IamEndPointPasswordPropertiesAllOf) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetNotificationPeriod

`func (o *IamEndPointPasswordPropertiesAllOf) GetNotificationPeriod() int64`

GetNotificationPeriod returns the NotificationPeriod field if non-nil, zero value otherwise.

### GetNotificationPeriodOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetNotificationPeriodOk() (*int64, bool)`

GetNotificationPeriodOk returns a tuple with the NotificationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPeriod

`func (o *IamEndPointPasswordPropertiesAllOf) SetNotificationPeriod(v int64)`

SetNotificationPeriod sets NotificationPeriod field to given value.

### HasNotificationPeriod

`func (o *IamEndPointPasswordPropertiesAllOf) HasNotificationPeriod() bool`

HasNotificationPeriod returns a boolean if a field has been set.

### GetPasswordExpiryDuration

`func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordExpiryDuration() int64`

GetPasswordExpiryDuration returns the PasswordExpiryDuration field if non-nil, zero value otherwise.

### GetPasswordExpiryDurationOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordExpiryDurationOk() (*int64, bool)`

GetPasswordExpiryDurationOk returns a tuple with the PasswordExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiryDuration

`func (o *IamEndPointPasswordPropertiesAllOf) SetPasswordExpiryDuration(v int64)`

SetPasswordExpiryDuration sets PasswordExpiryDuration field to given value.

### HasPasswordExpiryDuration

`func (o *IamEndPointPasswordPropertiesAllOf) HasPasswordExpiryDuration() bool`

HasPasswordExpiryDuration returns a boolean if a field has been set.

### GetPasswordHistory

`func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordHistory() int64`

GetPasswordHistory returns the PasswordHistory field if non-nil, zero value otherwise.

### GetPasswordHistoryOk

`func (o *IamEndPointPasswordPropertiesAllOf) GetPasswordHistoryOk() (*int64, bool)`

GetPasswordHistoryOk returns a tuple with the PasswordHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistory

`func (o *IamEndPointPasswordPropertiesAllOf) SetPasswordHistory(v int64)`

SetPasswordHistory sets PasswordHistory field to given value.

### HasPasswordHistory

`func (o *IamEndPointPasswordPropertiesAllOf) HasPasswordHistory() bool`

HasPasswordHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



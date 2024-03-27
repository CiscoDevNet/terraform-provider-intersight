# IamLocalUserPasswordPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LocalUserPasswordPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LocalUserPasswordPolicy"]
**EnableLockOutForAdminUser** | Pointer to **bool** | Determines if the user lock out feature must be enabled for the local admin user. | [optional] [default to false]
**FailedLoginTrackerWindow** | Pointer to **int64** | Seconds are tracked for consecutive incorrect login attempts. Users will be locked out if they exceed the max number of incorrect login attempts during this duration. | [optional] [default to 1800]
**LockOutTimePeriod** | Pointer to **int64** | The time period, in seconds, during which a user account will remain locked. | [optional] [default to 900]
**MaxFailedLoginsAllowed** | Pointer to **int64** | Users will be locked out after exceeding the max consecutive incorrect login attempts allowed within the configured time duration. | [optional] [default to 5]
**MinCharDifference** | Pointer to **int64** | Minimum number of characters different from previous password. | [optional] [default to 0]
**MinDaysBetweenPasswordChange** | Pointer to **int64** | Minimum Days allowed between password change. | [optional] [default to 0]
**MinLengthPassword** | Pointer to **int64** | Minimum length of password. | [optional] [default to 8]
**MinLowerCase** | Pointer to **int64** | Minimum number of required lower case characters. | [optional] [default to 1]
**MinNumeric** | Pointer to **int64** | Minimum number of required numeric characters. | [optional] [default to 1]
**MinSpecialChar** | Pointer to **int64** | Minimum number of required special characters. | [optional] [default to 0]
**MinUpperCase** | Pointer to **int64** | Minimum number of required upper case characters. | [optional] [default to 1]
**NumPreviousPasswordsDisallowed** | Pointer to **int64** | Number of previous passwords disallowed. | [optional] [default to 0]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamLocalUserPasswordPolicyAllOf

`func NewIamLocalUserPasswordPolicyAllOf(classId string, objectType string, ) *IamLocalUserPasswordPolicyAllOf`

NewIamLocalUserPasswordPolicyAllOf instantiates a new IamLocalUserPasswordPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLocalUserPasswordPolicyAllOfWithDefaults

`func NewIamLocalUserPasswordPolicyAllOfWithDefaults() *IamLocalUserPasswordPolicyAllOf`

NewIamLocalUserPasswordPolicyAllOfWithDefaults instantiates a new IamLocalUserPasswordPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLocalUserPasswordPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLocalUserPasswordPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLocalUserPasswordPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLocalUserPasswordPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableLockOutForAdminUser

`func (o *IamLocalUserPasswordPolicyAllOf) GetEnableLockOutForAdminUser() bool`

GetEnableLockOutForAdminUser returns the EnableLockOutForAdminUser field if non-nil, zero value otherwise.

### GetEnableLockOutForAdminUserOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetEnableLockOutForAdminUserOk() (*bool, bool)`

GetEnableLockOutForAdminUserOk returns a tuple with the EnableLockOutForAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLockOutForAdminUser

`func (o *IamLocalUserPasswordPolicyAllOf) SetEnableLockOutForAdminUser(v bool)`

SetEnableLockOutForAdminUser sets EnableLockOutForAdminUser field to given value.

### HasEnableLockOutForAdminUser

`func (o *IamLocalUserPasswordPolicyAllOf) HasEnableLockOutForAdminUser() bool`

HasEnableLockOutForAdminUser returns a boolean if a field has been set.

### GetFailedLoginTrackerWindow

`func (o *IamLocalUserPasswordPolicyAllOf) GetFailedLoginTrackerWindow() int64`

GetFailedLoginTrackerWindow returns the FailedLoginTrackerWindow field if non-nil, zero value otherwise.

### GetFailedLoginTrackerWindowOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetFailedLoginTrackerWindowOk() (*int64, bool)`

GetFailedLoginTrackerWindowOk returns a tuple with the FailedLoginTrackerWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginTrackerWindow

`func (o *IamLocalUserPasswordPolicyAllOf) SetFailedLoginTrackerWindow(v int64)`

SetFailedLoginTrackerWindow sets FailedLoginTrackerWindow field to given value.

### HasFailedLoginTrackerWindow

`func (o *IamLocalUserPasswordPolicyAllOf) HasFailedLoginTrackerWindow() bool`

HasFailedLoginTrackerWindow returns a boolean if a field has been set.

### GetLockOutTimePeriod

`func (o *IamLocalUserPasswordPolicyAllOf) GetLockOutTimePeriod() int64`

GetLockOutTimePeriod returns the LockOutTimePeriod field if non-nil, zero value otherwise.

### GetLockOutTimePeriodOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetLockOutTimePeriodOk() (*int64, bool)`

GetLockOutTimePeriodOk returns a tuple with the LockOutTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOutTimePeriod

`func (o *IamLocalUserPasswordPolicyAllOf) SetLockOutTimePeriod(v int64)`

SetLockOutTimePeriod sets LockOutTimePeriod field to given value.

### HasLockOutTimePeriod

`func (o *IamLocalUserPasswordPolicyAllOf) HasLockOutTimePeriod() bool`

HasLockOutTimePeriod returns a boolean if a field has been set.

### GetMaxFailedLoginsAllowed

`func (o *IamLocalUserPasswordPolicyAllOf) GetMaxFailedLoginsAllowed() int64`

GetMaxFailedLoginsAllowed returns the MaxFailedLoginsAllowed field if non-nil, zero value otherwise.

### GetMaxFailedLoginsAllowedOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMaxFailedLoginsAllowedOk() (*int64, bool)`

GetMaxFailedLoginsAllowedOk returns a tuple with the MaxFailedLoginsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedLoginsAllowed

`func (o *IamLocalUserPasswordPolicyAllOf) SetMaxFailedLoginsAllowed(v int64)`

SetMaxFailedLoginsAllowed sets MaxFailedLoginsAllowed field to given value.

### HasMaxFailedLoginsAllowed

`func (o *IamLocalUserPasswordPolicyAllOf) HasMaxFailedLoginsAllowed() bool`

HasMaxFailedLoginsAllowed returns a boolean if a field has been set.

### GetMinCharDifference

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinCharDifference() int64`

GetMinCharDifference returns the MinCharDifference field if non-nil, zero value otherwise.

### GetMinCharDifferenceOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinCharDifferenceOk() (*int64, bool)`

GetMinCharDifferenceOk returns a tuple with the MinCharDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCharDifference

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinCharDifference(v int64)`

SetMinCharDifference sets MinCharDifference field to given value.

### HasMinCharDifference

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinCharDifference() bool`

HasMinCharDifference returns a boolean if a field has been set.

### GetMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinDaysBetweenPasswordChange() int64`

GetMinDaysBetweenPasswordChange returns the MinDaysBetweenPasswordChange field if non-nil, zero value otherwise.

### GetMinDaysBetweenPasswordChangeOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinDaysBetweenPasswordChangeOk() (*int64, bool)`

GetMinDaysBetweenPasswordChangeOk returns a tuple with the MinDaysBetweenPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinDaysBetweenPasswordChange(v int64)`

SetMinDaysBetweenPasswordChange sets MinDaysBetweenPasswordChange field to given value.

### HasMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinDaysBetweenPasswordChange() bool`

HasMinDaysBetweenPasswordChange returns a boolean if a field has been set.

### GetMinLengthPassword

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinLengthPassword() int64`

GetMinLengthPassword returns the MinLengthPassword field if non-nil, zero value otherwise.

### GetMinLengthPasswordOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinLengthPasswordOk() (*int64, bool)`

GetMinLengthPasswordOk returns a tuple with the MinLengthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLengthPassword

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinLengthPassword(v int64)`

SetMinLengthPassword sets MinLengthPassword field to given value.

### HasMinLengthPassword

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinLengthPassword() bool`

HasMinLengthPassword returns a boolean if a field has been set.

### GetMinLowerCase

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinLowerCase() int64`

GetMinLowerCase returns the MinLowerCase field if non-nil, zero value otherwise.

### GetMinLowerCaseOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinLowerCaseOk() (*int64, bool)`

GetMinLowerCaseOk returns a tuple with the MinLowerCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLowerCase

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinLowerCase(v int64)`

SetMinLowerCase sets MinLowerCase field to given value.

### HasMinLowerCase

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinLowerCase() bool`

HasMinLowerCase returns a boolean if a field has been set.

### GetMinNumeric

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinNumeric() int64`

GetMinNumeric returns the MinNumeric field if non-nil, zero value otherwise.

### GetMinNumericOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinNumericOk() (*int64, bool)`

GetMinNumericOk returns a tuple with the MinNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumeric

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinNumeric(v int64)`

SetMinNumeric sets MinNumeric field to given value.

### HasMinNumeric

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinNumeric() bool`

HasMinNumeric returns a boolean if a field has been set.

### GetMinSpecialChar

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinSpecialChar() int64`

GetMinSpecialChar returns the MinSpecialChar field if non-nil, zero value otherwise.

### GetMinSpecialCharOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinSpecialCharOk() (*int64, bool)`

GetMinSpecialCharOk returns a tuple with the MinSpecialChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpecialChar

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinSpecialChar(v int64)`

SetMinSpecialChar sets MinSpecialChar field to given value.

### HasMinSpecialChar

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinSpecialChar() bool`

HasMinSpecialChar returns a boolean if a field has been set.

### GetMinUpperCase

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinUpperCase() int64`

GetMinUpperCase returns the MinUpperCase field if non-nil, zero value otherwise.

### GetMinUpperCaseOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetMinUpperCaseOk() (*int64, bool)`

GetMinUpperCaseOk returns a tuple with the MinUpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpperCase

`func (o *IamLocalUserPasswordPolicyAllOf) SetMinUpperCase(v int64)`

SetMinUpperCase sets MinUpperCase field to given value.

### HasMinUpperCase

`func (o *IamLocalUserPasswordPolicyAllOf) HasMinUpperCase() bool`

HasMinUpperCase returns a boolean if a field has been set.

### GetNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicyAllOf) GetNumPreviousPasswordsDisallowed() int64`

GetNumPreviousPasswordsDisallowed returns the NumPreviousPasswordsDisallowed field if non-nil, zero value otherwise.

### GetNumPreviousPasswordsDisallowedOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetNumPreviousPasswordsDisallowedOk() (*int64, bool)`

GetNumPreviousPasswordsDisallowedOk returns a tuple with the NumPreviousPasswordsDisallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicyAllOf) SetNumPreviousPasswordsDisallowed(v int64)`

SetNumPreviousPasswordsDisallowed sets NumPreviousPasswordsDisallowed field to given value.

### HasNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicyAllOf) HasNumPreviousPasswordsDisallowed() bool`

HasNumPreviousPasswordsDisallowed returns a boolean if a field has been set.

### GetAccount

`func (o *IamLocalUserPasswordPolicyAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamLocalUserPasswordPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamLocalUserPasswordPolicyAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamLocalUserPasswordPolicyAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamLocalUserPasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LocalUserPasswordPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LocalUserPasswordPolicy"]
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

### NewIamLocalUserPasswordPolicy

`func NewIamLocalUserPasswordPolicy(classId string, objectType string, ) *IamLocalUserPasswordPolicy`

NewIamLocalUserPasswordPolicy instantiates a new IamLocalUserPasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLocalUserPasswordPolicyWithDefaults

`func NewIamLocalUserPasswordPolicyWithDefaults() *IamLocalUserPasswordPolicy`

NewIamLocalUserPasswordPolicyWithDefaults instantiates a new IamLocalUserPasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLocalUserPasswordPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLocalUserPasswordPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLocalUserPasswordPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLocalUserPasswordPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLocalUserPasswordPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLocalUserPasswordPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMinCharDifference

`func (o *IamLocalUserPasswordPolicy) GetMinCharDifference() int64`

GetMinCharDifference returns the MinCharDifference field if non-nil, zero value otherwise.

### GetMinCharDifferenceOk

`func (o *IamLocalUserPasswordPolicy) GetMinCharDifferenceOk() (*int64, bool)`

GetMinCharDifferenceOk returns a tuple with the MinCharDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCharDifference

`func (o *IamLocalUserPasswordPolicy) SetMinCharDifference(v int64)`

SetMinCharDifference sets MinCharDifference field to given value.

### HasMinCharDifference

`func (o *IamLocalUserPasswordPolicy) HasMinCharDifference() bool`

HasMinCharDifference returns a boolean if a field has been set.

### GetMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicy) GetMinDaysBetweenPasswordChange() int64`

GetMinDaysBetweenPasswordChange returns the MinDaysBetweenPasswordChange field if non-nil, zero value otherwise.

### GetMinDaysBetweenPasswordChangeOk

`func (o *IamLocalUserPasswordPolicy) GetMinDaysBetweenPasswordChangeOk() (*int64, bool)`

GetMinDaysBetweenPasswordChangeOk returns a tuple with the MinDaysBetweenPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicy) SetMinDaysBetweenPasswordChange(v int64)`

SetMinDaysBetweenPasswordChange sets MinDaysBetweenPasswordChange field to given value.

### HasMinDaysBetweenPasswordChange

`func (o *IamLocalUserPasswordPolicy) HasMinDaysBetweenPasswordChange() bool`

HasMinDaysBetweenPasswordChange returns a boolean if a field has been set.

### GetMinLengthPassword

`func (o *IamLocalUserPasswordPolicy) GetMinLengthPassword() int64`

GetMinLengthPassword returns the MinLengthPassword field if non-nil, zero value otherwise.

### GetMinLengthPasswordOk

`func (o *IamLocalUserPasswordPolicy) GetMinLengthPasswordOk() (*int64, bool)`

GetMinLengthPasswordOk returns a tuple with the MinLengthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLengthPassword

`func (o *IamLocalUserPasswordPolicy) SetMinLengthPassword(v int64)`

SetMinLengthPassword sets MinLengthPassword field to given value.

### HasMinLengthPassword

`func (o *IamLocalUserPasswordPolicy) HasMinLengthPassword() bool`

HasMinLengthPassword returns a boolean if a field has been set.

### GetMinLowerCase

`func (o *IamLocalUserPasswordPolicy) GetMinLowerCase() int64`

GetMinLowerCase returns the MinLowerCase field if non-nil, zero value otherwise.

### GetMinLowerCaseOk

`func (o *IamLocalUserPasswordPolicy) GetMinLowerCaseOk() (*int64, bool)`

GetMinLowerCaseOk returns a tuple with the MinLowerCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLowerCase

`func (o *IamLocalUserPasswordPolicy) SetMinLowerCase(v int64)`

SetMinLowerCase sets MinLowerCase field to given value.

### HasMinLowerCase

`func (o *IamLocalUserPasswordPolicy) HasMinLowerCase() bool`

HasMinLowerCase returns a boolean if a field has been set.

### GetMinNumeric

`func (o *IamLocalUserPasswordPolicy) GetMinNumeric() int64`

GetMinNumeric returns the MinNumeric field if non-nil, zero value otherwise.

### GetMinNumericOk

`func (o *IamLocalUserPasswordPolicy) GetMinNumericOk() (*int64, bool)`

GetMinNumericOk returns a tuple with the MinNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumeric

`func (o *IamLocalUserPasswordPolicy) SetMinNumeric(v int64)`

SetMinNumeric sets MinNumeric field to given value.

### HasMinNumeric

`func (o *IamLocalUserPasswordPolicy) HasMinNumeric() bool`

HasMinNumeric returns a boolean if a field has been set.

### GetMinSpecialChar

`func (o *IamLocalUserPasswordPolicy) GetMinSpecialChar() int64`

GetMinSpecialChar returns the MinSpecialChar field if non-nil, zero value otherwise.

### GetMinSpecialCharOk

`func (o *IamLocalUserPasswordPolicy) GetMinSpecialCharOk() (*int64, bool)`

GetMinSpecialCharOk returns a tuple with the MinSpecialChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpecialChar

`func (o *IamLocalUserPasswordPolicy) SetMinSpecialChar(v int64)`

SetMinSpecialChar sets MinSpecialChar field to given value.

### HasMinSpecialChar

`func (o *IamLocalUserPasswordPolicy) HasMinSpecialChar() bool`

HasMinSpecialChar returns a boolean if a field has been set.

### GetMinUpperCase

`func (o *IamLocalUserPasswordPolicy) GetMinUpperCase() int64`

GetMinUpperCase returns the MinUpperCase field if non-nil, zero value otherwise.

### GetMinUpperCaseOk

`func (o *IamLocalUserPasswordPolicy) GetMinUpperCaseOk() (*int64, bool)`

GetMinUpperCaseOk returns a tuple with the MinUpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpperCase

`func (o *IamLocalUserPasswordPolicy) SetMinUpperCase(v int64)`

SetMinUpperCase sets MinUpperCase field to given value.

### HasMinUpperCase

`func (o *IamLocalUserPasswordPolicy) HasMinUpperCase() bool`

HasMinUpperCase returns a boolean if a field has been set.

### GetNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicy) GetNumPreviousPasswordsDisallowed() int64`

GetNumPreviousPasswordsDisallowed returns the NumPreviousPasswordsDisallowed field if non-nil, zero value otherwise.

### GetNumPreviousPasswordsDisallowedOk

`func (o *IamLocalUserPasswordPolicy) GetNumPreviousPasswordsDisallowedOk() (*int64, bool)`

GetNumPreviousPasswordsDisallowedOk returns a tuple with the NumPreviousPasswordsDisallowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicy) SetNumPreviousPasswordsDisallowed(v int64)`

SetNumPreviousPasswordsDisallowed sets NumPreviousPasswordsDisallowed field to given value.

### HasNumPreviousPasswordsDisallowed

`func (o *IamLocalUserPasswordPolicy) HasNumPreviousPasswordsDisallowed() bool`

HasNumPreviousPasswordsDisallowed returns a boolean if a field has been set.

### GetAccount

`func (o *IamLocalUserPasswordPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamLocalUserPasswordPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamLocalUserPasswordPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamLocalUserPasswordPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



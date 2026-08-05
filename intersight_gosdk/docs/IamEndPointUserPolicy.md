# IamEndPointUserPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUserPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUserPolicy"]
**AccountLockoutDuration** | Pointer to **int64** | Timeout duration specifies the duration (in seconds) after which a locked account is automatically unlocked. - Set to 0 when accountUnlockMode is Manual. - Set a value between 1 and 604800 when accountUnlockMode is Automatic. | [optional] [default to 300]
**AccountLockoutThreshold** | Pointer to **int64** | Set Account Lockout Threshold for endpoint users. | [optional] [default to 5]
**AccountUnlockMode** | Pointer to **string** | Account unlock method specifies how the account is unlocked after it is locked: - Manual: Account must be manually unlocked by an administrator. - Automatic: Account unlocks automatically after a timeout duration. * &#x60;Automatic&#x60; - Set Automatic on the selected end point. * &#x60;Manual&#x60; - Set Manual on the selected end point. | [optional] [default to "Automatic"]
**PasswordProperties** | Pointer to [**NullableIamEndPointPasswordProperties**](IamEndPointPasswordProperties.md) |  | [optional] 
**EndPointUserRoles** | Pointer to [**[]IamEndPointUserRoleRelationship**](IamEndPointUserRoleRelationship.md) | An array of relationships to iamEndPointUserRole resources. | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewIamEndPointUserPolicy

`func NewIamEndPointUserPolicy(classId string, objectType string, ) *IamEndPointUserPolicy`

NewIamEndPointUserPolicy instantiates a new IamEndPointUserPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserPolicyWithDefaults

`func NewIamEndPointUserPolicyWithDefaults() *IamEndPointUserPolicy`

NewIamEndPointUserPolicyWithDefaults instantiates a new IamEndPointUserPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountLockoutDuration

`func (o *IamEndPointUserPolicy) GetAccountLockoutDuration() int64`

GetAccountLockoutDuration returns the AccountLockoutDuration field if non-nil, zero value otherwise.

### GetAccountLockoutDurationOk

`func (o *IamEndPointUserPolicy) GetAccountLockoutDurationOk() (*int64, bool)`

GetAccountLockoutDurationOk returns a tuple with the AccountLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockoutDuration

`func (o *IamEndPointUserPolicy) SetAccountLockoutDuration(v int64)`

SetAccountLockoutDuration sets AccountLockoutDuration field to given value.

### HasAccountLockoutDuration

`func (o *IamEndPointUserPolicy) HasAccountLockoutDuration() bool`

HasAccountLockoutDuration returns a boolean if a field has been set.

### GetAccountLockoutThreshold

`func (o *IamEndPointUserPolicy) GetAccountLockoutThreshold() int64`

GetAccountLockoutThreshold returns the AccountLockoutThreshold field if non-nil, zero value otherwise.

### GetAccountLockoutThresholdOk

`func (o *IamEndPointUserPolicy) GetAccountLockoutThresholdOk() (*int64, bool)`

GetAccountLockoutThresholdOk returns a tuple with the AccountLockoutThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockoutThreshold

`func (o *IamEndPointUserPolicy) SetAccountLockoutThreshold(v int64)`

SetAccountLockoutThreshold sets AccountLockoutThreshold field to given value.

### HasAccountLockoutThreshold

`func (o *IamEndPointUserPolicy) HasAccountLockoutThreshold() bool`

HasAccountLockoutThreshold returns a boolean if a field has been set.

### GetAccountUnlockMode

`func (o *IamEndPointUserPolicy) GetAccountUnlockMode() string`

GetAccountUnlockMode returns the AccountUnlockMode field if non-nil, zero value otherwise.

### GetAccountUnlockModeOk

`func (o *IamEndPointUserPolicy) GetAccountUnlockModeOk() (*string, bool)`

GetAccountUnlockModeOk returns a tuple with the AccountUnlockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUnlockMode

`func (o *IamEndPointUserPolicy) SetAccountUnlockMode(v string)`

SetAccountUnlockMode sets AccountUnlockMode field to given value.

### HasAccountUnlockMode

`func (o *IamEndPointUserPolicy) HasAccountUnlockMode() bool`

HasAccountUnlockMode returns a boolean if a field has been set.

### GetPasswordProperties

`func (o *IamEndPointUserPolicy) GetPasswordProperties() IamEndPointPasswordProperties`

GetPasswordProperties returns the PasswordProperties field if non-nil, zero value otherwise.

### GetPasswordPropertiesOk

`func (o *IamEndPointUserPolicy) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool)`

GetPasswordPropertiesOk returns a tuple with the PasswordProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProperties

`func (o *IamEndPointUserPolicy) SetPasswordProperties(v IamEndPointPasswordProperties)`

SetPasswordProperties sets PasswordProperties field to given value.

### HasPasswordProperties

`func (o *IamEndPointUserPolicy) HasPasswordProperties() bool`

HasPasswordProperties returns a boolean if a field has been set.

### SetPasswordPropertiesNil

`func (o *IamEndPointUserPolicy) SetPasswordPropertiesNil(b bool)`

 SetPasswordPropertiesNil sets the value for PasswordProperties to be an explicit nil

### UnsetPasswordProperties
`func (o *IamEndPointUserPolicy) UnsetPasswordProperties()`

UnsetPasswordProperties ensures that no value is present for PasswordProperties, not even an explicit nil
### GetEndPointUserRoles

`func (o *IamEndPointUserPolicy) GetEndPointUserRoles() []IamEndPointUserRoleRelationship`

GetEndPointUserRoles returns the EndPointUserRoles field if non-nil, zero value otherwise.

### GetEndPointUserRolesOk

`func (o *IamEndPointUserPolicy) GetEndPointUserRolesOk() (*[]IamEndPointUserRoleRelationship, bool)`

GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRoles

`func (o *IamEndPointUserPolicy) SetEndPointUserRoles(v []IamEndPointUserRoleRelationship)`

SetEndPointUserRoles sets EndPointUserRoles field to given value.

### HasEndPointUserRoles

`func (o *IamEndPointUserPolicy) HasEndPointUserRoles() bool`

HasEndPointUserRoles returns a boolean if a field has been set.

### SetEndPointUserRolesNil

`func (o *IamEndPointUserPolicy) SetEndPointUserRolesNil(b bool)`

 SetEndPointUserRolesNil sets the value for EndPointUserRoles to be an explicit nil

### UnsetEndPointUserRoles
`func (o *IamEndPointUserPolicy) UnsetEndPointUserRoles()`

UnsetEndPointUserRoles ensures that no value is present for EndPointUserRoles, not even an explicit nil
### GetOrganization

`func (o *IamEndPointUserPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamEndPointUserPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamEndPointUserPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamEndPointUserPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *IamEndPointUserPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *IamEndPointUserPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *IamEndPointUserPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IamEndPointUserPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IamEndPointUserPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IamEndPointUserPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *IamEndPointUserPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *IamEndPointUserPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamGuestAccessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.GuestAccessSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.GuestAccessSettings"]
**AllowedDomainNames** | Pointer to **[]string** |  | [optional] 
**MaxGuestAccessLinkShelfLife** | Pointer to **int64** | Stores the maximum duration (in seconds) during which guest access link remains valid and accessible. It is the maximum value that is set  at the account level which account admin can configure. Any guest access link that is set with expiration time beyond this property will be disallowed. The default value is set to 604800 seconds (7 days). | [optional] [default to 604800]
**OverrideIpAccessRestriction** | Pointer to **bool** | Stores an option for Account Admin to override IP Access Restriction if it is enabled in the Account. This option is used to disable IP Access restrictions for guest users logging in to the account, while restriction is enforced for other normal users (who are authenticated via SAML or LDAP). | [optional] 
**ResourceLimits** | Pointer to [**NullableIamResourceLimitsRelationship**](IamResourceLimitsRelationship.md) |  | [optional] 

## Methods

### NewIamGuestAccessSettings

`func NewIamGuestAccessSettings(classId string, objectType string, ) *IamGuestAccessSettings`

NewIamGuestAccessSettings instantiates a new IamGuestAccessSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamGuestAccessSettingsWithDefaults

`func NewIamGuestAccessSettingsWithDefaults() *IamGuestAccessSettings`

NewIamGuestAccessSettingsWithDefaults instantiates a new IamGuestAccessSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamGuestAccessSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamGuestAccessSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamGuestAccessSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamGuestAccessSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamGuestAccessSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamGuestAccessSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedDomainNames

`func (o *IamGuestAccessSettings) GetAllowedDomainNames() []string`

GetAllowedDomainNames returns the AllowedDomainNames field if non-nil, zero value otherwise.

### GetAllowedDomainNamesOk

`func (o *IamGuestAccessSettings) GetAllowedDomainNamesOk() (*[]string, bool)`

GetAllowedDomainNamesOk returns a tuple with the AllowedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomainNames

`func (o *IamGuestAccessSettings) SetAllowedDomainNames(v []string)`

SetAllowedDomainNames sets AllowedDomainNames field to given value.

### HasAllowedDomainNames

`func (o *IamGuestAccessSettings) HasAllowedDomainNames() bool`

HasAllowedDomainNames returns a boolean if a field has been set.

### SetAllowedDomainNamesNil

`func (o *IamGuestAccessSettings) SetAllowedDomainNamesNil(b bool)`

 SetAllowedDomainNamesNil sets the value for AllowedDomainNames to be an explicit nil

### UnsetAllowedDomainNames
`func (o *IamGuestAccessSettings) UnsetAllowedDomainNames()`

UnsetAllowedDomainNames ensures that no value is present for AllowedDomainNames, not even an explicit nil
### GetMaxGuestAccessLinkShelfLife

`func (o *IamGuestAccessSettings) GetMaxGuestAccessLinkShelfLife() int64`

GetMaxGuestAccessLinkShelfLife returns the MaxGuestAccessLinkShelfLife field if non-nil, zero value otherwise.

### GetMaxGuestAccessLinkShelfLifeOk

`func (o *IamGuestAccessSettings) GetMaxGuestAccessLinkShelfLifeOk() (*int64, bool)`

GetMaxGuestAccessLinkShelfLifeOk returns a tuple with the MaxGuestAccessLinkShelfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGuestAccessLinkShelfLife

`func (o *IamGuestAccessSettings) SetMaxGuestAccessLinkShelfLife(v int64)`

SetMaxGuestAccessLinkShelfLife sets MaxGuestAccessLinkShelfLife field to given value.

### HasMaxGuestAccessLinkShelfLife

`func (o *IamGuestAccessSettings) HasMaxGuestAccessLinkShelfLife() bool`

HasMaxGuestAccessLinkShelfLife returns a boolean if a field has been set.

### GetOverrideIpAccessRestriction

`func (o *IamGuestAccessSettings) GetOverrideIpAccessRestriction() bool`

GetOverrideIpAccessRestriction returns the OverrideIpAccessRestriction field if non-nil, zero value otherwise.

### GetOverrideIpAccessRestrictionOk

`func (o *IamGuestAccessSettings) GetOverrideIpAccessRestrictionOk() (*bool, bool)`

GetOverrideIpAccessRestrictionOk returns a tuple with the OverrideIpAccessRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpAccessRestriction

`func (o *IamGuestAccessSettings) SetOverrideIpAccessRestriction(v bool)`

SetOverrideIpAccessRestriction sets OverrideIpAccessRestriction field to given value.

### HasOverrideIpAccessRestriction

`func (o *IamGuestAccessSettings) HasOverrideIpAccessRestriction() bool`

HasOverrideIpAccessRestriction returns a boolean if a field has been set.

### GetResourceLimits

`func (o *IamGuestAccessSettings) GetResourceLimits() IamResourceLimitsRelationship`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *IamGuestAccessSettings) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *IamGuestAccessSettings) SetResourceLimits(v IamResourceLimitsRelationship)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *IamGuestAccessSettings) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.

### SetResourceLimitsNil

`func (o *IamGuestAccessSettings) SetResourceLimitsNil(b bool)`

 SetResourceLimitsNil sets the value for ResourceLimits to be an explicit nil

### UnsetResourceLimits
`func (o *IamGuestAccessSettings) UnsetResourceLimits()`

UnsetResourceLimits ensures that no value is present for ResourceLimits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamUserSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserSetting"]
**AllowUiSessionRecording** | Pointer to **bool** | UI preference of the user for Session Recording. | [optional] [default to true]
**UserIdOrEmail** | Pointer to **string** | UserID or email as configured in the IdP. | [optional] [readonly] 
**UserUniqueIdentifier** | Pointer to **string** | Unique id of the user used by the identity provider to store the user. | [optional] [readonly] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**NullableIamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 

## Methods

### NewIamUserSetting

`func NewIamUserSetting(classId string, objectType string, ) *IamUserSetting`

NewIamUserSetting instantiates a new IamUserSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserSettingWithDefaults

`func NewIamUserSettingWithDefaults() *IamUserSetting`

NewIamUserSettingWithDefaults instantiates a new IamUserSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowUiSessionRecording

`func (o *IamUserSetting) GetAllowUiSessionRecording() bool`

GetAllowUiSessionRecording returns the AllowUiSessionRecording field if non-nil, zero value otherwise.

### GetAllowUiSessionRecordingOk

`func (o *IamUserSetting) GetAllowUiSessionRecordingOk() (*bool, bool)`

GetAllowUiSessionRecordingOk returns a tuple with the AllowUiSessionRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUiSessionRecording

`func (o *IamUserSetting) SetAllowUiSessionRecording(v bool)`

SetAllowUiSessionRecording sets AllowUiSessionRecording field to given value.

### HasAllowUiSessionRecording

`func (o *IamUserSetting) HasAllowUiSessionRecording() bool`

HasAllowUiSessionRecording returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *IamUserSetting) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *IamUserSetting) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *IamUserSetting) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *IamUserSetting) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetUserUniqueIdentifier

`func (o *IamUserSetting) GetUserUniqueIdentifier() string`

GetUserUniqueIdentifier returns the UserUniqueIdentifier field if non-nil, zero value otherwise.

### GetUserUniqueIdentifierOk

`func (o *IamUserSetting) GetUserUniqueIdentifierOk() (*string, bool)`

GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUniqueIdentifier

`func (o *IamUserSetting) SetUserUniqueIdentifier(v string)`

SetUserUniqueIdentifier sets UserUniqueIdentifier field to given value.

### HasUserUniqueIdentifier

`func (o *IamUserSetting) HasUserUniqueIdentifier() bool`

HasUserUniqueIdentifier returns a boolean if a field has been set.

### GetIdp

`func (o *IamUserSetting) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUserSetting) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUserSetting) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUserSetting) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *IamUserSetting) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *IamUserSetting) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetIdpReference

`func (o *IamUserSetting) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *IamUserSetting) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *IamUserSetting) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *IamUserSetting) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

### SetIdpReferenceNil

`func (o *IamUserSetting) SetIdpReferenceNil(b bool)`

 SetIdpReferenceNil sets the value for IdpReference to be an explicit nil

### UnsetIdpReference
`func (o *IamUserSetting) UnsetIdpReference()`

UnsetIdpReference ensures that no value is present for IdpReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



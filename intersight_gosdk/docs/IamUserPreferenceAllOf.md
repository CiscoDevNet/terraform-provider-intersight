# IamUserPreferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserPreference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserPreference"]
**Preference** | Pointer to **interface{}** | UI preferences of the user. | [optional] 
**UserUniqueIdentifier** | Pointer to **string** | Unique id of the user used by the identity provider to store the user. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**IamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 

## Methods

### NewIamUserPreferenceAllOf

`func NewIamUserPreferenceAllOf(classId string, objectType string, ) *IamUserPreferenceAllOf`

NewIamUserPreferenceAllOf instantiates a new IamUserPreferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserPreferenceAllOfWithDefaults

`func NewIamUserPreferenceAllOfWithDefaults() *IamUserPreferenceAllOf`

NewIamUserPreferenceAllOfWithDefaults instantiates a new IamUserPreferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserPreferenceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserPreferenceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserPreferenceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserPreferenceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserPreferenceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserPreferenceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPreference

`func (o *IamUserPreferenceAllOf) GetPreference() interface{}`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *IamUserPreferenceAllOf) GetPreferenceOk() (*interface{}, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *IamUserPreferenceAllOf) SetPreference(v interface{})`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *IamUserPreferenceAllOf) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreferenceNil

`func (o *IamUserPreferenceAllOf) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *IamUserPreferenceAllOf) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetUserUniqueIdentifier

`func (o *IamUserPreferenceAllOf) GetUserUniqueIdentifier() string`

GetUserUniqueIdentifier returns the UserUniqueIdentifier field if non-nil, zero value otherwise.

### GetUserUniqueIdentifierOk

`func (o *IamUserPreferenceAllOf) GetUserUniqueIdentifierOk() (*string, bool)`

GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUniqueIdentifier

`func (o *IamUserPreferenceAllOf) SetUserUniqueIdentifier(v string)`

SetUserUniqueIdentifier sets UserUniqueIdentifier field to given value.

### HasUserUniqueIdentifier

`func (o *IamUserPreferenceAllOf) HasUserUniqueIdentifier() bool`

HasUserUniqueIdentifier returns a boolean if a field has been set.

### GetIdp

`func (o *IamUserPreferenceAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUserPreferenceAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUserPreferenceAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUserPreferenceAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpReference

`func (o *IamUserPreferenceAllOf) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *IamUserPreferenceAllOf) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *IamUserPreferenceAllOf) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *IamUserPreferenceAllOf) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamUserPreferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | Pointer to **interface{}** | UI preferences of the user. | [optional] 
**UserUniqueIdentifier** | Pointer to **string** | Unique id of the user used by the identity provider to store the user. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](iam.Idp.Relationship.md) |  | [optional] 
**IdpReference** | Pointer to [**IamIdpReferenceRelationship**](iam.IdpReference.Relationship.md) |  | [optional] 

## Methods

### NewIamUserPreferenceAllOf

`func NewIamUserPreferenceAllOf() *IamUserPreferenceAllOf`

NewIamUserPreferenceAllOf instantiates a new IamUserPreferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserPreferenceAllOfWithDefaults

`func NewIamUserPreferenceAllOfWithDefaults() *IamUserPreferenceAllOf`

NewIamUserPreferenceAllOfWithDefaults instantiates a new IamUserPreferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



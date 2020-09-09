# IamQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion. | [optional] [readonly] 
**Value** | Pointer to **[]string** |  | [optional] 
**Usergroup** | Pointer to [**IamUserGroupRelationship**](iam.UserGroup.Relationship.md) |  | [optional] 

## Methods

### NewIamQualifier

`func NewIamQualifier() *IamQualifier`

NewIamQualifier instantiates a new IamQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamQualifierWithDefaults

`func NewIamQualifierWithDefaults() *IamQualifier`

NewIamQualifierWithDefaults instantiates a new IamQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamQualifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamQualifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamQualifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamQualifier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *IamQualifier) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IamQualifier) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IamQualifier) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IamQualifier) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUsergroup

`func (o *IamQualifier) GetUsergroup() IamUserGroupRelationship`

GetUsergroup returns the Usergroup field if non-nil, zero value otherwise.

### GetUsergroupOk

`func (o *IamQualifier) GetUsergroupOk() (*IamUserGroupRelationship, bool)`

GetUsergroupOk returns a tuple with the Usergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroup

`func (o *IamQualifier) SetUsergroup(v IamUserGroupRelationship)`

SetUsergroup sets Usergroup field to given value.

### HasUsergroup

`func (o *IamQualifier) HasUsergroup() bool`

HasUsergroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



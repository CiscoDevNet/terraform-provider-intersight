# IamQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Qualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Qualifier"]
**Name** | Pointer to **string** | The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion. | [optional] [readonly] 
**Value** | Pointer to **[]string** |  | [optional] 
**Usergroup** | Pointer to [**IamUserGroupRelationship**](IamUserGroupRelationship.md) |  | [optional] 

## Methods

### NewIamQualifier

`func NewIamQualifier(classId string, objectType string, ) *IamQualifier`

NewIamQualifier instantiates a new IamQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamQualifierWithDefaults

`func NewIamQualifierWithDefaults() *IamQualifier`

NewIamQualifierWithDefaults instantiates a new IamQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetValueNil

`func (o *IamQualifier) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IamQualifier) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
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



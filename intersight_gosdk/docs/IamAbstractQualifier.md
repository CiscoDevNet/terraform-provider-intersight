# IamAbstractQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Usergroup** | Pointer to [**NullableIamUserGroupRelationship**](IamUserGroupRelationship.md) |  | [optional] 

## Methods

### NewIamAbstractQualifier

`func NewIamAbstractQualifier(classId string, objectType string, ) *IamAbstractQualifier`

NewIamAbstractQualifier instantiates a new IamAbstractQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAbstractQualifierWithDefaults

`func NewIamAbstractQualifierWithDefaults() *IamAbstractQualifier`

NewIamAbstractQualifierWithDefaults instantiates a new IamAbstractQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAbstractQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAbstractQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAbstractQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAbstractQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAbstractQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAbstractQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUsergroup

`func (o *IamAbstractQualifier) GetUsergroup() IamUserGroupRelationship`

GetUsergroup returns the Usergroup field if non-nil, zero value otherwise.

### GetUsergroupOk

`func (o *IamAbstractQualifier) GetUsergroupOk() (*IamUserGroupRelationship, bool)`

GetUsergroupOk returns a tuple with the Usergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroup

`func (o *IamAbstractQualifier) SetUsergroup(v IamUserGroupRelationship)`

SetUsergroup sets Usergroup field to given value.

### HasUsergroup

`func (o *IamAbstractQualifier) HasUsergroup() bool`

HasUsergroup returns a boolean if a field has been set.

### SetUsergroupNil

`func (o *IamAbstractQualifier) SetUsergroupNil(b bool)`

 SetUsergroupNil sets the value for Usergroup to be an explicit nil

### UnsetUsergroup
`func (o *IamAbstractQualifier) UnsetUsergroup()`

UnsetUsergroup ensures that no value is present for Usergroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



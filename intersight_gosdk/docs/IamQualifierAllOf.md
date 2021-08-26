# IamQualifierAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Qualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Qualifier"]
**Name** | Pointer to **string** | The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion. | [optional] [readonly] 
**Value** | Pointer to **[]string** |  | [optional] 
**Usergroup** | Pointer to [**IamUserGroupRelationship**](IamUserGroupRelationship.md) |  | [optional] 

## Methods

### NewIamQualifierAllOf

`func NewIamQualifierAllOf(classId string, objectType string, ) *IamQualifierAllOf`

NewIamQualifierAllOf instantiates a new IamQualifierAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamQualifierAllOfWithDefaults

`func NewIamQualifierAllOfWithDefaults() *IamQualifierAllOf`

NewIamQualifierAllOfWithDefaults instantiates a new IamQualifierAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamQualifierAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamQualifierAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamQualifierAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamQualifierAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamQualifierAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamQualifierAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamQualifierAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamQualifierAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamQualifierAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamQualifierAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *IamQualifierAllOf) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IamQualifierAllOf) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IamQualifierAllOf) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IamQualifierAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IamQualifierAllOf) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IamQualifierAllOf) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetUsergroup

`func (o *IamQualifierAllOf) GetUsergroup() IamUserGroupRelationship`

GetUsergroup returns the Usergroup field if non-nil, zero value otherwise.

### GetUsergroupOk

`func (o *IamQualifierAllOf) GetUsergroupOk() (*IamUserGroupRelationship, bool)`

GetUsergroupOk returns a tuple with the Usergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroup

`func (o *IamQualifierAllOf) SetUsergroup(v IamUserGroupRelationship)`

SetUsergroup sets Usergroup field to given value.

### HasUsergroup

`func (o *IamQualifierAllOf) HasUsergroup() bool`

HasUsergroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



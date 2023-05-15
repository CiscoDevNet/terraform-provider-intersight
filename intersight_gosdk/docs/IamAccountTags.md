# IamAccountTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.AccountTags"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.AccountTags"]
**Key** | Pointer to **string** | The string representation of a tag key. | [optional] 
**Value** | Pointer to **string** | The string representation of a tag value. | [optional] 

## Methods

### NewIamAccountTags

`func NewIamAccountTags(classId string, objectType string, ) *IamAccountTags`

NewIamAccountTags instantiates a new IamAccountTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountTagsWithDefaults

`func NewIamAccountTagsWithDefaults() *IamAccountTags`

NewIamAccountTagsWithDefaults instantiates a new IamAccountTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAccountTags) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAccountTags) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAccountTags) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAccountTags) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAccountTags) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAccountTags) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *IamAccountTags) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IamAccountTags) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IamAccountTags) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IamAccountTags) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *IamAccountTags) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IamAccountTags) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IamAccountTags) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IamAccountTags) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



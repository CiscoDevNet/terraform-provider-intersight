# ResourceMembershipHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.MembershipHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.MembershipHolder"]
**Name** | Pointer to **string** | The name of this resource membership holder. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewResourceMembershipHolder

`func NewResourceMembershipHolder(classId string, objectType string, ) *ResourceMembershipHolder`

NewResourceMembershipHolder instantiates a new ResourceMembershipHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipHolderWithDefaults

`func NewResourceMembershipHolderWithDefaults() *ResourceMembershipHolder`

NewResourceMembershipHolderWithDefaults instantiates a new ResourceMembershipHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceMembershipHolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceMembershipHolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceMembershipHolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceMembershipHolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceMembershipHolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceMembershipHolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *ResourceMembershipHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceMembershipHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceMembershipHolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceMembershipHolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceMembershipHolder) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceMembershipHolder) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceMembershipHolder) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceMembershipHolder) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



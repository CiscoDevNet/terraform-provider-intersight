# IamSharingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SharingRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SharingRule"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**SharedResource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**SharedWithResource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamSharingRule

`func NewIamSharingRule(classId string, objectType string, ) *IamSharingRule`

NewIamSharingRule instantiates a new IamSharingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSharingRuleWithDefaults

`func NewIamSharingRuleWithDefaults() *IamSharingRule`

NewIamSharingRuleWithDefaults instantiates a new IamSharingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSharingRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSharingRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSharingRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSharingRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSharingRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSharingRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *IamSharingRule) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamSharingRule) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamSharingRule) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamSharingRule) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamSharingRule) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamSharingRule) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetSharedResource

`func (o *IamSharingRule) GetSharedResource() MoBaseMoRelationship`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *IamSharingRule) GetSharedResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *IamSharingRule) SetSharedResource(v MoBaseMoRelationship)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *IamSharingRule) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.

### SetSharedResourceNil

`func (o *IamSharingRule) SetSharedResourceNil(b bool)`

 SetSharedResourceNil sets the value for SharedResource to be an explicit nil

### UnsetSharedResource
`func (o *IamSharingRule) UnsetSharedResource()`

UnsetSharedResource ensures that no value is present for SharedResource, not even an explicit nil
### GetSharedWithResource

`func (o *IamSharingRule) GetSharedWithResource() MoBaseMoRelationship`

GetSharedWithResource returns the SharedWithResource field if non-nil, zero value otherwise.

### GetSharedWithResourceOk

`func (o *IamSharingRule) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedWithResourceOk returns a tuple with the SharedWithResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithResource

`func (o *IamSharingRule) SetSharedWithResource(v MoBaseMoRelationship)`

SetSharedWithResource sets SharedWithResource field to given value.

### HasSharedWithResource

`func (o *IamSharingRule) HasSharedWithResource() bool`

HasSharedWithResource returns a boolean if a field has been set.

### SetSharedWithResourceNil

`func (o *IamSharingRule) SetSharedWithResourceNil(b bool)`

 SetSharedWithResourceNil sets the value for SharedWithResource to be an explicit nil

### UnsetSharedWithResource
`func (o *IamSharingRule) UnsetSharedWithResource()`

UnsetSharedWithResource ensures that no value is present for SharedWithResource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



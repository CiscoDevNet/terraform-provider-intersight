# IamSharingRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SharingRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SharingRule"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**SharedResource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**SharedWithResource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamSharingRuleAllOf

`func NewIamSharingRuleAllOf(classId string, objectType string, ) *IamSharingRuleAllOf`

NewIamSharingRuleAllOf instantiates a new IamSharingRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSharingRuleAllOfWithDefaults

`func NewIamSharingRuleAllOfWithDefaults() *IamSharingRuleAllOf`

NewIamSharingRuleAllOfWithDefaults instantiates a new IamSharingRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSharingRuleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSharingRuleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSharingRuleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSharingRuleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSharingRuleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSharingRuleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *IamSharingRuleAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamSharingRuleAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamSharingRuleAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamSharingRuleAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSharedResource

`func (o *IamSharingRuleAllOf) GetSharedResource() MoBaseMoRelationship`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *IamSharingRuleAllOf) GetSharedResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *IamSharingRuleAllOf) SetSharedResource(v MoBaseMoRelationship)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *IamSharingRuleAllOf) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.

### GetSharedWithResource

`func (o *IamSharingRuleAllOf) GetSharedWithResource() MoBaseMoRelationship`

GetSharedWithResource returns the SharedWithResource field if non-nil, zero value otherwise.

### GetSharedWithResourceOk

`func (o *IamSharingRuleAllOf) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedWithResourceOk returns a tuple with the SharedWithResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithResource

`func (o *IamSharingRuleAllOf) SetSharedWithResource(v MoBaseMoRelationship)`

SetSharedWithResource sets SharedWithResource field to given value.

### HasSharedWithResource

`func (o *IamSharingRuleAllOf) HasSharedWithResource() bool`

HasSharedWithResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



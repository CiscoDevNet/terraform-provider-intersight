# IamEndPointUserPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUserPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUserPolicyInventory"]
**PasswordProperties** | Pointer to [**NullableIamEndPointPasswordProperties**](IamEndPointPasswordProperties.md) |  | [optional] 
**EndPointUserRoles** | Pointer to [**[]IamEndPointUserRoleInventoryRelationship**](IamEndPointUserRoleInventoryRelationship.md) | An array of relationships to iamEndPointUserRoleInventory resources. | [optional] [readonly] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamEndPointUserPolicyInventory

`func NewIamEndPointUserPolicyInventory(classId string, objectType string, ) *IamEndPointUserPolicyInventory`

NewIamEndPointUserPolicyInventory instantiates a new IamEndPointUserPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserPolicyInventoryWithDefaults

`func NewIamEndPointUserPolicyInventoryWithDefaults() *IamEndPointUserPolicyInventory`

NewIamEndPointUserPolicyInventoryWithDefaults instantiates a new IamEndPointUserPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPasswordProperties

`func (o *IamEndPointUserPolicyInventory) GetPasswordProperties() IamEndPointPasswordProperties`

GetPasswordProperties returns the PasswordProperties field if non-nil, zero value otherwise.

### GetPasswordPropertiesOk

`func (o *IamEndPointUserPolicyInventory) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool)`

GetPasswordPropertiesOk returns a tuple with the PasswordProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProperties

`func (o *IamEndPointUserPolicyInventory) SetPasswordProperties(v IamEndPointPasswordProperties)`

SetPasswordProperties sets PasswordProperties field to given value.

### HasPasswordProperties

`func (o *IamEndPointUserPolicyInventory) HasPasswordProperties() bool`

HasPasswordProperties returns a boolean if a field has been set.

### SetPasswordPropertiesNil

`func (o *IamEndPointUserPolicyInventory) SetPasswordPropertiesNil(b bool)`

 SetPasswordPropertiesNil sets the value for PasswordProperties to be an explicit nil

### UnsetPasswordProperties
`func (o *IamEndPointUserPolicyInventory) UnsetPasswordProperties()`

UnsetPasswordProperties ensures that no value is present for PasswordProperties, not even an explicit nil
### GetEndPointUserRoles

`func (o *IamEndPointUserPolicyInventory) GetEndPointUserRoles() []IamEndPointUserRoleInventoryRelationship`

GetEndPointUserRoles returns the EndPointUserRoles field if non-nil, zero value otherwise.

### GetEndPointUserRolesOk

`func (o *IamEndPointUserPolicyInventory) GetEndPointUserRolesOk() (*[]IamEndPointUserRoleInventoryRelationship, bool)`

GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRoles

`func (o *IamEndPointUserPolicyInventory) SetEndPointUserRoles(v []IamEndPointUserRoleInventoryRelationship)`

SetEndPointUserRoles sets EndPointUserRoles field to given value.

### HasEndPointUserRoles

`func (o *IamEndPointUserPolicyInventory) HasEndPointUserRoles() bool`

HasEndPointUserRoles returns a boolean if a field has been set.

### SetEndPointUserRolesNil

`func (o *IamEndPointUserPolicyInventory) SetEndPointUserRolesNil(b bool)`

 SetEndPointUserRolesNil sets the value for EndPointUserRoles to be an explicit nil

### UnsetEndPointUserRoles
`func (o *IamEndPointUserPolicyInventory) UnsetEndPointUserRoles()`

UnsetEndPointUserRoles ensures that no value is present for EndPointUserRoles, not even an explicit nil
### GetTargetMo

`func (o *IamEndPointUserPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *IamEndPointUserPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *IamEndPointUserPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *IamEndPointUserPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *IamEndPointUserPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *IamEndPointUserPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



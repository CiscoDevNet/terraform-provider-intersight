# IamEndPointUserInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUserInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUserInventory"]
**Name** | Pointer to **string** | Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters. | [optional] [readonly] 
**UserId** | Pointer to **string** | UserId for the end point user. | [optional] [readonly] 
**EndPointUserRole** | Pointer to [**[]IamEndPointUserRoleInventoryRelationship**](IamEndPointUserRoleInventoryRelationship.md) | An array of relationships to iamEndPointUserRoleInventory resources. | [optional] [readonly] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamEndPointUserInventory

`func NewIamEndPointUserInventory(classId string, objectType string, ) *IamEndPointUserInventory`

NewIamEndPointUserInventory instantiates a new IamEndPointUserInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserInventoryWithDefaults

`func NewIamEndPointUserInventoryWithDefaults() *IamEndPointUserInventory`

NewIamEndPointUserInventoryWithDefaults instantiates a new IamEndPointUserInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamEndPointUserInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointUserInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointUserInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointUserInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *IamEndPointUserInventory) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IamEndPointUserInventory) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IamEndPointUserInventory) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *IamEndPointUserInventory) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEndPointUserRole

`func (o *IamEndPointUserInventory) GetEndPointUserRole() []IamEndPointUserRoleInventoryRelationship`

GetEndPointUserRole returns the EndPointUserRole field if non-nil, zero value otherwise.

### GetEndPointUserRoleOk

`func (o *IamEndPointUserInventory) GetEndPointUserRoleOk() (*[]IamEndPointUserRoleInventoryRelationship, bool)`

GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRole

`func (o *IamEndPointUserInventory) SetEndPointUserRole(v []IamEndPointUserRoleInventoryRelationship)`

SetEndPointUserRole sets EndPointUserRole field to given value.

### HasEndPointUserRole

`func (o *IamEndPointUserInventory) HasEndPointUserRole() bool`

HasEndPointUserRole returns a boolean if a field has been set.

### SetEndPointUserRoleNil

`func (o *IamEndPointUserInventory) SetEndPointUserRoleNil(b bool)`

 SetEndPointUserRoleNil sets the value for EndPointUserRole to be an explicit nil

### UnsetEndPointUserRole
`func (o *IamEndPointUserInventory) UnsetEndPointUserRole()`

UnsetEndPointUserRole ensures that no value is present for EndPointUserRole, not even an explicit nil
### GetTargetMo

`func (o *IamEndPointUserInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *IamEndPointUserInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *IamEndPointUserInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *IamEndPointUserInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *IamEndPointUserInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *IamEndPointUserInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ResourcePermission"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ResourcePermission"]
**PermissionRoles** | Pointer to [**[]IamPermissionToRoles**](IamPermissionToRoles.md) |  | [optional] 
**TargetApp** | Pointer to **string** | Name of the service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**IamSecurityHolderRelationship**](IamSecurityHolderRelationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamResourcePermission

`func NewIamResourcePermission(classId string, objectType string, ) *IamResourcePermission`

NewIamResourcePermission instantiates a new IamResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourcePermissionWithDefaults

`func NewIamResourcePermissionWithDefaults() *IamResourcePermission`

NewIamResourcePermissionWithDefaults instantiates a new IamResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamResourcePermission) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamResourcePermission) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamResourcePermission) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamResourcePermission) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamResourcePermission) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamResourcePermission) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPermissionRoles

`func (o *IamResourcePermission) GetPermissionRoles() []IamPermissionToRoles`

GetPermissionRoles returns the PermissionRoles field if non-nil, zero value otherwise.

### GetPermissionRolesOk

`func (o *IamResourcePermission) GetPermissionRolesOk() (*[]IamPermissionToRoles, bool)`

GetPermissionRolesOk returns a tuple with the PermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionRoles

`func (o *IamResourcePermission) SetPermissionRoles(v []IamPermissionToRoles)`

SetPermissionRoles sets PermissionRoles field to given value.

### HasPermissionRoles

`func (o *IamResourcePermission) HasPermissionRoles() bool`

HasPermissionRoles returns a boolean if a field has been set.

### SetPermissionRolesNil

`func (o *IamResourcePermission) SetPermissionRolesNil(b bool)`

 SetPermissionRolesNil sets the value for PermissionRoles to be an explicit nil

### UnsetPermissionRoles
`func (o *IamResourcePermission) UnsetPermissionRoles()`

UnsetPermissionRoles ensures that no value is present for PermissionRoles, not even an explicit nil
### GetTargetApp

`func (o *IamResourcePermission) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *IamResourcePermission) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *IamResourcePermission) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *IamResourcePermission) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *IamResourcePermission) GetHolder() IamSecurityHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IamResourcePermission) GetHolderOk() (*IamSecurityHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IamResourcePermission) SetHolder(v IamSecurityHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IamResourcePermission) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetResource

`func (o *IamResourcePermission) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamResourcePermission) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamResourcePermission) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamResourcePermission) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



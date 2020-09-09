# ResourceMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupPermissionRoles** | Pointer to [**[]IamGroupPermissionToRoles**](iam.GroupPermissionToRoles.md) |  | [optional] 
**TargetApp** | Pointer to **string** | Name of the Service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**ResourceMembershipHolderRelationship**](resource.MembershipHolder.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewResourceMembership

`func NewResourceMembership() *ResourceMembership`

NewResourceMembership instantiates a new ResourceMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipWithDefaults

`func NewResourceMembershipWithDefaults() *ResourceMembership`

NewResourceMembershipWithDefaults instantiates a new ResourceMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupPermissionRoles

`func (o *ResourceMembership) GetGroupPermissionRoles() []IamGroupPermissionToRoles`

GetGroupPermissionRoles returns the GroupPermissionRoles field if non-nil, zero value otherwise.

### GetGroupPermissionRolesOk

`func (o *ResourceMembership) GetGroupPermissionRolesOk() (*[]IamGroupPermissionToRoles, bool)`

GetGroupPermissionRolesOk returns a tuple with the GroupPermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPermissionRoles

`func (o *ResourceMembership) SetGroupPermissionRoles(v []IamGroupPermissionToRoles)`

SetGroupPermissionRoles sets GroupPermissionRoles field to given value.

### HasGroupPermissionRoles

`func (o *ResourceMembership) HasGroupPermissionRoles() bool`

HasGroupPermissionRoles returns a boolean if a field has been set.

### GetTargetApp

`func (o *ResourceMembership) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *ResourceMembership) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *ResourceMembership) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *ResourceMembership) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *ResourceMembership) GetHolder() ResourceMembershipHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *ResourceMembership) GetHolderOk() (*ResourceMembershipHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *ResourceMembership) SetHolder(v ResourceMembershipHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *ResourceMembership) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetResource

`func (o *ResourceMembership) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceMembership) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceMembership) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceMembership) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



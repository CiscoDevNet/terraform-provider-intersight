# ResourceMembershipAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.Membership"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.Membership"]
**GroupPermissionRoles** | Pointer to [**[]IamGroupPermissionToRoles**](IamGroupPermissionToRoles.md) |  | [optional] 
**TargetApp** | Pointer to **string** | Name of the Service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**ResourceMembershipHolderRelationship**](ResourceMembershipHolderRelationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourceMembershipAllOf

`func NewResourceMembershipAllOf(classId string, objectType string, ) *ResourceMembershipAllOf`

NewResourceMembershipAllOf instantiates a new ResourceMembershipAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipAllOfWithDefaults

`func NewResourceMembershipAllOfWithDefaults() *ResourceMembershipAllOf`

NewResourceMembershipAllOfWithDefaults instantiates a new ResourceMembershipAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceMembershipAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceMembershipAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceMembershipAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceMembershipAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceMembershipAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceMembershipAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGroupPermissionRoles

`func (o *ResourceMembershipAllOf) GetGroupPermissionRoles() []IamGroupPermissionToRoles`

GetGroupPermissionRoles returns the GroupPermissionRoles field if non-nil, zero value otherwise.

### GetGroupPermissionRolesOk

`func (o *ResourceMembershipAllOf) GetGroupPermissionRolesOk() (*[]IamGroupPermissionToRoles, bool)`

GetGroupPermissionRolesOk returns a tuple with the GroupPermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPermissionRoles

`func (o *ResourceMembershipAllOf) SetGroupPermissionRoles(v []IamGroupPermissionToRoles)`

SetGroupPermissionRoles sets GroupPermissionRoles field to given value.

### HasGroupPermissionRoles

`func (o *ResourceMembershipAllOf) HasGroupPermissionRoles() bool`

HasGroupPermissionRoles returns a boolean if a field has been set.

### SetGroupPermissionRolesNil

`func (o *ResourceMembershipAllOf) SetGroupPermissionRolesNil(b bool)`

 SetGroupPermissionRolesNil sets the value for GroupPermissionRoles to be an explicit nil

### UnsetGroupPermissionRoles
`func (o *ResourceMembershipAllOf) UnsetGroupPermissionRoles()`

UnsetGroupPermissionRoles ensures that no value is present for GroupPermissionRoles, not even an explicit nil
### GetTargetApp

`func (o *ResourceMembershipAllOf) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *ResourceMembershipAllOf) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *ResourceMembershipAllOf) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *ResourceMembershipAllOf) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *ResourceMembershipAllOf) GetHolder() ResourceMembershipHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *ResourceMembershipAllOf) GetHolderOk() (*ResourceMembershipHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *ResourceMembershipAllOf) SetHolder(v ResourceMembershipHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *ResourceMembershipAllOf) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetResource

`func (o *ResourceMembershipAllOf) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceMembershipAllOf) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceMembershipAllOf) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceMembershipAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



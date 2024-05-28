# ResourceReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.Reservation"]
**MarkFail** | Pointer to **bool** | MarkFail is used to set the reservation status to Failed. | [optional] 
**ResourceMoids** | Pointer to **[]string** |  | [optional] 
**UserMoid** | Pointer to **string** | Moid of the user who created the reservation. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**CustomPermissionResources** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Groups** | Pointer to [**[]ResourceGroupRelationship**](ResourceGroupRelationship.md) | An array of relationships to resourceGroup resources. | [optional] 

## Methods

### NewResourceReservation

`func NewResourceReservation(classId string, objectType string, ) *ResourceReservation`

NewResourceReservation instantiates a new ResourceReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReservationWithDefaults

`func NewResourceReservationWithDefaults() *ResourceReservation`

NewResourceReservationWithDefaults instantiates a new ResourceReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMarkFail

`func (o *ResourceReservation) GetMarkFail() bool`

GetMarkFail returns the MarkFail field if non-nil, zero value otherwise.

### GetMarkFailOk

`func (o *ResourceReservation) GetMarkFailOk() (*bool, bool)`

GetMarkFailOk returns a tuple with the MarkFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkFail

`func (o *ResourceReservation) SetMarkFail(v bool)`

SetMarkFail sets MarkFail field to given value.

### HasMarkFail

`func (o *ResourceReservation) HasMarkFail() bool`

HasMarkFail returns a boolean if a field has been set.

### GetResourceMoids

`func (o *ResourceReservation) GetResourceMoids() []string`

GetResourceMoids returns the ResourceMoids field if non-nil, zero value otherwise.

### GetResourceMoidsOk

`func (o *ResourceReservation) GetResourceMoidsOk() (*[]string, bool)`

GetResourceMoidsOk returns a tuple with the ResourceMoids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMoids

`func (o *ResourceReservation) SetResourceMoids(v []string)`

SetResourceMoids sets ResourceMoids field to given value.

### HasResourceMoids

`func (o *ResourceReservation) HasResourceMoids() bool`

HasResourceMoids returns a boolean if a field has been set.

### SetResourceMoidsNil

`func (o *ResourceReservation) SetResourceMoidsNil(b bool)`

 SetResourceMoidsNil sets the value for ResourceMoids to be an explicit nil

### UnsetResourceMoids
`func (o *ResourceReservation) UnsetResourceMoids()`

UnsetResourceMoids ensures that no value is present for ResourceMoids, not even an explicit nil
### GetUserMoid

`func (o *ResourceReservation) GetUserMoid() string`

GetUserMoid returns the UserMoid field if non-nil, zero value otherwise.

### GetUserMoidOk

`func (o *ResourceReservation) GetUserMoidOk() (*string, bool)`

GetUserMoidOk returns a tuple with the UserMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMoid

`func (o *ResourceReservation) SetUserMoid(v string)`

SetUserMoid sets UserMoid field to given value.

### HasUserMoid

`func (o *ResourceReservation) HasUserMoid() bool`

HasUserMoid returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceReservation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceReservation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceReservation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceReservation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ResourceReservation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ResourceReservation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCustomPermissionResources

`func (o *ResourceReservation) GetCustomPermissionResources() []MoBaseMoRelationship`

GetCustomPermissionResources returns the CustomPermissionResources field if non-nil, zero value otherwise.

### GetCustomPermissionResourcesOk

`func (o *ResourceReservation) GetCustomPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPermissionResources

`func (o *ResourceReservation) SetCustomPermissionResources(v []MoBaseMoRelationship)`

SetCustomPermissionResources sets CustomPermissionResources field to given value.

### HasCustomPermissionResources

`func (o *ResourceReservation) HasCustomPermissionResources() bool`

HasCustomPermissionResources returns a boolean if a field has been set.

### SetCustomPermissionResourcesNil

`func (o *ResourceReservation) SetCustomPermissionResourcesNil(b bool)`

 SetCustomPermissionResourcesNil sets the value for CustomPermissionResources to be an explicit nil

### UnsetCustomPermissionResources
`func (o *ResourceReservation) UnsetCustomPermissionResources()`

UnsetCustomPermissionResources ensures that no value is present for CustomPermissionResources, not even an explicit nil
### GetGroups

`func (o *ResourceReservation) GetGroups() []ResourceGroupRelationship`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ResourceReservation) GetGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ResourceReservation) SetGroups(v []ResourceGroupRelationship)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ResourceReservation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ResourceReservation) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ResourceReservation) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


